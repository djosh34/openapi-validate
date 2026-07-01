package decode_tests

import (
	"testing"

	"decode_and_validate_generator/pkg/decode/example"

	"github.com/stretchr/testify/require"
)

func TestArrayNullableDecodeAllowedWays(t *testing.T) {
	t.Parallel()

	for name, tt := range map[string]struct {
		inputJson      string
		expectedStruct example.ArrayNullable
	}{
		"null array": {
			inputJson:      `null`,
			expectedStruct: example.ArrayNullable{},
		},
		"empty array": {
			inputJson: `[]`,
			expectedStruct: example.ArrayNullable{
				Value: []example.ArrayNullableItem{},
			},
		},
		"single item": {
			inputJson: `["one"]`,
			expectedStruct: example.ArrayNullable{
				Value: []example.ArrayNullableItem{"one"},
			},
		},
		"multiple items": {
			inputJson: `["one","two",""]`,
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
			err := actualStruct.UnmarshalJSON([]byte(tt.inputJson))

			// Assert
			require.NoError(t, err)
			require.Equal(t, tt.expectedStruct, actualStruct)
		})
	}
}

func TestArrayNullableDecodeRejectsInvalidShapes(t *testing.T) {
	t.Parallel()

	for name, tt := range map[string]struct {
		inputJson   string
		expectedErr error
	}{
		"not array string": {
			inputJson: `"not-array"`,
		},
		"not array object": {
			inputJson: `{}`,
		},
		"not array number": {
			inputJson: `123`,
		},
		"not array bool": {
			inputJson: `true`,
		},
		"null item": {
			inputJson:   `[null]`,
			expectedErr: example.NullForNotNullableStringError,
		},
		"number item": {
			inputJson:   `[123]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"bool item": {
			inputJson:   `[false]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"object item": {
			inputJson:   `[{}]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"array item": {
			inputJson:   `[[]]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"invalid json": {
			inputJson: `[`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			var actualStruct example.ArrayNullable

			// Act
			err := actualStruct.UnmarshalJSON([]byte(tt.inputJson))

			// Assert
			if tt.expectedErr == nil {
				require.Error(t, err)
			} else {
				require.ErrorIs(t, err, tt.expectedErr)
			}
		})
	}
}

func TestArrayNotNullableDecodeAllowedWays(t *testing.T) {
	t.Parallel()

	for name, tt := range map[string]struct {
		inputJson      string
		expectedStruct example.ArrayNotNullable
	}{
		"empty array": {
			inputJson:      `[]`,
			expectedStruct: example.ArrayNotNullable{},
		},
		"single item": {
			inputJson:      `["one"]`,
			expectedStruct: example.ArrayNotNullable{"one"},
		},
		"multiple items": {
			inputJson:      `["one","two",""]`,
			expectedStruct: example.ArrayNotNullable{"one", "two", ""},
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			var actualStruct example.ArrayNotNullable

			// Act
			err := actualStruct.UnmarshalJSON([]byte(tt.inputJson))

			// Assert
			require.NoError(t, err)
			require.Equal(t, tt.expectedStruct, actualStruct)
		})
	}
}

func TestArrayNotNullableDecodeRejectsInvalidShapes(t *testing.T) {
	t.Parallel()

	for name, tt := range map[string]struct {
		inputJson           string
		expectedErr         error
		expectedErrContains string
	}{
		"null array": {
			inputJson:           `null`,
			expectedErrContains: "null for not nullable array",
		},
		"not array string": {
			inputJson: `"not-array"`,
		},
		"not array object": {
			inputJson: `{}`,
		},
		"not array number": {
			inputJson: `123`,
		},
		"not array bool": {
			inputJson: `true`,
		},
		"null item": {
			inputJson:   `[null]`,
			expectedErr: example.NullForNotNullableStringError,
		},
		"number item": {
			inputJson:   `[123]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"bool item": {
			inputJson:   `[false]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"object item": {
			inputJson:   `[{}]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"array item": {
			inputJson:   `[[]]`,
			expectedErr: example.NonStringForStringSchemaError,
		},
		"invalid json": {
			inputJson: `[`,
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			var actualStruct example.ArrayNotNullable

			// Act
			err := actualStruct.UnmarshalJSON([]byte(tt.inputJson))

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

func TestOptionalArrayNullableDecodeAllowedWays(t *testing.T) {
	t.Parallel()

	for name, tt := range map[string]struct {
		inputJson      []byte
		expectedStruct example.OptionalArrayNullable
	}{
		"nil body": {
			expectedStruct: example.OptionalArrayNullable{},
		},
		"empty body": {
			inputJson:      []byte(""),
			expectedStruct: example.OptionalArrayNullable{},
		},
		"null array": {
			inputJson:      []byte(`null`),
			expectedStruct: example.OptionalArrayNullable{},
		},
		"empty array": {
			inputJson: []byte(`[]`),
			expectedStruct: example.OptionalArrayNullable{
				Value: []example.OptionalArrayNullableItem{},
			},
		},
		"single item": {
			inputJson: []byte(`["one"]`),
			expectedStruct: example.OptionalArrayNullable{
				Value: []example.OptionalArrayNullableItem{"one"},
			},
		},
		"multiple items": {
			inputJson: []byte(`["one","two",""]`),
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
			err := actualStruct.UnmarshalJSON(tt.inputJson)

			// Assert
			require.NoError(t, err)
			require.Equal(t, tt.expectedStruct, actualStruct)
		})
	}
}

func TestOptionalArrayNullableDecodeRejectsInvalidShapes(t *testing.T) {
	t.Parallel()

	for name, tt := range map[string]struct {
		inputJson   []byte
		expectedErr error
	}{
		"not array string": {
			inputJson: []byte(`"not-array"`),
		},
		"not array object": {
			inputJson: []byte(`{}`),
		},
		"not array number": {
			inputJson: []byte(`123`),
		},
		"not array bool": {
			inputJson: []byte(`true`),
		},
		"null item": {
			inputJson:   []byte(`[null]`),
			expectedErr: example.NullForNotNullableStringError,
		},
		"number item": {
			inputJson:   []byte(`[123]`),
			expectedErr: example.NonStringForStringSchemaError,
		},
		"bool item": {
			inputJson:   []byte(`[false]`),
			expectedErr: example.NonStringForStringSchemaError,
		},
		"object item": {
			inputJson:   []byte(`[{}]`),
			expectedErr: example.NonStringForStringSchemaError,
		},
		"array item": {
			inputJson:   []byte(`[[]]`),
			expectedErr: example.NonStringForStringSchemaError,
		},
		"invalid json": {
			inputJson: []byte(`[`),
		},
	} {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Arrange
			var actualStruct example.OptionalArrayNullable

			// Act
			err := actualStruct.UnmarshalJSON(tt.inputJson)

			// Assert
			if tt.expectedErr == nil {
				require.Error(t, err)
			} else {
				require.ErrorIs(t, err, tt.expectedErr)
			}
		})
	}
}
