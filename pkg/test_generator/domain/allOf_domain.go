package domain

import (
	"decode_and_validate_generator/pkg/test_generator/types"
	"encoding/json"
	"errors"
)

type AllOfDomain struct {
	Domains      []types.Domain
	MergedDomain types.Domain
}

func (a *AllOfDomain) ToHasher() (types.Hasher, error) {
	if a == nil {
		return nil, errors.New("domain of allOf cannot be nil")
	}

	return nil, errors.New("NOT IMPLEMENTED")
}

func (dc *DomainContext) ParseAllOf(node *json.RawMessage) (AllOfDomain, error) {

	return AllOfDomain{}, errors.New("NOT IMPLEMENTED")
}
