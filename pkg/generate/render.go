package generate

import (
	"bytes"
	"cmp"
	"embed"
	"fmt"
	"go/format"
	"maps"
	"slices"
	"strings"
	"text/template"
	"unicode"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

type fileTemplateContext struct {
	Schemas []SchemaObject
}

type objectProperty struct {
	JSONName  string
	Schema    SchemaObject
	Required  bool
	LocalName string
}

func renderModelsFile(schemas []SchemaObject) ([]byte, error) {
	templates, err := parseGenerateTemplates()
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer
	err = templates.ExecuteTemplate(&out, "file.tmpl", fileTemplateContext{Schemas: schemas})
	if err != nil {
		return nil, err
	}

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		return nil, fmt.Errorf("format generated models.go: %w", err)
	}

	return formatted, nil
}

func executeGoTemplate(name string, data any) (string, error) {
	templates, err := parseGenerateTemplates()
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	err = templates.ExecuteTemplate(&out, name, data)
	if err != nil {
		return "", err
	}

	return formatGoSnippet(out.String())
}

func parseGenerateTemplates() (*template.Template, error) {
	templates, err := template.ParseFS(templateFS, "templates/*.tmpl")
	if err != nil {
		return nil, fmt.Errorf("parse generate templates: %w", err)
	}

	return templates, nil
}

func formatGoSnippet(src string) (string, error) {
	const prefix = "package generated\n\n"

	formatted, err := format.Source([]byte(prefix + src))
	if err != nil {
		return "", fmt.Errorf("format generated snippet: %w", err)
	}

	return strings.TrimPrefix(string(formatted), prefix), nil
}

func namedOperationSchemas(operations map[string]SchemaObject) ([]SchemaObject, error) {
	operationIDs := slices.Sorted(maps.Keys(operations))
	schemas := make([]SchemaObject, 0, len(operationIDs))
	for _, operationID := range operationIDs {
		schema, err := namedSchemaObject(operations[operationID], exportedName(operationID))
		if err != nil {
			return nil, fmt.Errorf("operation %q schema: %w", operationID, err)
		}

		schemas = append(schemas, schema)
	}

	return schemas, nil
}

func namedSchemaObject(schemaObject SchemaObject, name string) (SchemaObject, error) {
	switch schema := schemaObject.(type) {
	case ObjectContext:
		properties, err := namedSchemaObjectMap(schema.Properties)
		if err != nil {
			return nil, err
		}

		schema.ContextName = exportedName(name)
		schema.Properties = properties
		if schema.AdditionalPropertiesSchema != nil {
			additionalPropertiesSchema, err := namedSchemaObject(schema.AdditionalPropertiesSchema, schema.ContextName+"AdditionalProperty")
			if err != nil {
				return nil, fmt.Errorf("additionalProperties schema: %w", err)
			}

			schema.AdditionalPropertiesSchema = additionalPropertiesSchema
		}

		return schema, nil
	case *ObjectContext:
		if schema == nil {
			return nil, fmt.Errorf("nil object schema context")
		}

		return namedSchemaObject(*schema, name)
	case StringContext:
		schema.ContextName = exportedName(name)
		return schema, nil
	case *StringContext:
		if schema == nil {
			return nil, fmt.Errorf("nil string schema context")
		}

		return namedSchemaObject(*schema, name)
	case ArrayContext:
		if schema.Items == nil {
			return nil, fmt.Errorf("array schema has nil items")
		}

		items, err := namedSchemaObject(schema.Items, exportedName(name)+"Item")
		if err != nil {
			return nil, fmt.Errorf("array items schema: %w", err)
		}

		schema.ContextName = exportedName(name)
		schema.Items = items
		return schema, nil
	case *ArrayContext:
		if schema == nil {
			return nil, fmt.Errorf("nil array schema context")
		}

		return namedSchemaObject(*schema, name)
	default:
		return nil, fmt.Errorf("unsupported schema context %T", schemaObject)
	}
}

func namedSchemaObjectMap(properties map[string]SchemaObject) (map[string]SchemaObject, error) {
	namedProperties := make(map[string]SchemaObject, len(properties))
	for propertyName, propertySchema := range properties {
		schema, err := namedSchemaObject(propertySchema, exportedName(propertyName))
		if err != nil {
			return nil, fmt.Errorf("property %q schema: %w", propertyName, err)
		}

		namedProperties[propertyName] = schema
	}

	return namedProperties, nil
}

func schemaDefinitions(schemas []SchemaObject) ([]SchemaObject, error) {
	var definitions []SchemaObject
	seen := map[string]struct{}{}
	for _, schema := range schemas {
		err := collectSchemaDefinitions(schema, seen, &definitions)
		if err != nil {
			return nil, err
		}
	}

	return definitions, nil
}

func collectSchemaDefinitions(schema SchemaObject, seen map[string]struct{}, definitions *[]SchemaObject) error {
	if schema == nil {
		return fmt.Errorf("nil schema context")
	}

	switch schema := schema.(type) {
	case ObjectContext:
		for _, property := range schema.OrderedProperties() {
			err := collectSchemaDefinitions(property.Schema, seen, definitions)
			if err != nil {
				return err
			}
		}
		if schema.AdditionalPropertiesSchema != nil {
			err := collectSchemaDefinitions(schema.AdditionalPropertiesSchema, seen, definitions)
			if err != nil {
				return err
			}
		}
	case *ObjectContext:
		if schema == nil {
			return fmt.Errorf("nil object schema context")
		}

		return collectSchemaDefinitions(*schema, seen, definitions)
	case ArrayContext:
		if schema.Items == nil {
			return fmt.Errorf("array schema %q has nil items", schema.Name())
		}

		err := collectSchemaDefinitions(schema.Items, seen, definitions)
		if err != nil {
			return err
		}
	case *ArrayContext:
		if schema == nil {
			return fmt.Errorf("nil array schema context")
		}

		return collectSchemaDefinitions(*schema, seen, definitions)
	case StringContext:
	case *StringContext:
		if schema == nil {
			return fmt.Errorf("nil string schema context")
		}
	default:
		return fmt.Errorf("unsupported schema context %T", schema)
	}

	if _, ok := seen[schema.Name()]; ok {
		return nil
	}

	seen[schema.Name()] = struct{}{}
	*definitions = append(*definitions, schema)
	return nil
}

func (c fileTemplateContext) UsesFmt() bool {
	for _, schema := range c.Schemas {
		if schemaUsesFmt(schema) {
			return true
		}
	}

	return false
}

func schemaUsesFmt(schema SchemaObject) bool {
	switch schema := schema.(type) {
	case ObjectContext:
		return !schema.AdditionalProperties || len(schema.Required) != 0
	case *ObjectContext:
		return schema != nil && schemaUsesFmt(*schema)
	case ArrayContext:
		return !schema.Nullable
	case *ArrayContext:
		return schema != nil && schemaUsesFmt(*schema)
	default:
		return false
	}
}

func (o ObjectContext) OrderedProperties() []objectProperty {
	requiredSet := make(map[string]struct{}, len(o.Required))
	properties := make([]objectProperty, 0, len(o.Properties))

	for _, propertyName := range o.Required {
		requiredSet[propertyName] = struct{}{}
		propertySchema, ok := o.Properties[propertyName]
		if !ok {
			continue
		}

		properties = append(properties, newObjectProperty(propertyName, propertySchema, true))
	}

	optionalNames := make([]string, 0, len(o.Properties))
	for propertyName := range o.Properties {
		if _, ok := requiredSet[propertyName]; ok {
			continue
		}

		optionalNames = append(optionalNames, propertyName)
	}
	slices.SortFunc(optionalNames, func(left string, right string) int {
		return compareOptionalProperties(o.Properties[left], o.Properties[right], left, right)
	})

	for _, propertyName := range optionalNames {
		properties = append(properties, newObjectProperty(propertyName, o.Properties[propertyName], false))
	}

	return properties
}

func (o ObjectContext) RequiredProperties() []objectProperty {
	requiredProperties := make([]objectProperty, 0, len(o.Required))
	for _, property := range o.OrderedProperties() {
		if property.Required {
			requiredProperties = append(requiredProperties, property)
		}
	}

	return requiredProperties
}

func newObjectProperty(jsonName string, schema SchemaObject, required bool) objectProperty {
	return objectProperty{
		JSONName:  jsonName,
		Schema:    schema,
		Required:  required,
		LocalName: unexportedName(schema.Name()),
	}
}

func (p objectProperty) FieldType() string {
	if p.Required {
		return p.Schema.Name()
	}

	return "*" + p.Schema.Name()
}

func (p objectProperty) JSONTag() string {
	if p.Required {
		return fmt.Sprintf("`json:%q`", p.JSONName)
	}

	return fmt.Sprintf("`json:%q`", p.JSONName+",omitzero")
}

func compareOptionalProperties(leftSchema SchemaObject, rightSchema SchemaObject, leftName string, rightName string) int {
	if nullableCompare := cmp.Compare(nullableRank(leftSchema), nullableRank(rightSchema)); nullableCompare != 0 {
		return nullableCompare
	}

	return cmp.Compare(leftName, rightName)
}

func nullableRank(schema SchemaObject) int {
	if schemaNullable(schema) {
		return 0
	}

	return 1
}

func schemaNullable(schema SchemaObject) bool {
	switch schema := schema.(type) {
	case ObjectContext:
		return schema.Nullable
	case *ObjectContext:
		return schema != nil && schema.Nullable
	case StringContext:
		return schema.Nullable
	case *StringContext:
		return schema != nil && schema.Nullable
	case ArrayContext:
		return schema.Nullable
	case *ArrayContext:
		return schema != nil && schema.Nullable
	default:
		return false
	}
}

func exportedName(name string) string {
	return identifierName(name, true)
}

func unexportedName(name string) string {
	return identifierName(name, false)
}

func identifierName(name string, exported bool) string {
	var out strings.Builder
	upperNext := exported
	for _, r := range name {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			upperNext = true
			continue
		}

		if out.Len() == 0 && unicode.IsDigit(r) {
			out.WriteString("Schema")
		}

		if upperNext {
			out.WriteRune(unicode.ToUpper(r))
			upperNext = false
			continue
		}

		out.WriteRune(r)
	}

	if out.Len() == 0 {
		if exported {
			return "Schema"
		}

		return "schema"
	}

	if exported {
		return out.String()
	}

	runes := []rune(out.String())
	runes[0] = unicode.ToLower(runes[0])
	return string(runes)
}
