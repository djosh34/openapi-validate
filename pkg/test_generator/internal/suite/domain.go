// Package suite compiles OpenAPI schemas into canonical constructive domains.
package suite

import (
	"fmt"

	//nolint:depguard // Internal suite architecture intentionally depends on internal/jsonvalue.
	"decode_and_validate_generator/pkg/test_generator/internal/jsonvalue"
)

// DomainID identifies one canonical Domain in a DomainRegistry.
type DomainID uint32

const (
	// NoDomain is returned when compilation did not produce a Domain.
	NoDomain DomainID = iota
	// AnyJSONDomainID identifies the Domain containing every JSON value.
	AnyJSONDomainID
	// EmptyDomainID identifies the Domain containing no JSON value.
	EmptyDomainID
)

// KindState says whether a JSON kind is absent, unconstrained, or constrained.
type KindState uint8

const (
	// KindExcluded means no value of the kind is reachable.
	KindExcluded KindState = iota
	// KindUnrestricted means every value of the kind is reachable.
	KindUnrestricted
	// KindRestricted means the kind's constraint set applies.
	KindRestricted
)

// DomainStatus describes whether a Domain can construct values.
type DomainStatus uint8

const (
	// DomainProductive contains at least one constructible JSON value.
	DomainProductive DomainStatus = iota
	// DomainEmpty contains no JSON value.
	DomainEmpty
	// DomainUnsupported uses understood OpenAPI behavior not implemented by this step.
	DomainUnsupported
	// DomainUnconstructible is understood but lacks enough trusted generation input.
	DomainUnconstructible
)

// Domain is one constructible set spanning all JSON value kinds.
type Domain struct {
	Null    KindState
	Boolean KindState
	Number  NumberConstraints
	String  StringConstraints
	Array   ArrayConstraints
	Object  ObjectConstraints
	Enum    *EnumSet
	Status  DomainStatus
}

// NumberBound is one exact inclusive or exclusive number bound.
type NumberBound struct {
	Value     jsonvalue.Number
	Exclusive bool
}

// NumberConstraints constrains JSON numbers.
type NumberConstraints struct {
	State        KindState
	IntegersOnly bool
	Minimum      *NumberBound
	Maximum      *NumberBound
	MultipleOf   *jsonvalue.Number
	Format       *string
}

// StringConstraints constrains JSON strings.
type StringConstraints struct {
	State     KindState
	MinLength int
	MaxLength *int
	Patterns  []string
	Formats   []string
}

// ArrayConstraints constrains JSON arrays and references their item Domain.
type ArrayConstraints struct {
	State    KindState
	Items    DomainID
	MinItems int
	MaxItems *int
}

// PropertyState says whether a named object property may occur.
type PropertyState uint8

const (
	// PropertyAllowed permits the property with Values.
	PropertyAllowed PropertyState = iota
	// PropertyForbidden rejects the property.
	PropertyForbidden
)

// NamedProperty constrains one object property.
type NamedProperty struct {
	Name     string
	Required bool
	State    PropertyState
	Values   DomainID
}

// AdditionalProperties constrains object names without a NamedProperty.
type AdditionalProperties struct {
	Values DomainID
}

// ObjectConstraints constrains JSON objects and references property Domains.
type ObjectConstraints struct {
	State      KindState
	Properties []NamedProperty
	Additional AdditionalProperties
	MinProps   int
	MaxProps   *int
}

// EnumSet is a canonical finite set of exact semantic JSON values.
type EnumSet struct {
	Values []jsonvalue.Value
}

// DomainPair is an unordered pair reserved for the step-4 intersection cache.
type DomainPair struct {
	First  DomainID
	Second DomainID
}

// ConstraintSource locates one compiled schema keyword.
type ConstraintSource struct {
	Pointer string
	Keyword string
}

// GenerationExamples are trusted generation inputs and not Domain identity.
type GenerationExamples struct {
	Valid   []jsonvalue.Value
	Invalid []jsonvalue.Value
}

// SchemaUse preserves source metadata separately from a canonical Domain.
type SchemaUse struct {
	Pointer     string
	Domain      DomainID
	Constraints []ConstraintSource
	Examples    GenerationExamples
}

// Error reports a stable compilation outcome with source context.
type Error struct {
	Phase   string
	Code    string
	Pointer string
	Keyword string
	Cause   error
}

// Error formats compilation context.
func (compileError *Error) Error() string {
	location := compileError.Pointer
	if compileError.Keyword != "" {
		location += "/" + compileError.Keyword
	}

	return fmt.Sprintf("%s %s at %s: %v", compileError.Phase, compileError.Code, location, compileError.Cause)
}

// Unwrap returns the underlying failure.
func (compileError *Error) Unwrap() error {
	return compileError.Cause
}
