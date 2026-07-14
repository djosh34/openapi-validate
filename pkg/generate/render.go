//nolint:godoclint // The private template contexts are local implementation details.
package generate

import (
	"bytes"
	"embed"
	"fmt"
	"strconv"
	"text/template"

	"github.com/djosh34/decode_and_validate_generator/pkg/validation"
	"golang.org/x/tools/imports"
)

//go:embed templates/*.go.tmpl
var templateFiles embed.FS

type sourceTemplate struct {
	Package     string
	Assignments []assignmentTemplate
}

type assignmentTemplate struct {
	Name  string
	Nodes []string
	Links []string
}

type testTemplate struct {
	Package    string
	OpenAPI    string
	Operations []Operation
}

func render(
	packageName string,
	openAPI []byte,
	operations []Operation,
	parsed []*validation.Validation,
) (map[string][]byte, error) {
	templates, err := template.ParseFS(templateFiles, "templates/*.go.tmpl")
	if err != nil {
		return nil, fmt.Errorf("parse templates: %w", err)
	}

	source := sourceTemplate{Package: packageName}

	for index, compiled := range parsed {
		source.Assignments = append(source.Assignments, renderAssignment(operations[index].Variable, compiled))
	}

	validate, err := executeTemplate(templates, "validate.go.tmpl", source)
	if err != nil {
		return nil, err
	}

	validateTest, err := executeTemplate(templates, "validate_test.go.tmpl", testTemplate{
		Package:    packageName,
		OpenAPI:    strconv.Quote(string(openAPI)),
		Operations: operations,
	})
	if err != nil {
		return nil, err
	}

	return map[string][]byte{
		"validate.go":      validate,
		"validate_test.go": validateTest,
	}, nil
}

func executeTemplate(templates *template.Template, name string, data any) ([]byte, error) {
	var output bytes.Buffer
	if err := templates.ExecuteTemplate(&output, name, data); err != nil {
		return nil, fmt.Errorf("execute %s: %w", name, err)
	}

	formatted, err := imports.Process(name, output.Bytes(), nil)
	if err != nil {
		return nil, fmt.Errorf("format %s: %w", name, err)
	}

	return formatted, nil
}
