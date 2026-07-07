package domain

import (
	"bytes"
	"decode_and_validate_generator/pkg/test_generator/types"
	"encoding/json"
	"errors"
)

func parseEnums(jsonKV JSONKV) ([]types.Enum, bool, error) {
	enumRaw, ok := jsonKV["enum"]
	if !ok {
		return nil, false, nil
	}

	var enumValues []json.RawMessage
	if err := json.Unmarshal(enumRaw, &enumValues); err != nil {
		return nil, true, errors.New("enum must be array")
	}
	if enumValues == nil {
		return nil, true, errors.New("enum cannot be null")
	}
	if len(enumValues) == 0 {
		return nil, true, errors.New("enum cannot be empty")
	}

	enums := make([]types.Enum, 0, len(enumValues))
	for _, enumValue := range enumValues {
		enums = append(enums, types.Enum(enumValue))
	}

	return enums, true, nil
}

func mergeEnums(left []types.Enum, right []types.Enum) ([]types.Enum, error) {
	if left == nil && right == nil {
		return nil, nil
	}
	if err := validateEnums(left); err != nil {
		return nil, err
	}
	if err := validateEnums(right); err != nil {
		return nil, err
	}
	if left == nil {
		return append([]types.Enum(nil), right...), nil
	}
	if right == nil {
		return append([]types.Enum(nil), left...), nil
	}

	merged := make([]types.Enum, 0, len(left))
	for _, leftEnum := range left {
		for _, rightEnum := range right {
			if bytes.Equal(leftEnum, rightEnum) {
				merged = append(merged, leftEnum)
				break
			}
		}
	}
	if len(merged) == 0 {
		return nil, errors.New("enum intersection is empty")
	}
	return merged, nil
}

func validateEnums(enums []types.Enum) error {
	for _, enumValue := range enums {
		if enumValue == nil {
			return errors.New("enum raw value cannot be nil")
		}
		if !json.Valid(enumValue) {
			return errors.New("enum raw value must be valid JSON")
		}
	}
	return nil
}

type EnumDomain struct {
	RawMessage *json.RawMessage
}

var _ types.Domain = new(EnumDomain)

func (e *EnumDomain) AllOfMerge(domain types.Domain) (types.Domain, error) {
	if e == nil {
		return nil, errors.New("enum domain cannot be nil")
	}
	if e.RawMessage == nil {
		return nil, errors.New("enum raw value cannot be nil")
	}
	if !json.Valid(*e.RawMessage) {
		return nil, errors.New("enum raw value must be valid JSON")
	}
	otherEnum, ok := domain.(*EnumDomain)
	if !ok || otherEnum == nil {
		return nil, errors.New("domain is not EnumDomain")
	}
	if otherEnum.RawMessage == nil {
		return nil, errors.New("enum raw value cannot be nil")
	}
	if !json.Valid(*otherEnum.RawMessage) {
		return nil, errors.New("enum raw value must be valid JSON")
	}
	if !bytes.Equal(*e.RawMessage, *otherEnum.RawMessage) {
		return nil, errors.New("enum values do not intersect")
	}
	raw := append(json.RawMessage(nil), (*e.RawMessage)...)
	return &EnumDomain{RawMessage: &raw}, nil
}

func (e *EnumDomain) ToHasher() (types.Hasher, error) {
	if e == nil {
		return nil, errors.New("enum domain cannot be nil")
	}
	if e.RawMessage == nil {
		return nil, errors.New("enum raw value cannot be nil")
	}
	if !json.Valid(*e.RawMessage) {
		return nil, errors.New("enum raw value must be valid JSON")
	}
	return types.Enum(*e.RawMessage), nil
}
