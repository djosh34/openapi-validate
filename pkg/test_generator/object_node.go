package testgenerator

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
)

const additionalPropertyCaseKey = "__decode_and_validate_generator_additional_property_61f80f12f8b14e2f__"

var _ Caseable = new(ObjectNode)

type ObjectNode struct {
	BaseNode             `yaml:",inline"`
	Required             []string                 `yaml:"required"`
	AdditionalProperties AdditionalPropertiesNode `yaml:"additionalProperties"`
	Properties           map[string]SchemaNode    `yaml:"properties"`
}

func (o *ObjectNode) ValidCases() []Case {
	_, optionalNames, baselineRequiredObject, requiredObjects := o.objectCaseData()

	cases := append([]Case{}, o.BaseNode.ValidCases()...)
	for _, requiredObject := range requiredObjects {
		cases = append(cases, objectCase("required properties", requiredObject))
	}

	for _, name := range optionalNames {
		schema := o.Properties[name]
		for _, validCase := range schema.ValidCases() {
			object := cloneObject(baselineRequiredObject)
			object[name] = validCase.Value

			cases = append(cases, objectCase(
				"optional property "+name+" "+validCase.Name,
				object,
			))
		}
	}

	return append(cases, o.additionalPropertyValidCases(baselineRequiredObject)...)
}

func (o *ObjectNode) InvalidCases() []Case {
	requiredNames, optionalNames, baselineRequiredObject, _ := o.objectCaseData()

	cases := append([]Case{}, o.BaseNode.InvalidCases()...)
	cases = append(cases,
		Case{Name: "string", Value: json.RawMessage(`"not-object"`)},
		Case{Name: "number", Value: json.RawMessage(`123`)},
		Case{Name: "boolean", Value: json.RawMessage(`true`)},
		Case{Name: "array", Value: json.RawMessage(`[]`)},
	)

	if len(requiredNames) > 0 {
		cases = append(cases, objectCase("missing required properties", nil))
	}

	if len(requiredNames) > 1 {
		for _, missingName := range requiredNames {
			object := cloneObject(baselineRequiredObject)
			delete(object, missingName)
			cases = append(cases, objectCase("missing required property "+missingName, object))
		}
	}

	for _, name := range append(requiredNames, optionalNames...) {
		schema := o.Properties[name]
		for _, invalidCase := range schema.InvalidCases() {
			object := cloneObject(baselineRequiredObject)
			object[name] = invalidCase.Value

			cases = append(cases, objectCase(
				"invalid property "+name+" "+invalidCase.Name,
				object,
			))
		}
	}

	return append(cases, o.additionalPropertyInvalidCases(baselineRequiredObject)...)
}

func (o *ObjectNode) Merge(schema SchemaNode) (SchemaNode, error) {
	if schema.Type != "object" {
		return SchemaNode{}, fmt.Errorf("cannot merge schema type %q with %q", "object", schema.Type)
	}
	if schema.Object == nil {
		return SchemaNode{}, fmt.Errorf("object schema is missing object node")
	}

	var properties map[string]SchemaNode
	if len(o.Properties)+len(schema.Object.Properties) > 0 {
		properties = make(map[string]SchemaNode, len(o.Properties)+len(schema.Object.Properties))
		for name, property := range o.Properties {
			properties[name] = property
		}
		for name, property := range schema.Object.Properties {
			leftProperty, ok := properties[name]
			if !ok {
				properties[name] = property
				continue
			}

			merged, err := leftProperty.Merge(property)
			if err != nil {
				return SchemaNode{}, fmt.Errorf("property %q: %w", name, err)
			}
			properties[name] = merged
		}
	}

	required := make([]string, 0, len(o.Required)+len(schema.Object.Required))
	seenRequired := map[string]struct{}{}
	for _, names := range [][]string{o.Required, schema.Object.Required} {
		for _, name := range names {
			if _, ok := seenRequired[name]; ok {
				continue
			}

			seenRequired[name] = struct{}{}
			required = append(required, name)
		}
	}

	var additionalProperties AdditionalPropertiesNode
	leftAdditionalPropertiesFalse := o.AdditionalProperties.Allowed != nil && !*o.AdditionalProperties.Allowed
	rightAdditionalPropertiesFalse := schema.Object.AdditionalProperties.Allowed != nil && !*schema.Object.AdditionalProperties.Allowed

	switch {
	case leftAdditionalPropertiesFalse || rightAdditionalPropertiesFalse:
		additionalProperties.Allowed = new(false)
	case o.AdditionalProperties.Schema != nil && schema.Object.AdditionalProperties.Schema != nil:
		merged, err := o.AdditionalProperties.Schema.Merge(*schema.Object.AdditionalProperties.Schema)
		if err != nil {
			return SchemaNode{}, fmt.Errorf("additionalProperties: %w", err)
		}
		additionalProperties.Schema = &merged
	case o.AdditionalProperties.Schema != nil:
		additionalProperties.Schema = o.AdditionalProperties.Schema
	case schema.Object.AdditionalProperties.Schema != nil:
		additionalProperties.Schema = schema.Object.AdditionalProperties.Schema
	case o.AdditionalProperties.Allowed != nil || schema.Object.AdditionalProperties.Allowed != nil:
		additionalProperties.Allowed = new(true)
	}

	return SchemaNode{
		Type: "object",
		Object: &ObjectNode{
			BaseNode:             mergeBaseNode(o.BaseNode, schema.Object.BaseNode),
			Required:             required,
			AdditionalProperties: additionalProperties,
			Properties:           properties,
		},
	}, nil
}

type AdditionalPropertiesNode struct {
	Allowed *bool
	Schema  *SchemaNode
}

func objectCase(name string, object map[string]json.RawMessage) Case {
	if object == nil {
		object = map[string]json.RawMessage{}
	}

	value, err := json.Marshal(object)
	if err != nil {
		panic(err)
	}

	return Case{
		Name:  name,
		Value: value,
	}
}

func (o *ObjectNode) additionalPropertyValidCases(requiredObject map[string]json.RawMessage) []Case {
	additionalPropertyKey := o.additionalPropertyKey()

	switch {
	case o.AdditionalProperties.Schema != nil:
		var cases []Case
		schema := *o.AdditionalProperties.Schema
		for _, validCase := range schema.ValidCases() {
			object := cloneObject(requiredObject)
			object[additionalPropertyKey] = validCase.Value

			cases = append(cases, objectCase(
				"additional property "+validCase.Name,
				object,
			))
		}
		return cases
	case o.AdditionalProperties.Allowed != nil && !*o.AdditionalProperties.Allowed:
		return nil
	default:
		object := cloneObject(requiredObject)
		object[additionalPropertyKey] = json.RawMessage(`"additional-property"`)

		return []Case{objectCase("additional property", object)}
	}
}

func (o *ObjectNode) additionalPropertyInvalidCases(requiredObject map[string]json.RawMessage) []Case {
	additionalPropertyKey := o.additionalPropertyKey()

	switch {
	case o.AdditionalProperties.Schema != nil:
		var cases []Case
		schema := *o.AdditionalProperties.Schema
		for _, invalidCase := range schema.InvalidCases() {
			object := cloneObject(requiredObject)
			object[additionalPropertyKey] = invalidCase.Value

			cases = append(cases, objectCase(
				"invalid additional property "+invalidCase.Name,
				object,
			))
		}
		return cases
	case o.AdditionalProperties.Allowed != nil && !*o.AdditionalProperties.Allowed:
		object := cloneObject(requiredObject)
		object[additionalPropertyKey] = json.RawMessage(`"not-allowed"`)

		return []Case{objectCase("additional property not allowed", object)}
	default:
		return nil
	}
}

func (o *ObjectNode) objectCaseData() ([]string, []string, map[string]json.RawMessage, []map[string]json.RawMessage) {
	requiredNames := make([]string, 0, len(o.Required))
	requiredSet := map[string]struct{}{}
	for _, name := range o.Required {
		if _, ok := requiredSet[name]; ok {
			continue
		}

		requiredSet[name] = struct{}{}
		requiredNames = append(requiredNames, name)
	}

	baselineRequiredObject := map[string]json.RawMessage{}
	requiredValidNames := make([]string, 0, len(requiredNames))
	var requiredValidCases [][]Case
	for _, name := range requiredNames {
		schema, ok := o.Properties[name]
		if !ok {
			continue
		}

		validCases := schema.ValidCases()
		if len(validCases) == 0 {
			continue
		}

		baselineRequiredObject[name] = validCases[0].Value
		requiredValidNames = append(requiredValidNames, name)
		requiredValidCases = append(requiredValidCases, validCases)
	}

	requiredObjects := []map[string]json.RawMessage{cloneObject(baselineRequiredObject)}
	for propertyIndex, validCases := range requiredValidCases {
		name := requiredValidNames[propertyIndex]
		for _, validCase := range validCases[1:] {
			object := cloneObject(baselineRequiredObject)
			object[name] = validCase.Value
			requiredObjects = append(requiredObjects, object)
		}
	}

	optionalNames := make([]string, 0, len(o.Properties))
	for name := range o.Properties {
		if _, ok := requiredSet[name]; !ok {
			optionalNames = append(optionalNames, name)
		}
	}
	sort.Strings(optionalNames)

	return requiredNames, optionalNames, baselineRequiredObject, requiredObjects
}

func (o *ObjectNode) additionalPropertyKey() string {
	if _, ok := o.Properties[additionalPropertyCaseKey]; !ok {
		return additionalPropertyCaseKey
	}

	for suffix := 2; ; suffix++ {
		key := additionalPropertyCaseKey + "_" + strconv.Itoa(suffix)
		if _, ok := o.Properties[key]; !ok {
			return key
		}
	}
}

func cloneObject(object map[string]json.RawMessage) map[string]json.RawMessage {
	clone := make(map[string]json.RawMessage, len(object))
	for name, value := range object {
		clone[name] = value
	}

	return clone
}
