package set1

import (
	"fmt"
	"log"
)

func Hex2Base64(hexString string) (output string) {
	b := byte('0')
	bytes := []byte("")
	for i := 0; i < len(hexString); i += 2 {
		b = 16 * HexDigitValue(hexString[i])
		b += HexDigitValue(hexString[i+1])
		bytes = append(bytes, b)
	}

	// Pad the bytes to ensure divisible by 3
	for true {
		if float32(len(bytes)/3) == float32(len(bytes))/3 {
			break
		} else {
			bytes = append(bytes, 0)
		}
	}

	fmt.Print("Bytes are: ")
	for _, element := range bytes {
		fmt.Print(element, ", ")
	}
	fmt.Println()

	return hexString
}

func HexDigitValue(hexDigit byte) (output byte) {
	if 48 <= hexDigit && hexDigit <= 57 {
		return hexDigit - 48
	} else if 97 <= hexDigit && hexDigit <= 102 {
		return hexDigit - 87
	} else {
		log.Fatal("Unable to parse hex string")
		return 0
	}
}
