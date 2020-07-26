package challenge

import (
	"testing"
)

func TestFoobar(t *testing.T) {
	testCases := []struct {
		input          int
		expectedOutput int
	}{
		{0, 10},
		{3, 10},
		{4, 10},
		{27, 270},
		{30, 300},
		{1000, 10000},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := Solution(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, Foobar(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestFoobar2(t *testing.T) {
	testCases := []struct {
		input          int
		input2         int
		expectedOutput int
	}{
		{3, 7, 3},
		{0, 0, 0},
		{0, 1, 0},
		{895137, 147963, 25},
		{100000000, 100000000, 20},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := Solution2(testCase.input, testCase.input2)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, Solution2(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestFoobar3(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput int
	}{
		{"", 0},
		{"baaaaa", 1},
		{"baaabbaabbba", 2},
		{"baabab", 0},
		{"baaabbab", 1},
		{"baaaaaaaaaaaaa", 4},
		{"aaaaaaaaaaaaaaa", 5},
		{"babababababaababababa", 0},
		{"aaabbbaaabbbaaabbb", 6},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := Solution3(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, Solution3(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}
