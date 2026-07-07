package domain

import (
	"decode_and_validate_generator/pkg/test_generator/types"
	"encoding/json"
	"errors"
)

type ArrayDomain struct {
	Nullable bool `json:"nullable"`

	Items types.Domain `json:"items"`

	MinItems int  `json:"minItems"`
	MaxItems *int `json:"maxItems"`
}

func (a *ArrayDomain) ToHasher() (types.Hasher, error) {
	if a == nil {
		return nil, errors.New("domain of array cannot be nil")
	}

	return nil, errors.New("NOT IMPLEMENTED")
}

func (dc *DomainContext) ParseArray(node *json.RawMessage) (ArrayDomain, error) {
	return ArrayDomain{}, errors.New("NOT IMPLEMENTED")
}
