// Package generate writes compiled request-body validations and query decoders as Go source.
package generate

import (
	"errors"
	"fmt"
	"go/token"
	"os"
	"path/filepath"

	"github.com/djosh34/klopt/pkg/patternvalidator"
	"github.com/djosh34/klopt/pkg/validation"
)

const (
	// directoryMode is used for the generated directory.
	directoryMode = 0o755
	// fileMode is used for generated Go files.
	fileMode = 0o644
)

var (
	// ErrNilPatternOption reports a nil pattern option.
	ErrNilPatternOption = errors.New("generate: nil pattern option")
	// ErrUnsafeOperationID reports an operation ID that cannot name generated Go state.
	ErrUnsafeOperationID = errors.New("generate: unsafe operation ID")
)

// Generate parses one OpenAPI document and writes validate.go and validate_test.go.
func Generate(
	dir string,
	packageName string,
	openAPI []byte,
	patternOption patternvalidator.Option,
) error {
	files, err := GenerateInMemory(packageName, openAPI, patternOption)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(dir, directoryMode); err != nil {
		return err
	}

	for name, contents := range files {
		if err := os.WriteFile(filepath.Join(dir, name), contents, fileMode); err != nil {
			return err
		}
	}

	return nil
}

// GenerateInMemory parses one OpenAPI document and returns validate.go and validate_test.go.
//
//nolint:revive // GenerateInMemory is the required public API name.
func GenerateInMemory(
	packageName string,
	openAPI []byte,
	patternOption patternvalidator.Option,
) (map[string][]byte, error) {
	if patternOption == nil {
		return nil, ErrNilPatternOption
	}

	settings := patternSettings{}
	captureSettings := patternvalidator.Option(func(compiled *patternvalidator.PatternValidation) {
		patternOption(compiled)
		settings.RejectNonASCII = compiled.RejectsNonASCII()
		settings.UseRE2 = compiled.UsesRE2()
	})

	parsed, queryDecoders, err := validation.Parse(openAPI, captureSettings)
	if err != nil {
		return nil, err
	}

	for operationID := range parsed {
		if !isSafeOperationIdentifier(operationID) {
			return nil, fmt.Errorf(
				"%w: operation ID %q cannot be used as a generated Go identifier",
				ErrUnsafeOperationID,
				operationID,
			)
		}
	}

	files, err := render(packageName, openAPI, parsed, queryDecoders, settings)
	if err != nil {
		return nil, err
	}

	return files, nil
}

// isSafeOperationIdentifier reports whether an operation ID can name a generated package variable.
func isSafeOperationIdentifier(operationID string) bool {
	if !token.IsIdentifier(operationID) || operationID == "_" || operationID == "init" {
		return false
	}

	switch operationID {
	case "byte", "error", "errors", "json", "jsonvalue", "openAPI", "patternvalidator", "string", "testing",
		"testgenerator", "TestValidations", "true", "validation", "validations", "queryDecoders", "mustQueryDecoder":
		return false
	default:
		return true
	}
}
