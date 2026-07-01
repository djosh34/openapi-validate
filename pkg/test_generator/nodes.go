package testgenerator

var _ Caseable = new(SchemaNode)

type SchemaNode struct {
	Type   string `yaml:"type"`
	Object *ObjectNode
	String *StringNode
}

func (s *SchemaNode) ValidCases() []Case {
	switch {
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
	case s.Object != nil:
		return s.Object.InvalidCases()
	case s.String != nil:
		return s.String.InvalidCases()
	default:
		return nil
	}
}

type BaseNode struct {
	Nullable bool `yaml:"nullable"`
}
