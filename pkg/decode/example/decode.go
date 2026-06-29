// Package example contains generated decode examples.
package example

import (
	"decode_and_validate_generator/pkg/peekjson"
	"encoding/json"
)

type Decoder interface {
	Decode(decoder *peekjson.Decoder) error
	json.Marshaler
}
