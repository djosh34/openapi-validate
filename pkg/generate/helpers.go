package generate

import (
	"context"
	"fmt"
	"maps"
	"os"
	"slices"

	"github.com/getkin/kin-openapi/openapi3"
)

// LoadOpenapi loads and validates an OpenAPI document for generation.
func LoadOpenapi(ctx context.Context, path string) (*GenerateContext, error) {
	openAPISource, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	loader := &openapi3.Loader{Context: ctx, IsExternalRefsAllowed: false}

	doc, err := loader.LoadFromFile(path)
	if err != nil {
		return nil, err
	}

	err = doc.Validate(ctx)
	if err != nil {
		return nil, err
	}

	return &GenerateContext{
		Document:      doc,
		OpenAPISource: openAPISource,
	}, nil
}

// FilterOperations removes operations not named in operation.
func (c *GenerateContext) FilterOperations(operation ...string) error {
	if len(operation) == 0 {
		return nil
	}

	if c.Document == nil || c.Document.Paths == nil {
		return fmt.Errorf("openapi document has no paths")
	}

	required := make(map[string]struct{}, len(operation))
	for _, operationID := range operation {
		required[operationID] = struct{}{}
	}

	for _, path := range c.Document.Paths.InMatchingOrder() {
		pathItem := c.Document.Paths.Value(path)
		if pathItem == nil {
			continue
		}

		filterPathItemOperations(pathItem, required)

		if len(pathItem.Operations()) == 0 {
			c.Document.Paths.Delete(path)
		}
	}

	if len(required) != 0 {
		return fmt.Errorf("operation not found: %v", slices.Sorted(maps.Keys(required)))
	}

	return nil
}

// filterPathItemOperations removes operations absent from required.
func filterPathItemOperations(pathItem *openapi3.PathItem, required map[string]struct{}) {
	for method, operation := range pathItem.Operations() {
		if operation == nil {
			continue
		}

		if _, ok := required[operation.OperationID]; ok {
			delete(required, operation.OperationID)

			continue
		}

		pathItem.SetOperation(method, nil)
	}
}
