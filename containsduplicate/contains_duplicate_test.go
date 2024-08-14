package containsduplicate

import "testing"

func TestContainsDuplicate(t *testing.T) {
	for _, test := range []struct {
		Name     string
		Nums     []int
		Expected bool
	}{
		{
			Name:     "[1, 2, 3, 1] should return true",
			Nums:     []int{1, 2, 3, 1},
			Expected: true,
		},
		{
			Name:     "[1, 2, 3, 4] should return false",
			Nums:     []int{1, 2, 3, 4},
			Expected: false,
		},
		{
			Name:     "[1, 1, 1, 3, 3, 4, 3, 2, 4, 2] should return true",
			Nums:     []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			Expected: true,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			assertEqual(t, ContainsDuplicate(test.Nums), test.Expected)
		})
	}
}

func assertEqual(t testing.TB, actual, expected bool) {
	if actual != expected {
		t.Errorf("FAILED: Expected - %v - Actual - %v", actual, expected)
	}
}
