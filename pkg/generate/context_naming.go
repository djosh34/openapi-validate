package generate

import (
	"fmt"
	"unicode"
)

// namedSchemaDefinitions assigns names and returns every schema definition.
func namedSchemaDefinitions(schema Schema) ([]Schema, error) {
	var definitions []Schema

	err := nameSchema(schema, "", &definitions)
	if err != nil {
		return nil, err
	}

	return definitions, nil
}

// nameSchema assigns a parent-qualified name and recursively collects children.
func nameSchema(schema Schema, parentName string, definitions *[]Schema) error {
	if schema == nil {
		return fmt.Errorf("nil schema")
	}

	base := schema.Base()
	if base == nil {
		return fmt.Errorf("schema %T has nil base", schema)
	}

	if base.Name == "" {
		return fmt.Errorf("schema %T has no type name", schema)
	}

	base.Name = exportedName(parentName + base.Name)

	*definitions = append(*definitions, schema)

	for _, child := range schema.ChildSchemas() {
		err := nameSchema(child, base.Name, definitions)
		if err != nil {
			return fmt.Errorf("child schema: %w", err)
		}
	}

	return nil
}

// exportedName converts name to an exported Go identifier.
func exportedName(name string) string {
	return identifierName(name, true)
}

// unexportedName converts name to an unexported Go identifier.
func unexportedName(name string) string {
	return identifierName(name, false)
}

// identifierName converts arbitrary text to a Go identifier.
func identifierName(name string, exported bool) string {
	var out []rune

	upperNext := exported

	for _, character := range name {
		if !unicode.IsLetter(character) && !unicode.IsDigit(character) {
			upperNext = true

			continue
		}

		if len(out) == 0 && unicode.IsDigit(character) {
			out = append(out, []rune("Schema")...)
		}

		if upperNext {
			out = append(out, unicode.ToUpper(character))
			upperNext = false

			continue
		}

		out = append(out, character)
	}

	if len(out) == 0 {
		if exported {
			return "Schema"
		}

		return "schema"
	}

	if exported {
		return string(out)
	}

	out[0] = unicode.ToLower(out[0])

	return string(out)
}
