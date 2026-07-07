package domain

import (
	"decode_and_validate_generator/pkg/test_generator/hashables"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

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
