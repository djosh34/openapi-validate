package testgenerator

import "gopkg.in/yaml.v3"

type Hash [32]byte

type Domain interface {
	Hash() Hash
	MergeAllOf(domain Domain) Domain
	Parse(node yaml.Node) error
}

func Parse(node yaml.Node) error {

	return nil
}
