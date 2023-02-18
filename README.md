# oapi-generator 

oapi-generator is a command-line tool that generates Go code from an OpenAPI specification. via configuration file 

## Installation

Use the go package install the oapi-codegen and oapi-generator 

```bash
go install https://github.com/deepmap/oapi-codegen
go install github.com/mrbryside/oapi-generator
```


## Create your oapi-cfg.yaml (configuration file)
After create oapi-cfg.yaml in project directory add the content below

```bash
#oapi-cfg.yaml
gen-dir: your/generated/output/path #generated output path
spec-dir: your/spec/openapi/path #spec open-api path
```

## Create spec
create spec folder 
```bash
mkdir your/spec/openapi/path/spec
```
create server.cfg.yaml file content below you can decide to use additional-imports
*** do not change #name in this file (it's use for replace the package name)
```bash
package: #name
#additional-imports:
#  - package: github.com/[repo]/[gen-dir]/oapi/#name
#    alias: .
generate:
  echo-server: true
  embedded-spec: true
```
create [your-openapi-name].yaml file content below
```bash
openapi: 3.0.0

info:
  title: pocket app OAS
  description: OpenApi specification for a pocket api
  version: 1.0.0

servers:
  - url: http://localhost:8080/

......
....
```

## Usage
```bash
oapi-generator generate-server --name=[your-openapi-name]
```