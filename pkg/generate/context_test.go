package generate

import (
	"net/http"
	"testing"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/stretchr/testify/require"
)

// TestJSONRequestBodySchemasKeepsOnlyOperationsWithJSONBodySchema exercises the named generator behavior.
func TestJSONRequestBodySchemasKeepsOnlyOperationsWithJSONBodySchema(t *testing.T) {
	t.Parallel()

	jsonSchema := openapi3.NewStringSchema()
	jsonOperation := operationWithContent("jsonBody", openapi3.NewContentWithJSONSchema(jsonSchema))
	noRequestBodyOperation := &openapi3.Operation{OperationID: "noRequestBody"}
	noJSONOperation := operationWithContent(
		"noJSON",
		openapi3.NewContentWithSchema(openapi3.NewStringSchema(), []string{"text/plain"}),
	)
	noSchemaOperation := operationWithContent("noSchema", openapi3.Content{
		"application/json": openapi3.NewMediaType(),
	})

	generateContext := &GenerateContext{
		Document: &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/json", &openapi3.PathItem{Post: jsonOperation}),
				openapi3.WithPath("/no-request-body", &openapi3.PathItem{Post: noRequestBodyOperation}),
				openapi3.WithPath("/no-json", &openapi3.PathItem{Post: noJSONOperation}),
				openapi3.WithPath("/no-schema", &openapi3.PathItem{Post: noSchemaOperation}),
			),
		},
	}

	schemas, err := generateContext.JSONRequestBodySchemas()
	require.NoError(t, err)

	require.Equal(t, map[*openapi3.Operation]*openapi3.Schema{
		jsonOperation: jsonSchema,
	}, schemas)
}

// TestJSONRequestBodyModelSchemasConvertsRequestBodySchemas exercises the named generator behavior.
func TestJSONRequestBodyModelSchemasConvertsRequestBodySchemas(t *testing.T) {
	t.Parallel()

	openapiExamplePath := exampleOpenAPIPath(t)
	generateContext, err := LoadOpenapi(t.Context(), openapiExamplePath)
	require.NoError(t, err)

	err = generateContext.FilterOperations("objectKeysAdditionalPropertiesFalse")
	require.NoError(t, err)

	schemas, err := generateContext.JSONRequestBodyModelSchemas()
	require.NoError(t, err)

	optionalNotNullableString := &StringSchema{
		BaseSchema: BaseSchema{Name: "ObjectKeysAdditionalPropertiesFalseOptionalNotNullableString"},
	}
	optionalNullableString := &StringSchema{
		BaseSchema: BaseSchema{Name: "ObjectKeysAdditionalPropertiesFalseOptionalNullableString", Nullable: true},
	}
	requiredNotNullableString := &StringSchema{
		BaseSchema: BaseSchema{Name: "ObjectKeysAdditionalPropertiesFalseRequiredNotNullableString"},
	}
	requiredNullableString := &StringSchema{
		BaseSchema: BaseSchema{Name: "ObjectKeysAdditionalPropertiesFalseRequiredNullableString", Nullable: true},
	}

	require.Equal(t, []Schema{
		&ObjectSchema{
			BaseSchema:           BaseSchema{Name: "ObjectKeysAdditionalPropertiesFalse"},
			AdditionalProperties: false,
			Properties: []ObjectFieldContext{
				{
					PropertyName: "optionalNotNullableString",
					Schema:       optionalNotNullableString,
				},
				{
					PropertyName: "optionalNullableString",
					Schema:       optionalNullableString,
				},
				{
					PropertyName: "requiredNotNullableString",
					Schema:       requiredNotNullableString,
					Required:     true,
				},
				{
					PropertyName: "requiredNullableString",
					Schema:       requiredNullableString,
					Required:     true,
				},
			},
		},
		optionalNotNullableString,
		optionalNullableString,
		requiredNotNullableString,
		requiredNullableString,
	}, schemas)
}

// TestJSONRequestBodyModelSchemasAllowsEmptyBodyWhenRequestBodyIsOptional exercises the named generator behavior.
func TestJSONRequestBodyModelSchemasAllowsEmptyBodyWhenRequestBodyIsOptional(t *testing.T) {
	t.Parallel()

	schema := openapi3.NewArraySchema()
	schema.WithItems(openapi3.NewStringSchema())

	generateContext := &GenerateContext{
		Document: &openapi3.T{
			Paths: openapi3.NewPaths(
				openapi3.WithPath("/optional-body", &openapi3.PathItem{
					Post: operationWithContent("optionalBody", openapi3.NewContentWithJSONSchema(schema)),
				}),
			),
		},
	}

	schemas, err := generateContext.JSONRequestBodyModelSchemas()
	require.NoError(t, err)

	require.Equal(t, []Schema{
		&ArraySchema{
			BaseSchema: BaseSchema{Name: "OptionalBody", EmptyBodyAllowed: true},
			Items: &StringSchema{
				BaseSchema: BaseSchema{Name: "OptionalBodyItem"},
			},
		},
		&StringSchema{
			BaseSchema: BaseSchema{Name: "OptionalBodyItem"},
		},
	}, schemas)
}

// TestSchemaFromOpenAPISchemaRecursesObjectProperties exercises the named generator behavior.
func TestSchemaFromOpenAPISchemaRecursesObjectProperties(t *testing.T) {
	t.Parallel()

	nestedSchema := openapi3.NewObjectSchema()
	nestedSchema.WithProperty("child", openapi3.NewStringSchema().WithNullable())

	schema := openapi3.NewObjectSchema()
	schema.Required = []string{"name", "nested"}
	schema.WithoutAdditionalProperties()
	schema.WithProperty("name", openapi3.NewStringSchema())
	schema.WithProperty("nested", nestedSchema)

	generatedSchema, err := SchemaFromOpenAPISchema(schema)
	require.NoError(t, err)

	require.Equal(t, &ObjectSchema{
		AdditionalProperties: false,
		Properties: []ObjectFieldContext{
			{
				PropertyName: "name",
				Schema: &StringSchema{
					BaseSchema: BaseSchema{Name: "Name"},
				},
				Required: true,
			},
			{
				PropertyName: "nested",
				Required:     true,
				Schema: &ObjectSchema{
					BaseSchema:           BaseSchema{Name: "Nested"},
					AdditionalProperties: true,
					Properties: []ObjectFieldContext{
						{
							PropertyName: "child",
							Schema: &StringSchema{
								BaseSchema: BaseSchema{Name: "Child", Nullable: true},
							},
						},
					},
				},
			},
		},
	}, generatedSchema)
}

// TestSchemaFromOpenAPISchemaConvertsArrayItems exercises the named generator behavior.
func TestSchemaFromOpenAPISchemaConvertsArrayItems(t *testing.T) {
	t.Parallel()

	schema := openapi3.NewArraySchema().WithNullable()
	schema.WithItems(openapi3.NewStringSchema())

	generatedSchema, err := SchemaFromOpenAPISchema(schema)
	require.NoError(t, err)

	require.Equal(t, &ArraySchema{
		BaseSchema: BaseSchema{Nullable: true},
		Items: &StringSchema{
			BaseSchema: BaseSchema{Name: "Item"},
		},
	}, generatedSchema)
}

// TestSchemaFromOpenAPISchemaConvertsAdditionalPropertiesSchema exercises the named generator behavior.
func TestSchemaFromOpenAPISchemaConvertsAdditionalPropertiesSchema(t *testing.T) {
	t.Parallel()

	schema := openapi3.NewObjectSchema()
	schema.WithAdditionalProperties(openapi3.NewStringSchema().WithNullable())

	generatedSchema, err := SchemaFromOpenAPISchema(schema)
	require.NoError(t, err)

	require.Equal(t, &ObjectSchema{
		AdditionalProperties: true,
		AdditionalPropertiesSchema: &StringSchema{
			BaseSchema: BaseSchema{Name: "AdditionalProperty", Nullable: true},
		},
		Properties: []ObjectFieldContext{},
	}, generatedSchema)
}

// operationWithContent supports generator tests.
func operationWithContent(operationID string, content openapi3.Content) *openapi3.Operation {
	return &openapi3.Operation{
		OperationID: operationID,
		RequestBody: &openapi3.RequestBodyRef{
			Value: &openapi3.RequestBody{
				Content: content,
			},
		},
	}
}

// TestJSONRequestBodySchemasSupportsHTTPMethods exercises the named generator behavior.
func TestJSONRequestBodySchemasSupportsHTTPMethods(t *testing.T) {
	t.Parallel()

	schema := openapi3.NewStringSchema()
	operation := operationWithContent("putBody", openapi3.NewContentWithJSONSchema(schema))
	pathItem := new(openapi3.PathItem)
	pathItem.SetOperation(http.MethodPut, operation)

	generateContext := &GenerateContext{
		Document: &openapi3.T{
			Paths: openapi3.NewPaths(openapi3.WithPath("/put", pathItem)),
		},
	}

	schemas, err := generateContext.JSONRequestBodySchemas()
	require.NoError(t, err)

	require.Equal(t, map[*openapi3.Operation]*openapi3.Schema{
		operation: schema,
	}, schemas)
}

// TestJSONRequestBodyModelSchemasOrdersMethods keeps generated output deterministic.
func TestJSONRequestBodyModelSchemasOrdersMethods(t *testing.T) {
	t.Parallel()

	pathItem := new(openapi3.PathItem)

	for method, operationID := range map[string]string{
		http.MethodPut:  "putBody",
		http.MethodGet:  "getBody",
		http.MethodPost: "postBody",
	} {
		operation := operationWithContent(
			operationID,
			openapi3.NewContentWithJSONSchema(openapi3.NewStringSchema()),
		)
		pathItem.SetOperation(method, operation)
	}

	generateContext := &GenerateContext{
		Document: &openapi3.T{
			Paths: openapi3.NewPaths(openapi3.WithPath("/multi-method", pathItem)),
		},
	}

	_, err := generateContext.JSONRequestBodyModelSchemas()
	require.NoError(t, err)
	require.Equal(t, []JSONRequestBodyOperation{
		{OperationID: "getBody", TypeName: "GetBody"},
		{OperationID: "postBody", TypeName: "PostBody"},
		{OperationID: "putBody", TypeName: "PutBody"},
	}, generateContext.JSONRequestBodyOperations)
}
