package insertionsort

import "fmt"

func InsertionSort(numbers []int) []int {
	for i := 0; i < len(numbers); i++ {
		for j := i; j > 0 && numbers[j-1] > numbers[j]; {
			fmt.Printf("Swapping index %d value %d with index %d value %d\n", j, numbers[j], j-1, numbers[j-1])
			numbers[j], numbers[j-1] = numbers[j-1], numbers[j]
			j--
		}
	}
	return numbers
}
