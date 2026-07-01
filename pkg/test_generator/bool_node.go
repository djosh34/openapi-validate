package testgenerator

import (
	"encoding/json"
	"fmt"
)

var _ Caseable = new(BoolNode)

type BoolNode struct {
	BaseNode `yaml:",inline"`
}

func (b *BoolNode) ValidCases() []Case {
	cases := []Case{
		{Name: "true", Value: json.RawMessage(`true`)},
		{Name: "false", Value: json.RawMessage(`false`)},
	}

	return append(cases, b.BaseNode.ValidCases()...)
}

func (b *BoolNode) InvalidCases() []Case {
	cases := append([]Case{}, b.BaseNode.InvalidCases()...)
	return append(cases,
		Case{Name: "string", Value: json.RawMessage(`"not-boolean"`)},
		Case{Name: "string true", Value: json.RawMessage(`"true"`)},
		Case{Name: "string false", Value: json.RawMessage(`"false"`)},
		Case{Name: "number", Value: json.RawMessage(`123`)},
		Case{Name: "zero", Value: json.RawMessage(`0`)},
		Case{Name: "object", Value: json.RawMessage(`{}`)},
		Case{Name: "array", Value: json.RawMessage(`[]`)},
	)
}

func (b *BoolNode) Merge(schema SchemaNode) (SchemaNode, error) {
	if schema.Type != "boolean" {
		return SchemaNode{}, fmt.Errorf("cannot merge schema type %q with %q", "boolean", schema.Type)
	}
	if schema.Bool == nil {
		return SchemaNode{}, fmt.Errorf("boolean schema is missing boolean node")
	}

	return SchemaNode{
		Type: "boolean",
		Bool: &BoolNode{
			BaseNode: mergeBaseNode(b.BaseNode, schema.Bool.BaseNode),
		},
	}, nil
}
