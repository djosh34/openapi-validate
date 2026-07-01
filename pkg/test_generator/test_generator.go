package testgenerator

import "fmt"

type Case struct {
	Name      string
	JSON      string
	WantValid bool
}

func GenerateCasesFromOpenAPIFile(openapiPath string) ([]Case, error) {
	return nil, fmt.Errorf("generate cases from %q: not implemented", openapiPath)
}
