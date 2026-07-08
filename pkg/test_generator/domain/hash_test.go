package domain

import (
	"crypto/sha256"
	"encoding/json"
	"testing"

	"decode_and_validate_generator/pkg/test_generator/types"

	"github.com/stretchr/testify/require"
)

func requireGeneratedHash(t *testing.T, hashType string, value any) types.Hash {
	t.Helper()

	jsonBytes, err := json.Marshal(domainHashJSON{Type: hashType, Value: value})
	require.NoError(t, err)

	return sha256.Sum256(jsonBytes)
}
