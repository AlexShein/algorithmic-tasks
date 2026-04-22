package integerpartitions

var partitionsCache = map[int][][]int{
	1: {{1}},
	2: {{2}, {1, 1}},
}

// Returns ordered integer partitions with largest number in the first position
// Not used in the kata, written for own reference
func partitions(n int) [][]int {
	if cachedPartitions, cacheHit := partitionsCache[n]; cacheHit {
		return cachedPartitions
	}
	result := [][]int{{n}}
	for k := n - 1; k > 0; k-- {
		for _, subPartition := range partitions(n - k) {
			if subPartition[0] <= k {
				result = append(result, append([]int{k}, subPartition...))
			}
		}
	}
	partitionsCache[n] = result
	return result
}

func partCombinations(n, m int, cache [][]int) int {
	if n == 0 {
		return 1 // No more "sum" left
	}
	if n < 0 {
		return 0 // Dead end, we can't reach the sum with integer that big
	}
	if m == 0 {
		return 0 // No parts allowed to fulfill the sum
	}
	if cachedPartitionsNumber := cache[n][m]; cachedPartitionsNumber != -1 {
		return cachedPartitionsNumber
	}
	// Main decision: use m or don't use m at all
	res := partCombinations(n-m, m, cache) + partCombinations(n, m-1, cache)
	cache[n][m] = res
	return res
}

// Partitions returns number of integer partitions for positive integers between 1 and 100 as per task
func Partitions(n int) int {

	cache := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		cache[i] = make([]int, n+1)
		for j := 0; j < n+1; j++ {
			cache[i][j] = -1 // Indicates that the value has not been cached
		}
	}

	return partCombinations(n, n, cache)
}
