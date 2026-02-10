package skyscrappers

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
}

func TestSolvePuzzle(t *testing.T) {
	for _, testCase := range solvePuzzleTestCases {
		if funcResult := SolvePuzzle(testCase.input); !reflect.DeepEqual(funcResult, testCase.output) {
			t.Errorf("Test %s\nInput: %v, Expected %v, Got %v", testCase.name, testCase.input, testCase.output, funcResult)
		}
	}
}

var testInitTestCase = []TestCase{
	{
		name: "Init a trivial values: ones",
		input: []int{
			0, 1, 0, 0,
			0, 1, 0, 0,
			0, 1, 0, 0,
			0, 1, 0, 0},
		output: [][]int{
			{0, 4, 0, 0},
			{0, 0, 0, 4},
			{4, 0, 0, 0},
			{0, 0, 4, 0}},
	},
	{
		name: "Init a trivial values: fours - columns",
		input: []int{
			4, 0, 0, 0,
			0, 0, 0, 0,
			4, 0, 0, 0,
			0, 0, 0, 0},
		output: [][]int{
			{1, 0, 0, 4},
			{2, 0, 0, 3},
			{3, 0, 0, 2},
			{4, 0, 0, 1}},
	},
	{
		name: "Init a trivial values: fours - rows",
		input: []int{
			0, 0, 0, 0,
			4, 0, 0, 0,
			0, 0, 0, 0,
			4, 0, 0, 0},
		output: [][]int{
			{4, 3, 2, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{1, 2, 3, 4}},
	},
}

func TestNewPuzzleSolver(t *testing.T) {
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
		if solver.ROW_CLUES != testCase.expectedRowClues {
			t.Errorf("Test for the initializeClues failed:\n Expected row clues %v,  got %v\n",
				testCase.expectedRowClues, solver.ROW_CLUES)
		}
		if solver.COL_CLUES != testCase.expectedColClues {
			t.Errorf("Test for the initializeClues failed:\n Expected col clues %v,  got %v\n",
				testCase.expectedColClues, solver.COL_CLUES)
		}
	}
}

func TestSetUnsetSolutionValue(t *testing.T) {
	solver := NewPuzzleSolver([]int{
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0})
	firstRow := 2
	firstCol := 3
	firstValue := 3
	expectedRowBitmaskValue := 1 << firstValue
	expectedColBitmaskValue := expectedRowBitmaskValue
	solver.setSolutionValue(firstRow, firstCol, firstValue)
	if solver.COL_BITMASKS[firstCol] != expectedColBitmaskValue || solver.ROW_BITMASKS[firstRow] != expectedRowBitmaskValue {
		solver.printSolution()
		t.Errorf("Test for the first call of setSolutionValue failed:\n"+
			"Expected bitmask %b, Got firstRow bitmask %b, firstCol bitmask %b\n",
			expectedRowBitmaskValue, solver.ROW_BITMASKS[firstRow], solver.COL_BITMASKS[firstCol])
	}
	expectedFilledCellsBitmaskValue := uint64(1 << (firstRow*PUZZLE_SIZE + firstCol))
	if solver.FILLED_CELLS_BITMASK != expectedFilledCellsBitmaskValue {
		solver.printSolution()
		t.Errorf("Test for the first call of setSolutionValue failed:\n"+
			"Expected filled cell bitmask %b, Got row bitmask %b\n",
			expectedFilledCellsBitmaskValue, solver.FILLED_CELLS_BITMASK)
	}

	secondRow := 3
	secondCol := 3
	value := 4
	expectedRowBitmaskValue = 1 << value
	expectedColBitmaskValue = expectedColBitmaskValue | 1<<value
	solver.setSolutionValue(secondRow, secondCol, value)
	if solver.COL_BITMASKS[secondCol] != expectedColBitmaskValue || solver.ROW_BITMASKS[secondRow] != expectedRowBitmaskValue {
		solver.printSolution()
		t.Errorf("Test for the second call of setSolutionValue failed:\n"+
			"Expected bitmask %b, Got secondRow bitmask %b, secondCol bitmask %b\n",
			expectedRowBitmaskValue, solver.ROW_BITMASKS[secondRow], solver.COL_BITMASKS[secondCol])

	}

	expectedFilledCellsBitmaskValue = expectedFilledCellsBitmaskValue | uint64(1<<(secondRow*PUZZLE_SIZE+secondCol))
	if solver.FILLED_CELLS_BITMASK != expectedFilledCellsBitmaskValue {
		solver.printSolution()
		t.Errorf("Test for the first call of setSolutionValue failed:\n"+
			"Expected filled cell bitmask %b, Got secondRow bitmask %b\n",
			expectedFilledCellsBitmaskValue, solver.FILLED_CELLS_BITMASK)
	}

	// === unSetSolutionValue ===
	solver.unSetSolutionValue(secondRow, secondCol, value)

	expectedRowBitmaskValue = 1 << firstValue
	expectedColBitmaskValue = expectedRowBitmaskValue
	expectedFilledCellsBitmaskValue = uint64(1 << (firstRow*PUZZLE_SIZE + firstCol))
	if solver.COL_BITMASKS[firstCol] != expectedColBitmaskValue || solver.ROW_BITMASKS[firstRow] != expectedRowBitmaskValue {
		solver.printSolution()
		t.Errorf("Test for the call of unSetSolutionValue failed:\n"+
			"Expected bitmask %b, Got firstRow bitmask %b, firstCol bitmask %b\n",
			expectedRowBitmaskValue, solver.ROW_BITMASKS[firstRow], solver.COL_BITMASKS[firstCol])
	}
	expectedFilledCellsBitmaskValue = uint64(1 << (firstRow*PUZZLE_SIZE + firstCol))
	if solver.FILLED_CELLS_BITMASK != expectedFilledCellsBitmaskValue {
		solver.printSolution()
		t.Errorf("Test for the call of unSetSolutionValue failed:\n"+
			"Expected filled cell bitmask %b, Got firstRow bitmask %b\n",
			expectedFilledCellsBitmaskValue, solver.FILLED_CELLS_BITMASK)
	}

}

func TestGetNextMostConstrainedCell(t *testing.T) {
	solver := NewPuzzleSolver([]int{
		0, 1, 0, 0,
		0, 2, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0})
	expectedRow, expectedCol := 1, 1
	row, col := solver.getNextMostConstrainedCell()
	if row != expectedRow || col != expectedCol {
		t.Errorf("Test for the getNextMostConstrainedCell failed:\n Expected row %d, col %d, got %d, %d\n",
			expectedRow, expectedCol, row, col)
	}
}

func TestGetPossibleValuesForPosition(t *testing.T) {
	solver := NewPuzzleSolver([]int{
		0, 1, 0, 0,
		0, 2, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0})
	row, col := 1, 1
	values := solver.getPossibleValuesForPosition(row, col)
	solver.printSolution()
	expectedValues := []int{1, 2, 3}
	solver.printSolution()
	if !reflect.DeepEqual(values, expectedValues) {
		t.Errorf("Test for the getPossibleValuesForPosition failed:\n Expected row %v,  got %v\n",
			expectedValues, values)
	}
}

type TestCaseDoValuesFulfillClue struct {
	values   []int
	clue     int
	dir      direction
	expected bool
}

func TestDoValuesFulfillClue(t *testing.T) {
	testCases := []TestCaseDoValuesFulfillClue{
		{
			[]int{1, 2, 3, 4},
			4,
			FORWARD,
			true,
		},
		{
			[]int{1, 2, 3, 4},
			1,
			BACKWARD,
			true,
		},
		{
			[]int{2, 1, 3, 4},
			3,
			FORWARD,
			true,
		},
		{
			[]int{2, 3, 1, 4},
			2,
			BACKWARD,
			false,
		},
		{
			[]int{0, 4, 0, 2},
			2,
			BACKWARD,
			true,
		},
	}

	for _, testCase := range testCases {
		result := doValuesFulfillClue(testCase.values, testCase.clue, testCase.dir)
		if result != testCase.expected {
			t.Errorf("Test for the doValuesFulfillClue failed:\n"+
				"Values: %v, clue %d, direction %v\n"+
				"Expected %v, got %v\n",
				testCase.values, testCase.clue, testCase.dir,
				testCase.expected, result)
		}
	}
}

var doesCellValueMatchCluesTestCases = []struct {
	row      int
	col      int
	value    int
	expected bool
}{
	{
		1,
		1,
		2,
		true,
	},
	{
		1,
		2,
		3,
		true,
	},
	{
		1,
		3,
		4,
		false,
	},
}

func TestDoesCellValueMatchClues(t *testing.T) {
	solver := NewPuzzleSolver([]int{
		0, 1, 0, 0,
		0, 2, 0, 0,
		0, 0, 0, 0,
		0, 0, 0, 0})

	for _, testCase := range doesCellValueMatchCluesTestCases {

		result := solver.doesCellValueMatchClues(testCase.row, testCase.col, testCase.value)
		if result != testCase.expected {
			t.Errorf("Test for the doesCellValueMatchClues failed:\n"+
				"Row: %d, col %d, value %d\n"+
				"Expected %v, got %v\n",
				testCase.row, testCase.col, testCase.value,
				testCase.expected, result)
		}
	}
}
