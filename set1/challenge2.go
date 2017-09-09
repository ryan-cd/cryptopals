package set1

import (
	"bytes"
	"log"
)

//FixedXOR - Challenge 2
func FixedXOR(buffer1 string, buffer2 string) (output string) {
	byteSplice1 := HexString2ByteSplice(buffer1)
	byteSplice2 := HexString2ByteSplice(buffer2)
	outputByteSlice := []byte("")
	var outputBuffer bytes.Buffer

	if len(byteSplice1) != len(byteSplice2) {
		log.Fatal("Input string lengths do not match")
		return ""
	}

	for i := 0; i < len(byteSplice1); i++ {
		outputByteSlice = append(outputByteSlice, xor(byteSplice1[i], byteSplice2[i]))
	}

	for i := 0; i < len(outputByteSlice); i++ {
		outputBuffer.WriteString(byte2HexPair(outputByteSlice[i]))
	}

	return outputBuffer.String()
}

func xor(operand1 byte, operand2 byte) (output byte) {
	return ^(operand1 & operand2) & (operand1 | operand2)
}

func byte2HexPair(input byte) (output string) {
	firstLetter := "0"
	secondLetter := "0"
	inputDiv16 := byte(input / 16)
	remainder := input - inputDiv16*16
	var buffer bytes.Buffer
	firstLetter = val2Hex(inputDiv16)
	secondLetter = val2Hex(remainder)
	buffer.WriteString(firstLetter)
	buffer.WriteString(secondLetter)

	return buffer.String()
}

func val2Hex(input byte) (output string) {
	if 0 <= input && input <= 9 {
		return string(input + 48)
	} else if 10 <= input && input <= 15 {
		return string(input + 87)
	} else {
		log.Fatal("Invalid val2hex conversion")
		return "0"
	}
}
