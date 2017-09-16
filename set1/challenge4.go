package set1

import (
	"strings"
)

// DetectSingleCharacterXOR - Challenge 4
func DetectSingleCharacterXOR(file string) (key string, ascii string) {
	lines := strings.Split(file, "\n")
	invalidLine := false

	for _, item := range lines {
		invalidLine = false
		key, ascii = SingleByteXORCipher(item)

		for _, character := range ascii {
			// Check if the string contained a character outside of the alphabet, a space, or a newline
			if (character < 65 || character > 122) && (character != 10 && character != 32) {
				invalidLine = true
				break
			}
		}
		if !invalidLine {
			break
		}
	}

	return key, ascii
}
