package skyscrappers

import "fmt"

const PUZZLE_SIZE = 4

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
}

func (puzzle *PuzzleSolver) setSolutionValue(row, col, value int) {
	// Track that value is already used for the row
	puzzle.ROW_BITMASKS[row] = puzzle.ROW_BITMASKS[row] | (1 << value)
	// Track that value is already used for the column
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

func (puzzle PuzzleSolver) printSolution() {
	fmt.Println("= Solution =")
	for i := 0; i < PUZZLE_SIZE; i++ {
		fmt.Printf("%v\n", puzzle.solution[i])
	}
	fmt.Println()
	fmt.Printf("Filled cells bitmask %b\n", puzzle.FILLED_CELLS_BITMASK)
	fmt.Println()
	fmt.Printf("Row bitmasks %v\n", puzzle.ROW_BITMASKS)
	fmt.Printf("Col bitmasks %v\n", puzzle.COL_BITMASKS)
	fmt.Println()
	fmt.Printf("Row clues %v\n", puzzle.ROW_CLUES)
	fmt.Printf("Col clues %v\n", puzzle.COL_CLUES)
	fmt.Println()
}

// Sets whole row or column to [1..PUZZLE_SIZE]
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

func (puzzle *PuzzleSolver) initializeClues() {
	for i := 0; i < PUZZLE_SIZE; i++ {
		puzzle.ROW_CLUES[i][0], puzzle.ROW_CLUES[i][1] = puzzle.clues[PUZZLE_SIZE*4-1-i], puzzle.clues[PUZZLE_SIZE+i]
		puzzle.COL_CLUES[i][0], puzzle.COL_CLUES[i][1] = puzzle.clues[i], puzzle.clues[3*PUZZLE_SIZE-1-i]

		puzzle.handleTrivialValues(ROW, i, puzzle.ROW_CLUES[i])
		puzzle.handleTrivialValues(COLUMN, i, puzzle.COL_CLUES[i])
	}
}

// Initializes clues array and covers trivial cases of clues 1 and PUZZLE_SIZE
// func (puzzle *PuzzleSolver) initializeCluesv1() {

// 	for i, value := range puzzle.clues {
// 		if value != 0 {
// 			j := i / PUZZLE_SIZE
// 			if j == 0 {
// 				columnIndex := i
// 				puzzle.COL_CLUES[columnIndex][0] = value
// 				if value == 1 {
// 					puzzle.setSolutionValue(0, columnIndex, PUZZLE_SIZE)
// 				}

// 				if value == PUZZLE_SIZE {
// 					puzzle.setRange(COLUMN, columnIndex, true)
// 				}

// 			}
// 			if j == 1 {
// 				rowIndex := i - PUZZLE_SIZE
// 				puzzle.ROW_CLUES[rowIndex][1] = value
// 				if value == 1 {
// 					puzzle.setSolutionValue(rowIndex, PUZZLE_SIZE-1, PUZZLE_SIZE)
// 				}
// 				if value == PUZZLE_SIZE {
// 					puzzle.setRange(ROW, rowIndex, false)
// 				}
// 			}
// 			if j == 2 {
// 				columIndex := PUZZLE_SIZE - (i - PUZZLE_SIZE*j) - 1
// 				puzzle.COL_CLUES[columIndex][1] = value
// 				if value == 1 {
// 					puzzle.setSolutionValue(PUZZLE_SIZE-1, columIndex, PUZZLE_SIZE)
// 				}

// 				if value == PUZZLE_SIZE {
// 					puzzle.setRange(COLUMN, columIndex, false)
// 				}

// 			}
// 			if j == 3 {
// 				rowIndex := PUZZLE_SIZE - (i - PUZZLE_SIZE*j) - 1
// 				puzzle.ROW_CLUES[rowIndex][0] = value
// 				if value == 1 {
// 					puzzle.setSolutionValue(rowIndex, 0, PUZZLE_SIZE)
// 				}
// 				if value == PUZZLE_SIZE {
// 					puzzle.setRange(ROW, rowIndex, true)
// 				}
// 			}
// 		}
// 	}

// }

// Constraint score is used to define the next cell algorithm will use
// The more clues and already set cells in the cell's row / col, the higher the score will be
func (puzzle *PuzzleSolver) getCellConstraintScore(row, col int) int {
	score := 0

	score += puzzle.ROW_CLUES[row][0] + puzzle.ROW_CLUES[row][1]
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

func (puzzle *PuzzleSolver) getNextMostConstrainedCell() (int, int) {
	row, col := -1, -1
	constraintScore := 0
	// puzzle.printSolution()
	for i := 0; i < PUZZLE_SIZE*PUZZLE_SIZE; i++ {
		if puzzle.FILLED_CELLS_BITMASK&(1<<i) == 0 {
			// fmt.Printf("Cell %d is free\n", i)
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

func doValuesFulfillClue(values []int, clue int, dir direction) bool {
	if clue == 0 {
		return true
	}

	count := 1
	zeros_count := 0
	switch dir {
	case FORWARD:
		max := values[0]
		for i := 1; i < PUZZLE_SIZE; i++ {
			if values[i-1] == PUZZLE_SIZE {
				break
			}
			if values[i] == PUZZLE_SIZE { // No need to iterate further as we won't see other skyscrapers after the tallest one.
				count += 1
				break // Check first one
			}
			if values[i] != 0 {
				if values[i] > values[i-1] && values[i] >= max {
					count += 1
				}
				if values[i] > max {
					max = values[i]
				}
			} else {
				zeros_count += 1
			}
		}
	case BACKWARD:
		max := values[PUZZLE_SIZE-1]
		for i := PUZZLE_SIZE - 2; i >= 0; i-- {
			if values[i+1] == PUZZLE_SIZE {
				break
			}
			if values[i] == PUZZLE_SIZE { // No need to iterate further as we won't see other skyscrapers after the tallest one.
				count += 1
				break
			}
			if values[i] != 0 {
				if values[i] > values[i+1] && values[i] >= max {
					count += 1
				}
				if values[i] > max {
					max = values[i]
				}
			} else {
				zeros_count += 1
			}
		}
	}

	fmt.Printf("doValuesFulfillClue\n Values %v, clue %d, direction %d, count %d, zeros %d\n", values, clue, dir, count, zeros_count)
	fmt.Printf("doValuesFulfillClue returning %v\n", clue-count <= zeros_count)
	// # TODO (Alexander Shein)
	return clue-count <= zeros_count // It should be within count of zeros from the clue
}

func (puzzle *PuzzleSolver) doesCellValueMatchClues(row, col, value int) bool {

	fmt.Printf("doesCellValueMatchClues\n"+
		"row %d col %d value %d\n", row, col, value)
	if puzzle.ROW_CLUES[row][0] != 0 || puzzle.ROW_CLUES[row][1] != 0 {
		fmt.Printf("Row clues: %v\n", puzzle.ROW_CLUES[row])
		row_values := make([]int, PUZZLE_SIZE)
		copy(row_values, puzzle.solution[row])
		row_values[col] = value

		if !doValuesFulfillClue(row_values, puzzle.ROW_CLUES[row][0], FORWARD) ||
			!doValuesFulfillClue(row_values, puzzle.ROW_CLUES[row][1], BACKWARD) {
			return false
		}
	}
	if puzzle.COL_CLUES[col][0] != 0 || puzzle.COL_CLUES[col][1] != 0 {
		fmt.Printf("Col clues: %v\n", puzzle.ROW_CLUES[row])
		col_values := make([]int, PUZZLE_SIZE)
		for i := 0; i < PUZZLE_SIZE; i++ {
			col_values[i] = puzzle.solution[i][col]
		}
		col_values[row] = value
		if !doValuesFulfillClue(col_values, puzzle.COL_CLUES[col][0], FORWARD) ||
			!doValuesFulfillClue(col_values, puzzle.COL_CLUES[col][1], BACKWARD) {
			return false
		}
	}

	return true
}

func (puzzle *PuzzleSolver) getPossibleValuesForPosition(row, col int) []int {
	// Go through bitmasks for row, col
	// Check each number to see if it is compliant with clues
	result := []int{}

	// fmt.Printf("row %d, col %d\n", row, col)
	// fmt.Printf("row bitmask %b, col bitmask %b\n", puzzle.ROW_BITMASKS[row], puzzle.COL_BITMASKS[col])

	for value := 1; value <= PUZZLE_SIZE; value++ {
		// fmt.Printf("Col bitmask check for %b, res %b\n", (1 << value), puzzle.COL_BITMASKS[col]&(1<<value))
		// fmt.Printf("puzzle.ROW_BITMASKS[row]&(1<<value) %v\n", puzzle.ROW_BITMASKS[row]&(1<<value) == 0)
		// fmt.Printf("puzzle.COL_BITMASKS[col]&(1<<value) %v\n", puzzle.COL_BITMASKS[col]&(1<<value) == 0)

		alreadyUsedValuesBitmask := puzzle.ROW_BITMASKS[row] ^ puzzle.COL_BITMASKS[col]

		if alreadyUsedValuesBitmask&(1<<value) == 0 && puzzle.doesCellValueMatchClues(row, col, value) {
			result = append(result, value)
			// fmt.Printf("Updated res %v\n", result)
		}
	}

	return result
}

// Recursively trying to solve the puzzle, one point / value at a time
func (puzzle *PuzzleSolver) solutionStep() bool {

	fmt.Printf("Entering the solution step\n")
	puzzle.printSolution()
	return false // # TODO (Alexander Shein)

	row, col := puzzle.getNextMostConstrainedCell()
	if row != -1 {
		possibleValues := puzzle.getPossibleValuesForPosition(row, col)
		for _, value := range possibleValues {
			puzzle.setSolutionValue(row, col, value)
			if puzzle.solutionStep() {
				return true
			}
			puzzle.unSetSolutionValue(row, col, value)
		}
	}
	return false
}

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

func SolvePuzzle(clues []int) [][]int {
	puzzleSolver := NewPuzzleSolver((clues))
	puzzleSolver.solutionStep() // The task states we always receive a valid input with exactly one solution
	return puzzleSolver.solution
}
