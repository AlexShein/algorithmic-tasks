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
		{
			"aab",
			[]string{"aab", "aba", "baa"},
		},
		{
			"abcd",
			[]string{"abcd", "abdc", "acbd", "acdb", "adbc", "adcb",
				"bacd", "badc", "bcad", "bcda", "bdac", "bdca",
				"cabd", "cadb", "cbad", "cbda", "cdab", "cdba",
				"dabc", "dacb", "dbac", "dbca", "dcab", "dcba"},
		},
		{
			"baab",
			[]string{"aabb", "abab", "abba", "baab", "baba", "bbaa"},
		},
	}

	for _, test_case := range tests {
		res := PermutationsRecursive(test_case.input)
		if !areStringSlicesEqual(res, test_case.expected) {
			t.Errorf("PermutationsRecursive test %s failed. Got %v, expected %v", test_case.input, res, test_case.expected)
		}
	}

	for _, test_case := range tests {
		res := Permutations(test_case.input)
		if !areStringSlicesEqual(res, test_case.expected) {
			t.Errorf("Permutations test %s failed. Got %v, expected %v", test_case.input, res, test_case.expected)
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

func TestNextPermutation(t *testing.T) {

	tests := []struct {
		input    string
		expected bool
	}{
		{
			"ab",
			true,
		},
		{
			"ba",
			false,
		},
		{
			"aabac",
			true,
		},
		{
			"cba",
			false,
		},
	}

	for _, test_case := range tests {
		input := []rune(test_case.input)
		res := nextPermutation(input)

		if res != test_case.expected {
			t.Errorf("Test %s failed. Got %v, expected %v", test_case.input, res, test_case.expected)
		}

	}
}
