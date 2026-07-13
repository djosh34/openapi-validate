package generate

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"sync"
	"text/template"
)

// templateFS contains every generator template.
//
//go:embed templates/*.go.tmpl
var templateFS embed.FS

// templatePattern selects every embedded generator template.
const templatePattern = "templates/*.go.tmpl"

var (
	// generateTemplatesOnce ensures templates are parsed once.
	generateTemplatesOnce sync.Once
	// generateTemplates contains the parsed templates.
	generateTemplates *template.Template
	// errGenerateTemplates records a template parsing failure.
	errGenerateTemplates error
)

// fileTemplateContext contains data rendered into models.go.
type fileTemplateContext struct {
	Schemas     []Schema
	UsesRFC3339 bool
}

// modelsTestTemplateContext contains data rendered into models_test.go.
type modelsTestTemplateContext struct {
	OpenAPI    string
	Operations []JSONRequestBodyOperation
}

// renderModelsFile renders and formats models.go.
func renderModelsFile(schemas []Schema) ([]byte, error) {
	templates, err := parsedGenerateTemplates()
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer

	err = templates.ExecuteTemplate(&out, "file.go.tmpl", fileTemplateContext{
		Schemas:     schemas,
		UsesRFC3339: usesRFC3339(schemas),
	})
	if err != nil {
		return nil, err
	}

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		return nil, fmt.Errorf("format generated models.go: %w", err)
	}

	return formatted, nil
}

// usesRFC3339 reports whether generated models need the time package.
func usesRFC3339(schemas []Schema) bool {
	for _, schema := range schemas {
		stringSchema, ok := schema.(*StringSchema)
		if ok && stringSchema.Format == "date-time" {
			return true
		}
	}

	return false
}

// renderModelsTestFile renders and formats models_test.go.
func renderModelsTestFile(openAPI []byte, operations []JSONRequestBodyOperation) ([]byte, error) {
	if bytes.Contains(openAPI, []byte("`")) {
		return nil, fmt.Errorf("openapi source contains backtick")
	}

	templates, err := parsedGenerateTemplates()
	if err != nil {
		return nil, err
	}

	var out bytes.Buffer

	err = templates.ExecuteTemplate(&out, "models_test.go.tmpl", modelsTestTemplateContext{
		OpenAPI:    string(openAPI),
		Operations: operations,
	})
	if err != nil {
		return nil, err
	}

	formatted, err := format.Source(out.Bytes())
	if err != nil {
		return nil, fmt.Errorf("format generated models_test.go: %w", err)
	}

	return formatted, nil
}

// executeGoTemplate executes one named generator template.
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

// parsedGenerateTemplates returns the shared parsed templates.
func parsedGenerateTemplates() (*template.Template, error) {
	generateTemplatesOnce.Do(func() {
		generateTemplates, errGenerateTemplates = template.ParseFS(templateFS, templatePattern)
	})

	if errGenerateTemplates != nil {
		return nil, fmt.Errorf("parse generate templates: %w", errGenerateTemplates)
	}

	return generateTemplates, nil
}
