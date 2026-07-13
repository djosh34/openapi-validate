package decode_tests

import (
	"testing"

	"decode_and_validate_generator/pkg/decode/example"

	"github.com/stretchr/testify/require"
)

// TestArrayNullableDecodeAllowedWays checks accepted nullable array forms.
func TestArrayNullableDecodeAllowedWays(t *testing.T) {
	t.Parallel()

	for name, tt := range map[string]struct {
		inputJSON      string
		expectedStruct example.ArrayNullable
	}{
		"null array": {
			inputJSON:      `null`,
			expectedStruct: example.ArrayNullable{},
		},
		"empty array": {
			inputJSON: `[]`,
			expectedStruct: example.ArrayNullable{
				Value: []example.ArrayNullableItem{},
			},
		},
		"single item": {
			inputJSON: `["one"]`,
			expectedStruct: example.ArrayNullable{
				Value: []example.ArrayNullableItem{"one"},
			},
		},
		"multiple items": {
			inputJSON: `["one","two",""]`,
			expectedStruct: example.ArrayNullable{
				Value: []example.ArrayNullableItem{"one", "two", ""},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			var actualStruct example.ArrayNullable

			// Act
			err := actualStruct.UnmarshalJSON([]byte(tt.inputJSON))

			// Assert
			require.NoError(t, err)
			require.Equal(t, tt.expectedStruct, actualStruct)
		})
	}
}

// TestArrayNullableDecodeRejectsInvalidShapes checks rejected nullable array forms.
func TestArrayNullableDecodeRejectsInvalidShapes(t *testing.T) {
	t.Parallel()

	for name, tt := range map[string]struct {
		inputJSON   string
		expectedErr error
	}{
		"not array string": {
			inputJSON: `"not-array"`,
		},
		"not array object": {
			inputJSON: `{}`,
		},
		"not array number": {
			inputJSON: `123`,
		},
		"not array bool": {
			inputJSON: `true`,
		},
		"null item": {
			inputJSON:   `[null]`,
			expectedErr: example.NullForNotNullableStringError,
		},
		"number item": {
			inputJSON:   `[123]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"bool item": {
			inputJSON:   `[false]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"object item": {
			inputJSON:   `[{}]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"array item": {
			inputJSON:   `[[]]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"invalid json": {
			inputJSON: `[`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			var actualStruct example.ArrayNullable

			// Act
			err := actualStruct.UnmarshalJSON([]byte(tt.inputJSON))

			// Assert
			if tt.expectedErr == nil {
				require.Error(t, err)
			} else {
				require.ErrorIs(t, err, tt.expectedErr)
			}
		})
	}
}

// TestArrayNotNullableDecodeAllowedWays checks accepted non-nullable array forms.
func TestArrayNotNullableDecodeAllowedWays(t *testing.T) {
	t.Parallel()

	for name, tt := range map[string]struct {
		inputJSON      string
		expectedStruct example.ArrayNotNullable
	}{
		"empty array": {
			inputJSON:      `[]`,
			expectedStruct: example.ArrayNotNullable{},
		},
		"single item": {
			inputJSON:      `["one"]`,
			expectedStruct: example.ArrayNotNullable{"one"},
		},
		"multiple items": {
			inputJSON:      `["one","two",""]`,
			expectedStruct: example.ArrayNotNullable{"one", "two", ""},
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			var actualStruct example.ArrayNotNullable

			// Act
			err := actualStruct.UnmarshalJSON([]byte(tt.inputJSON))

			// Assert
			require.NoError(t, err)
			require.Equal(t, tt.expectedStruct, actualStruct)
		})
	}
}

// TestArrayNotNullableDecodeRejectsInvalidShapes checks rejected non-nullable array forms.
func TestArrayNotNullableDecodeRejectsInvalidShapes(t *testing.T) {
	t.Parallel()

	for name, tt := range map[string]struct {
		inputJSON           string
		expectedErr         error
		expectedErrContains string
	}{
		"null array": {
			inputJSON:           `null`,
			expectedErrContains: "null for not nullable array",
		},
		"not array string": {
			inputJSON: `"not-array"`,
		},
		"not array object": {
			inputJSON: `{}`,
		},
		"not array number": {
			inputJSON: `123`,
		},
		"not array bool": {
			inputJSON: `true`,
		},
		"null item": {
			inputJSON:   `[null]`,
			expectedErr: example.NullForNotNullableStringError,
		},
		"number item": {
			inputJSON:   `[123]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"bool item": {
			inputJSON:   `[false]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"object item": {
			inputJSON:   `[{}]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"array item": {
			inputJSON:   `[[]]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"invalid json": {
			inputJSON: `[`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			var actualStruct example.ArrayNotNullable

			// Act
			err := actualStruct.UnmarshalJSON([]byte(tt.inputJSON))

			// Assert
			switch {
			case tt.expectedErr != nil:
				require.ErrorIs(t, err, tt.expectedErr)
			case tt.expectedErrContains != "":
				require.ErrorContains(t, err, tt.expectedErrContains)
			default:
				require.Error(t, err)
			}
		})
	}
}

// TestOptionalArrayNullableDecodeAllowedWays checks accepted optional array forms.
func TestOptionalArrayNullableDecodeAllowedWays(t *testing.T) {
	t.Parallel()

	for name, tt := range map[string]struct {
		inputJSON      []byte
		expectedStruct example.OptionalArrayNullable
	}{
		"nil body": {
			expectedStruct: example.OptionalArrayNullable{},
		},
		"empty body": {
			inputJSON:      []byte(""),
			expectedStruct: example.OptionalArrayNullable{},
		},
		"null array": {
			inputJSON:      []byte(`null`),
			expectedStruct: example.OptionalArrayNullable{},
		},
		"empty array": {
			inputJSON: []byte(`[]`),
			expectedStruct: example.OptionalArrayNullable{
				Value: []example.OptionalArrayNullableItem{},
			},
		},
		"single item": {
			inputJSON: []byte(`["one"]`),
			expectedStruct: example.OptionalArrayNullable{
				Value: []example.OptionalArrayNullableItem{"one"},
			},
		},
		"multiple items": {
			inputJSON: []byte(`["one","two",""]`),
			expectedStruct: example.OptionalArrayNullable{
				Value: []example.OptionalArrayNullableItem{"one", "two", ""},
			},
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			var actualStruct example.OptionalArrayNullable

			// Act
			err := actualStruct.UnmarshalJSON(tt.inputJSON)

			// Assert
			require.NoError(t, err)
			require.Equal(t, tt.expectedStruct, actualStruct)
		})
	}
}

// TestOptionalArrayNullableDecodeRejectsInvalidShapes checks rejected optional array forms.
func TestOptionalArrayNullableDecodeRejectsInvalidShapes(t *testing.T) {
	t.Parallel()

	for name, tt := range map[string]struct {
		inputJSON   []byte
		expectedErr error
	}{
		"not array string": {
			inputJSON: []byte(`"not-array"`),
		},
		"not array object": {
			inputJSON: []byte(`{}`),
		},
		"not array number": {
			inputJSON: []byte(`123`),
		},
		"not array bool": {
			inputJSON: []byte(`true`),
		},
		"null item": {
			inputJSON:   []byte(`[null]`),
			expectedErr: example.NullForNotNullableStringError,
		},
		"number item": {
			inputJSON:   []byte(`[123]`),
			expectedErr: example.NonStringForStringSchemaError,
		},
		"bool item": {
			inputJSON:   []byte(`[false]`),
			expectedErr: example.NonStringForStringSchemaError,
		},
		"object item": {
			inputJSON:   []byte(`[{}]`),
			expectedErr: example.NonStringForStringSchemaError,
		},
		"array item": {
			inputJSON:   []byte(`[[]]`),
			expectedErr: example.NonStringForStringSchemaError,
		},
		"invalid json": {
			inputJSON: []byte(`[`),
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			var actualStruct example.OptionalArrayNullable

			// Act
			err := actualStruct.UnmarshalJSON(tt.inputJSON)

			// Assert
			if tt.expectedErr == nil {
				require.Error(t, err)
			} else {
				require.ErrorIs(t, err, tt.expectedErr)
			}
		})
	}
}
