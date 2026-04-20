package lastdigit

import "testing"

func TestLastDigit(t *testing.T) {
	testCases := []struct {
		n1             string
		n2             string
		expectedResult int
	}{
		{n1: "4", n2: "1", expectedResult: 4},
		{n1: "4", n2: "2", expectedResult: 6},
		{n1: "9", n2: "7", expectedResult: 9},
		{n1: "10", n2: "10000000000", expectedResult: 0},
		{n1: "13", n2: "1", expectedResult: 3},
		{n1: "13", n2: "2", expectedResult: 9},
		{n1: "13", n2: "3", expectedResult: 7},
		{n1: "13", n2: "4", expectedResult: 1},
		{n1: "13", n2: "5", expectedResult: 3},
		{n1: "3", n2: "51", expectedResult: 7},
		{n1: "2", n2: "807", expectedResult: 8},
		{n1: "2", n2: "108", expectedResult: 6},
		{n1: "1606938044258990275541962092341162602522202993782792835301376", n2: "2037035976334486086268445688409378161051468393665936250636140449354381299763336706183397376", expectedResult: 6},
		{n1: "3715290469715693021198967285016729344580685479654510946723", n2: "68819615221552997273737174557165657483427362207517952651", expectedResult: 7},
	}

	for _, testCase := range testCases {

		if res := LastDigit(testCase.n1, testCase.n2); res != testCase.expectedResult {
			t.Errorf(
				"Test failed.\nN1:%v\nN2:%v\nExpected %d, got %d",
				testCase.n1, testCase.n2,
				testCase.expectedResult, res)
		}
	}
}
