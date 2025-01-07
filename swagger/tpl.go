package swagger

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io"
	"slices"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v2"
)

var (
	//go:embed default_tpl.md
	DefaultTelContent string
)

type ExecuteConfig struct {
	TemplateFile          string
	IgnoreDefaultResponse bool
}

func Execute(swaggerData []byte, format string, output io.Writer, cfg ExecuteConfig) (err error) {
	tplEvaluator := tplEvaluator{
		ExecuteConfig: cfg,
	}

	if err = tplEvaluator.ParseSwagger(swaggerData, format); err != nil {
		return fmt.Errorf("parsing swagger file: %w", err)
	}

	var tpl = template.New("swagger-tpl").Funcs(tplEvaluator.FuncMap())
	if cfg.TemplateFile == "" {
		tpl, err = tpl.Parse(DefaultTelContent)
	} else {
		tpl, err = tpl.ParseFiles(cfg.TemplateFile)
	}
	if err != nil {
		return fmt.Errorf("parse tpl %w", err)
	}

	return tpl.Execute(output, tplEvaluator.Swagger)
}

type tplEvaluator struct {
	ExecuteConfig
	Swagger
}

func (t *tplEvaluator) ParseSwagger(swaggerData []byte, format string) (err error) {
	switch format {
	case "yaml", "yml":
		err = yaml.Unmarshal(swaggerData, &t.Swagger)
	default:
		err = json.Unmarshal(swaggerData, &t.Swagger)
	}
	if t.IgnoreDefaultResponse {
		for k := range t.Swagger.Paths {
			for p := range t.Swagger.Paths[k] {
				delete(t.Swagger.Paths[k][p].Responses, "default")
			}
		}
	}
	return
}

func (t *tplEvaluator) FuncMap() template.FuncMap {
	return template.FuncMap{
		"Upper":           strings.ToUpper,
		"Lower":           strings.ToLower,
		"ExpressParam":    t.ExpressParam,
		"ExpressResponse": t.ExpressResponse,
	}
}

func (t *tplEvaluator) GetRef(ref string) Schema {
	return t.Swagger.Definitions[strings.TrimPrefix(ref, "#/definitions/")]
}

func (t *tplEvaluator) ParamType(p Param) string {
	if p.Type == "array" {
		return "[]" + p.Items.Type
	}
	return p.Type
}

type Express struct {
	Name        string
	Type        string
	Required    bool
	Description string `yaml:"description"`
}

func (t *tplEvaluator) ExpressParam(p Param) []Express {
	expressParams := []Express{}

	if p.Name == "body" {
		p.Name = ""
	}

	if p.Ref != nil {
		schema := t.GetRef(*p.Ref)
		expressParams = append(expressParams, t.ExpressSchema(schema, p.Required, appendName(nil, p.Name)...)...)
	}

	if p.Schema != nil {
		expressParams = append(expressParams, t.ExpressSchema(*p.Schema, p.Required, appendName(nil, p.Name)...)...)
	}

	if p.Name != "" {
		expressParams = append(expressParams, Express{Name: p.Name, Type: t.ParamType(p), Required: p.Required, Description: Escape(p.Description)})
	}

	return expressParams
}

func (t *tplEvaluator) ExpressResponse(r Response) []Express {
	if r.Schema == nil {
		return []Express{}
	}
	return t.ExpressSchema(*r.Schema, true)
}

func (t *tplEvaluator) ExpressSchema(s Schema, required bool, parent ...string) (expressParams []Express) {
	if s.Ref != "" {
		schema := t.GetRef(s.Ref)
		schema.Format = strings.TrimPrefix(s.Ref, "#/definitions/")

		// 判断循环引用类型
		ps := s
		for {
			if ps.parent == nil {
				break
			}
			ps = *ps.parent
			if ps.Ref != "" && ps.Ref == s.Ref {
				jt := fmt.Sprintf("%s(%s)", schema.Type, schema.Format)
				expressParams = append(expressParams, Express{Name: joinName(parent), Type: jt, Required: required, Description: s.BuildDescription()})
				return
			}
		}

		schema.parent = &s
		expressParams = append(expressParams, t.ExpressSchema(schema, required, parent...)...)
		return
	}

	if len(parent) > 0 {
		jt := s.Type
		if s.Format != "" {
			jt = fmt.Sprintf("%s(%s)", s.Type, s.Format)
		}
		expressParams = append(expressParams, Express{Name: joinName(parent), Type: jt, Required: required, Description: s.BuildDescription()})
	}

	if s.Type == "array" {
		if s.Items != nil {
			if len(parent) > 0 {
				parent[len(parent)-1] += "[]"
			} else {
				parent = append(parent, "[]")
			}
			s.Items.parent = &s
			expressParams = append(expressParams, t.ExpressSchema(*s.Items, required, parent...)...)
		}

		return
	}

	if s.Type == "object" {
		for k, v := range s.Properties {
			v.parent = &s
			expressParams = append(expressParams, t.ExpressSchema(v, slices.Contains(s.Required, k), appendName(parent, k)...)...)
		}
		sort.Slice(expressParams, func(i, j int) bool {
			return expressParams[i].Name < expressParams[j].Name
		})
		return
	}

	return
}

func appendName(a []string, n string) []string {
	if n != "" {
		return append(a, n)
	}
	return a
}

func joinName(names []string) string {
	bs := strings.Builder{}
	for _, v := range names {
		if v != "" {
			bs.WriteString(v)
			bs.WriteString(".")
		}
	}
	return strings.TrimRight(bs.String(), ".")
}
