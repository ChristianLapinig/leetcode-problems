package removeelement

import (
	"slices"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	for _, test := range []struct {
		Name     string
		Nums     []int
		Val      int
		Expected []int
	}{
		{
			Name:     "should return [2, 2, _, _]",
			Nums:     []int{3, 2, 2, 3},
			Val:      3,
			Expected: []int{2, 2},
		},
		{
			Name:     "should return [0,1,4,0,3,_,_,_]",
			Nums:     []int{0, 1, 2, 2, 3, 0, 4, 2},
			Val:      2,
			Expected: []int{0, 0, 1, 3, 4},
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			k, actual := RemoveElement(test.Nums, test.Val)
			judge(t, actual, test.Expected, k)
		})
	}
}

func judge(t testing.TB, actual, expected []int, k int) {
	if k != len(expected) {
		t.Errorf("FAILED - Actual length does not match the expected length")
		return
	}

	slices.Sort(actual)

	for i := 0; i < k; i++ {
		if actual[i] != expected[i] {
			t.Errorf("FAILED - Actual: %v - Expected: %v", actual, expected)
			return
		}
	}
}
