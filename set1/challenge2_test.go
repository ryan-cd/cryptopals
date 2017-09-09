package set1

import (
	"testing"
)

func TestFixedXOR(t *testing.T) {
	cases := []struct {
		input1, input2, want string
	}{
		{
			"1c0111001f010100061a024b53535009181c",
			"686974207468652062756c6c277320657965",
			"746865206b696420646f6e277420706c6179",
		},
	}
	for _, testCase := range cases {
		got := FixedXOR(testCase.input1, testCase.input2)
		if got != testCase.want {
			t.Errorf("Hex2Base64(%q, %q) == \n%q, want \n%q", testCase.input1, testCase.input2, got, testCase.want)
		}
	}
}
