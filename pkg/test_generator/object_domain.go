package testgenerator

type AdditionalPolicyKind int

const (
	AdditionalTrue AdditionalPolicyKind = iota
	AdditionalFalse
	AdditionalSchema
)

type Property struct {
	Domain
	Required bool
}

type ObjectDomain struct {
	Properties map[string]Property

	AdditionalPropertyKind   AdditionalPolicyKind
	AdditionalPropertyDomain Domain

	MinProps int
	MaxProps *int
}
