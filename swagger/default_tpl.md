# {{.Info.Title}} ({{.Info.Version}})

**Schemes:** {{range .Schemes}}{{.}} {{end}}

---
{{- range $path, $pathVal := .Paths}}
  {{- range $method, $operation := $pathVal}}

## {{Upper $method}} {{$path}}
    {{- with $operation}}
{{.Summary}}

{{- if .Parameters}}
### Parameters
{{- $hp := $operation.HeaderParams}}
{{- $dp := $operation.PathParams}}
{{- $qp := $operation.QueryParams}}
{{- $bp := $operation.BodyParams}}

      {{- if $hp}}
#### Header parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
        {{- range $param := $hp}}
        {{- $exps := ExpressParam $param }}
          {{- range $exp := $exps}}
| {{$exp.Name}} | {{$exp.Type}} | {{$exp.Required}} | {{$exp.Description}} |
          {{- end}}
        {{- end}}
       {{- end}}

      {{- if $dp}}
#### Path parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
        {{- range $param := $dp}}
        {{- $exps := ExpressParam $param }}
          {{- range $exp := $exps}}
| {{$exp.Name}} | {{$exp.Type}} | {{$exp.Required}} | {{$exp.Description}} |
          {{- end}}
        {{- end}}
       {{- end}}

      {{- if $qp}}
#### Query parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
        {{- range $param := $qp}}
        {{- $exps := ExpressParam $param }}
          {{- range $exp := $exps}}
| {{$exp.Name}} | {{$exp.Type}} | {{$exp.Required}} | {{$exp.Description}} |
          {{- end}}
        {{- end}}
       {{- end}}

      {{- if $bp}}
#### Body parameters
| Name | type | Required | Description |
| ---- |  ---- | -------- | ----------- |
        {{- range $param := $bp}}
        {{- $exps := ExpressParam $param }}
          {{- range $exp := $exps}}
| {{$exp.Name}} | {{$exp.Type}} | {{$exp.Required}} | {{$exp.Description}} |
          {{- end}}
        {{- end}}
       {{- end}}
{{- end}}

### Responses
      {{- range $respName, $resp := .Responses}}

#### {{$respName}}
{{$resp.Description}}
        {{- $exps := ExpressResponse $resp }}
        {{- if $exps}}
| Name | type | Required | Description |
| ---- | ---- | -------- | ----------- |
          {{- range $exp := $exps}}
| {{$exp.Name}} | {{$exp.Type}} | {{$exp.Required}} | {{$exp.Description}} |
          {{- end}}
        {{- end}}
      {{- end}}
    {{- end}}
  {{- end}}
{{- end}}
