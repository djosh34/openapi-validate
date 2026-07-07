package testgenerator

import "gopkg.in/yaml.v3"

type Hash [32]byte

type Hasher interface {
	GenerateHash() (Hash, error)
}

//type AllOfMerger interface {
//	MergeAllOf(domain Domain) Domain
//}

type Domain interface {
	Hasher
	//AllOfMerger
}

type DomainContext struct {
	// Each Domain that is created, must be added here
	domainStore map[Hash]Domain
	// Exists only for testing, to 'mock'/'inject' wanted parse outputs
	parse func(node yaml.Node) (*Hash, error)
}

func (dc *DomainContext) Parse(node yaml.Node) (*Hash, error) {
	if dc.parse != nil {
		return dc.parse(node)
	}

	return dc.ParseDefault(node)
}

func (dc *DomainContext) ParseDefault(node yaml.Node) (*Hash, error) {

	return nil, nil
}
