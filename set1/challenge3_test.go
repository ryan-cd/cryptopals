package set1

import "testing"

func TestSingleByteXORCipher(t *testing.T) {
	cases := []struct {
		input, wantKey, wantASCII string
	}{
		{
			"1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736",
			"58",
			"Cooking MC's like a pound of bacon",
		},
	}
	for _, testCase := range cases {
		gotKey, gotASCII := SingleByteXORCipher(testCase.input)
		if gotKey != testCase.wantKey || gotASCII != testCase.wantASCII {
			t.Errorf("SingleByteXORCipher(%q) == %q, %q, want %q, %q\n", testCase.input, gotKey, gotASCII, testCase.wantKey, testCase.wantASCII)
		}
	}
}
