package lastdigit

import (
	"strconv"
)

var possiblePowerLastDigits = [][]int{
	{0},
	{1},
	{2, 4, 8, 6},
	{3, 9, 7, 1},
	{4, 6},
	{5},
	{6},
	{7, 9, 3, 1},
	{8, 4, 2, 6},
	{9, 1},
}

func LastDigit(n1, n2 string) int {
	lastBaseDigit, _ := strconv.Atoi(string(n1[len(n1)-1]))
	last2ExpDigits, _ := strconv.Atoi(string(n2[len(n2)-1]))
	if len(n2) > 1 {
		last2ExpDigits, _ = strconv.Atoi(string(n2[len(n2)-2:]))
	}
	period := len(possiblePowerLastDigits[lastBaseDigit]) // How often last digits of powers repeat
	index := last2ExpDigits%period - 1                    // Take the value from maps by index
	if index < 0 {
		index = period + index // In case the dvision remainder is 0, take the last element from the slice
	}
	return possiblePowerLastDigits[lastBaseDigit][index]
}
