package lib

import "testing"

func Test_IsMatrixEqual(t *testing.T) {
	testCases := []struct {
		input1         [][]int
		input2         [][]int
		expectedOutput bool
	}{
		{
			CreateMatrix([]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3}, []int{2, 3, 1}, []int{3, 1, 2}, []int{3, 2, 1}),
			CreateMatrix([]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3}, []int{2, 3, 1}, []int{3, 1, 2}, []int{3, 2, 1}),
			true,
		},
		{
			CreateMatrix([]int{1, 2, 3}),
			CreateMatrix([]int{1, 2, 3}),
			true,
		},
		{
			CreateMatrix([]int{}),
			CreateMatrix([]int{}),
			true,
		},
		{
			CreateMatrix(),
			CreateMatrix(),
			true,
		},
		{
			CreateMatrix([]int{1, 2, 3}),
			CreateMatrix([]int{1, 2, 4}),
			false,
		},
		{
			CreateMatrix([]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3}, []int{2, 3, 4}, []int{3, 1, 2}, []int{3, 2, 1}),
			CreateMatrix([]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3}, []int{2, 3, 1}, []int{3, 1, 2}, []int{3, 2, 1}),
			false,
		},
	}

	for _, testCase := range testCases {
		output := IsMatrixEqual(testCase.input1, testCase.input2)

		if output != testCase.expectedOutput {
			t.Errorf("Ouput was incorrect, IsMatrixEqual()= %v, want: %v.", output, testCase.expectedOutput)
		}
	}

}
