package strenco

import (
	"fmt"
	"math"
	"strings"

	"strenco/utilities"
)

func Decode(value string) (string, error) {
	adjustedString, validationError := utilities.DecodeValidation(value)
	if validationError != nil {
		return "", validationError
	}

	runes := []rune(adjustedString)

	STRING_LENGTH := len(runes)
	SHIFT := utilities.GetShift(STRING_LENGTH)

	var result []string
	for index, symbol := range runes {
		alphabetIndex := strings.Index(utilities.ALPHABET, string(symbol))
		if alphabetIndex < 0 {
			result = append(result, string(symbol))
			continue
		}
		resultIndex := alphabetIndex - SHIFT - index
		if resultIndex < 0 {
			abs := int(math.Abs(float64(resultIndex)))
			if abs > utilities.ALPHABET_LENGTH {
				resultIndex = utilities.ALPHABET_LENGTH - (abs % utilities.ALPHABET_LENGTH)
				if resultIndex == utilities.ALPHABET_LENGTH {
					resultIndex = 0
				}
			} else {
				resultIndex = utilities.ALPHABET_LENGTH - abs
			}
		}
		result = append(result, string(utilities.ALPHABET[resultIndex]))
	}
	return strings.Join(result, ""), nil
}

func Encode(value string) (string, error) {
	validationError := utilities.EncodeValidation(value)
	if validationError != nil {
		return "", validationError
	}

	runes := []rune(value)

	STRING_LENGTH := len(runes)
	SHIFT := utilities.GetShift(STRING_LENGTH)

	var result []string
	for index, symbol := range runes {
		alphabetIndex := strings.Index(utilities.ALPHABET, string(symbol))
		if alphabetIndex < 0 {
			result = append(result, string(symbol))
			continue
		}
		resultIndex := (alphabetIndex + index + SHIFT) % utilities.ALPHABET_LENGTH
		result = append(result, string(utilities.ALPHABET[resultIndex]))
	}

	return fmt.Sprintf(
		"%s%s%s",
		utilities.WRAP,
		strings.Join(result, ""),
		utilities.WRAP,
	), nil
}
