package removeelement

// 27. Remove Element
// Description: https://leetcode.com/problems/remove-element/description/
// Editorial/Solution: https://leetcode.com/problems/remove-element/editorial/
func RemoveElement(arr []int, val int) (int, []int) {
	i := 0
	k := len(arr)

	for i < k {
		// relative order does not matter here
		// move occurrence of val to a place in the array we will ignore
		if arr[i] == val {
			arr[i] = arr[k-1]
			k--
		} else {
			// non duplicate found
			i++
		}
	}

	return k, arr
}
