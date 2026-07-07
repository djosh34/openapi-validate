package hashables

import (
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

	expectedHash := types.Hash{0xd5, 0x74, 0x90, 0x77, 0xb6, 0xab, 0x1, 0x4a, 0x7c, 0x59, 0xa3, 0x1b, 0xeb, 0x8a, 0x8d, 0x61, 0xed, 0x67, 0x99, 0x6a, 0xcd, 0x76, 0x4, 0xa8, 0xef, 0x9c, 0x1d, 0x69, 0x47, 0x5d, 0xd9, 0x61}

	gotHash, err := hashable.GenerateHash()
	require.NoError(t, err)
	require.Equal(t, expectedHash, gotHash)
}

func TestEnumHashableGenerateHashNil(t *testing.T) {
	_, err := (*EnumHashable)(nil).GenerateHash()
	require.Error(t, err)
}
