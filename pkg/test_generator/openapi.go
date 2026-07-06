package testgenerator

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type openAPIDocument struct {
	Paths map[string]openAPIPathItem `yaml:"paths"`
}

type openAPIPathItem struct {
	Get     *openAPIOperation `yaml:"get"`
	Put     *openAPIOperation `yaml:"put"`
	Post    *openAPIOperation `yaml:"post"`
	Delete  *openAPIOperation `yaml:"delete"`
	Options *openAPIOperation `yaml:"options"`
	Head    *openAPIOperation `yaml:"head"`
	Patch   *openAPIOperation `yaml:"patch"`
	Trace   *openAPIOperation `yaml:"trace"`
}

type openAPIOperation struct {
	OperationID string              `yaml:"operationId"`
	RequestBody *openAPIRequestBody `yaml:"requestBody"`
}

type openAPIRequestBody struct {
	Content map[string]openAPIMediaType `yaml:"content"`
}

type openAPIMediaType struct {
	Schema yaml.Node `yaml:"schema"`
}

func OpenAPIRequestBodySchemaNode(openAPIYAMLSpec []byte, operationID string) (yaml.Node, error) {
	var document openAPIDocument
	err := yaml.Unmarshal(openAPIYAMLSpec, &document)
	if err != nil {
		return yaml.Node{}, fmt.Errorf("parse openapi yaml spec: %w", err)
	}

	var matches []*openAPIOperation
	for _, pathItem := range document.Paths {
		for _, operation := range []*openAPIOperation{
			pathItem.Get,
			pathItem.Put,
			pathItem.Post,
			pathItem.Delete,
			pathItem.Options,
			pathItem.Head,
			pathItem.Patch,
			pathItem.Trace,
		} {
			if operation == nil {
				continue
			}
			if operation.OperationID == operationID {
				matches = append(matches, operation)
			}
		}
	}

	switch len(matches) {
	case 0:
		return yaml.Node{}, fmt.Errorf("operationId %q not found", operationID)
	case 1:
	default:
		return yaml.Node{}, fmt.Errorf("operationId %q found multiple times", operationID)
	}

	operation := matches[0]
	if operation.RequestBody == nil || len(operation.RequestBody.Content) == 0 {
		return yaml.Node{}, fmt.Errorf("operationId %q request body content type does not exist", operationID)
	}

	mediaType, ok := operation.RequestBody.Content["application/json"]
	if !ok {
		return yaml.Node{}, fmt.Errorf("operationId %q request body content type is not json", operationID)
	}
	if mediaType.Schema.Kind == 0 {
		return yaml.Node{}, fmt.Errorf("operationId %q application/json schema does not exist", operationID)
	}

	return mediaType.Schema, nil
}
