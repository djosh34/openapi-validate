package testgenerator

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func TestParseObjectWithAllObjectSchemaFields(t *testing.T) {
	const objectSchemaYAML = `
type: object
required:
  - name
properties:
  name:
    type: string
additionalProperties:
  type: string
minProperties: 1
maxProperties: 3
`

	var node yaml.Node
	err := yaml.Unmarshal([]byte(objectSchemaYAML), &node)
	require.NoError(t, err)
	require.Len(t, node.Content, 1)

	dc := DomainContext{}
	_, err = dc.ParseObject(*node.Content[0])
	require.NoError(t, err)
}
