package testgenerator

import "encoding/json"

type StringNode struct {
	BaseNode `yaml:",inline"`
}

func (s *StringNode) ValidCases() []Case {
	cases := []Case{
		{
			GenerateValid: func(valid, invalid map[string]SchemaNode) json.RawMessage {
				return json.RawMessage(`"valid-string"`)
			},
		},
	}

	if s.Nullable {
		cases = append(cases, nullCase())
	}

	return cases
}

func (s *StringNode) InvalidCases() []Case {
	if s.Nullable {
		return nil
	}

	return []Case{nullCase()}
}

func nullCase() Case {
	return Case{
		GenerateValid: func(valid, invalid map[string]SchemaNode) json.RawMessage {
			return json.RawMessage(`null`)
		},
	}
}
