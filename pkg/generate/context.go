package generate

import "github.com/getkin/kin-openapi/openapi3"

type GenerateContext struct {
	Document *openapi3.T
}

type SchemaObject interface{}

type ObjectContext struct {
	Nullable             bool
	AdditionalProperties bool
	Required             []string
	Properties           map[string]SchemaObject
}

type StringContext struct {
	Nullable bool
}
