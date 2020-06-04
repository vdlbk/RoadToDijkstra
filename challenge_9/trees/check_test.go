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

func TestIsValidBST(t *testing.T) {
	testCases := []struct {
		input          *TreeNode
		expectedOutput bool
	}{
		{GenerateTreeNode([]int{}), true},
		{GenerateTreeNode([]int{1}), true},
		{GenerateTreeNode([]int{2, 1, 3}), true},
		{GenerateTreeNode([]int{0, -1, 1}), true},
		{GenerateTreeNode([]int{10, 5, 15, 2, 7, 12, 17}), true},
		{GenerateTreeNode([]int{8, 4, 12, null, 6, null, 14}), true},
		{GenerateTreeNode([]int{8, 4, 12, 2, null, 10, null}), true},
		{GenerateTreeNode([]int{10, 5, 15, 2, 7, 12, 17, 1, 3, 6, 8, 11, 14, 16, 21}), true},
		{GenerateTreeNode([]int{10, 5, 15, 2, 7, 12, 17, 1, 3, 6, 8, 11, 14, 36, 21}), false},
		{GenerateTreeNode([]int{10, 8, 9, 7, 11}), false},
		{GenerateTreeNode([]int{5, 1, 4, null, null, 3, 6}), false},
		{GenerateTreeNode([]int{10, 5, 15, null, null, 6, 20}), false},
		{GenerateTreeNode([]int{0, 0, 0}), false},
		{GenerateTreeNode([]int{1, 0, 1}), false},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := isValidBST(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, isValidBST(%v) = %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestIsSymmetric(t *testing.T) {
	testCases := []struct {
		input          *TreeNode
		expectedOutput bool
	}{
		{GenerateTreeNode([]int{}), true},
		{GenerateTreeNode([]int{1}), true},
		{GenerateTreeNode([]int{1, 2, 2}), true},
		{GenerateTreeNode([]int{1, 2, 2, 3, 4, 4, 3}), true},
		{GenerateTreeNode([]int{0, 1, 1, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3}), true},
		{GenerateTreeNode([]int{0, 1, 1, 2, 18, 18, 2, null, 3, 4, 5, 5, 4, 3, null}), true},
		{GenerateTreeNode([]int{0, 1, 1, 2, 18, 18, 2, 3, null, 4, 5, 5, 4, null, 3}), true},
		{GenerateTreeNode([]int{1, 2, null}), false},
		{GenerateTreeNode([]int{1, null, 2}), false},
		{GenerateTreeNode([]int{1, 2, 3}), false},
		{GenerateTreeNode([]int{1, 2, 2, null, 3, null, 3}), false},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := isSymmetric(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, isSymmetric(%v) = %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}
