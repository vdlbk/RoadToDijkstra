package challenge

import (
	"testing"
)

func TestCountRecurrentCharacter(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"ABCA", "A"},
		{"BCABA", "B"},
		{"flovol", "l"},
		{"ABC", ""},
		{"AAAAA", "A"},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := CountRecurrentCharacter(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, CountRecurrentCharacter(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		input          string
		input2         string
		expectedOutput bool
	}{
		{"ABCD", "DABC", true},
		{"AABB", "BBAA", true},
		{"aABB", "BBaA", true},
		{"", "", true},
		{"A", "A", true},
		{"@#", "#@", true},
		{"AAAB", "AAB", false},
		{"AAB", "AAAB", false},
		{"AAAA", "BBBB", false},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := IsAnagram(testCase.input, testCase.input2)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, IsAnagram(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}
