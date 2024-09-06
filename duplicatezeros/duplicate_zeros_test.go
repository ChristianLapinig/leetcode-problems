package duplicatezeros

import "testing"

func TestDuplicateZero(t *testing.T) {
	for _, test := range []struct {
		Name     string
		Nums     []int
		Expected []int
	}{
		{
			Name:     "should return [1,0,0,2,3,0,0,4]",
			Nums:     []int{1, 0, 2, 3, 0, 4, 5, 0},
			Expected: []int{1, 0, 0, 2, 3, 0, 0, 4},
		},
		{
			Name:     "should return [1,2,3] since there are not 0s",
			Nums:     []int{1, 2, 3},
			Expected: []int{1, 2, 3},
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			DuplicateZeros(test.Nums)
			assertArrayEquals(t, test.Nums, test.Expected)
		})
	}
}

func assertArrayEquals(t testing.TB, actual, expected []int) {
	if len(actual) != len(expected) {
		t.Errorf("FAILED - Actual length does not match the expected length")
		return
	}

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("FAILED - Actual: %q - Expected: %q", actual, expected)
			return
		}
	}
}
