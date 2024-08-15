package validanagram

import "testing"

func TestIsAnagram(t *testing.T) {
	for _, test := range []struct {
		Name     string
		S        string
		T        string
		Expected bool
	}{
		{
			Name:     "anagram and nagaram should return true",
			S:        "anagram",
			T:        "nagaram",
			Expected: true,
		},
		{
			Name:     "rat and car should return false",
			S:        "rat",
			T:        "car",
			Expected: false,
		},
		{
			Name:     "aa and bb should return false",
			S:        "aa",
			T:        "bb",
			Expected: false,
		},
		{
			Name:     "bad and dab should return true",
			S:        "bad",
			T:        "dab",
			Expected: true,
		},
		{
			Name:     "bark and bar should return false",
			S:        "bark",
			T:        "bar",
			Expected: false,
		},
	} {
		t.Run(test.Name, func(t *testing.T) {
			assertEqual(t, IsAnagram(test.S, test.T), test.Expected)
		})
	}
}

func assertEqual(t testing.TB, actual, expected bool) {
	if actual != expected {
		t.Errorf("FAILED: Expected - %v Actual - %v", expected, actual)
	}
}
