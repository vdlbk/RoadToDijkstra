package challenge_7

import (
	"testing"
)

func Test_SingleNumber(t *testing.T) {
	testCases := []struct {
		id             int
		input          []int
		expectedOutput int
	}{
		{0, []int{2, 2, 1}, 1},
		{1, []int{4, 1, 2, 1, 2}, 4},
		{2, []int{3, 3}, -1},
		{3, []int{4, 1, 2, 6, 3, 5, 1, 2, 3, 5, 4}, 6},

	}



	for _, testCase := range testCases {
		output := SingleNumber(testCase.input)

		if output != testCase.expectedOutput {
			t.Errorf("[%d] Ouput was incorrect, SingleNumber()= %v, want: %v.", testCase.id, output, testCase.expectedOutput)
		}
	}
}

func Test_SingleNumberBis(t *testing.T) {
	testCases := []struct {
		id             int
		input          []int
		expectedOutput int
	}{
		{0, []int{2, 2, 1}, 1},
		{1, []int{4, 1, 2, 1, 2}, 4},
		{2, []int{3, 3}, -1},
		{3, []int{4, 1, 2, 6, 3, 5, 1, 2, 3, 5, 4}, 6},
	}

	for _, testCase := range testCases {
		output := SingleNumberBis(testCase.input)

		if output != testCase.expectedOutput {
			t.Errorf("[%d] Ouput was incorrect, SingleNumberBis()= %v, want: %v.", testCase.id, output, testCase.expectedOutput)
		}
	}
}