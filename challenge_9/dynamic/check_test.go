package dynamic

import (
	"testing"
)

func TestClimbStairs(t *testing.T) {
	testCases := []struct {
		input          int
		expectedOutput int
	}{
		{-1, 0},
		{0, 0},
		{1, 1},
		{2, 2},
		{3, 3},
		{35, 14930352},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			result := climbStairs(testCase.input)

			if result != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, climbStairs(%v)= %v, want: %v.", i,
					testCase.input,
					result,
					testCase.expectedOutput,
				)
			}
		})
	}
}
