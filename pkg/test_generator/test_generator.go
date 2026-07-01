package testgenerator

import "encoding/json"

type Caseable interface {
	ValidCases() []Case
	InvalidCases() []Case
}

type Case struct {
	Name  string
	Value json.RawMessage
}
