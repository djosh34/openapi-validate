package hashables

import (
	"crypto/sha256"
	"decode_and_validate_generator/pkg/test_generator/types"
	"encoding/json"
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

	jsonBytes, err := json.Marshal(propertyHashableHashJSON{Type: "property", Value: hashable})
	require.NoError(t, err)
	expectedHash := types.Hash(sha256.Sum256(jsonBytes))

	gotHash, err := hashable.GenerateHash()
	require.NoError(t, err)
	require.Equal(t, expectedHash, gotHash)
}

func TestObjectHashableGenerateHash(t *testing.T) {
	hashable := ObjectHashable{
		Properties:               []types.Hasher{fakeHasher{}},
		AdditionalPropertyKind:   AdditionalSchema,
		AdditionalPropertyDomain: fakeHasher{},
		MinProps:                 1,
		MaxProps:                 new(3),
	}

	jsonBytes, err := json.Marshal(objectHashableHashJSON{Type: "object", Value: hashable})
	require.NoError(t, err)
	expectedHash := types.Hash(sha256.Sum256(jsonBytes))

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
