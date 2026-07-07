package testgenerator

import (
	"crypto/sha256"
	"encoding/json"
)

type StringDomain struct {
	Nullable bool `json:"nullable"`

	Enum []string `json:"enum"`

	Pattern *string `json:"pattern"`
	Format  *string `json:"format"`

	MinLength int  `json:"minLength"`
	MaxLength *int `json:"maxLength"`
}

var _ Hasher = new(StringDomain)

func (domain *StringDomain) Hash() Hash {
	return sha256.Sum256(mustMarshalJSON(json.Marshal(domain)))
}

func mustMarshalJSON(jsonBytes []byte, err error) []byte {
	if err != nil {
		panic(err)
	}

	return jsonBytes
}
