package example

import (
	"encoding/json"

	"decode_and_validate_generator/pkg/peekjson"
)

// Decoder decodes a JSON value and supports standard JSON marshaling.
type Decoder interface {
	Decode(decoder *peekjson.Decoder) error
	json.Marshaler
}
