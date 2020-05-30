package trees

import (
	"testing"
)

func TestMaxDepth(t *testing.T) {
	testCases := []struct {
		input          *TreeNode
		expectedOutput int
	}{
		{GenerateTreeNode([]int{}), 0},
		{GenerateTreeNode([]int{3}), 1},
		{GenerateTreeNode([]int{3, 9, 20}), 2},
		{GenerateTreeNode([]int{3, 9}), 2},
		{GenerateTreeNode([]int{3, 9, 20, null, null, 15, 7}), 3},
		{GenerateTreeNode([]int{0, 1, 2, 3, 5, 4, 6, 7, 9, 11, 13, 8, 10, 12, 14}), 4},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := maxDepth(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, maxDepth(%v) = %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}
