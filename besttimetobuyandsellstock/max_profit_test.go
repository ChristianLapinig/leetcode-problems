package besttimetobuyandsellstock

import "testing"

func TestMaxProfit(t *testing.T) {
	for _, test := range []struct {
		Name     string
		Prices   []int
		Expected int
	}{
		{
			Name:     "returns 5",
			Prices:   []int{7, 1, 5, 3, 6, 4},
			Expected: 5,
		},
		{
			Name:     "returns 0",
			Prices:   []int{7, 6, 4, 3, 1},
			Expected: 0,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			actual := maxProfit(test.Prices)
			assertEquals(t, actual, test.Expected)
		})
	}
}

func assertEquals(t testing.TB, actual, expected int) {
	if actual != expected {
		t.Errorf("FAILED: Expected - %q - Actual - %q", expected, actual)
	}
}
