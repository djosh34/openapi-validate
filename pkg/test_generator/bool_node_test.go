package testgenerator

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBoolNodeValidCasesIncludeTrueFalseAndNullableNull(t *testing.T) {
	node := BoolNode{BaseNode: BaseNode{Nullable: true}}

	require.Equal(t, []Case{
		{Name: "true", Value: json.RawMessage(`true`)},
		{Name: "false", Value: json.RawMessage(`false`)},
		{Name: "null", Value: json.RawMessage(`null`)},
	}, node.ValidCases())
}

func TestBoolNodeInvalidCasesRejectCoercionShapes(t *testing.T) {
	node := BoolNode{}

	require.Equal(t, []string{
		"null",
		"string",
		"string true",
		"string false",
		"number",
		"zero",
		"object",
		"array",
	}, caseNames(node.InvalidCases()))
}
