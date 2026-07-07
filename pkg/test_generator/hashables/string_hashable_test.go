package hashables

import (
	"crypto/sha256"
	"decode_and_validate_generator/pkg/test_generator/types"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestStringHashableImplementsHasher(t *testing.T) {
	require.Implements(t, (*types.Hasher)(nil), new(StringHashable))
}

func TestStringHashableGenerateHash(t *testing.T) {
	hashable := StringHashable{
		Nullable:  true,
		Enum:      []string{"alpha", "beta"},
		Pattern:   new("^[a-z]+$"),
		Format:    new("email"),
		MinLength: 2,
		MaxLength: new(5),
	}

	jsonBytes, err := json.Marshal(stringHashableHashJSON{Type: "string", Value: hashable})
	require.NoError(t, err)
	expectedHash := types.Hash(sha256.Sum256(jsonBytes))

	gotHash, err := hashable.GenerateHash()
	require.NoError(t, err)
	require.NotEmpty(t, jsonBytes)
	require.Equal(t, expectedHash, gotHash)
}

func TestStringHashableGenerateHashNil(t *testing.T) {
	_, err := (*StringHashable)(nil).GenerateHash()
	require.Error(t, err)
}
