package swagger

import (
	"fmt"
	"strings"
)

// swagger: "2.0"
// info:
//   title: Sample API
//   description: API description in Markdown.
//   version: 1.0.0

// host: api.example.com
// basePath: /v1
// schemes:
//   - https

// paths:
//   /users:
//     get:
//       summary: Returns a list of users.
//       description: Optional extended description in Markdown.
//       produces:
//         - application/json
//       responses:
//         200:
//           description: OK

type Ref = string

type Swagger struct {
	Swagger     string            `json:"swagger" yaml:"swagger"`
	Info        Info              `json:"info" yaml:"info"`
	Host        string            `json:"host" yaml:"host"`
	BasePath    string            `json:"basePath" yaml:"basePath"`
	Schemes     []string          `json:"schemes" yaml:"schemes"`
	Paths       map[string]Path   `json:"paths" yaml:"paths"`
	Definitions map[string]Schema `json:"definitions" yaml:"definitions"`
}

type Info struct {
	Title       string `json:"title" yaml:"title"`
	Description string `json:"description" yaml:"description"`
	Version     string `json:"version" yaml:"version"`
}

type Path map[string]Operation

type Operation struct {
	Summary      string              `json:"summary" yaml:"summary"`
	Description  string              `json:"description" yaml:"description"`
	OperationId  string              `json:"operationId" yaml:"operationId"`
	Tags         []string            `json:"tags" yaml:"tags"`
	Produces     []string            `json:"produces" yaml:"produces"`
	Parameters   []Param             `json:"parameters" yaml:"parameters"`
	Responses    map[string]Response `json:"responses" yaml:"responses"`
	ExternalDocs map[string]string   `json:"externalDocs" yaml:"externalDocs"`
	Deprecated   bool                `json:"deprecated" yaml:"deprecated"`
}

type Param struct {
	Ref    *Ref    `json:"$ref" yaml:"$ref"`
	Schema *Schema `json:"schema" yaml:"schema"`

	Name        string      `json:"name" yaml:"name"`
	In          string      `json:"in" yaml:"in"`
	Description string      `json:"description" yaml:"description"`
	Type        string      `json:"type" yaml:"type"`
	Required    bool        `json:"required" yaml:"required"`
	Default     string      `json:"default" yaml:"default"`
	Minimum     any         `json:"minimum" yaml:"minimum"`
	Maximum     any         `json:"maximum" yaml:"maximum"`
	Enum        []any       `json:"enum" yaml:"enum"`
	Items       *ParamItems `json:"items" yaml:"items"`
}

func (p Param) BuildDescription() string {
	desc := []string{strings.ReplaceAll(p.Description, "\n", " ")}
	if p.Default != "" {
		desc = append(desc, "default:"+p.Default)
	}
	if p.Enum != nil {
		desc = append(desc, "enum:"+fmt.Sprint(p.Enum))
	}
	if p.Minimum != nil {
		desc = append(desc, "min:"+fmt.Sprint(p.Minimum))
	}
	if p.Maximum != nil {
		desc = append(desc, "max:"+fmt.Sprint(p.Maximum))
	}
	return strings.Join(desc, "; ")
}

type ParamItems struct {
	Type string `json:"type" yaml:"type"`
}

type Response struct {
	Description string  `json:"description" yaml:"description"`
	Schema      *Schema `json:"schema" yaml:"schema"`
}

type Schema struct {
	Ref Ref `json:"$ref" yaml:"$ref"`

	Type       string            `json:"type" yaml:"type"`
	Format     string            `json:"format" yaml:"format"`
	Title      string            `json:"title" yaml:"title"`
	Required   []string          `json:"required" yaml:"required"`
	Items      *Schema           `json:"items" yaml:"items"`
	Properties map[string]Schema `json:"properties" yaml:"properties"`
}
