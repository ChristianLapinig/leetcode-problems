package validmountainarray

// 941. Valid Mountain Array
// Description: https://leetcode.com/problems/valid-mountain-array/description/
// Solution: https://leetcode.com/problems/valid-mountain-array/editorial/
func ValidMountainArray(arr []int) bool {
	// constraint not met
	if len(arr) < 3 {
		return false
	}

	i := 0
	n := len(arr)

	// go up the mountain
	for i < n-1 && arr[i] < arr[i+1] {
		i++
	}

	// peak not found
	if i == 0 || i == n-1 {
		return false
	}

	// go down the mountain
	for i+1 < n && arr[i] > arr[i+1] {
		i++
	}

	return i == len(arr)-1
}
