package testgenerator

import "gopkg.in/yaml.v3"

var _ Hasher = new(Property)
var _ YamlParser = new(ObjectDomain)

type AdditionalPolicyKind int

const (
	AdditionalTrue AdditionalPolicyKind = iota
	AdditionalFalse
	AdditionalSchema
)

type Property struct {
	Key string
	*Hash
	Required bool
}

func (p *Property) GenerateHash() (Hash, error) {
	//TODO implement me
	panic("implement me")
}

type ObjectDomain struct {
	Properties []*Hash

	AdditionalPropertyKind   AdditionalPolicyKind
	AdditionalPropertyDomain *Hash

	MinProps int
	MaxProps *int
}

func (o *ObjectDomain) Parse(node yaml.Node) error {
	// Parse Properties

	// Parse AdditionalProperties

	// Parse MinProps, MaxProps

	//

	//TODO implement me
	panic("implement me")
}
