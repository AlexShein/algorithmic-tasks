package permutations

// Permutations returns a list of all possible symbol permutations for a given input, without duplicates
func Permutations(s string) []string {

	if len(s) == 1 {
		return []string{s}
	}

	result := []string{}
	uniquePermutations := map[string]bool{} // Used for quick duplicates check
	for position, character := range s {
		for _, lowerOrderPermutation := range Permutations(s[:position] + s[position+1:]) {
			permutation := string(character) + lowerOrderPermutation
			if !uniquePermutations[permutation] {
				result = append(result, permutation)
				uniquePermutations[permutation] = true
			}
		}
	}
	return result
}
