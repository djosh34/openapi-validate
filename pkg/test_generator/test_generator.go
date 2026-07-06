package testgenerator

import "fmt"

func GenerateValid(openAPIYAMLSpec []byte, operationID string, unmarshal func([]byte) error) error {
	_ = unmarshal

	_, err := OpenAPIRequestBodySchemaNode(openAPIYAMLSpec, operationID)
	if err != nil {
		return fmt.Errorf("openapi yaml spec parse failed: %w", err)
	}

	return nil
}

func GenerateInvalid(openAPIYAMLSpec []byte, operationID string, unmarshal func([]byte) error) error {
	_ = unmarshal

	_, err := OpenAPIRequestBodySchemaNode(openAPIYAMLSpec, operationID)
	if err != nil {
		return fmt.Errorf("openapi yaml spec parse failed: %w", err)
	}

	return nil
}
