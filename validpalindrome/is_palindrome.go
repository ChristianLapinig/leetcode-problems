package validpalindrome

import (
	"regexp"
	"strings"
)

func IsPalindrome(s string) bool {
	// strip all non-alphanumeric characters
	regex := regexp.MustCompile("[^a-zA-Z0-9]+")
	s = strings.ToLower(regex.ReplaceAllString(s, ""))

	left := 0
	right := len(s) - 1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}
