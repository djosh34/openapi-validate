package testgenerator

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func (s *SchemaNode) UnmarshalYAML(value *yaml.Node) error {
	if value == nil || value.Kind == 0 {
		return fmt.Errorf("missing schema")
	}

	for _, keyword := range []string{"anyOf", "oneOf"} {
		if schemaMappingValue(value, keyword) != nil {
			return fmt.Errorf("unsupported schema composition %q", keyword)
		}
	}

	allOfNode := schemaMappingValue(value, "allOf")
	if allOfNode != nil {
		var allOf struct {
			Schemas []SchemaNode `yaml:"allOf"`
		}
		err := value.Decode(&allOf)
		if err != nil {
			return err
		}
		if len(allOf.Schemas) == 0 {
			return fmt.Errorf("allOf must contain at least one schema")
		}

		merged := allOf.Schemas[0]
		for i, schema := range allOf.Schemas[1:] {
			mergedSchema, err := merged.Merge(schema)
			if err != nil {
				return fmt.Errorf("merge allOf schema %d: %w", i+2, err)
			}
			merged = mergedSchema
		}

		*s = merged
		return nil
	}

	var schema struct {
		Type string `yaml:"type"`
	}
	err := value.Decode(&schema)
	if err != nil {
		return err
	}

	switch schema.Type {
	case "array":
		var arrayNode ArrayNode
		err = value.Decode(&arrayNode)
		if err != nil {
			return err
		}
		s.Type = schema.Type
		s.Array = &arrayNode
		s.Bool = nil
		s.Number = nil
		s.Object = nil
		s.String = nil
		return nil
	case "boolean":
		var boolNode BoolNode
		err = value.Decode(&boolNode)
		if err != nil {
			return err
		}
		s.Type = schema.Type
		s.Array = nil
		s.Bool = &boolNode
		s.Number = nil
		s.Object = nil
		s.String = nil
		return nil
	case "number":
		var numberNode NumberNode
		err = value.Decode(&numberNode)
		if err != nil {
			return err
		}
		s.Type = schema.Type
		s.Array = nil
		s.Bool = nil
		s.Number = &numberNode
		s.Object = nil
		s.String = nil
		return nil
	case "object":
		var objectNode ObjectNode
		err = value.Decode(&objectNode)
		if err != nil {
			return err
		}
		s.Type = schema.Type
		s.Array = nil
		s.Bool = nil
		s.Number = nil
		s.Object = &objectNode
		s.String = nil
		return nil
	case "string":
		var stringNode StringNode
		err = value.Decode(&stringNode)
		if err != nil {
			return err
		}
		s.Type = schema.Type
		s.Array = nil
		s.Bool = nil
		s.Number = nil
		s.Object = nil
		s.String = &stringNode
		return nil
	default:
		return fmt.Errorf("unsupported schema type %q", schema.Type)
	}
}

func schemaMappingValue(node *yaml.Node, key string) *yaml.Node {
	if node == nil || node.Kind != yaml.MappingNode {
		return nil
	}

	for i := 0; i < len(node.Content)-1; i += 2 {
		if node.Content[i].Value == key {
			return node.Content[i+1]
		}
	}

	return nil
}

func (a *AdditionalPropertiesNode) UnmarshalYAML(value *yaml.Node) error {
	if value == nil || value.Kind == 0 {
		return nil
	}

	switch value.Kind {
	case yaml.ScalarNode:
		if value.Tag != "!!bool" {
			return fmt.Errorf("unsupported scalar %s", value.Tag)
		}

		var allowed bool
		err := value.Decode(&allowed)
		if err != nil {
			return err
		}

		a.Allowed = &allowed
		a.Schema = nil
		return nil
	case yaml.MappingNode:
		var schema SchemaNode
		err := value.Decode(&schema)
		if err != nil {
			return err
		}

		a.Allowed = nil
		a.Schema = &schema
		return nil
	default:
		return fmt.Errorf("unsupported yaml node kind %d", value.Kind)
	}
}
