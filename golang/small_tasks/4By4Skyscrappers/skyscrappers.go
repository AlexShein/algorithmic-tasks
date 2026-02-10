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

	stepCounter int // To be removed
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
	// fmt.Printf("doValuesFulfillClue values %v clue %d dir %d\n", values, clue, dir)
	if clue == 0 {
		return true
	}

	count, zeros_count, max, index := 0, 0, 0, 0

	for i := 0; i < PUZZLE_SIZE; i++ {
		switch dir {
		case FORWARD:
			index = i
		case BACKWARD:
			index = PUZZLE_SIZE - 1 - i
		}
		if values[index] == 0 {
			zeros_count += 1
			break
		} else if values[index] > max {
			count += 1
			max = values[index]
		}
		if values[index] == PUZZLE_SIZE { // No need to iterate further as we won't see other skyscrapers after the tallest one.
			break
		}
	}

	// fmt.Printf("doValuesFulfillClue\n Values %v, clue %d, direction %d, count %d, zeros %d\n", values, clue, dir, count, zeros_count)
	// fmt.Printf("doValuesFulfillClue returning %v\n", count <= clue && clue-count <= zeros_count)
	if zeros_count == 0 { // Only return false if values clearly violate the clue
		return count == clue
	}
	return true
}

func (puzzle *PuzzleSolver) doesCellValueMatchClues(row, col, value int) bool {

	// fmt.Printf("doesCellValueMatchClues\n"+
	// 	"row %d col %d value %d\n", row, col, value)
	if puzzle.ROW_CLUES[row][0] != 0 || puzzle.ROW_CLUES[row][1] != 0 {
		// fmt.Printf("Row clues: %v\n", puzzle.ROW_CLUES[row])
		row_values := make([]int, PUZZLE_SIZE)
		copy(row_values, puzzle.solution[row])
		row_values[col] = value

		// if row == 2 && col == 1 {

		// 	// fmt.Printf("row_values %v clues %v\n", row_values, puzzle.ROW_CLUES[row])
		// 	// fmt.Printf("!doValuesFulfillClue(row_values, puzzle.ROW_CLUES[row][0], FORWARD) %v\n", !doValuesFulfillClue(row_values, puzzle.ROW_CLUES[row][0], FORWARD))
		// 	// fmt.Printf("!doValuesFulfillClue(row_values, puzzle.ROW_CLUES[row][1], BACKWARD) %v\n", !doValuesFulfillClue(row_values, puzzle.ROW_CLUES[row][1], BACKWARD))
		// }
		if !doValuesFulfillClue(row_values, puzzle.ROW_CLUES[row][0], FORWARD) ||
			!doValuesFulfillClue(row_values, puzzle.ROW_CLUES[row][1], BACKWARD) {
			// fmt.Printf("doesCellValueMatchClues returns false because values don't fulfill a *row* clue\n")
			return false
		}
	}
	if puzzle.COL_CLUES[col][0] != 0 || puzzle.COL_CLUES[col][1] != 0 {
		// fmt.Printf("doesCellValueMatchClues: Col clues: %v\n", puzzle.ROW_CLUES[row])
		col_values := make([]int, PUZZLE_SIZE)
		for i := 0; i < PUZZLE_SIZE; i++ {
			col_values[i] = puzzle.solution[i][col]
		}
		col_values[row] = value
		if !doValuesFulfillClue(col_values, puzzle.COL_CLUES[col][0], FORWARD) ||
			!doValuesFulfillClue(col_values, puzzle.COL_CLUES[col][1], BACKWARD) {
			// fmt.Printf("doesCellValueMatchClues returns false because values don't fulfill a *column* clue\n")

			return false
		}
	}

	return true
}

func (puzzle *PuzzleSolver) getPossibleValuesForPosition(row, col, skipBitmask int) int {
	// # TODO (Alexander Shein) Take used numbers as an arg to prevent checking all of them

	// Go through bitmasks for row, col
	// Check each number to see if it is compliant with clues

	// fmt.Printf("==getPossibleValuesForPosition is called with row %d, col %d\n", row, col)
	// fmt.Printf("row %d, col %d\n", row, col)
	// fmt.Printf("row bitmask %b, col bitmask %b skip bitmask %b\n", puzzle.ROW_BITMASKS[row], puzzle.COL_BITMASKS[col], skipBitmask)
	alreadyUsedValuesBitmask := puzzle.ROW_BITMASKS[row] | puzzle.COL_BITMASKS[col] | skipBitmask
	// if col == 0 {
	// 	fmt.Printf("Row bitmask %b\n", puzzle.ROW_BITMASKS[row])
	// 	fmt.Printf("Col bitmask %b\n", puzzle.COL_BITMASKS[col])
	// 	fmt.Printf("Skip bitmask %b\n", skipBitmask)
	// 	fmt.Printf("alreadyUsedValuesBitmask %b\n", alreadyUsedValuesBitmask)
	// }

	for value := 1; value <= PUZZLE_SIZE; value++ {
		// fmt.Printf("Bitmask check for %b, res %v\n", (1 << value), alreadyUsedValuesBitmask&(1<<value) == 0)
		// fmt.Printf("puzzle.ROW_BITMASKS[row]&(1<<value) %v\n", puzzle.ROW_BITMASKS[row]&(1<<value) == 0)
		// fmt.Printf("puzzle.COL_BITMASKS[col]&(1<<value) %v\n", puzzle.COL_BITMASKS[col]&(1<<value) == 0)

		if alreadyUsedValuesBitmask&(1<<value) == 0 && puzzle.doesCellValueMatchClues(row, col, value) {
			// fmt.Printf("==getPossibleValuesForPosition returns %d\n", value)
			return value
			// fmt.Printf("Updated res %v\n", result)
		}
	}
	// fmt.Printf("==getPossibleValuesForPosition returns %d\n", -1)
	return -1
}

// Recursively trying to solve the puzzle, one point / value at a time
func (puzzle *PuzzleSolver) solutionStep() bool {

	// if puzzle.stepCounter > 3 {
	// 	return false
	// }

	puzzle.stepCounter++

	// fmt.Printf("Entering the solution step\n")
	// puzzle.printSolution(true)

	row, col := puzzle.getNextMostConstrainedCell()
	skipBitmask := 0
	if row != -1 {
		for nextVal := puzzle.getPossibleValuesForPosition(row, col, skipBitmask); nextVal != -1; {
			// nextVal := puzzle.getPossibleValuesForPosition(row, col, skipBitmask)
			// if row == 2 {
			// }
			// fmt.Printf("Most constrained cell is %d %d\n nextVal is %v\n skipBitmask is %b\n", row, col, nextVal, skipBitmask)
			skipBitmask |= (1 << nextVal)

			// nextVal = puzzle.getPossibleValuesForPosition(row, col, skipBitmask)
			puzzle.setSolutionValue(row, col, nextVal)
			// fmt.Printf("Set the new value, printout after\n")
			// puzzle.printSolution(true)

			if puzzle.solutionStep() {
				return true
			}
			puzzle.unSetSolutionValue(row, col, nextVal)

			// This is critical for the loop to work
			nextVal = puzzle.getPossibleValuesForPosition(row, col, skipBitmask)
		}
		return false
	}
	return true
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

// The task states we always receive a valid input with exactly one solution
// Hence, we do no validation for clues
func SolvePuzzle(clues []int) [][]int {
	fmt.Printf("Received clues %v\n", clues)
	puzzleSolver := NewPuzzleSolver((clues))
	puzzleSolver.solutionStep()
	puzzleSolver.printSolution(false)
	return puzzleSolver.solution
}
