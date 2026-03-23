package skyscrapersV2

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
	// Each row and column could have 0 - 2 clues.
	rowClues [PUZZLE_SIZE][2]int
	colClues [PUZZLE_SIZE][2]int

	// 1 at n-th bit of k-th value indicates the n could be set as a value for the k cell
	possibleCellValuesBitmask [PUZZLE_SIZE * PUZZLE_SIZE]int

	// The solution is transformed into 2d array when returned
	solution [PUZZLE_SIZE * PUZZLE_SIZE]int

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
	c. Set the value to the cell. If this leaves no possible values to some cell - return.
	d. Call the solve function recursively.
	e. If solution is found - *return true*.
	f. If the solution has not been reached, unset the cell value and continue.
	g. If out of possible values, go back to the previous cell by returning false.
3. Return the solution.

The task states we always receive a valid input with exactly one solution.
Hence, there's no validation for clues.
*/
func SolvePuzzle(clues []int) [][]int {
	fmt.Printf("Received clues %v\n", clues)
	puzzleSolver := NewPuzzleSolver((clues))
	puzzleSolver.solutionStep()
	puzzleSolver.printSolution(false)

	solution := make([][]int, PUZZLE_SIZE)
	for i := 0; i < PUZZLE_SIZE; i++ {
		solution[i] = make([]int, PUZZLE_SIZE)
		for j := 0; j < PUZZLE_SIZE; j++ {
			solution[i][j] = puzzleSolver.solution[i*PUZZLE_SIZE+j]
		}
	}

	return solution
}

// Factory function for the solver. Initializes internally used arrays
// and takes care of trivial cell values with clues 1 and PUZZLE_SIZE
func NewPuzzleSolver(clues []int) *PuzzleSolver {
	puzzle := PuzzleSolver{
		rowClues:                  [PUZZLE_SIZE][2]int{},
		colClues:                  [PUZZLE_SIZE][2]int{},
		possibleCellValuesBitmask: [PUZZLE_SIZE * PUZZLE_SIZE]int{},
		solution:                  [PUZZLE_SIZE * PUZZLE_SIZE]int{},
	}

	defaultPossibleCellValues := 0
	for i := 1; i <= PUZZLE_SIZE; i++ {
		defaultPossibleCellValues |= 1 << i
	}

	for i := 0; i < PUZZLE_SIZE*PUZZLE_SIZE; i++ {
		puzzle.possibleCellValuesBitmask[i] = defaultPossibleCellValues
	}
	puzzle.initializeClues(clues)

	return &puzzle

}

func (puzzle PuzzleSolver) printSolution(detailed bool) {
	fmt.Println("= Solution =")
	fmt.Print("{\n")
	for i := 0; i < PUZZLE_SIZE; i++ {
		fmt.Print("{")
		for j := 0; j < PUZZLE_SIZE; j++ {
			fmt.Print(puzzle.solution[i*PUZZLE_SIZE+j], ",")
		}
		fmt.Print("},")

		fmt.Println()
	}
	fmt.Print("}")
	fmt.Println()
	fmt.Printf("Recursion counter %d\n", puzzle.stepCounter)
	fmt.Println()
	if detailed {
		fmt.Print("Possible cell values bitmask: \n{")
		for i, val := range puzzle.possibleCellValuesBitmask {
			if i%PUZZLE_SIZE == 0 {
				fmt.Println()
			}
			fmt.Printf("%d,", val)
		}
		fmt.Print("\n}\n")

		fmt.Print("Row clues: {")
		for _, val := range puzzle.rowClues {
			fmt.Printf("{%d, %d},", val[0], val[1])
		}
		fmt.Print("}\n")
		fmt.Print("Col clues: {")
		for _, val := range puzzle.colClues {
			fmt.Printf("{%d, %d},", val[0], val[1])
		}
		fmt.Print("}\n")
		fmt.Println()
	}
}

// Recursively solves the puzzle, one cell at a time, starting with the most constrained cell.
func (puzzle *PuzzleSolver) solutionStep() bool {
	puzzle.stepCounter++
	skipBitmask := 0
	row, col := puzzle.getNextMostConstrainedCell()
	if row != -1 {
		for nextVal := puzzle.getPossibleValuesForPosition(row, col, skipBitmask); nextVal != -1; {
			skipBitmask |= (1 << nextVal) // Marking the current value to be skipped in the next iteration
			oldCellPossibleValuesBimask := puzzle.possibleCellValuesBitmask[row*PUZZLE_SIZE+col]
			if puzzle.setSolutionValue(row, col, nextVal) {
				if puzzle.solutionStep() { // Check if current value leads to a solution
					return true
				}
			}
			// Resets the cell in case the value did not lead to the solution
			puzzle.unSetSolutionValue(row, col, nextVal, oldCellPossibleValuesBimask)
			nextVal = puzzle.getPossibleValuesForPosition(row, col, skipBitmask)
		}
		return false
	}
	return true
}

// Goes through possible values for the column and returns first one that does fulfill all conditions
// Returns -1 if there's none
func (puzzle *PuzzleSolver) getPossibleValuesForPosition(row, col, skipBitmask int) int {
	cellIndex := row*PUZZLE_SIZE + col
	for value := 1; value <= PUZZLE_SIZE; value++ {
		valueBitmask := (1 << value)
		// Check if the value is allowed for the cell
		if valueBitmask&puzzle.possibleCellValuesBitmask[cellIndex]&(^skipBitmask) != 0 {
			if puzzle.doesCellValueMatchClues(row, col, value) {
				return value
			}
		}
	}
	return -1
}

// setSolutionValue returns true if other cells still have values to be set and false otherwise
func (puzzle *PuzzleSolver) setSolutionValue(row, col, value int) bool {

	isSuccess := true

	valueBitmask := ^(1 << value)
	currentCellIndex := row*PUZZLE_SIZE + col
	for i := 0; i < PUZZLE_SIZE; i++ {
		// Cover the column
		colIndex := i*PUZZLE_SIZE + col
		if colIndex != currentCellIndex {
			if puzzle.possibleCellValuesBitmask[colIndex]&valueBitmask == 0 {
				isSuccess = false
			}
			if puzzle.solution[colIndex] == 0 {
				// Set 0 to corresponding bit of the possible cell values bitmask
				puzzle.possibleCellValuesBitmask[colIndex] &= valueBitmask
			}
		}
		// Cover the row
		rowIndex := row*PUZZLE_SIZE + i
		if rowIndex != currentCellIndex {

			if puzzle.possibleCellValuesBitmask[rowIndex]&valueBitmask == 0 {
				isSuccess = false
			}
			if puzzle.solution[rowIndex] == 0 {
				// Set 0 to corresponding bit of the possible cell values bitmask
				puzzle.possibleCellValuesBitmask[rowIndex] &= valueBitmask
			}
		}
	}

	puzzle.solution[currentCellIndex] = value
	puzzle.possibleCellValuesBitmask[currentCellIndex] = 1 << value

	return isSuccess
}

func (puzzle *PuzzleSolver) unSetSolutionValue(row, col, oldValue, oldCellPossibleValuesBimask int) {
	currentCellIndex := row*PUZZLE_SIZE + col
	valueBitmask := 1 << oldValue
	for i := 0; i < PUZZLE_SIZE; i++ {
		// Cover the column
		colIndex := i*PUZZLE_SIZE + col
		if colIndex != currentCellIndex {
			// Set 1 to corresponding bit of the possible cell values bitmask
			if puzzle.solution[colIndex] == 0 {
				puzzle.possibleCellValuesBitmask[colIndex] |= valueBitmask
			}
		}
		// Cover the row
		rowIndex := row*PUZZLE_SIZE + i
		if rowIndex != currentCellIndex {
			// Set 1 to corresponding bit of the possible cell values bitmask
			if puzzle.solution[rowIndex] == 0 {
				puzzle.possibleCellValuesBitmask[rowIndex] |= valueBitmask
			}
		}
	}

	puzzle.solution[currentCellIndex] = 0
	puzzle.possibleCellValuesBitmask[currentCellIndex] = oldCellPossibleValuesBimask

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
func (puzzle *PuzzleSolver) initializeClues(clues []int) {
	for i := 0; i < PUZZLE_SIZE; i++ {
		puzzle.rowClues[i][0], puzzle.rowClues[i][1] = clues[PUZZLE_SIZE*4-1-i], clues[PUZZLE_SIZE+i]
		puzzle.colClues[i][0], puzzle.colClues[i][1] = clues[i], clues[3*PUZZLE_SIZE-1-i]

		puzzle.handleTrivialValues(ROW, i, puzzle.rowClues[i])
		puzzle.handleTrivialValues(COLUMN, i, puzzle.colClues[i])
	}
}

// Constraint score is used to define the next cell algorithm will use
// The more clues and already set cells in the cell's row or col, the higher the score will be
func (puzzle *PuzzleSolver) getCellConstraintScore(row, col int) int {
	score := puzzle.rowClues[row][0] + puzzle.rowClues[row][1]
	score += puzzle.colClues[col][0] + puzzle.colClues[col][1]

	cellIndex := row*PUZZLE_SIZE + col

	for i := 1; i <= PUZZLE_SIZE; i++ {
		// Increasing the score by 1 for every number unavailable for the cell
		if (^puzzle.possibleCellValuesBitmask[cellIndex])&(1<<i) != 0 {
			score += 1
		}

	}

	return score
}

// Finds the cell with least number of possible values to put into
func (puzzle *PuzzleSolver) getNextMostConstrainedCell() (int, int) {
	row, col := -1, -1
	constraintScore := 0
	for cellRow := 0; cellRow < PUZZLE_SIZE; cellRow++ {
		for cellCol := 0; cellCol < PUZZLE_SIZE; cellCol++ {
			index := cellRow*PUZZLE_SIZE + cellCol
			if puzzle.solution[index] == 0 {
				newConstraintScore := puzzle.getCellConstraintScore(cellRow, cellCol)
				if newConstraintScore > constraintScore {
					constraintScore = newConstraintScore
					row, col = cellRow, cellCol
				}
			}
		}
	}
	return row, col
}

// Checks if values in row / column fulfill clues. Returns true if at least one number is unset.
func (puzzle *PuzzleSolver) doValuesFulfillClue(row, col, clue int, axis axis, dir direction) bool {
	if clue == 0 {
		return true // No clue is given for this column / row
	}
	// If curr == PUZZLE_SIZE and # of cells before + 1 < clue - return false (We won't get to a solution)

	visibleCount, zerosCount, maxHeight, index := 0, 0, 0, 0

	for i := 0; i < PUZZLE_SIZE; i++ {
		switch dir {
		case FORWARD:
			index = i
		case BACKWARD:
			index = PUZZLE_SIZE - 1 - i
		}

		currentValue := 0
		if axis == ROW {
			currentValue = puzzle.solution[row*PUZZLE_SIZE+index]
		}
		if axis == COLUMN {
			currentValue = puzzle.solution[index*PUZZLE_SIZE+col]
		}

		if i+1+(PUZZLE_SIZE-currentValue) < clue { // Number of cells we went through + how much more we *might* see.
			return false // Total number of observable buildings is less than <clue>
		}

		if currentValue == 0 {
			zerosCount += 1
			break
		} else if currentValue > maxHeight {
			visibleCount += 1
			maxHeight = currentValue
		}
		if currentValue == PUZZLE_SIZE { // No need to iterate further as we won't see other skyscrapers after the tallest one.
			break
		}
	}

	if zerosCount == 0 { // Only return false if values clearly violate the clue
		return visibleCount == clue
	}
	return true
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

		currentValue := values[index]

		if currentValue == 0 {
			zerosCount += 1
			break
		} else if currentValue > maxHeight {
			visibleCount += 1
			maxHeight = currentValue
		}
		if currentValue == PUZZLE_SIZE { // No need to iterate further as we won't see other skyscrapers after the tallest one.
			break
		}
	}

	if zerosCount == 0 { // Only return false if values clearly violate the clue
		return visibleCount == clue
	}
	return true
}

func (puzzle *PuzzleSolver) doesCellValueMatchClues(row, col, value int) bool {
	cellIndex := row*PUZZLE_SIZE + col
	oldValue := puzzle.solution[cellIndex]
	puzzle.solution[cellIndex] = value

	if puzzle.rowClues[row][0] != 0 || puzzle.rowClues[row][1] != 0 {
		if !puzzle.doValuesFulfillClue(row, col, puzzle.rowClues[row][0], ROW, FORWARD) ||
			!puzzle.doValuesFulfillClue(row, col, puzzle.rowClues[row][1], ROW, BACKWARD) {
			puzzle.solution[cellIndex] = oldValue
			return false
		}
	}
	if puzzle.colClues[col][0] != 0 || puzzle.colClues[col][1] != 0 {
		if !puzzle.doValuesFulfillClue(row, col, puzzle.colClues[col][0], COLUMN, FORWARD) ||
			!puzzle.doValuesFulfillClue(row, col, puzzle.colClues[col][1], COLUMN, BACKWARD) {

			puzzle.solution[cellIndex] = oldValue
			return false
		}
	}

	puzzle.solution[cellIndex] = oldValue
	return true
}
