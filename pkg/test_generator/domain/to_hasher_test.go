package domain

import (
	"decode_and_validate_generator/pkg/test_generator/hashables"
	"decode_and_validate_generator/pkg/test_generator/types"
	"encoding/json"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

type failingToHasherDomain struct{}

func (f failingToHasherDomain) ToHasher() (types.Hasher, error) {
	return nil, errors.New("to hasher failed")
}

func TestStringDomainToHasher(t *testing.T) {
	domain := StringDomain{Nullable: true, Enum: []string{"alpha"}, Pattern: new("x"), Format: new("email"), MinLength: 1, MaxLength: new(5)}

	hasher, err := domain.ToHasher()
	require.NoError(t, err)
	require.Equal(t, &hashables.StringHashable{Nullable: true, Enum: []string{"alpha"}, Pattern: new("x"), Format: new("email"), MinLength: 1, MaxLength: new(5)}, hasher)
}

func TestStringDomainToHasherNil(t *testing.T) {
	_, err := (*StringDomain)(nil).ToHasher()
	require.Error(t, err)
}

func TestEnumDomainToHasher(t *testing.T) {
	raw := json.RawMessage(`"alpha"`)
	domain := EnumDomain{RawMessage: &raw}

	hasher, err := domain.ToHasher()
	require.NoError(t, err)
	require.Equal(t, &hashables.EnumHashable{RawMessage: &raw}, hasher)
}

func TestEnumDomainToHasherNil(t *testing.T) {
	_, err := (*EnumDomain)(nil).ToHasher()
	require.Error(t, err)
}

func TestPropertyToHasher(t *testing.T) {
	property := Property{Key: "name", Domain: &StringDomain{}, Required: true}

	hasher, err := property.ToHasher()
	require.NoError(t, err)
	require.Equal(t, &hashables.PropertyHashable{Key: "name", Hasher: &hashables.StringHashable{}, Required: true}, hasher)
}

func TestPropertyToHasherErrors(t *testing.T) {
	_, err := (*Property)(nil).ToHasher()
	require.Error(t, err)

	_, err = (&Property{Domain: failingToHasherDomain{}}).ToHasher()
	require.Error(t, err)
}

func TestObjectDomainToHasher(t *testing.T) {
	maxProps := new(3)
	object := ObjectDomain{
		Enum:                     []types.Domain{&EnumDomain{}},
		Properties:               []types.Domain{&Property{Key: "name", Domain: &StringDomain{}, Required: true}},
		AdditionalPropertyKind:   AdditionalSchema,
		AdditionalPropertyDomain: &StringDomain{},
		MinProps:                 1,
		MaxProps:                 maxProps,
	}

	hasher, err := object.ToHasher()
	require.NoError(t, err)
	require.Equal(t, &hashables.ObjectHashable{
		Enum:                     []types.Hasher{&hashables.EnumHashable{}},
		Properties:               []types.Hasher{&hashables.PropertyHashable{Key: "name", Hasher: &hashables.StringHashable{}, Required: true}},
		AdditionalPropertyKind:   hashables.AdditionalSchema,
		AdditionalPropertyDomain: &hashables.StringHashable{},
		MinProps:                 1,
		MaxProps:                 maxProps,
	}, hasher)
}

func TestObjectDomainToHasherErrors(t *testing.T) {
	_, err := (*ObjectDomain)(nil).ToHasher()
	require.Error(t, err)

	_, err = (&ObjectDomain{Enum: []types.Domain{failingToHasherDomain{}}}).ToHasher()
	require.Error(t, err)

	_, err = (&ObjectDomain{Properties: []types.Domain{failingToHasherDomain{}}}).ToHasher()
	require.Error(t, err)

	_, err = (&ObjectDomain{AdditionalPropertyDomain: failingToHasherDomain{}}).ToHasher()
	require.Error(t, err)
}
