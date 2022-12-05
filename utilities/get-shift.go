package utilities

func GetShift(length int) int {
	if length > ALPHABET_LENGTH {
		return length % ALPHABET_LENGTH
	}
	return ALPHABET_LENGTH % length
}
