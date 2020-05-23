package linked_list

import (
	"reflect"
	"testing"
)

func TestDeleteNode(t *testing.T) {
	testCases := []struct {
		input          *ListNode
		node           int
		expectedOutput *ListNode
	}{
		{generateListFromSlice([]int{4, 5, 1, 9}), 5, generateListFromSlice([]int{4, 1, 9})},
		{generateListFromSlice([]int{4, 5, 1, 9}), 4, generateListFromSlice([]int{5, 1, 9})},
		{generateListFromSlice([]int{4, 5, 1, 9}), 1, generateListFromSlice([]int{4, 5, 9})},
		{generateListFromSlice([]int{1, 9}), 1, generateListFromSlice([]int{9})},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			deleteNode(testCase.input, testCase.node)

			if !reflect.DeepEqual(testCase.input, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, deleteNode(%v) = %v, want: %v.", i,
					testCase.node,
					testCase.input,
					testCase.expectedOutput,
				)
			}
		})
	}
}
