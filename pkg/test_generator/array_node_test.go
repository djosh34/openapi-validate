package testgenerator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestArrayNodeValidCasesIncludeEmptySingleAndMultipleItems(t *testing.T) {
	node := ArrayNode{
		BaseNode: BaseNode{Nullable: true},
		Items:    stringSchema(true),
	}

	require.Equal(t, []string{
		`null`,
		`[]`,
		`["valid-string"]`,
		`[null]`,
		`["valid-string",null]`,
	}, rawMessages(node.ValidCases()))
	require.Contains(t, caseNames(node.ValidCases()), "multiple items")
}

func TestArrayNodeInvalidCasesIncludeInvalidItemPositions(t *testing.T) {
	node := ArrayNode{
		Items: stringSchema(false),
	}

	invalidCases := rawMessages(node.InvalidCases())

	require.Contains(t, invalidCases, `null`)
	require.Contains(t, invalidCases, `"not-array"`)
	require.Contains(t, invalidCases, `[null]`)
	require.Contains(t, invalidCases, `["valid-string",null]`)
	require.Contains(t, invalidCases, `[123]`)
	require.Contains(t, invalidCases, `["valid-string",123]`)
}
