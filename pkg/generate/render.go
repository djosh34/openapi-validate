package generate

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"maps"
	"slices"
	"sync"
	"text/template"
)

//go:embed templates/*.tmpl
var templateFS embed.FS

var (
	generateTemplatesOnce sync.Once
	generateTemplates     *template.Template
	generateTemplatesErr  error
)

type fileTemplateContext struct {
	Schemas []SchemaObject
}

func renderModelsFile(schemas []SchemaObject) ([]byte, error) {
	templates, err := parsedGenerateTemplates()
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
	templates, err := parsedGenerateTemplates()
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	err = templates.ExecuteTemplate(&out, name, data)
	if err != nil {
		return "", err
	}

	return out.String(), nil
}

func parsedGenerateTemplates() (*template.Template, error) {
	generateTemplatesOnce.Do(func() {
		generateTemplates, generateTemplatesErr = template.ParseFS(templateFS, "templates/*.tmpl")
	})
	if generateTemplatesErr != nil {
		return nil, fmt.Errorf("parse generate templates: %w", generateTemplatesErr)
	}

	return generateTemplates, nil
}

func schemaDefinitions(schemas []SchemaObject) ([]SchemaObject, error) {
	definitionsByName := map[string]SchemaObject{}
	for _, schema := range schemas {
		err := collectSchemaDefinitions(schema, definitionsByName)
		if err != nil {
			return nil, err
		}
	}

	definitions := make([]SchemaObject, 0, len(definitionsByName))
	for _, name := range slices.Sorted(maps.Keys(definitionsByName)) {
		definitions = append(definitions, definitionsByName[name])
	}

	return definitions, nil
}

func collectSchemaDefinitions(schema SchemaObject, definitions map[string]SchemaObject) error {
	if schema == nil {
		return fmt.Errorf("nil schema context")
	}

	switch schema := schema.(type) {
	case ObjectContext:
		for _, property := range schema.Properties {
			err := collectSchemaDefinitions(property.Schema, definitions)
			if err != nil {
				return err
			}
		}
		if schema.AdditionalPropertiesSchema != nil {
			err := collectSchemaDefinitions(schema.AdditionalPropertiesSchema, definitions)
			if err != nil {
				return err
			}
		}
	case *ObjectContext:
		if schema == nil {
			return fmt.Errorf("nil object schema context")
		}

		return collectSchemaDefinitions(*schema, definitions)
	case ArrayContext:
		if schema.Items == nil {
			return fmt.Errorf("array schema %q has nil items", schema.Name())
		}

		err := collectSchemaDefinitions(schema.Items, definitions)
		if err != nil {
			return err
		}
	case *ArrayContext:
		if schema == nil {
			return fmt.Errorf("nil array schema context")
		}

		return collectSchemaDefinitions(*schema, definitions)
	case StringContext:
	case *StringContext:
		if schema == nil {
			return fmt.Errorf("nil string schema context")
		}
	default:
		return fmt.Errorf("unsupported schema context %T", schema)
	}

	definitions[schema.Name()] = schema
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
		return !schema.AdditionalProperties || schema.HasRequiredProperty()
	case *ObjectContext:
		return schema != nil && schemaUsesFmt(*schema)
	case ArrayContext:
		return !schema.IsNullable()
	case *ArrayContext:
		return schema != nil && schemaUsesFmt(*schema)
	default:
		return false
	}
}
