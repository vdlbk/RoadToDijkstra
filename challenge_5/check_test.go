package challenge_5

import (
	"testing"
)

func Test_ValidateStackSequences(t *testing.T) {
	type input struct {
		pushed []int
		popped []int
	}
	testCases := []struct {
		id             int
		input          input
		expectedOutput bool
	}{
		{0, input{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}}, true},
		{1, input{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}}, false},
		{2, input{[]int{0}, []int{0}}, true},
		{3, input{[]int{}, []int{}}, true},
		{4, input{[]int{2, 3}, []int{2, 3}}, true},
		{5, input{[]int{2, 3}, []int{3, 2}}, true},
		{6, input{[]int{0, 2, 1}, []int{0, 1, 2}}, true},
		{7, input{[]int{2, 1, 0}, []int{2, 1, 0}}, true},
	}

	for _, testCase := range testCases {
		output := ValidateStackSequences(testCase.input.pushed, testCase.input.popped)
		//output := ValidateStackSequencesBis(testCase.input.pushed, testCase.input.popped)

		if output != testCase.expectedOutput {
			t.Errorf("[%d] Ouput was incorrect, ValidateStackSequences()= %v, want: %v.", testCase.id, output, testCase.expectedOutput)
		}
	}
}