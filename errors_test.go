package strenco

import (
	"testing"

	"github.com/julyskies/strenco/utilities"
)

var INVALID_FORMAT = utilities.WRAP + "isinvalid%"

const SHORT_STRING = "$Ij"

func TestDecodingErrors(t *testing.T) {
	result, decodingError := Decode("")
	if result != "" {
		t.Error("Expected to return an empty string when decoding an empty string!")
	}
	if !(decodingError != nil &&
		decodingError.Error() == utilities.ERRORS.ProvidedStringIsEmpty) {
		t.Errorf(
			`Expected to return an error "%s" if provided string is empty!`,
			utilities.ERRORS.ProvidedStringIsEmpty,
		)
	}

	result, decodingError = Decode(SHORT_STRING)
	if result != "" {
		t.Error("Expected to return an empty string when decoding a string that is too short!")
	}
	if !(decodingError != nil &&
		decodingError.Error() == utilities.ERRORS.InvalidStringFormat) {
		t.Errorf(
			`Expected to return an error "%s" if provided string is too short!`,
			utilities.ERRORS.InvalidStringFormat,
		)
	}

	result, decodingError = Decode(INVALID_FORMAT)
	if result != "" {
		t.Error("Expected to return an empty string when decoding a string with invalid format!")
	}
	if !(decodingError != nil &&
		decodingError.Error() == utilities.ERRORS.InvalidStringFormat) {
		t.Errorf(
			`Expected to return an error "%s" if provided string has invalid format!`,
			utilities.ERRORS.InvalidStringFormat,
		)
	}
}

func TestEncodingErrors(t *testing.T) {
	result, encodingError := Encode("")
	if result != "" {
		t.Error("Expected to return an empty string when encoding an empty string!")
	}
	if !(encodingError != nil &&
		encodingError.Error() == utilities.ERRORS.ProvidedStringIsEmpty) {
		t.Errorf(
			`Expected to return an error "%s" if provided string is empty!`,
			utilities.ERRORS.ProvidedStringIsEmpty,
		)
	}
}
