package domain

import (
	"crypto/sha256"
	"decode_and_validate_generator/pkg/test_generator/types"
	"encoding/json"
	"errors"
)

type StringDomain struct {
	Nullable bool `json:"nullable"`

	Enum []string `json:"enum"`

	Pattern *string `json:"pattern"`
	Format  *string `json:"format"`

	MinLength int  `json:"minLength"`
	MaxLength *int `json:"maxLength"`
}

type stringDomainHashJson struct {
	Type  string       `json:"type"`
	Value StringDomain `json:"value"`
}

var _ types.Hasher = new(StringDomain)

func (domain *StringDomain) GenerateHash() (types.Hash, error) {
	if domain == nil {
		return types.Hash{}, errors.New("domain of string cannot be nil")
	}

	sdJson := stringDomainHashJson{
		Type:  "string",
		Value: *domain,
	}

	jsonBytes, err := json.Marshal(&sdJson)
	if err != nil {
		return types.Hash{}, err
	}

	return sha256.Sum256(jsonBytes), nil
}
