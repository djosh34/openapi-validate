package domain

import (
	"decode_and_validate_generator/pkg/test_generator/types"
	"encoding/json"
	"errors"
)

type NumberDomain struct {
	Nullable bool     `json:"nullable"`
	Enum     []Number `json:"enum"`

	Minimum          *Number `json:"minimum"`
	Maximum          *Number `json:"maximum"`
	ExclusiveMinimum bool    `json:"exclusiveMinimum"`
	ExclusiveMaximum bool    `json:"exclusiveMaximum"`
	MultipleOf       *Number `json:"multipleOf"`
	Format           *string `json:"format"`
}

func (n *NumberDomain) ToHasher() (types.Hasher, error) {
	if n == nil {
		return nil, errors.New("domain of number cannot be nil")
	}

	panic("TO DO")
}

func (dc *DomainContext) ParseNumber(node *json.RawMessage) (NumberDomain, error) {
	panic("TO DO")
}
