package example

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	NotAnObjectError              = errors.New("not an object")
	AdditionalPropertyError       = errors.New("additional property")
	MissingRequiredPropertyError  = errors.New("missing required property")
	NullForNotNullableStringError = errors.New("null for not nullable string")
	NonStringForStringSchemaError = errors.New("non-string for string schema")
)

var jsonNull = []byte("null")

type ObjectKeysAdditionalPropertiesFalse struct {
	OptionalNotNullableString *ObjectKeysAdditionalPropertiesFalseOptionalNotNullableString `json:"optionalNotNullableString,omitzero"`
	OptionalNullableString    *ObjectKeysAdditionalPropertiesFalseOptionalNullableString    `json:"optionalNullableString,omitzero"`
	RequiredNotNullableString ObjectKeysAdditionalPropertiesFalseRequiredNotNullableString  `json:"requiredNotNullableString"`
	RequiredNullableString    ObjectKeysAdditionalPropertiesFalseRequiredNullableString     `json:"requiredNullableString"`
}

var _ json.Unmarshaler = (*ObjectKeysAdditionalPropertiesFalse)(nil)

func (o *ObjectKeysAdditionalPropertiesFalse) UnmarshalJSON(data []byte) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()

	tok, err := d.Token()
	if err != nil {
		return err
	}
	if tok != json.Delim('{') {
		return NotAnObjectError
	}
	var hasRequiredNotNullableString bool
	var hasRequiredNullableString bool

	for d.More() {
		nameTok, nameErr := d.Token()
		if nameErr != nil {
			return nameErr
		}

		name, ok := nameTok.(string)
		if !ok {
			return NotAnObjectError
		}

		var value json.RawMessage
		err = d.Decode(&value)
		if err != nil {
			return err
		}

		switch name {
		case "optionalNotNullableString":

			var optionalNotNullableString ObjectKeysAdditionalPropertiesFalseOptionalNotNullableString
			err = json.Unmarshal(value, &optionalNotNullableString)
			if err != nil {
				return err
			}
			o.OptionalNotNullableString = &optionalNotNullableString
		case "optionalNullableString":

			var optionalNullableString ObjectKeysAdditionalPropertiesFalseOptionalNullableString
			err = json.Unmarshal(value, &optionalNullableString)
			if err != nil {
				return err
			}
			o.OptionalNullableString = &optionalNullableString
		case "requiredNotNullableString":
			hasRequiredNotNullableString = true

			var requiredNotNullableString ObjectKeysAdditionalPropertiesFalseRequiredNotNullableString
			err = json.Unmarshal(value, &requiredNotNullableString)
			if err != nil {
				return err
			}
			o.RequiredNotNullableString = requiredNotNullableString
		case "requiredNullableString":
			hasRequiredNullableString = true

			var requiredNullableString ObjectKeysAdditionalPropertiesFalseRequiredNullableString
			err = json.Unmarshal(value, &requiredNullableString)
			if err != nil {
				return err
			}
			o.RequiredNullableString = requiredNullableString
		default:
			return fmt.Errorf("%w: %s", AdditionalPropertyError, name)
		}
	}
	if !hasRequiredNotNullableString {
		return fmt.Errorf("%w: %s", MissingRequiredPropertyError, "requiredNotNullableString")
	}
	if !hasRequiredNullableString {
		return fmt.Errorf("%w: %s", MissingRequiredPropertyError, "requiredNullableString")
	}

	return nil
}

type ObjectKeysAdditionalPropertiesFalseOptionalNotNullableString string

var _ json.Unmarshaler = new(ObjectKeysAdditionalPropertiesFalseOptionalNotNullableString)

func (s *ObjectKeysAdditionalPropertiesFalseOptionalNotNullableString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		return NullForNotNullableStringError
	}

	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return NonStringForStringSchemaError
	}
	*s = ObjectKeysAdditionalPropertiesFalseOptionalNotNullableString(value)
	return nil
}

type ObjectKeysAdditionalPropertiesFalseOptionalNullableString struct {
	Value *string
}

var _ json.Unmarshaler = new(ObjectKeysAdditionalPropertiesFalseOptionalNullableString)

func (s *ObjectKeysAdditionalPropertiesFalseOptionalNullableString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		s.Value = nil
		return nil
	}

	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return NonStringForStringSchemaError
	}
	s.Value = new(value)
	return nil
}

type ObjectKeysAdditionalPropertiesFalseRequiredNotNullableString string

var _ json.Unmarshaler = new(ObjectKeysAdditionalPropertiesFalseRequiredNotNullableString)

func (s *ObjectKeysAdditionalPropertiesFalseRequiredNotNullableString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		return NullForNotNullableStringError
	}

	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return NonStringForStringSchemaError
	}
	*s = ObjectKeysAdditionalPropertiesFalseRequiredNotNullableString(value)
	return nil
}

type ObjectKeysAdditionalPropertiesFalseRequiredNullableString struct {
	Value *string
}

var _ json.Unmarshaler = new(ObjectKeysAdditionalPropertiesFalseRequiredNullableString)

func (s *ObjectKeysAdditionalPropertiesFalseRequiredNullableString) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, jsonNull) {
		s.Value = nil
		return nil
	}

	var value string
	err := json.Unmarshal(data, &value)
	if err != nil {
		return NonStringForStringSchemaError
	}
	s.Value = new(value)
	return nil
}
