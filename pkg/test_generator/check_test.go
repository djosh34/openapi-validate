package testgenerator

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestCheckJSONRequestBodyRunsEveryExampleAsValidJSON verifies scheduling and JSON syntax.
func TestCheckJSONRequestBodyRunsEveryExampleAsValidJSON(t *testing.T) {
	t.Parallel()

	spec := requestBodyExampleSpec(`{
		"type":"object",
		"x-valid-examples":[
			{"child":"baseline"},
			{"child":"alternate","unicode":"λ"}
		],
		"x-invalid-examples":[
			{},
			{"child":42}
		]
	}`)

	calls := make(map[string]int)

	CheckJSONRequestBody(t, spec, "checkThing", func(body []byte) error {
		require.True(t, json.Valid(body))

		compact, err := compactJSON(body)
		require.NoError(t, err)

		calls[compact]++

		if strings.Contains(compact, `"child":"baseline"`) ||
			strings.Contains(compact, `"child":"alternate"`) {
			return nil
		}

		return errors.New("rejected")
	})

	require.Len(t, calls, 4)

	for body, count := range calls {
		require.Positive(t, count, body)
	}
}

// TestCompileCasePlansRejectsMissingOrMalformedExamples verifies the incremental example contract.
func TestCompileCasePlansRejectsMissingOrMalformedExamples(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		schema    string
		wantError string
	}{
		"valid examples missing": {
			schema:    `{"type":"string","x-invalid-examples":[1]}`,
			wantError: "x-valid-examples is required",
		},
		"invalid examples missing": {
			schema:    `{"type":"string","x-valid-examples":["ok"]}`,
			wantError: "x-invalid-examples is required",
		},
		"examples are not array": {
			schema:    `{"type":"string","x-valid-examples":"ok","x-invalid-examples":["bad"]}`,
			wantError: "x-valid-examples must be an array",
		},
		"examples are empty": {
			schema:    `{"type":"string","x-valid-examples":[],"x-invalid-examples":["bad"]}`,
			wantError: "x-valid-examples must not be empty",
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			plans, err := compileCasePlans(requestBodyExampleSpec(tt.schema), "checkThing")
			require.ErrorContains(t, err, tt.wantError)
			require.Nil(t, plans)
		})
	}
}

// TestCheckJSONRequestBodyFindsBuggyValidatorsByKeywordFamily proves known bugs reach the SUT.
func TestCheckJSONRequestBodyFindsBuggyValidatorsByKeywordFamily(t *testing.T) {
	t.Parallel()

	families := []string{"string", "number", "array", "object", "boolean-enum"}
	for _, family := range families {
		t.Run(family, func(t *testing.T) {
			t.Parallel()

			command := exec.Command(
				os.Args[0],
				"-test.run=^TestCheckJSONRequestBodyBuggyValidatorHelper$",
				"-test.v",
				"-rapid.checks=1",
				"-rapid.nofailfile",
			)

			command.Env = append(os.Environ(), "TEST_GENERATOR_BUG_FAMILY="+family)

			output, err := command.CombinedOutput()
			require.Error(t, err)
			require.Contains(t, string(output), "x-invalid-examples_/_0")
			require.Contains(t, string(output), "invalid JSON accepted")
		})
	}
}

// TestCheckJSONRequestBodyBuggyValidatorHelper runs one intentionally failing subprocess case.
func TestCheckJSONRequestBodyBuggyValidatorHelper(t *testing.T) {
	t.Parallel()

	family := os.Getenv("TEST_GENERATOR_BUG_FAMILY")
	if family == "" {
		t.Skip("subprocess helper")
	}

	fixture, ok := buggyValidatorFixtures()[family]
	require.True(t, ok)

	CheckJSONRequestBody(t, requestBodyExampleSpec(fixture.schema), "checkThing", func(body []byte) error {
		compact, err := compactJSON(body)
		if err != nil {
			return err
		}

		if compact == fixture.buggyInvalid {
			return nil
		}

		for _, valid := range fixture.valid {
			if compact == valid {
				return nil
			}
		}

		return errors.New("rejected")
	})
}

// compactJSON removes insignificant JSON whitespace for fixture comparisons.
func compactJSON(body []byte) (string, error) {
	var compact bytes.Buffer
	if err := json.Compact(&compact, body); err != nil {
		return "", err
	}

	return compact.String(), nil
}

// buggyValidatorFixture describes one schema family and its fake validator bug.
type buggyValidatorFixture struct {
	schema       string
	valid        []string
	buggyInvalid string
}

// buggyValidatorFixtures returns one deliberately faulty validator input per supported family.
func buggyValidatorFixtures() map[string]buggyValidatorFixture {
	return map[string]buggyValidatorFixture{
		"string": {
			schema:       `{"type":"string","minLength":3,"x-valid-examples":["good"],"x-invalid-examples":["xy"]}`,
			valid:        []string{`"good"`},
			buggyInvalid: `"xy"`,
		},
		"number": {
			schema:       `{"type":"integer","minimum":10,"x-valid-examples":[10],"x-invalid-examples":[9]}`,
			valid:        []string{`10`},
			buggyInvalid: `9`,
		},
		"array": {
			schema: `{"type":"array","items":{"type":"string"},"minItems":1,` +
				`"x-valid-examples":[["ok"]],"x-invalid-examples":[[]]}`,
			valid:        []string{`["ok"]`},
			buggyInvalid: `[]`,
		},
		"object": {
			schema: `{"type":"object","required":["child"],` +
				`"properties":{"child":{"type":"string"}},` +
				`"x-valid-examples":[{"child":"baseline"},{"child":"alternate"}],` +
				`"x-invalid-examples":[{"child":42}]}`,
			valid:        []string{`{"child":"baseline"}`, `{"child":"alternate"}`},
			buggyInvalid: `{"child":42}`,
		},
		"boolean-enum": {
			schema:       `{"type":"boolean","enum":[true],"x-valid-examples":[true],"x-invalid-examples":[false]}`,
			valid:        []string{`true`},
			buggyInvalid: `false`,
		},
	}
}

// requestBodyExampleSpec embeds a request Schema Object in a minimal OpenAPI document.
func requestBodyExampleSpec(schema string) []byte {
	return fmt.Appendf(nil, `
openapi: 3.0.3
info:
  title: contract test
  version: 1.0.0
paths:
  /things:
    post:
      operationId: checkThing
      requestBody:
        content:
          application/json:
            schema: %s
`, schema)
}
