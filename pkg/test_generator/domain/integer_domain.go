package domain

import (
	"decode_and_validate_generator/pkg/test_generator/types"
	"encoding/json"
	"errors"
)

type IntegerDomain struct {
	Nullable bool     `json:"nullable"`
	Enum     []Number `json:"enum"`

	Minimum          *Number `json:"minimum"`
	Maximum          *Number `json:"maximum"`
	ExclusiveMinimum bool    `json:"exclusiveMinimum"`
	ExclusiveMaximum bool    `json:"exclusiveMaximum"`
	MultipleOf       *Number `json:"multipleOf"`
	Format           *string `json:"format"`
}

func (i *IntegerDomain) ToHasher() (types.Hasher, error) {
	if i == nil {
		return nil, errors.New("domain of integer cannot be nil")
	}

	panic("TO DO")
}

func (dc *DomainContext) ParseInteger(node *json.RawMessage) (IntegerDomain, error) {
	panic("TO DO")
}
