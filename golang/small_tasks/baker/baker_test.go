package baker

import "testing"

func TestCakes(t *testing.T) {
	testCases := []struct {
		recipe         map[string]int
		ingredients    map[string]int
		expectedResult int
	}{{
		ingredients:    map[string]int{"flour": 1200, "sugar": 1200, "eggs": 5, "milk": 200},
		recipe:         map[string]int{"flour": 500, "sugar": 200, "eggs": 1},
		expectedResult: 2,
	}, {
		ingredients:    map[string]int{"apples": 3, "flour": 300, "sugar": 150, "milk": 100, "oil": 100},
		recipe:         map[string]int{"sugar": 500, "flour": 2000, "milk": 2000},
		expectedResult: 0,
	},
	}

	for _, testCase := range testCases {

		if res := Cakes(testCase.recipe, testCase.ingredients); res != testCase.expectedResult {
			t.Errorf(
				"Test failed.\nIngredients:\n%v\nRecipe:\n%v.\nExpected %d, got %d",
				testCase.ingredients, testCase.recipe,
				testCase.expectedResult, res)
		}
	}
}
