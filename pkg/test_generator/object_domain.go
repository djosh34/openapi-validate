package testgenerator

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

var _ Hasher = new(Property)
var _ Hasher = new(ObjectDomain)

type AdditionalPropertyKind int

const (
	AdditionalTrue AdditionalPropertyKind = iota
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
	Enum []*Hash

	Properties []*Hash

	AdditionalPropertyKind
	AdditionalPropertyDomain *Hash

	MinProps int
	MaxProps *int
}

func (o *ObjectDomain) GenerateHash() (Hash, error) {
	//TODO implement me
	panic("implement me")
}

type YamlKV map[string]yaml.Node
type YamlObject struct {
	Type                 string      `yaml:"type"`
	Required             []string    `yaml:"required"`
	Properties           YamlKV      `yaml:"properties"`
	AdditionalProperties yaml.Node   `yaml:"additionalProperties"`
	MinProperties        *int        `yaml:"minProperties"`
	MaxProperties        *int        `yaml:"maxProperties"`
	Enum                 []yaml.Node `yaml:"enum"`
}

type PropertyAlreadyExistsError struct {
	Key string
}

func (p *PropertyAlreadyExistsError) Error() string {
	return fmt.Sprintf("property %q already exists in object", p.Key)
}

func (dc *DomainContext) ParseObject(node yaml.Node) (ObjectDomain, error) {
	yamlKV := make(YamlKV)

	decodeKVErr := node.Decode(yamlKV)
	if decodeKVErr != nil {
		return ObjectDomain{}, decodeKVErr
	}

	yamlObject := YamlObject{}
	decodeErr := node.Decode(&yamlObject)
	if decodeErr != nil {
		return ObjectDomain{}, decodeErr
	}

	objectDomain := ObjectDomain{}

	// Parse Enums early, and if it exists, return early (we will not check that enum is valid, and only populate enum field of ObjectDomain)

	properties := make(map[string]Property, len(yamlObject.Properties))

	// Parse Properties
	if _, propertiesOk := yamlKV["properties"]; propertiesOk {
		delete(yamlKV, "properties")

		for propertyKey, propertyValue := range yamlObject.Properties {
			if _, propertyOk := properties[propertyKey]; propertyOk {
				return objectDomain, &PropertyAlreadyExistsError{
					Key: propertyKey,
				}
			}

			propertyHash, propertyErr := dc.Parse(propertyValue)
			if propertyErr != nil {
				return ObjectDomain{}, propertyErr
			}

			property := Property{
				Key:  propertyKey,
				Hash: propertyHash,
			}

			properties[propertyKey] = property
		}

	}

	// Parse required
	if _, requiredOk := yamlKV["required"]; requiredOk {
		delete(yamlKV, "required")

		for _, requiredKey := range yamlObject.Required {
			property, propertyOk := properties[requiredKey]
			if !propertyOk {
				property = Property{
					Key:      requiredKey,
					Required: true,
				}
			} else {
				property.Required = true
			}

			properties[requiredKey] = property
		}
	}

	// Parse AdditionalProperties

	// Parse MinProps, MaxProps

	// Reject if any other keys are left in node?

	return objectDomain, nil
}
