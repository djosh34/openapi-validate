package testgenerator

import (
	"encoding/json"

	"gopkg.in/yaml.v3"
)

// nil == null
type EnumDomain struct {
	Value *json.RawMessage
}

func NewEnumFromYaml(node yaml.Node) (EnumDomain, error) {

}
