package hashables

import (
	"decode_and_validate_generator/pkg/test_generator/types"
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

	expectedHash := types.Hash{0x19, 0xd0, 0x2d, 0x93, 0x89, 0xa0, 0xe3, 0x29, 0x81, 0x6b, 0x4a, 0x6b, 0x71, 0x74, 0x79, 0xf5, 0xf7, 0x49, 0xdc, 0x34, 0x84, 0xbe, 0x6b, 0xd, 0x6e, 0x65, 0x97, 0xb3, 0x31, 0xf8, 0xa0, 0x9c}

	gotHash, err := hashable.GenerateHash()
	require.NoError(t, err)
	require.Equal(t, expectedHash, gotHash)
}

func TestStringHashableGenerateHashNil(t *testing.T) {
	_, err := (*StringHashable)(nil).GenerateHash()
	require.Error(t, err)
}
