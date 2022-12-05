package utilities

import "errors"

func baseValidation(value string) error {
	if value == "" {
		return errors.New(ERRORS.ProvidedStringIsEmpty)
	}
	return nil
}

func DecodeValidation(value string) (string, error) {
	baseError := baseValidation(value)
	if baseError != nil {
		return "", baseError
	}

	runes := []rune(value)

	ORIGINAL_STRING_LENGTH := len(runes)
	if ORIGINAL_STRING_LENGTH < 4 {
		return "", errors.New(ERRORS.InvalidStringFormat)
	}

	endSymbols := string(runes[:2])
	startSymbols := string(runes[len(runes)-2:])
	if endSymbols != WRAP || startSymbols != WRAP {
		return "", errors.New(ERRORS.InvalidStringFormat)
	}
	return string(runes[2 : len(runes)-2]), nil
}

func EncodeValidation(value string) error {
	baseError := baseValidation(value)
	if baseError != nil {
		return baseError
	}
	return nil
}
