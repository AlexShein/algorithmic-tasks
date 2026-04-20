package baker

func Cakes(recipe, available map[string]int) int {
	maxCakes := int(^uint(0) >> 1) // Max possible int value, just to be safe
	for ingredient, required := range recipe {
		available := available[ingredient]
		if divRes := available / required; divRes < maxCakes {
			maxCakes = divRes
		}
	}
	return maxCakes
}
