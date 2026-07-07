package types

//	type AllOfMerger interface {
//		MergeAllOf(domain Domain) Domain
//	}
type Domain interface {
	//Hasher
	//AllOfMerger
}

type Hash [32]byte

type Hasher interface {
	GenerateHash() (Hash, error)
}
