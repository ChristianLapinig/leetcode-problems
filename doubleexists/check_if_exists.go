package doubleexists

// 1346. Check If N and Its Double Exist
// Description: https://leetcode.com/problems/check-if-n-and-its-double-exist/
// Editorial/Solution: https://leetcode.com/problems/check-if-n-and-its-double-exist/solutions
func CheckIfExists(arr []int) bool {
	// "have I seen you before" problem
	table := make(map[int]int)

	for _, val := range arr {
		if table[val*2] != 0 || (val%2 == 0 && table[val/2] != 0) {
			return true
		}
		table[val] = 1
	}

	return false
}
