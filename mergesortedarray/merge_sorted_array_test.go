package mergesortedarray

import "testing"

func TestMergeSortedArray(t *testing.T) {
	for _, test := range []struct {
		Name     string
		Nums1    []int
		Nums2    []int
		Expected []int
		M        int
		N        int
	}{
		{
			Name:     "returns [1,2,2,3,5,6]",
			Nums1:    []int{1, 2, 3, 0, 0, 0},
			Nums2:    []int{2, 5, 6},
			Expected: []int{1, 2, 2, 3, 5, 6},
			M:        3,
			N:        3,
		},
		{
			Name:     "returns [1]",
			Nums1:    []int{1},
			Nums2:    []int{},
			Expected: []int{1},
			M:        1,
			N:        0,
		},
		{
			Name:     "returns [1]",
			Nums1:    []int{0},
			Nums2:    []int{1},
			Expected: []int{1},
			M:        0,
			N:        1,
		},
		{
			Name:  "returns an empty array nums1 and nums2 are empty",
			Nums1: []int{0},
			Nums2: []int{},
			M:     0,
			N:     0,
		},
		{
			Name:     "returns [1,2,3,4,5,6,7]",
			Nums1:    []int{1, 7, 0, 0, 0, 0, 0},
			Nums2:    []int{2, 3, 4, 5, 6},
			Expected: []int{1, 2, 3, 4, 5, 6, 7},
			M:        2,
			N:        5,
		},
		{
			Name:     "returns [10, 30, 50, 60, 90, 100]",
			Nums1:    []int{30, 50, 90, 100, 0, 0},
			Nums2:    []int{10, 60},
			Expected: []int{10, 30, 50, 60, 90, 100},
			M:        4,
			N:        2,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			merge(test.Nums1, test.M, test.Nums2, test.N)
			assertArrayEquals(t, test.Nums1, test.Expected)
		})
	}
}

func assertArrayEquals(t testing.TB, actual, expected []int) {
	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("FAILED: Expected %q - Actual %q", expected, actual)
		}
	}
}
