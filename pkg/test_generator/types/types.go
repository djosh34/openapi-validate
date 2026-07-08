package types

type AllOfMerger interface {
	AllOfMerge(domain Domain) (Domain, error)
}

type Domain interface {
	Hasher
	AllOfMerger
}

type Hash [32]byte

type Hasher interface {
	GenerateHash() (Hash, error)
}
