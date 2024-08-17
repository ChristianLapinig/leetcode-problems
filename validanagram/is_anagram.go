package validanagram

/*
* 242. Valid Anagram
* Description: https://leetcode.com/problems/valid-anagram/
* Editorial/Solution: https://leetcode.com/problems/valid-anagram/editorial/
 */
func IsAnagram(s, t string) bool {
	// anagrams must have the same length
	if len(s) != len(t) {
		return false
	}

	// an anagram is a word or phrase formed by rearranging the letters of
	// a different word or phrase using all of the characters of the original
	// word exactly once. this means the count of each character in t is the same
	// as s, or s[char] - t[char] == 0.

	// we can use a table that keeps track of the
	// number of instances of each character of each string
	var table [26]int

	// we can determine which character count to increment by s[i] - 'a' (ASCII math)
	// we can use 's' to define the count of each character, and 't' to check if
	// the anagram is valid according to s[char] - t[char] == 0
	for i := 0; i < len(s); i++ {
		// if we have an anagram, table[char] == 0
		table[s[i]-'a']++
		table[t[i]-'a']--
	}

	// count cannot be greater than 0
	for _, count := range table {
		if count > 0 {
			return false
		}
	}

	// time complexity: O(n)
	// space complexity: O(1) - table is constant
	return true
}
