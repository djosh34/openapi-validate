package testgenerator

type ObjectDomain struct {
	Properties map[string]Domain
	Required   map[string]bool

	Additional AdditionalPolicy

	MinProps int
	MaxProps *int
}

type AdditionalPolicyKind int

const (
	AdditionalTrue AdditionalPolicyKind = iota
	AdditionalFalse
	AdditionalSchema
)

type AdditionalPolicy struct {
	Kind   AdditionalPolicyKind
	Domain Domain
}
