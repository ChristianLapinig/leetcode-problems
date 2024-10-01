package doubleexists

import (
	"testing"
)

func TestCheckIfExists(t *testing.T) {
	for _, test := range []struct {
		Name     string
		Arr      []int
		Expected bool
	}{
		{
			Name:     "should return true",
			Arr:      []int{10, 2, 5, 3},
			Expected: true,
		},
		{
			Name:     "should return false",
			Arr:      []int{3, 1, 7, 11},
			Expected: false,
		},
		{
			Name:     "should return true",
			Arr:      []int{5, 2, 10, 3},
			Expected: true,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			actual := CheckIfExists(test.Arr)

			if actual != test.Expected {
				t.Errorf("FAILED - Expected: %v - Actual: %v", test.Expected, actual)
			}
		})
	}
}
