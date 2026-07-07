package domain

import "encoding/json"

// nil == null
type EnumDomain struct {
	*json.RawMessage
}

func NewEnumFromJSON(node *json.RawMessage) (EnumDomain, error) {
	return EnumDomain{node}, nil
}
