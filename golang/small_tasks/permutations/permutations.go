package permutations

// Permutations returns a list of all possible symbol permutations for a given input, without duplicates
func PermutationsRecursive(s string) []string {

	if len(s) == 1 {
		return []string{s}
	}

	result := []string{}
	uniquePermutations := map[string]bool{} // Used for quick duplicates check
	for position, character := range s {
		for _, lowerOrderPermutation := range PermutationsRecursive(s[:position] + s[position+1:]) {
			permutation := string(character) + lowerOrderPermutation
			if !uniquePermutations[permutation] {
				result = append(result, permutation)
				uniquePermutations[permutation] = true
			}
		}
	}
	return result
}

// Permutations returns a list of all possible symbol permutations for a given input, without duplicates
func Permutations(s string) []string {

	if len(s) == 1 {
		return []string{s}
	}

	runes := []rune(s)
	InsertionSort(runes)

	result := []string{string(runes)}
	for nextPermutation(runes) {
		result = append(result, string(runes))
	}
	return result
}

// nextPermutation moves the rune array into next permutaiton lexicographically returning false if it is not possilbe
func nextPermutation(runes []rune) bool {

	for i := len(runes) - 1; i > 0; i-- {
		if runes[i] > runes[i-1] { // A pivot value is found
			// Take smallest item to the right from the pivot one
			itemIndexToSwapPivotWith := i
			for j := i + 1; j < len(runes); j++ { // Finding a smallest item to the right from pivot
				if runes[j] <= runes[itemIndexToSwapPivotWith] && runes[j] > runes[i-1] {
					itemIndexToSwapPivotWith = j
				}
			}
			runes[i-1], runes[itemIndexToSwapPivotWith] = runes[itemIndexToSwapPivotWith], runes[i-1]
			Reverse(runes[i:])

			return true
		}
	}
	return false
}

// InsertionSort sorts items inplace. O(n^2) worst case, O(n) best case.
// Implemented instead of using sorting package purely for fun.
func InsertionSort(runes []rune) {
	for i := 0; i < len(runes); {
		for j := i; j > 0 && runes[j-1] > runes[j]; {
			runes[j], runes[j-1] = runes[j-1], runes[j]
			j--
		}
		i++
	}
}

// Reverse reverses a slice inplace. Kata uses go 1.20, i.e. slices package is not there yet.
func Reverse(runes []rune) {
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
}
