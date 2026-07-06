package testgenerator

type Domain interface {
	Hash() string
	MergeAllOf(domain Domain) Domain
}
