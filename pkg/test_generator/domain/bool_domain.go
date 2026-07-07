package domain

import (
	"decode_and_validate_generator/pkg/test_generator/types"
	"encoding/json"
	"errors"
)

type BoolDomain struct {
	Nullable bool   `json:"nullable"`
	Enum     []bool `json:"enum"`
}

func (b *BoolDomain) ToHasher() (types.Hasher, error) {
	if b == nil {
		return nil, errors.New("domain of bool cannot be nil")
	}

	panic("TO DO")
}

func (dc *DomainContext) ParseBool(node *json.RawMessage) (BoolDomain, error) {
	panic("TO DO")
}
