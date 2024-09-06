package duplicatezeros

// 1089. Duplicate Zeros
// Description: https://leetcode.com/problems/duplicate-zeros/description/
// Editorial/Solution: https://leetcode.com/problems/duplicate-zeros/editorial/
func DuplicateZeros(arr []int) {
	// TODO: review the optimal two-pass solution
	// naive solution
	// this will be O(n^2) but currently going through arrays 101 and
	// learning about insertions
	i := 0

	for i < len(arr) {
		if arr[i] == 0 {
			// shift elements to the right
			for j := len(arr) - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			// move pointer to next non-zero
			i += 2
		} else {
			i++
		}
	}
}
