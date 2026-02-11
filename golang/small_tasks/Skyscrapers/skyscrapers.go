package skyscrapers

import "fmt"

const PUZZLE_SIZE = 4 // N

type axis = int

const (
	ROW axis = iota
	COLUMN
)

type direction = int

const (
	FORWARD direction = iota
	BACKWARD
)

type PuzzleSolver struct {
	// 1 at n-th position indicates that n has already been used for a given row.
	ROW_BITMASKS [PUZZLE_SIZE]int
	COL_BITMASKS [PUZZLE_SIZE]int

	// Each row and column could have 0 - 2 clues.
	ROW_CLUES [PUZZLE_SIZE][2]int
	COL_CLUES [PUZZLE_SIZE][2]int

	// bit in n-th position indicates that n-th cell already has a value.
	// using uint64 to accomodate up to 8x8 puzzles
	FILLED_CELLS_BITMASK uint64

	clues    []int
	solution [][]int

	stepCounter int // Used to track performance
}

/*
== TASK ==
In a grid of N by N squares you want to place a skyscraper in each square with only some clues:

The height of the skyscrapers is between 1 and N
No two skyscrapers in a row or column may have the same number of floors
A clue is the number of skyscrapers that you can see in a row or column from the outside
Higher skyscrapers block the view of lower skyscrapers located behind them

== Solution ==
The solution presented here uses a backtracking mechanism to find a solution.
The goal is to reduce the possible solutions space as much as possible with each step.
Bitmasks are used for fast number / cell usage checks.

The algorithm works as following:
0. Initialize the solver with bitmasks and internal arrays.
1. Initialize clues and put numbers to cells / rows / cols that have exactly 1 possible value (trivial cells).
2. Iterate while there are empty cells:
	a. Find most constrained cell.
	b. Find first possible value to be used in the cell.
	c. Set the value to the cell and call the solve function recursively.
	d. If solution is found - *return true*.
	e. If the solution has not been reached, unset the cell value and continue.
	f. If out of possible values, go back to the previous cell by returning false.
3. Return the solution.

The task states we always receive a valid input with exactly one solution
Hence, there's no validation for clues
*/
func SolvePuzzle(clues []int) [][]int {
	fmt.Printf("Received clues %v\n", clues)
	puzzleSolver := NewPuzzleSolver((clues))
	puzzleSolver.solutionStep()
	puzzleSolver.printSolution(false)
	return puzzleSolver.solution
}

// Factory function for the solver. Initializes internally used slices, arrays
// and takes care of trivial cell values with clues 1 and PUZZLE_SIZE
func NewPuzzleSolver(clues []int) *PuzzleSolver {
	puzzle := PuzzleSolver{
		ROW_BITMASKS: [PUZZLE_SIZE]int{},
		COL_BITMASKS: [PUZZLE_SIZE]int{},
		ROW_CLUES:    [PUZZLE_SIZE][2]int{},
		COL_CLUES:    [PUZZLE_SIZE][2]int{},
		clues:        clues}

	puzzle.solution = make([][]int, PUZZLE_SIZE)
	for i := 0; i < PUZZLE_SIZE; i++ {
		puzzle.solution[i] = make([]int, PUZZLE_SIZE)
	}
	puzzle.initializeClues()

	return &puzzle

}

// Recursively solves the puzzle, one cell at a time, starting with the most constrained cell.
func (puzzle *PuzzleSolver) solutionStep() bool {
	puzzle.stepCounter++
	skipBitmask := 0
	row, col := puzzle.getNextMostConstrainedCell()
	if row != -1 {
		for nextVal := puzzle.getPossibleValuesForPosition(row, col, skipBitmask); nextVal != -1; {
			skipBitmask |= (1 << nextVal) // Marking the current value to be skipped in the next iteration
			puzzle.setSolutionValue(row, col, nextVal)
			if puzzle.solutionStep() { // Check if current value leads to a solution
				return true
			}
			puzzle.unSetSolutionValue(row, col, nextVal) // Resets the cell in case the value did not lead to the solution
			nextVal = puzzle.getPossibleValuesForPosition(row, col, skipBitmask)
		}
		return false
	}
	return true
}

// Goes through possible values for the column and returns first one that does fulfill all conditions
// Returns -1 if there's none
func (puzzle *PuzzleSolver) getPossibleValuesForPosition(row, col, skipBitmask int) int {
	alreadyUsedValuesBitmask := puzzle.ROW_BITMASKS[row] | puzzle.COL_BITMASKS[col] | skipBitmask
	for value := 1; value <= PUZZLE_SIZE; value++ {
		if alreadyUsedValuesBitmask&(1<<value) == 0 && puzzle.doesCellValueMatchClues(row, col, value) {
			return value
		}
	}
	return -1
}

func (puzzle *PuzzleSolver) setSolutionValue(row, col, value int) {
	// Mark the value as already used for the row
	puzzle.ROW_BITMASKS[row] = puzzle.ROW_BITMASKS[row] | (1 << value)
	// Mark the value as already used for the column
	puzzle.COL_BITMASKS[col] = puzzle.COL_BITMASKS[col] | (1 << value)
	// Mark cell as filled
	puzzle.FILLED_CELLS_BITMASK = puzzle.FILLED_CELLS_BITMASK | uint64(1<<(row*PUZZLE_SIZE+col))

	puzzle.solution[row][col] = value
}

func (puzzle *PuzzleSolver) unSetSolutionValue(row, col, oldValue int) {
	// Mark old value as 0 in row bitmask
	puzzle.ROW_BITMASKS[row] = puzzle.ROW_BITMASKS[row] & ^(1 << oldValue)
	// Mark old value as 0 in column bitmask
	puzzle.COL_BITMASKS[col] = puzzle.COL_BITMASKS[col] & ^(1 << oldValue)
	// Mark cell as unfilled
	puzzle.FILLED_CELLS_BITMASK = puzzle.FILLED_CELLS_BITMASK & ^uint64(1<<(row*PUZZLE_SIZE+col))

	puzzle.solution[row][col] = 0
}

func (puzzle PuzzleSolver) printSolution(detailed bool) {
	fmt.Println("= Solution =")
	for i := 0; i < PUZZLE_SIZE; i++ {
		fmt.Printf("%v\n", puzzle.solution[i])
	}
	fmt.Println()
	fmt.Printf("Recursion counter %d\n", puzzle.stepCounter)
	fmt.Println()
	if detailed {
		fmt.Printf("Filled cells bitmask %b\n", puzzle.FILLED_CELLS_BITMASK)
		fmt.Println()
		fmt.Printf("Row bitmasks %v\n", puzzle.ROW_BITMASKS)
		fmt.Printf("Col bitmasks %v\n", puzzle.COL_BITMASKS)
		fmt.Println()
		fmt.Printf("Row clues %v\n", puzzle.ROW_CLUES)
		fmt.Printf("Col clues %v\n", puzzle.COL_CLUES)
		fmt.Println()
	}
}

// Sets whole row or column to [1..PUZZLE_SIZE] in ascending or descending order
func (puzzle *PuzzleSolver) setRange(d axis, index int, ascending bool) {
	for k := 0; k < PUZZLE_SIZE; k++ {
		value := k + 1
		if !ascending {
			value = PUZZLE_SIZE - k
		}
		switch d {
		case ROW:
			puzzle.setSolutionValue(index, k, value)
		case COLUMN:
			puzzle.setSolutionValue(k, index, value)
		}
	}
}

// Sets values to cells, rows and cols that have invariant clues 1 and PUZZLE_SIZE
func (puzzle *PuzzleSolver) handleTrivialValues(d axis, index int, clues [2]int) {
	for i, value := range clues {
		ascending := true
		secondIndex := 0
		if i == 1 {
			ascending = false
			secondIndex = PUZZLE_SIZE - 1
		}
		if value == 1 {
			switch d {
			case ROW:
				puzzle.setSolutionValue(index, secondIndex, PUZZLE_SIZE)
			case COLUMN:
				puzzle.setSolutionValue(secondIndex, index, PUZZLE_SIZE)
			}
		}
		if value == PUZZLE_SIZE {
			puzzle.setRange(d, index, ascending)
		}
	}
}

// Initializes clues array and covers trivial cases of clues 1 and PUZZLE_SIZE
func (puzzle *PuzzleSolver) initializeClues() {
	for i := 0; i < PUZZLE_SIZE; i++ {
		puzzle.ROW_CLUES[i][0], puzzle.ROW_CLUES[i][1] = puzzle.clues[PUZZLE_SIZE*4-1-i], puzzle.clues[PUZZLE_SIZE+i]
		puzzle.COL_CLUES[i][0], puzzle.COL_CLUES[i][1] = puzzle.clues[i], puzzle.clues[3*PUZZLE_SIZE-1-i]

		puzzle.handleTrivialValues(ROW, i, puzzle.ROW_CLUES[i])
		puzzle.handleTrivialValues(COLUMN, i, puzzle.COL_CLUES[i])
	}
}

// Constraint score is used to define the next cell algorithm will use
// The more clues and already set cells in the cell's row or col, the higher the score will be
func (puzzle *PuzzleSolver) getCellConstraintScore(row, col int) int {
	score := puzzle.ROW_CLUES[row][0] + puzzle.ROW_CLUES[row][1]
	score += puzzle.COL_CLUES[col][0] + puzzle.COL_CLUES[col][1]

	for i := 1; i <= PUZZLE_SIZE; i++ {
		// Increasing the score by 1 for every number already used in cell's row and column
		if puzzle.ROW_BITMASKS[row]&(1<<i) != 0 {
			score += 1
		}
		if puzzle.COL_BITMASKS[col]&(1<<i) != 0 {
			score += 1
		}
	}

	return score
}

// Finds the cell with least number of possible values to put into
func (puzzle *PuzzleSolver) getNextMostConstrainedCell() (int, int) {
	row, col := -1, -1
	constraintScore := 0
	for i := 0; i < PUZZLE_SIZE*PUZZLE_SIZE; i++ {
		if puzzle.FILLED_CELLS_BITMASK&(1<<i) == 0 {
			cellRow, cellCol := i/PUZZLE_SIZE, i%PUZZLE_SIZE
			newConstraintScore := puzzle.getCellConstraintScore(cellRow, cellCol)
			if newConstraintScore > constraintScore {
				constraintScore = newConstraintScore
				row, col = cellRow, cellCol
			}
		}
	}
	return row, col
}

// Checks if values in row / column fulfill clues. Returns true if at least one number is unset.
func doValuesFulfillClue(values []int, clue int, dir direction) bool {
	if clue == 0 {
		return true // No clue is given for this column / row
	}

	visibleCount, zerosCount, maxHeight, index := 0, 0, 0, 0

	for i := 0; i < PUZZLE_SIZE; i++ {
		switch dir {
		case FORWARD:
			index = i
		case BACKWARD:
			index = PUZZLE_SIZE - 1 - i
		}
		if values[index] == 0 {
			zerosCount += 1
			break
		} else if values[index] > maxHeight {
			visibleCount += 1
			maxHeight = values[index]
		}
		if values[index] == PUZZLE_SIZE { // No need to iterate further as we won't see other skyscrapers after the tallest one.
			break
		}
	}

	if zerosCount == 0 { // Only return false if values clearly violate the clue
		return visibleCount == clue
	}
	return true
}

func (puzzle *PuzzleSolver) doesCellValueMatchClues(row, col, value int) bool {
	if puzzle.ROW_CLUES[row][0] != 0 || puzzle.ROW_CLUES[row][1] != 0 {
		row_values := make([]int, PUZZLE_SIZE)
		copy(row_values, puzzle.solution[row])
		row_values[col] = value // Set the value to the cell and validate the row
		if !doValuesFulfillClue(row_values, puzzle.ROW_CLUES[row][0], FORWARD) ||
			!doValuesFulfillClue(row_values, puzzle.ROW_CLUES[row][1], BACKWARD) {
			return false
		}
	}
	if puzzle.COL_CLUES[col][0] != 0 || puzzle.COL_CLUES[col][1] != 0 {
		col_values := make([]int, PUZZLE_SIZE)
		for i := 0; i < PUZZLE_SIZE; i++ {
			col_values[i] = puzzle.solution[i][col]
		}
		col_values[row] = value // Set the value to the cell and validate the column
		if !doValuesFulfillClue(col_values, puzzle.COL_CLUES[col][0], FORWARD) ||
			!doValuesFulfillClue(col_values, puzzle.COL_CLUES[col][1], BACKWARD) {
			return false
		}
	}

	return true
}
