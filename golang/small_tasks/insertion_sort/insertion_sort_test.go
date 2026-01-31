package insertionsort

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {

	tests := []struct {
		input    []int
		expected []int
	}{
		{
			[]int{3, 2, 1},
			[]int{1, 2, 3},
		},
		{
			[]int{6, 5, 3, 1, 8, 7, 2, 4},
			[]int{1, 2, 3, 4, 5, 6, 7, 8},
		},
	}

	for _, test_case := range tests {
		res := InsertionSort(test_case.input)
		if !reflect.DeepEqual(res, test_case.expected) {
			t.Errorf("InsertionSort test %v failed. Got %v, expected %v", test_case.input, res, test_case.expected)
		}
	}
}
