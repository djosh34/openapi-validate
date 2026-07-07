package hashables

import (
	"decode_and_validate_generator/pkg/test_generator/types"
	"testing"

	"github.com/stretchr/testify/require"
)

type fakeHasher struct{}

func (f fakeHasher) GenerateHash() (types.Hash, error) {
	return types.Hash{1}, nil
}

func TestObjectHashablesImplementHasher(t *testing.T) {
	require.Implements(t, (*types.Hasher)(nil), new(PropertyHashable))
	require.Implements(t, (*types.Hasher)(nil), new(ObjectHashable))
}

func TestPropertyHashableGenerateHash(t *testing.T) {
	hashable := PropertyHashable{Key: "name", Hasher: fakeHasher{}, Required: true}

	expectedHash := types.Hash{0x80, 0xb6, 0x8e, 0x84, 0xf0, 0xbb, 0x33, 0xbf, 0xff, 0xd6, 0x4, 0xb1, 0x74, 0xae, 0x2c, 0x3e, 0x86, 0x81, 0x70, 0x23, 0x27, 0xc8, 0xfa, 0xf1, 0x6b, 0xbd, 0x90, 0x53, 0x38, 0xfe, 0xa, 0x58}

	gotHash, err := hashable.GenerateHash()
	require.NoError(t, err)
	require.Equal(t, expectedHash, gotHash)
}

func TestObjectHashableGenerateHash(t *testing.T) {
	hashable := ObjectHashable{
		Nullable:                 true,
		Properties:               []types.Hasher{fakeHasher{}},
		AdditionalPropertyKind:   AdditionalSchema,
		AdditionalPropertyDomain: fakeHasher{},
		MinProps:                 1,
		MaxProps:                 new(3),
	}

	expectedHash := types.Hash{0x9a, 0x14, 0x79, 0x7e, 0xbd, 0x1c, 0x40, 0x10, 0xbb, 0x6, 0x91, 0x23, 0x6c, 0xaa, 0xca, 0x69, 0x9c, 0x87, 0x21, 0x5e, 0x70, 0x71, 0x93, 0x1d, 0x22, 0xe6, 0xb7, 0x8, 0x48, 0x6, 0x3a, 0xf8}

	gotHash, err := hashable.GenerateHash()
	require.NoError(t, err)
	require.Equal(t, expectedHash, gotHash)
}

func TestObjectHashableGenerateHashNil(t *testing.T) {
	_, err := (*PropertyHashable)(nil).GenerateHash()
	require.Error(t, err)

	_, err = (*ObjectHashable)(nil).GenerateHash()
	require.Error(t, err)
}
