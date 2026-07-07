package hashables

import (
	"crypto/sha256"
	"decode_and_validate_generator/pkg/test_generator/types"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnumHashableImplementsHasher(t *testing.T) {
	require.Implements(t, (*types.Hasher)(nil), new(EnumHashable))
}

func TestEnumHashableGenerateHash(t *testing.T) {
	raw := json.RawMessage(`"alpha"`)
	hashable := EnumHashable{RawMessage: &raw}

	jsonBytes, err := json.Marshal(enumHashableHashJSON{Type: "enum", Value: hashable})
	require.NoError(t, err)
	expectedHash := types.Hash(sha256.Sum256(jsonBytes))

	gotHash, err := hashable.GenerateHash()
	require.NoError(t, err)
	require.Equal(t, expectedHash, gotHash)
}

func TestEnumHashableGenerateHashNil(t *testing.T) {
	_, err := (*EnumHashable)(nil).GenerateHash()
	require.Error(t, err)
}
