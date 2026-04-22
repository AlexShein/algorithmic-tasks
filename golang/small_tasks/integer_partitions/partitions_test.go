package integerpartitions

import "testing"

func TestPartitions(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{1, 1},
		{2, 2},
		{3, 3},
		{4, 5},
		{5, 7},
		{6, 11},
		{7, 15},
		{8, 22},
		{10, 42},
		{25, 1958},
		{100, 190569292},
	}

	for _, testCase := range testCases {

		if res := Partitions(testCase.input); res != testCase.expected {
			t.Errorf(
				"Test failed. Input: %d, expected %d, got %d",
				testCase.input,
				testCase.expected, res)
		}
	}
}

func TestPartitionCombinations(t *testing.T) {
	testCases := []struct {
		n        int
		m        int
		expected int
	}{
		{1, 6, 1},
		{2, 5, 2},
		{3, 4, 3},
		{4, 3, 4},
		{5, 2, 3},
		{6, 1, 1},
	}

	for _, testCase := range testCases {
		n := 10
		cache := make([][]int, n+1)
		for i := 0; i < n+1; i++ {
			cache[i] = make([]int, n+1)
			for j := 0; j < n+1; j++ {
				cache[i][j] = -1
			}
		}

		if res := partCombinations(testCase.n, testCase.m, cache); res != testCase.expected {
			t.Errorf(
				"Test failed. Input: n=%d, m=%d, expected %d, got %d",
				testCase.n, testCase.m,
				testCase.expected, res)
		}
	}
}
