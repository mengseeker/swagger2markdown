# swagger2markdown

## Install
```shell
go install github.com/mengseeker/swagger2markdown@latest
```

## Usage
```text
transform swagger into markdown.
only support swagger 2.0

Usage:
  swagger2markdown [flags]

Flags:
  -h, --help                 help for swagger2markdown
  -i, --input string         input file, can be json or yaml format, default read from stdin
  -f, --inputFormat string   input file format, json or yaml, default auto detect
  -o, --output string        output file, default print to stdout
  -m, --template string      custom template file
  -t, --toggle               Help message for toggle
```

## Example

```bash
# read from url
swagger2markdown -i https://petstore.swagger.io/v2/swagger.json

# read from file
swagger2markdown -i test.swagger.json

# read from stdin
curl https://petstore.swagger.io/v2/swagger.json | swagger2markdown
```

## output
[e1.md](example/e1.md)

[petstore.swagger.md](example/petstore.swagger.md)