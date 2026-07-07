package domain

import "decode_and_validate_generator/pkg/test_generator/hashables"

func toHashableNumberPtr(number *Number) *hashables.Number {
	if number == nil {
		return nil
	}

	hashableNumber := hashables.Number(*number)

	return &hashableNumber
}
