package domain

import (
	"crypto/sha256"
	"encoding/json"

	"decode_and_validate_generator/pkg/test_generator/types"
)

type domainHashJSON struct {
	Type  string `json:"type"`
	Value any    `json:"value"`
}

func generateHash(hashType string, value any) (types.Hash, error) {
	jsonBytes, err := json.Marshal(domainHashJSON{Type: hashType, Value: value})
	if err != nil {
		return types.Hash{}, err
	}

	return sha256.Sum256(jsonBytes), nil
}
