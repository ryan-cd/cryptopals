package set1

import (
	"bytes"
)

//RepeatingKeyXOR - Challenge 5
func RepeatingKeyXOR(input string) (output string) {
	var buffer bytes.Buffer
	keyElement := byte(73)
	for i, elem := range input {
		if i%3 == 0 {
			keyElement = byte(73) // "I"
		} else if (i-1)%3 == 0 {
			keyElement = byte(67) // "C"

		} else if (i-2)%3 == 0 {
			keyElement = byte(69) // "E"
		}
		buffer.WriteString(Byte2HexPair(Xor(byte(elem), keyElement)))
	}
	return buffer.String()
}
