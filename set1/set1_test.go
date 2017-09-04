package set1

import "testing"

func TestHex2Base64(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"4d", "TQ=="},
		{"4d61", "TWE="},
		{"4d616e", "TWFu"},
		{"49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d",
			"SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"},
	}
	for _, testCase := range cases {
		got := Hex2Base64(testCase.in)
		if got != testCase.want {
			t.Errorf("Hex2Base64(%q) == %q, want %q", testCase.in, got, testCase.want)
		}
	}
}
