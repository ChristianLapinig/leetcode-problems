package validmountainarray

import "testing"

func TestValidMountainArray(t *testing.T) {
	t.Run("[2, 1] should return false", func(t *testing.T) {
		actual := ValidMountainArray([]int{2, 1})
		assertFalse(t, actual)
	})

	t.Run("[3, 5, 5] should return false", func(t *testing.T) {
		actual := ValidMountainArray([]int{3, 5, 5})
		assertFalse(t, actual)
	})

	t.Run("[0, 3, 2, 1] should return true", func(t *testing.T) {
		actual := ValidMountainArray([]int{0, 3, 2, 1})
		assertTrue(t, actual)
	})

	t.Run("[0, 2, 1], should return true", func(t *testing.T) {
		actual := ValidMountainArray([]int{0, 2, 1})
		assertTrue(t, actual)
	})

	t.Run("[0, 1, 2, 3, 4, 5, 6, 7, 8, 9] should return false", func(t *testing.T) {
		actual := ValidMountainArray([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		assertFalse(t, actual)
	})
}

func assertTrue(t testing.TB, actual bool) {
	if !actual {
		t.Errorf("FAILED: Expected true - Actual %v", actual)
	}
}

func assertFalse(t testing.TB, actual bool) {
	if actual {
		t.Errorf("FAILED: Expected false - Actual %v", actual)
	}
}
