package mergesortedarray

/*
* 88. Merge Sorted Array
* Description: https://leetcode.com/problems/merge-sorted-array
* Editorial: https://leetcode.com/problems/merge-sorted-array/editorial
 */
func merge(nums1 []int, m int, nums2 []int, n int) {
	/*
	* using example nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3,
	* we can see that the we can go backwards through nums1 and nums2 to
	* to fill out the larger elements from n to the end of nums1.
	*
	* we can compare nums1[i] to nums2[j] to determine which element goes
	* at the end of nums1.
	*
	* we would need three pointers to make this work. one to iterate over
	* the valid elements of nums1, a second one to iterate over nums2, and
	* a third to iterate over the full "capacity" of nums1, m + n
	 */
	p1 := m - 1 // pointer to the elements of nums1
	p2 := n - 1 // pointer to the elements of nums2

	/*
	* the size of each array will not always be the same. if that happens
	* then we've reached the end of our sorting
	 */
	for i := m + n - 1; i >= 0; i-- {
		// done if we have processed all of the elements in nums2
		if p2 < 0 {
			break
		}

		// ensure that we are within the bounds of nums1
		if p1 >= 0 && nums1[p1] > nums2[p2] {
			nums1[i] = nums1[p1]
			p1--
		} else {
			nums1[i] = nums2[p2]
			p2--
		}
	}
}
