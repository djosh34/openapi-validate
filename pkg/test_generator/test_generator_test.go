package testgenerator

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestNullableObjectNodeHasOneNullableStringProperty(t *testing.T) {
	openapiPath := testdataOpenAPIPath("nullable_object_with_nullable_string.yaml")

	content, err := os.ReadFile(openapiPath)
	require.NoError(t, err)

	var doc OpenAPINode
	err = yaml.Unmarshal(content, &doc)
	require.NoError(t, err)

	require.Len(t, doc.Paths, 1)

	path := doc.Paths["/nullable-object-keys-additional-properties-false"]
	require.NotNil(t, path.Post)
	require.Equal(t, "nullableObjectKeysAdditionalPropertiesFalse", path.Post.OperationID)
	require.True(t, path.Post.RequestBody.Required)

	mediaTypeNode := path.Post.RequestBody.Content["application/json"]

	objectNode := mediaTypeNode.Schema.Object
	require.NotNil(t, objectNode)
	require.Nil(t, mediaTypeNode.Schema.String)

	require.Equal(t, BaseNode{Type: "object", Nullable: true}, objectNode.BaseNode)
	require.Equal(t, []string{"requiredNullableString"}, objectNode.Required)
	require.NotNil(t, objectNode.AdditionalProperties.Allowed)
	require.False(t, *objectNode.AdditionalProperties.Allowed)
	require.Nil(t, objectNode.AdditionalProperties.Schema)
	require.Len(t, objectNode.Properties, 1)

	propertyNode := objectNode.Properties["requiredNullableString"].String
	require.NotNil(t, propertyNode)
	require.Nil(t, objectNode.Properties["requiredNullableString"].Object)

	require.Equal(t, BaseNode{Type: "string", Nullable: true}, propertyNode.BaseNode)
}

func TestAdditionalPropertiesNodeCanBeSchema(t *testing.T) {
	content := []byte(`
type: object
additionalProperties:
  type: string
  nullable: true
`)

	var objectNode ObjectNode
	err := yaml.Unmarshal(content, &objectNode)
	require.NoError(t, err)

	require.Nil(t, objectNode.AdditionalProperties.Allowed)
	require.NotNil(t, objectNode.AdditionalProperties.Schema)

	additionalPropertiesNode := objectNode.AdditionalProperties.Schema.String
	require.NotNil(t, additionalPropertiesNode)
	require.Equal(t, BaseNode{Type: "string", Nullable: true}, additionalPropertiesNode.BaseNode)
}

func TestGenerateCasesFromOpenAPIFileStartsWithNullableStringProperty(t *testing.T) {
	t.Skip("starting point only; traversal and case generation design comes next")

	openapiPath := testdataOpenAPIPath("nullable_object_with_nullable_string.yaml")

	cases, err := GenerateCasesFromOpenAPIFile(openapiPath)
	require.NoError(t, err)

	require.Equal(t, []Case{
		{
			Name:      "required nullable string value",
			JSON:      `{"requiredNullableString":"required-nullable"}`,
			WantValid: true,
		},
		{
			Name:      "required nullable string null",
			JSON:      `{"requiredNullableString":null}`,
			WantValid: true,
		},
	}, cases)
}

func testdataOpenAPIPath(name string) string {
	return filepath.Join("testdata", name)
}
