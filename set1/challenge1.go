package set1

import (
	"bytes"
	"log"
)

//Hex2Base64 - Challenge 1
func Hex2Base64(hexString string) (output string) {
	var base64Buffer bytes.Buffer
	byteSplice := HexString2ByteSplice(hexString)
	paddingCount := 0

	// Pad the byteSplice to ensure divisible by 3
	for true {
		if float32(len(byteSplice)/3) == float32(len(byteSplice))/3 {
			break
		} else {
			byteSplice = append(byteSplice, 0)
			paddingCount++
		}
	}

	var quadruple [4]byte // Store the 4 values that the 3 byteSplice convert to
	for i := 0; i < len(byteSplice); i += 3 {
		quadruple[0] = byteSplice[i] >> 2
		quadruple[1] = ((byteSplice[i] & 3) << 4) | (byteSplice[i+1] >> 4)
		if paddingCount >= 2 {
			quadruple[2] = 64 // '=' character
		} else {
			quadruple[2] = ((byteSplice[i+1] & 15) << 2) | (byteSplice[i+2] >> 6)
		}
		if paddingCount >= 1 {
			quadruple[3] = 64 // '=' character
		} else {
			quadruple[3] = byteSplice[i+2] & 63
		}

		for _, item := range quadruple {
			base64Buffer.WriteString(index2Base64(item))
		}
	}

	return base64Buffer.String()
}

func HexString2ByteSplice(hexString string) (byteSplice []byte) {
	b := byte('0')
	byteSplice = []byte("")

	for i := 0; i < len(hexString); i += 2 {
		b = 16 * hexDigitValue(hexString[i])
		b += hexDigitValue(hexString[i+1])
		byteSplice = append(byteSplice, b)
	}
	return byteSplice
}

func hexDigitValue(hexDigit byte) (output byte) {
	if 48 <= hexDigit && hexDigit <= 57 {
		return hexDigit - 48
	} else if 97 <= hexDigit && hexDigit <= 102 {
		return hexDigit - 87
	} else {
		log.Fatal("Unable to parse hex string")
		return 0
	}
}

func index2Base64(index byte) (base64String string) {
	if index < 0 || index > 64 {
		log.Fatal("Invalid index to convert to base 64")
		return ""
	}
	dictionary := map[int]string{
		0: "A", 1: "B", 2: "C", 3: "D", 4: "E", 5: "F",
		6: "G", 7: "H", 8: "I", 9: "J", 10: "K", 11: "L",
		12: "M", 13: "N", 14: "O", 15: "P", 16: "Q", 17: "R",
		18: "S", 19: "T", 20: "U", 21: "V", 22: "W", 23: "X",
		24: "Y", 25: "Z", 26: "a", 27: "b", 28: "c", 29: "d",
		30: "e", 31: "f", 32: "g", 33: "h", 34: "i", 35: "j",
		36: "k", 37: "l", 38: "m", 39: "n", 40: "o", 41: "p",
		42: "q", 43: "r", 44: "s", 45: "t", 46: "u", 47: "v",
		48: "w", 49: "x", 50: "y", 51: "z", 52: "0", 53: "1",
		54: "2", 55: "3", 56: "4", 57: "5", 58: "6", 59: "7",
		60: "8", 61: "9", 62: "+", 63: "/", 64: "="}
	return dictionary[int(index)]
}
