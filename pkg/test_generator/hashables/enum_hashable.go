package hashables

import (
	"crypto/sha256"
	"decode_and_validate_generator/pkg/test_generator/types"
	"encoding/json"
	"errors"
)

type EnumHashable struct {
	*json.RawMessage
}

type enumHashableHashJSON struct {
	Type  string       `json:"type"`
	Value EnumHashable `json:"value"`
}

var _ types.Hasher = new(EnumHashable)

func (e *EnumHashable) GenerateHash() (types.Hash, error) {
	if e == nil {
		return types.Hash{}, errors.New("hashable of enum cannot be nil")
	}

	jsonBytes, err := json.Marshal(enumHashableHashJSON{Type: "enum", Value: *e})
	if err != nil {
		return types.Hash{}, err
	}

	return sha256.Sum256(jsonBytes), nil
}
