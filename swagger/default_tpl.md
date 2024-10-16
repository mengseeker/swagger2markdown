# {{.Info.Title}} ({{.Info.Version}})

**Schemes:** {{range .Schemes}}{{.}} {{end}}

---
{{- range $path, $pathVal := .Paths}}
  {{- range $method, $operation := $pathVal}}

## {{Upper $method}} {{$path}}
    {{- with $operation}}

### Summary
{{.Summary}}

### Parameters
| Name | Located in | type | Required | Description |
| ---- | ---------- |  ---- | -------- | ----------- |
      {{- range $param := .Parameters}}
      {{- $exps := ExpressParam $param }}
        {{- range $exp := $exps}}
| {{$exp.Name}} | {{$param.In}} | {{$exp.Type}} | {{$exp.Required}} | {{$exp.Description}} |
        {{- end}}
      {{- end}}

### Responses
      {{- range $respName, $resp := .Responses}}

#### {{$respName}}
{{$resp.Description}}

| Name | type | Required | Description |
| ---- | ---- | -------- | ----------- |
        {{- $exps := ExpressResponse $resp }}
        {{- range $exp := $exps}}
| {{$exp.Name}} | {{$exp.Type}} | {{$exp.Required}} | {{$exp.Description}} |
        {{- end}}

      {{- end}}

    {{- end}}


  {{- end}}


{{- end}}
