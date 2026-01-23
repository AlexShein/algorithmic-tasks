package vigenerecipher

import "testing"

type operation int

const (
	Encode operation = iota
	Decode
)

func TestVigenèreCipher(t *testing.T) {
	tests := []struct {
		name           string
		operation      operation
		input          string
		expected       string
		cipherKey      string
		cipherAplhabet string
	}{
		{"Basic case encode", Encode, "codewars", "rovwsoiv", "password", "abcdefghijklmnopqrstuvwxyz"},
		{"Basic case decode", Decode, "rovwsoiv", "codewars", "password", "abcdefghijklmnopqrstuvwxyz"},
		{"Second basic case encode", Encode, "waffles", "laxxhsj", "password", "abcdefghijklmnopqrstuvwxyz"},
		{"Second basic case decode", Decode, "laxxhsj", "waffles", "password", "abcdefghijklmnopqrstuvwxyz"},
		{"Input is out of alphabet: encode", Encode, "CODEWARS", "CODEWARS", "password", "abcdefghijklmnopqrstuvwxyz"},
		{"Input is out of alphabet: decode", Decode, "CODEWARS", "CODEWARS", "password", "abcdefghijklmnopqrstuvwxyz"},
		{"katakata: encode", Encode, "カタカナ", "タモタワ", "カタカナ", "アイウエオァィゥェォカキクケコサシスセソタチツッテトナニヌネノハヒフヘホマミムメモヤャユュヨョラリルレロワヲンー"},
		{"katakata: decode", Decode, "タモタワ", "カタカナ", "カタカナ", "アイウエオァィゥェォカキクケコサシスセソタチツッテトナニヌネノハヒフヘホマミムメモヤャユュヨョラリルレロワヲンー"},
	}

	for _, test_case := range tests {
		t.Run(test_case.name, func(t *testing.T) {

			cipher := VigenèreCipher{
				Key:   test_case.cipherKey,
				Alpha: test_case.cipherAplhabet,
			}
			var result string
			switch test_case.operation {
			case Decode:
				result = cipher.Decode(test_case.input)
			case Encode:
				result = cipher.Encode(test_case.input)
			default:
				t.Errorf("Invalid operation in the test case")
			}
			if result != test_case.expected {
				t.Errorf("Test %s failed: input %s, expected %s, got %s instead.", test_case.name, test_case.input, test_case.expected, result)
			}
		},
		)
	}

}
