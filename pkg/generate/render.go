package generate

import (
	"bytes"
	"embed"
	"fmt"
	"go/format"
	"sync"
	"text/template"
)

//go:embed templates/*.go.tmpl
var templateFS embed.FS

const templatePattern = "templates/*.go.tmpl"

var (
	generateTemplatesOnce sync.Once
	generateTemplates     *template.Template
	generateTemplatesErr  error
)

type fileTemplateContext struct {
	Schemas     []Schema
	UsesRFC3339 bool
}

type modelsTestTemplateContext struct {
	OpenAPI    string
	Operations []JSONRequestBodyOperation
}

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

func usesRFC3339(schemas []Schema) bool {
	for _, schema := range schemas {
		stringSchema, ok := schema.(*StringSchema)
		if ok && stringSchema.Format == "date-time" {
			return true
		}
	}

	return false
}

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
		generateTemplates, generateTemplatesErr = template.ParseFS(templateFS, templatePattern)
	})

	if generateTemplatesErr != nil {
		return nil, fmt.Errorf("parse generate templates: %w", generateTemplatesErr)
	}

	return generateTemplates, nil
}
