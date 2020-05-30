package trees

import (
	"reflect"
	"testing"
)

func TestGenerateTreeNode(t *testing.T) {
	null := -1 << 31
	testCases := []struct {
		input          []int
		expectedOutput *TreeNode
	}{
		{[]int{}, nil},
		{[]int{2, 1, 3}, &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}},
		{[]int{3, 9, 20, null, null, 15, 7},
			&TreeNode{Val: 3,
				Left: &TreeNode{Val: 9},
				Right: &TreeNode{Val: 20,
					Left:  &TreeNode{Val: 15},
					Right: &TreeNode{Val: 7}}},
		},
		{[]int{0, 1, 2, 3, null, 4, 6, 7, 9, null, null, 8, 10, 12, 14},
			&TreeNode{Val: 0,
				Left: &TreeNode{Val: 1,
					Left: &TreeNode{Val: 3,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 9}},
				},
				Right: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 4,
						Left:  &TreeNode{Val: 8},
						Right: &TreeNode{Val: 10}},
					Right: &TreeNode{Val: 6,
						Left:  &TreeNode{Val: 12},
						Right: &TreeNode{Val: 14}}},
			},
		},
		{[]int{0, 1, 2, 3, 5, 4, 6, 7, 9, 11, 13, 8, 10, 12, 14},
			&TreeNode{Val: 0,
				Left: &TreeNode{Val: 1,
					Left: &TreeNode{Val: 3,
						Left:  &TreeNode{Val: 7},
						Right: &TreeNode{Val: 9}},
					Right: &TreeNode{Val: 5,
						Left:  &TreeNode{Val: 11},
						Right: &TreeNode{Val: 13}}},
				Right: &TreeNode{Val: 2,
					Left: &TreeNode{Val: 4,
						Left:  &TreeNode{Val: 8},
						Right: &TreeNode{Val: 10}},
					Right: &TreeNode{Val: 6,
						Left:  &TreeNode{Val: 12},
						Right: &TreeNode{Val: 14}}},
			},
		},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := GenerateTreeNode(testCase.input)

			if !reflect.DeepEqual(output, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, GenerateTreeNode(%v) = %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}
