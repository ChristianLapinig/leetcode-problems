package validpalindrome

import "testing"

func TestIsPalindrome(t *testing.T) {
	for _, test := range []struct {
		Name     string
		S        string
		Expected bool
	}{
		{
			Name:     "A man, a plan, a canal: Panama should return true",
			S:        "A man, a plan, a canal: Panama",
			Expected: true,
		},
		{
			Name:     "race car should return true",
			S:        "race car",
			Expected: true,
		},
		{
			Name:     "race a car should return false",
			S:        "race a car",
			Expected: false,
		},
		{
			Name:     "empty string should return true",
			S:        " ",
			Expected: true,
		},
		{
			Name:     "A should return true",
			S:        "A",
			Expected: true,
		},
		{
			Name:     "AbCdE should return false",
			S:        "AbCdE",
			Expected: false,
		},
		{
			Name:     "AbCdA should return false",
			S:        "AbCdA",
			Expected: false,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			actual := IsPalindrome(test.S)
			assertEqual(t, actual, test.Expected)
		})
	}
}

func assertEqual(t testing.TB, actual, expected bool) {
	if actual != expected {
		t.Errorf("FAILED: Expected - %v - Actual - %v", actual, expected)
	}
}
