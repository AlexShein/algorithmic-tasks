package permutations

import (
	"reflect"
	"testing"
)

func TestPermutations(t *testing.T) {

	tests := []struct {
		input    string
		expected []string
	}{
		{
			"a",
			[]string{"a"},
		},
		{
			"ab",
			[]string{"ab", "ba"},
		},
		{
			"abc",
			[]string{"abc", "acb", "bac", "bca", "cab", "cba"},
		},
	}

	for _, test_case := range tests {
		res := Permutations(test_case.input)
		if !areStringSlicesEqual(res, test_case.expected) {
			t.Errorf("Test %s failed. Got %v, expected %v", test_case.input, res, test_case.expected)
		}
	}

}

// areStringSlicesEqual returns true if both input slices have same elements, order is ignored.
func areStringSlicesEqual(first []string, second []string) bool {
	firstMap := map[string]bool{}
	for _, str := range first {
		firstMap[str] = true
	}
	secondMap := map[string]bool{}
	for _, str := range second {
		secondMap[str] = true
	}
	return reflect.DeepEqual(firstMap, secondMap)
}
