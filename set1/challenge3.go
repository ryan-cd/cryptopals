package set1

import (
	"bytes"
	"log"
)

// SingleByteXORCipher - Challenge 3
func SingleByteXORCipher(hexString string) (key string, ascii string) {
	currentScore, maxScore := 0, 0
	maxScoreKey := byte(0)
	var fullKey string

	for i := byte(0); i < 255; i++ {
		currentScore = getScore(Byte2HexPair(i), hexString)
		if currentScore > maxScore {
			maxScore, maxScoreKey = currentScore, i
		}
	}
	fullKey = padHexPair(Byte2HexPair(maxScoreKey), len(hexString))

	return Byte2HexPair(maxScoreKey), string(HexString2ByteSplice(FixedXOR(hexString, fullKey)))
}

// Duplicates the hexPair enough times to make it have specified length
func padHexPair(hexPair string, length int) (hexString string) {
	var buffer bytes.Buffer
	if len(hexPair) != 2 || length%2 != 0 {
		log.Fatal("Invalid operators for padHexPair")
	}
	for len(buffer.String()) < length {
		buffer.WriteString(hexPair)
	}
	return buffer.String()
}

func getScore(hexPairKey string, hexString string) (score int) {
	var countArray [256]byte
	fullKey := padHexPair(hexPairKey, len(hexString))
	score = 0
	commonLetters := []byte{
		'e', 't', 'a', 'o', 'i', 'n'}
	uncommonLetters := []byte{
		'v', 'k', 'j', 'x', 'q', 'z'}

	byteSlice := HexString2ByteSplice(FixedXOR(hexString, fullKey))

	for i := 0; i < len(byteSlice); i++ {
		countArray[byteSlice[i]]++
	}

	skipMaxIndexes := []int{-1, -1, -1, -1, -1, -1}
	skipMinIndexes := []int{-1, -1, -1, -1, -1, -1}

	for i := 0; i < len(commonLetters); i++ {
		minIndex, maxIndex := getExtremeIndexes(countArray, skipMinIndexes, skipMaxIndexes)
		for j := 0; j < len(commonLetters); j++ {
			if maxIndex == commonLetters[j] {
				score++
			}
			if minIndex == uncommonLetters[j] {
				score++
			}
		}
		skipMaxIndexes[i] = int(maxIndex)
		skipMinIndexes[i] = int(minIndex)
	}

	return score
}

func getExtremeIndexes(countArray [256]byte, skipMinIndexes []int, skipMaxIndexes []int) (minIndex byte, maxIndex byte) {
	skipThisIndex := false
	maxValue, minValue, maxIndex, minIndex := 0, 255, byte(0), byte(0)

	for i := 0; i < len(countArray); i++ {
		if countArray[i] == 0 {
			continue
		}
		skipThisIndex = false
		for j := 0; j < len(skipMaxIndexes); j++ {
			if i == skipMinIndexes[j] || i == skipMaxIndexes[j] {
				skipThisIndex = true
			}
		}
		if skipThisIndex {
			continue
		}
		if int(countArray[i]) > maxValue {
			maxValue, maxIndex = int(countArray[i]), byte(i)
		}
		if int(countArray[i]) < minValue {
			minValue, minIndex = int(countArray[i]), byte(i)
		}
	}
	return minIndex, maxIndex
}
