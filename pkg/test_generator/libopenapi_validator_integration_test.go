package testgenerator

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"testing"

	//nolint:depguard // This test-only dependency loads documents for the independent integration SUT.
	"github.com/pb33f/libopenapi"
	//nolint:depguard // This test-only dependency is the independent integration SUT.
	validator "github.com/pb33f/libopenapi-validator"
	"github.com/stretchr/testify/require"
)

// TestCheckJSONRequestBodyAgainstLibopenapiValidator gives generated CasePlans to an independent validator.
func TestCheckJSONRequestBodyAgainstLibopenapiValidator(t *testing.T) {
	t.Parallel()

	fixtures := []struct {
		name       string
		schema     string
		components string
	}{
		{
			name: "realistic nested order",
			schema: `
      type: object
      required: [customer, lines, state]
      properties:
        customer:
          type: object
          required: [id, contact]
          properties:
            id: {type: integer, minimum: 1}
            # v0.13.13 treats format as annotation, so only its trusted valid example is cross-checked.
            contact: {type: string, format: email, x-valid-examples: [buyer@example.com]}
            note: {type: string, nullable: true, maxLength: 12}
          additionalProperties: false
        lines:
          type: array
          minItems: 1
          maxItems: 3
          items:
            type: object
            required: [sku, quantity]
            properties:
              sku: {type: string, minLength: 1, maxLength: 8}
              quantity: {type: integer, minimum: 1, maximum: 20}
            additionalProperties: false
        state: {type: string, enum: [draft, submitted]}
        tags:
          type: array
          maxItems: 2
          items: {type: string, minLength: 1}
      additionalProperties: false
`,
		},
		{
			name: "array item constraints",
			schema: `
      type: array
      minItems: 1
      maxItems: 3
      items:
        type: integer
        minimum: -2
        maximum: 2
`,
		},
		{
			name: "additional properties true",
			schema: `
      type: object
      required: [known]
      properties:
        known: {type: boolean}
      additionalProperties: true
`,
		},
		{
			name: "additional properties schema",
			schema: `
      type: object
      properties:
        fixed: {type: string}
      additionalProperties:
        type: integer
        minimum: 0
`,
		},
		{
			name: "local references",
			schema: `
      $ref: '#/components/schemas/Envelope'
`,
			components: `
components:
  schemas:
    Identifier:
      type: string
      minLength: 2
      maxLength: 6
    Envelope:
      type: object
      required: [id, payload]
      properties:
        id: {$ref: '#/components/schemas/Identifier'}
        payload:
          type: array
          minItems: 1
          items: {$ref: '#/components/schemas/Identifier'}
      additionalProperties: false
`,
		},
		{
			name: "allOf object intersection",
			schema: `
      type: object
      properties:
        id: {type: integer, minimum: 1}
        score: {type: number, minimum: 0, maximum: 10}
      additionalProperties: false
      allOf:
        - type: object
          required: [id]
          properties:
            id: {type: integer, minimum: 1}
        - type: object
          required: [score]
          properties:
            score: {type: number, minimum: 0, maximum: 10}
`,
		},
		{
			name: "allOf number intersection",
			schema: `
      allOf:
        - {type: number, minimum: -2, exclusiveMinimum: true}
        - {maximum: 2, exclusiveMaximum: true}
`,
		},
		{
			name: "trusted string languages and unicode lengths",
			schema: `
      type: object
      required: [code, email]
      properties:
        code:
          type: string
          minLength: 2
          maxLength: 3
          pattern: '^λ[0-9]$'
          x-valid-examples: [λ7]
          x-invalid-examples: [λ, xx]
        email:
          type: string
          format: email
          # v0.13.13 treats format as annotation, so no rejection is expected from this independent SUT.
          x-valid-examples: [a@example.com]
      additionalProperties: false
`,
		},
		{
			name: "typeless mixed keyword families",
			schema: `
      minLength: 1
      maxLength: 2
      minimum: -1
      maximum: 1
      minItems: 1
      maxItems: 2
      minProperties: 1
      maxProperties: 2
`,
		},
		{
			name: "zero collection bounds",
			schema: `
      maxLength: 0
      maxItems: 0
      maxProperties: 0
`,
		},
		{
			name: "decimal multipleOf",
			schema: `
      type: number
      minimum: -2
      maximum: 2
      multipleOf: 0.25
`,
		},
		{
			name: "null only",
			schema: `
      type: string
      nullable: true
      enum: [null]
`,
		},
		{
			name: "optional impossible property",
			schema: `
      type: object
      properties:
        impossible: {type: string, minLength: 2, maxLength: 1}
        live: {type: boolean}
      additionalProperties: false
`,
		},
		{
			name: "empty array with impossible items",
			schema: `
      type: array
      minItems: 0
      maxItems: 0
      items: {type: string, minLength: 2, maxLength: 1}
`,
		},
		{
			name: "adversarial property names and exact edges",
			schema: `
      type: object
      required: ['a/b', 't~n', empty]
      minProperties: 3
      maxProperties: 3
      properties:
        'a/b': {type: string, minLength: 0, maxLength: 0}
        't~n': {type: array, minItems: 0, maxItems: 0, items: {}}
        empty: {type: object, minProperties: 0, maxProperties: 0, additionalProperties: false}
      additionalProperties: false
`,
		},
	}

	for _, fixture := range fixtures {
		t.Run(fixture.name, func(t *testing.T) {
			t.Parallel()

			spec := append(requestBodySpec(fixture.schema), fixture.components...)
			validate := newLibopenapiRequestBodyValidator(t, spec)

			CheckJSONRequestBody(t, spec, "checkThing", validate)
		})
	}
}

// TestLibopenapiValidatorLargeNumberLimitations isolates the independent SUT's float64 JSON decoding limits.
// Keeping these characterizations separate avoids weakening any assertion in the broad CasePlan matrix.
func TestLibopenapiValidatorLargeNumberLimitations(t *testing.T) {
	t.Parallel()

	t.Run("adjacent integers above exact float64 range are conflated", func(t *testing.T) {
		t.Parallel()

		spec := requestBodySpec(`
      type: integer
      minimum: 9007199254740993
      maximum: 9007199254740993
`)
		validate := newLibopenapiRequestBodyValidator(t, spec)

		// 9007199254740992 violates minimum, but v0.13.13 rounds both it and the bound to the same float64.
		require.NoError(t, validate([]byte(`9007199254740992`)))
	})

	t.Run("finite JSON number outside float64 range is rejected", func(t *testing.T) {
		t.Parallel()

		spec := requestBodySpec(`
      type: number
`)
		validate := newLibopenapiRequestBodyValidator(t, spec)

		// JSON and OpenAPI numbers have no float64 range limit.
		err := validate([]byte(`1e400`))
		require.Error(t, err)
		require.Contains(t, err.Error(), "cannot unmarshal number 1e400 into Go value of type float64")
	})
}

// newLibopenapiRequestBodyValidator builds the independent request-body SUT once per schema.
func newLibopenapiRequestBodyValidator(t *testing.T, spec []byte) func([]byte) error {
	t.Helper()

	document, err := libopenapi.NewDocument(spec)
	require.NoError(t, err)

	requestValidator, buildErrs := validator.NewValidator(document)
	require.Empty(t, buildErrs)
	require.NotNil(t, requestValidator)
	t.Cleanup(requestValidator.Release)

	return func(body []byte) error {
		request, err := http.NewRequest(http.MethodPost, "/things", bytes.NewReader(body))
		if err != nil {
			return fmt.Errorf("construct integration request: %w", err)
		}

		request.Header.Set("Content-Type", "application/json")

		valid, validationErrs := requestValidator.ValidateHttpRequestSync(request)
		if valid && len(validationErrs) != 0 {
			return errors.New("libopenapi-validator returned valid with validation errors")
		}

		if valid {
			return nil
		}

		if len(validationErrs) == 0 {
			return errors.New("libopenapi-validator rejected request without an error")
		}

		messages := make([]string, len(validationErrs))
		for index, validationErr := range validationErrs {
			if validationErr == nil {
				return errors.New("libopenapi-validator returned a nil validation error")
			}

			messages[index] = validationErr.Error()
		}

		return errors.New(strings.Join(messages, "; "))
	}
}
