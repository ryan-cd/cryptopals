package set1

import (
	"io/ioutil"
	"log"
	"testing"
)

func TestDetectSingleCharacterXOR(t *testing.T) {
	file, err := ioutil.ReadFile("4.txt")
	if err != nil {
		log.Fatal(err)
	}
	var key string
	var ascii string
	key, ascii = DetectSingleCharacterXOR(string(file))
	if key != "35" || ascii != "Now that the party is jumping\n" {
		t.Error("Incorrect result. Received ", key, " ", ascii)
	}
}
