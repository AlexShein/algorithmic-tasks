package skyscrapersV2

import (
	"reflect"
	"testing"
)

type TestCase struct {
	name   string
	input  []int
	output [][]int
}

var solvePuzzleTestCases = []TestCase{
	{
		name: "Basic puzzle",
		input: []int{
			0, 0, 1, 2,
			0, 2, 0, 0,
			0, 3, 0, 0,
			0, 1, 0, 0},
		output: [][]int{
			{2, 1, 4, 3},
			{3, 4, 1, 2},
			{4, 2, 3, 1},
			{1, 3, 2, 4}},
	},
	{
		name: "Second puzzle with less clues",
		input: []int{
			0, 3, 0, 0,
			0, 1, 0, 0,
			0, 0, 1, 2,
			0, 2, 0, 0,
		},
		output: [][]int{
			{4, 2, 3, 1},
			{1, 3, 2, 4},
			{2, 1, 4, 3},
			{3, 4, 1, 2}},
	},
	{
		name: "All clues are provided",
		input: []int{
			2, 1, 3, 2,
			3, 1, 2, 3,
			3, 2, 2, 1,
			1, 2, 4, 2,
		},
		output: [][]int{
			{3, 4, 2, 1},
			{1, 2, 3, 4},
			{2, 1, 4, 3},
			{4, 3, 1, 2},
		},
	},
	{
		name: "All clues are provided 2",
		input: []int{
			2, 2, 1, 3,
			2, 2, 3, 1,
			1, 2, 2, 3,
			3, 2, 1, 3,
		},
		output: [][]int{
			{1, 3, 4, 2},
			{4, 2, 1, 3},
			{3, 4, 2, 1},
			{2, 1, 3, 4},
		},
	},
}

func TestSolvePuzzle(t *testing.T) {
	for _, testCase := range solvePuzzleTestCases {
		if funcResult := SolvePuzzle(testCase.input); !reflect.DeepEqual(funcResult, testCase.output) {
			t.Errorf("Test %s\nInput: %v, Expected %v, Got %v", testCase.name, testCase.input, testCase.output, funcResult)
		}
	}
}

func TestNewPuzzleSolver(t *testing.T) {
	testInitTestCase := []struct {
		name   string
		input  []int
		output [PUZZLE_SIZE * PUZZLE_SIZE]int
	}{
		{
			name: "Init a trivial values: ones",
			input: []int{
				0, 1, 0, 0,
				0, 1, 0, 0,
				0, 1, 0, 0,
				0, 1, 0, 0},
			output: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				0, 4, 0, 0,
				0, 0, 0, 4,
				4, 0, 0, 0,
				0, 0, 4, 0},
		},
		{
			name: "Init a trivial values: fours - columns",
			input: []int{
				4, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0},
			output: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 0, 0, 4,
				2, 0, 0, 3,
				3, 0, 0, 2,
				4, 0, 0, 1},
		},
		{
			name: "Init a trivial values: fours - rows",
			input: []int{
				0, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0,
				4, 0, 0, 0},
			output: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				4, 3, 2, 1,
				0, 0, 0, 0,
				0, 0, 0, 0,
				1, 2, 3, 4},
		},
	}

	for _, testCase := range testInitTestCase {
		if funcResult := NewPuzzleSolver(testCase.input); !reflect.DeepEqual(funcResult.solution, testCase.output) {
			t.Errorf("Test %s\nInput: %v, Expected %v, Got %v", testCase.name, testCase.input, testCase.output, funcResult.solution)
		}
	}

}

var initializeCluesTestCases = []struct {
	clues            []int
	expectedRowClues [4][2]int
	expectedColClues [4][2]int
}{
	{
		clues: []int{
			0, 0, 1, 2,
			0, 2, 0, 0,
			0, 3, 0, 0,
			0, 1, 0, 0},
		expectedRowClues: [4][2]int{
			{0, 0},
			{0, 2},
			{1, 0},
			{0, 0},
		},
		expectedColClues: [4][2]int{
			{0, 0},
			{0, 0},
			{1, 3},
			{2, 0},
		},
	},
}

func TestInitializeClues(t *testing.T) {
	for _, testCase := range initializeCluesTestCases {
		solver := NewPuzzleSolver(testCase.clues)
		if solver.rowClues != testCase.expectedRowClues {
			t.Errorf("Test for the initializeClues failed:\n Expected row clues %v,  got %v\n",
				testCase.expectedRowClues, solver.rowClues)
		}
		if solver.colClues != testCase.expectedColClues {
			t.Errorf("Test for the initializeClues failed:\n Expected col clues %v,  got %v\n",
				testCase.expectedColClues, solver.colClues)
		}
	}
}

func TestInitPossibleCellValuesBitmask(t *testing.T) {
	testCases := []struct {
		name                              string
		input                             []int
		expectedPossibleCellValuesBitmask [PUZZLE_SIZE * PUZZLE_SIZE]int
	}{
		{
			name: "Init with the only clue: 1",
			input: []int{
				0, 1, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0},
			expectedPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				14, 16, 14, 14,
				30, 14, 30, 30,
				30, 14, 30, 30,
				30, 14, 30, 30},
		},
		{
			name: "Init with the only clue: 4",
			input: []int{
				4, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0},
			expectedPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 28, 28, 28,
				4, 26, 26, 26,
				8, 22, 22, 22,
				16, 14, 14, 14},
		},
	}

	for _, testCase := range testCases {
		if funcResult := NewPuzzleSolver(testCase.input); !reflect.DeepEqual(funcResult.possibleCellValuesBitmask, testCase.expectedPossibleCellValuesBitmask) {
			t.Errorf("Test %s\nInput: %v, Expected %v, Got %v", testCase.name, testCase.input, testCase.expectedPossibleCellValuesBitmask, funcResult.possibleCellValuesBitmask)
		}
	}

}

func TestSetSolutionValue(t *testing.T) {
	testCases := []struct {
		name                              string
		currentSolution                   [PUZZLE_SIZE * PUZZLE_SIZE]int
		currentPossibleCellValuesBitmask  [PUZZLE_SIZE * PUZZLE_SIZE]int
		row                               int
		col                               int
		value                             int
		expectedIsSuccess                 bool
		expectedPossibleCellValuesBitmask [PUZZLE_SIZE * PUZZLE_SIZE]int
		expectedSolution                  [PUZZLE_SIZE * PUZZLE_SIZE]int
	}{
		{
			name: "Success: Set the last value in the first row",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 2, 3, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0},
			currentPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 4, 8, 16,
				28, 26, 22, 30,
				28, 26, 22, 30,
				28, 26, 22, 30},
			row:               0,
			col:               3,
			value:             4,
			expectedIsSuccess: true,
			expectedPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 4, 8, 16,
				28, 26, 22, 14,
				28, 26, 22, 14,
				28, 26, 22, 14},
			expectedSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 2, 3, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0},
		},
		{
			name: "Fail: Set the last value in the second row",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 2, 3, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0},
			currentPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 4, 8, 16,
				28, 26, 22, 30,
				28, 26, 22, 30,
				28, 26, 22, 30},
			row:                               1,
			col:                               3,
			value:                             4,
			expectedIsSuccess:                 false,
			expectedPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{},
			expectedSolution:                  [PUZZLE_SIZE * PUZZLE_SIZE]int{},
		},
		{
			name: "Success: More complex case with semi-full field",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 3, 4, 2,
				4, 0, 0, 0,
				3, 0, 0, 0,
				2, 0, 0, 4,
			},
			currentPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 8, 16, 4,
				16, 22, 14, 10,
				8, 22, 22, 26,
				4, 22, 14, 16,
			},
			row:               2,
			col:               1,
			value:             4,
			expectedIsSuccess: true,
			expectedPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 8, 16, 4,
				16, 6, 14, 10,
				8, 16, 6, 10,
				4, 6, 14, 16,
			},
			expectedSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 3, 4, 2,
				4, 0, 0, 0,
				3, 4, 0, 0,
				2, 0, 0, 4,
			}},
	}

	for _, testCase := range testCases {

		solver := PuzzleSolver{
			solution:                  testCase.currentSolution,
			possibleCellValuesBitmask: testCase.currentPossibleCellValuesBitmask,
		}
		isSuccess := solver.setSolutionValue(testCase.row, testCase.col, testCase.value)
		if isSuccess != testCase.expectedIsSuccess {
			t.Errorf("Test %s\nInput: %v, Expected %v, Got %v",
				testCase.name,
				[]int{testCase.row, testCase.col, testCase.value},
				testCase.expectedIsSuccess,
				isSuccess)
			return
		}
		if isSuccess {
			if solver.possibleCellValuesBitmask != testCase.expectedPossibleCellValuesBitmask {
				t.Errorf("Test %s\nExpected possible values bitmask %v, Got %v",
					testCase.name,
					testCase.expectedPossibleCellValuesBitmask,
					solver.possibleCellValuesBitmask)
			}
			if solver.solution != testCase.expectedSolution {
				t.Errorf("Test %s\nExpected solution %v, Got %v",
					testCase.name,
					testCase.expectedSolution,
					solver.solution)
			}
		}
	}
}

func TestUnSetSolutionValue(t *testing.T) {
	testCases := []struct {
		name                              string
		currentSolution                   [PUZZLE_SIZE * PUZZLE_SIZE]int
		currentPossibleCellValuesBitmask  [PUZZLE_SIZE * PUZZLE_SIZE]int
		row                               int
		col                               int
		oldValue                          int
		oldCellPossibleValuesBimask       int
		expectedPossibleCellValuesBitmask [PUZZLE_SIZE * PUZZLE_SIZE]int
		expectedSolution                  [PUZZLE_SIZE * PUZZLE_SIZE]int
	}{
		{
			name: "Success: unset the last value in the first row",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 2, 3, 4,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0},
			currentPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 4, 8, 16,
				28, 26, 22, 14,
				28, 26, 22, 14,
				28, 26, 22, 14},
			row:                         0,
			col:                         3,
			oldValue:                    4,
			oldCellPossibleValuesBimask: 16,
			expectedPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 4, 8, 16,
				28, 26, 22, 30,
				28, 26, 22, 30,
				28, 26, 22, 30},
			expectedSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 2, 3, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0},
		},
		{
			name: "More complex case with semi-full field",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 3, 4, 2,
				4, 0, 0, 0,
				3, 4, 0, 0,
				2, 0, 0, 4,
			},
			currentPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 8, 16, 4,
				16, 6, 14, 10,
				8, 16, 6, 10,
				4, 6, 14, 16,
			},
			row:                         2,
			col:                         1,
			oldValue:                    4,
			oldCellPossibleValuesBimask: 22,
			expectedPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 8, 16, 4,
				16, 22, 14, 10,
				8, 22, 22, 26,
				4, 22, 14, 16,
			},
			expectedSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 3, 4, 2,
				4, 0, 0, 0,
				3, 0, 0, 0,
				2, 0, 0, 4,
			}},
	}

	for _, testCase := range testCases {

		solver := PuzzleSolver{
			solution:                  testCase.currentSolution,
			possibleCellValuesBitmask: testCase.currentPossibleCellValuesBitmask,
		}
		solver.unSetSolutionValue(testCase.row, testCase.col, testCase.oldValue, testCase.oldCellPossibleValuesBimask)
		if solver.possibleCellValuesBitmask != testCase.expectedPossibleCellValuesBitmask {
			t.Errorf("Test %s\nExpected possible values bitmask %v, Got %v",
				testCase.name,
				testCase.expectedPossibleCellValuesBitmask,
				solver.possibleCellValuesBitmask)
		}
		if solver.solution != testCase.expectedSolution {
			t.Errorf("Test %s\nExpected solution %v, Got %v",
				testCase.name,
				testCase.expectedSolution,
				solver.solution)
		}
	}
}

func TestGetNextMostConstrainedCell(t *testing.T) {
	testCases := []struct {
		name                             string
		currentSolution                  [PUZZLE_SIZE * PUZZLE_SIZE]int
		currentPossibleCellValuesBitmask [PUZZLE_SIZE * PUZZLE_SIZE]int
		expectedRow                      int
		expectedCol                      int
	}{
		{
			name: "Get the most constrained cell: one left in a row",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 2, 3, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0},
			currentPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 4, 8, 16,
				28, 26, 22, 30,
				28, 26, 22, 30,
				28, 26, 22, 30},
			expectedRow: 0,
			expectedCol: 3,
		},
		{
			name: "Get the most constrained cell: one left in a row",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 2, 0, 0,
				2, 0, 0, 0,
				3, 0, 0, 0,
				0, 0, 0, 0},
			currentPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 4, 24, 24,
				4, 26, 26, 26,
				8, 18, 22, 22,
				16, 24, 30, 30},
			expectedRow: 3,
			expectedCol: 0,
		},
	}

	for _, testCase := range testCases {
		solver := PuzzleSolver{
			solution:                  testCase.currentSolution,
			possibleCellValuesBitmask: testCase.currentPossibleCellValuesBitmask,
		}

		row, col := solver.getNextMostConstrainedCell()
		if row != testCase.expectedRow || col != testCase.expectedCol {
			t.Errorf("Test for the getNextMostConstrainedCell failed:\n Expected row %d, col %d, got %d, %d\n",
				testCase.expectedRow, testCase.expectedCol, row, col)
		}
	}
}

func TestGetPossibleValuesForPosition(t *testing.T) {
	testCases := []struct {
		name                             string
		currentSolution                  [PUZZLE_SIZE * PUZZLE_SIZE]int
		currentPossibleCellValuesBitmask [PUZZLE_SIZE * PUZZLE_SIZE]int
		row                              int
		col                              int
		skipBitmask                      int
		expectedValue                    int
	}{
		{
			name: "Get the last value in the first row",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 2, 3, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0},
			currentPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 4, 8, 16,
				28, 26, 22, 30,
				28, 26, 22, 30,
				28, 26, 22, 30},
			row:           0,
			col:           3,
			skipBitmask:   0,
			expectedValue: 4,
		},
		{
			name: "Get a value considering the skipBitmask",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 2, 0, 0,
				2, 0, 0, 0,
				3, 0, 0, 0,
				0, 0, 0, 0},
			currentPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 4, 24, 24,
				4, 26, 26, 26,
				8, 18, 22, 22,
				16, 24, 30, 30},
			row:           0,
			col:           2,
			skipBitmask:   1 << 3,
			expectedValue: 4,
		},
		{
			name: "No possible values left",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 2, 0, 0,
				2, 0, 0, 0,
				3, 0, 0, 0,
				0, 4, 0, 0},
			currentPossibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				2, 4, 24, 24,
				4, 10, 26, 26,
				8, 2, 22, 22,
				0, 16, 14, 14},
			row:           3,
			col:           0,
			skipBitmask:   0,
			expectedValue: -1,
		},
	}

	for _, testCase := range testCases {

		solver := PuzzleSolver{
			solution:                  testCase.currentSolution,
			possibleCellValuesBitmask: testCase.currentPossibleCellValuesBitmask,
		}
		expectedValue := solver.getPossibleValuesForPosition(testCase.row, testCase.col, testCase.skipBitmask)
		if expectedValue != testCase.expectedValue {
			t.Errorf("Test %s\nInput: %v, Expected %v, Got %v",
				testCase.name,
				[]int{testCase.row, testCase.col, testCase.skipBitmask},
				testCase.expectedValue,
				expectedValue)
			return
		}
	}

}

func TestDoValuesFulfillClue(t *testing.T) {
	testCases := []struct {
		name            string
		currentSolution [PUZZLE_SIZE * PUZZLE_SIZE]int
		row             int
		col             int
		clue            int
		dir             direction
		axis            axis
		expected        bool
	}{
		{
			name: "Matches: Almost full first row",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 2, 3, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0},
			row:      0,
			col:      3,
			clue:     4,
			dir:      FORWARD,
			axis:     ROW,
			expected: true,
		},
		{
			name: "Matches: Almost full first col",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 0, 0, 0,
				2, 0, 0, 0,
				3, 0, 0, 0,
				0, 0, 0, 0},
			row:      3,
			col:      0,
			clue:     4,
			dir:      FORWARD,
			axis:     COLUMN,
			expected: true,
		},
		{
			name: "Doesn't match: Met 4 before the end",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
				0, 0, 0, 0},
			row:      3,
			col:      0,
			clue:     4,
			dir:      FORWARD,
			axis:     COLUMN,
			expected: false,
		},
		{
			name: "Doesn't match: First col backward",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 0, 0, 0,
				2, 0, 0, 0,
				4, 0, 0, 0,
				3, 0, 0, 0},
			row:      3,
			col:      0,
			clue:     3,
			dir:      BACKWARD,
			axis:     COLUMN,
			expected: false,
		},
		{
			name: "Matches: Second row backward",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				0, 0, 4, 0,
				0, 0, 3, 0,
				4, 0, 0, 0,
				0, 0, 0, 0},
			row:      1,
			col:      2,
			clue:     2,
			dir:      BACKWARD,
			axis:     ROW,
			expected: true,
		},
		{
			name: "Matches: Third col forward",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				0, 0, 4, 0,
				0, 0, 3, 0,
				4, 0, 0, 0,
				0, 0, 0, 0},
			row:      1,
			col:      2,
			clue:     1,
			dir:      FORWARD,
			axis:     COLUMN,
			expected: true,
		},
		{
			name: "Matches: Third col backward",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				0, 0, 4, 0,
				0, 0, 3, 0,
				4, 0, 0, 0,
				0, 0, 0, 0},
			row:      1,
			col:      2,
			clue:     3,
			dir:      BACKWARD,
			axis:     COLUMN,
			expected: true,
		},
	}

	for _, testCase := range testCases {

		solver := PuzzleSolver{
			solution: testCase.currentSolution,
		}
		res := solver.doValuesFulfillClue(testCase.row, testCase.col, testCase.clue, testCase.axis, testCase.dir)
		if res != testCase.expected {
			t.Errorf("Test %s\nInput: %v, Expected %v, Got %v",
				testCase.name,
				[]int{testCase.row, testCase.col, testCase.clue, testCase.axis, testCase.dir},
				testCase.expected,
				res)
			return
		}
	}

}

func TestDoesCellValueMatchClues(t *testing.T) {
	testCases := []struct {
		name                      string
		currentSolution           [PUZZLE_SIZE * PUZZLE_SIZE]int
		rowClues                  [PUZZLE_SIZE][2]int
		colClues                  [PUZZLE_SIZE][2]int
		possibleCellValuesBitmask [PUZZLE_SIZE * PUZZLE_SIZE]int
		row                       int
		col                       int
		value                     int
		expected                  bool
	}{
		{
			name: "Matches: Almost full first row",
			currentSolution: [PUZZLE_SIZE * PUZZLE_SIZE]int{
				1, 3, 4, 2,
				4, 0, 0, 0,
				3, 0, 0, 0,
				2, 0, 0, 4,
			},
			rowClues:                  [PUZZLE_SIZE][2]int{{3, 2}, {1, 2}, {2, 3}, {3, 1}},
			colClues:                  [PUZZLE_SIZE][2]int{{2, 3}, {2, 2}, {1, 2}, {3, 1}},
			possibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{2, 8, 16, 4, 16, 20, 14, 10, 8, 2, 20, 24, 4, 20, 14, 16},
			row:                       2,
			col:                       3,
			value:                     4,
			expected:                  true,
		},
	}
	for _, testCase := range testCases {

		solver := PuzzleSolver{
			solution: testCase.currentSolution,
		}
		res := solver.doesCellValueMatchClues(testCase.row, testCase.col, testCase.value)
		if res != testCase.expected {
			t.Errorf("Test %s\nInput: %v, Expected %v, Got %v",
				testCase.name,
				[]int{testCase.row, testCase.col, testCase.value},
				testCase.expected,
				res)
			return
		}
	}

}
