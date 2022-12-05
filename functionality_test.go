package strenco

import (
	"testing"

	"github.com/julyskies/strenco/utilities"
)

const LARGE_STRING = `
This is a pretty large string, it is used in the tests.

The length of this string is bigger than the length of the alphabet used for encoding!

123 321 123
345 543 321
098 980 890
000 000 000

And some special symbols: $%^:[]@()_ !,.=+

Encoding should ignore all of the symbols that are not in the alphabet:
\
Ê
µ

That's it, should be fine...
`

const SMALL_STRING = "This is a small string, it is used in the tests..."

func TestFunctionalityLargeString(t *testing.T) {
	encoded, encodingError := Encode(LARGE_STRING)
	if encoded == "" {
		t.Error("Expected to return an encoded string when encoding a valid short string!")
	}
	if encodingError != nil {
		t.Error("Expected to return a nil error when encoding a valid short string!")
	}

	runes := []rune(encoded)
	endSymbols := string(runes[:2])
	startSymbols := string(runes[len(runes)-2:])
	if endSymbols != utilities.WRAP || startSymbols != utilities.WRAP {
		t.Error("Expected returned encoded string to be wrapped!")
	}

	decoded, decodingError := Decode(encoded)
	if decoded == "" {
		t.Error("Expected to return a decoded string when decoding a valid short string!")
	}
	if decodingError != nil {
		t.Error("Expected to return a nil error when decoding a valid short string!")
	}

	if decoded != LARGE_STRING {
		t.Error("Expected decoded string to be equal to the original!")
	}
}

func TestFunctionalityShortString(t *testing.T) {
	encoded, encodingError := Encode(SHORT_STRING)
	if encoded == "" {
		t.Error("Expected to return an encoded string when encoding a valid short string!")
	}
	if encodingError != nil {
		t.Error("Expected to return a nil error when encoding a valid short string!")
	}

	decoded, decodingError := Decode(encoded)
	if decoded == "" {
		t.Error("Expected to return a decoded string when decoding a valid short string!")
	}
	if decodingError != nil {
		t.Error("Expected to return a nil error when decoding a valid short string!")
	}

	if decoded != SHORT_STRING {
		t.Error("Expected decoded string to be equal to the original!")
	}
}
