package testgenerator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

const openAPIObjectSchemaSpec = `
openapi: 3.0.3
info:
  title: test
  version: 1.0.0
paths:
  /things:
    post:
      operationId: createThing
      requestBody:
        content:
          application/json:
            schema:
              type: object
              required:
                - name
              properties:
                name:
                  type: string
`

func TestOpenAPIRequestBodySchemaNodeReturnsObjectSchemaNode(t *testing.T) {
	schemaNode, err := OpenAPIRequestBodySchemaNode([]byte(openAPIObjectSchemaSpec), "createThing")
	require.NoError(t, err)

	require.Equal(t, yaml.MappingNode, schemaNode.Kind)
	require.Len(t, schemaNode.Content, 6)
	require.Equal(t, "type", schemaNode.Content[0].Value)
	require.Equal(t, "object", schemaNode.Content[1].Value)
	require.Equal(t, "required", schemaNode.Content[2].Value)
	require.Equal(t, yaml.SequenceNode, schemaNode.Content[3].Kind)
	require.Equal(t, "properties", schemaNode.Content[4].Value)
	require.Equal(t, yaml.MappingNode, schemaNode.Content[5].Kind)
}

func TestOpenAPIRequestBodySchemaNodeReturnsRefSchemaNode(t *testing.T) {
	openAPISpec := []byte(`
openapi: 3.0.3
info:
  title: test
  version: 1.0.0
paths:
  /things/{id}:
    put:
      operationId: updateThing
      requestBody:
        content:
          application/json:
            schema:
              $ref: '#/components/schemas/Thing'
components:
  schemas:
    Thing:
      type: object
`)

	schemaNode, err := OpenAPIRequestBodySchemaNode(openAPISpec, "updateThing")
	require.NoError(t, err)

	require.Equal(t, yaml.MappingNode, schemaNode.Kind)
	require.Len(t, schemaNode.Content, 2)
	require.Equal(t, "$ref", schemaNode.Content[0].Value)
	require.Equal(t, "#/components/schemas/Thing", schemaNode.Content[1].Value)
}

func TestOpenAPIRequestBodySchemaNodeErrors(t *testing.T) {
	tests := map[string]struct {
		openAPISpec string
		operationID string
		wantError   string
	}{
		"invalid yaml": {
			openAPISpec: "paths: [",
			operationID: "createThing",
			wantError:   "parse openapi yaml spec",
		},
		"operation not found": {
			openAPISpec: openAPIObjectSchemaSpec,
			operationID: "missingThing",
			wantError:   `operationId "missingThing" not found`,
		},
		"duplicate operation id": {
			openAPISpec: `
openapi: 3.0.3
paths:
  /first:
    post:
      operationId: duplicateThing
      requestBody:
        content:
          application/json:
            schema:
              type: string
  /second:
    patch:
      operationId: duplicateThing
      requestBody:
        content:
          application/json:
            schema:
              type: object
`,
			operationID: "duplicateThing",
			wantError:   `operationId "duplicateThing" found multiple times`,
		},
		"content type does not exist": {
			openAPISpec: `
openapi: 3.0.3
paths:
  /things:
    post:
      operationId: createThing
      requestBody: {}
`,
			operationID: "createThing",
			wantError:   `operationId "createThing" request body content type does not exist`,
		},
		"content type is not json": {
			openAPISpec: `
openapi: 3.0.3
paths:
  /things:
    post:
      operationId: createThing
      requestBody:
        content:
          text/plain:
            schema:
              type: string
`,
			operationID: "createThing",
			wantError:   `operationId "createThing" request body content type is not json`,
		},
		"schema does not exist": {
			openAPISpec: `
openapi: 3.0.3
paths:
  /things:
    post:
      operationId: createThing
      requestBody:
        content:
          application/json: {}
`,
			operationID: "createThing",
			wantError:   `operationId "createThing" application/json schema does not exist`,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := OpenAPIRequestBodySchemaNode([]byte(tt.openAPISpec), tt.operationID)
			require.Error(t, err)
			require.ErrorContains(t, err, tt.wantError)
		})
	}
}

func TestGenerateFunctionsParseSchemaAndDoNothing(t *testing.T) {
	tests := map[string]struct {
		generate func([]byte, string, func([]byte) error) error
	}{
		"valid":   {generate: GenerateValid},
		"invalid": {generate: GenerateInvalid},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			called := false
			err := tt.generate([]byte(openAPIObjectSchemaSpec), "createThing", func(_ []byte) error {
				called = true
				return errors.New("unmarshal must not be called yet")
			})
			require.NoError(t, err)
			require.False(t, called)
		})
	}
}

func TestGenerateFunctionsWrapOpenAPIParseErrors(t *testing.T) {
	tests := map[string]struct {
		generate func([]byte, string, func([]byte) error) error
	}{
		"valid":   {generate: GenerateValid},
		"invalid": {generate: GenerateInvalid},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			err := tt.generate([]byte(openAPIObjectSchemaSpec), "missingThing", nil)
			require.Error(t, err)
			require.ErrorContains(t, err, "openapi yaml spec parse failed")
			require.ErrorContains(t, err, `operationId "missingThing" not found`)
		})
	}
}
