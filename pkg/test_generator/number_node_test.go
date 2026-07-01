package testgenerator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNumberNodeValidCasesIncludeWideNumericRange(t *testing.T) {
	node := NumberNode{BaseNode: BaseNode{Nullable: true}}

	require.Equal(t, []string{
		`0`,
		`123.45`,
		`-123.45`,
		`2147483647`,
		`2147483648`,
		`-2147483648`,
		`-2147483649`,
		`3.4028236e38`,
		`-3.4028236e38`,
		`1e-39`,
		`null`,
	}, rawMessages(node.ValidCases()))
	require.Contains(t, caseNames(node.ValidCases()), "above int32 max")
	require.Contains(t, caseNames(node.ValidCases()), "below int32 min")
	require.Contains(t, caseNames(node.ValidCases()), "above float32 max")
	require.Contains(t, caseNames(node.ValidCases()), "below negative float32 max")
}

func TestNumberNodeInvalidCasesRejectNonNumbersAndNumericStrings(t *testing.T) {
	node := NumberNode{}

	require.Equal(t, []string{
		"null",
		"string",
		"numeric string",
		"boolean",
		"object",
		"array",
	}, caseNames(node.InvalidCases()))
}
