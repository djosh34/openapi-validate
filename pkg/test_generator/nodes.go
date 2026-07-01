package testgenerator

import "fmt"

var _ Caseable = new(SchemaNode)

type SchemaNode struct {
	Type   string `yaml:"type"`
	Array  *ArrayNode
	Bool   *BoolNode
	Number *NumberNode
	Object *ObjectNode
	String *StringNode
}

func (s *SchemaNode) ValidCases() []Case {
	switch {
	case s.Array != nil:
		return s.Array.ValidCases()
	case s.Bool != nil:
		return s.Bool.ValidCases()
	case s.Number != nil:
		return s.Number.ValidCases()
	case s.Object != nil:
		return s.Object.ValidCases()
	case s.String != nil:
		return s.String.ValidCases()
	default:
		return nil
	}
}

func (s *SchemaNode) InvalidCases() []Case {
	switch {
	case s.Array != nil:
		return s.Array.InvalidCases()
	case s.Bool != nil:
		return s.Bool.InvalidCases()
	case s.Number != nil:
		return s.Number.InvalidCases()
	case s.Object != nil:
		return s.Object.InvalidCases()
	case s.String != nil:
		return s.String.InvalidCases()
	default:
		return nil
	}
}

func (s SchemaNode) Merge(other SchemaNode) (SchemaNode, error) {
	if s.Type != other.Type {
		return SchemaNode{}, fmt.Errorf("cannot merge schema type %q with %q", s.Type, other.Type)
	}

	switch s.Type {
	case "array":
		if s.Array == nil || other.Array == nil {
			return SchemaNode{}, fmt.Errorf("array schema is missing array node")
		}

		return s.Array.Merge(other)
	case "boolean":
		if s.Bool == nil || other.Bool == nil {
			return SchemaNode{}, fmt.Errorf("boolean schema is missing boolean node")
		}

		return s.Bool.Merge(other)
	case "number":
		if s.Number == nil || other.Number == nil {
			return SchemaNode{}, fmt.Errorf("number schema is missing number node")
		}

		return s.Number.Merge(other)
	case "object":
		if s.Object == nil || other.Object == nil {
			return SchemaNode{}, fmt.Errorf("object schema is missing object node")
		}

		return s.Object.Merge(other)
	case "string":
		if s.String == nil || other.String == nil {
			return SchemaNode{}, fmt.Errorf("string schema is missing string node")
		}

		return s.String.Merge(other)
	default:
		return SchemaNode{}, fmt.Errorf("unsupported schema type %q", s.Type)
	}
}

type BaseNode struct {
	Nullable bool `yaml:"nullable"`
}

func mergeBaseNode(left BaseNode, right BaseNode) BaseNode {
	return BaseNode{Nullable: left.Nullable && right.Nullable}
}
