package containsduplicate

/*
* 217. Contains Duplicate
* Description: https://leetcode.com/problems/contains-duplicate/description/
* Editorial: https://leetcode.com/problems/contains-duplicate/editorial/
 */
func ContainsDuplicate(nums []int) bool {
	// the easiest way to determine if a duplicate exists
	// is to have a map or set where we check if a value
	// already exists. if a value exists in the set, return
	// true. if not, add nums[i] to the set
	existingValues := make(map[int]bool) // turns this into O(n) space

	// O(n) time complexity
	for _, num := range nums {
		if _, ok := existingValues[num]; ok {
			return true
		}
		existingValues[num] = true
	}

	// no duplicates found
	return false
}
