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
		{GenerateListFromSlice([]int{4, 5, 1, 9}), 5, GenerateListFromSlice([]int{4, 1, 9})},
		{GenerateListFromSlice([]int{4, 5, 1, 9}), 4, GenerateListFromSlice([]int{5, 1, 9})},
		{GenerateListFromSlice([]int{4, 5, 1, 9}), 1, GenerateListFromSlice([]int{4, 5, 9})},
		{GenerateListFromSlice([]int{1, 9}), 1, GenerateListFromSlice([]int{9})},
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

func TestRemoveNthFromEnd(t *testing.T) {
	testCases := []struct {
		input          *ListNode
		node           int
		expectedOutput *ListNode
	}{
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5}), 1, GenerateListFromSlice([]int{1, 2, 3, 4})},
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5}), 2, GenerateListFromSlice([]int{1, 2, 3, 5})},
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5}), 3, GenerateListFromSlice([]int{1, 2, 4, 5})},
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5}), 4, GenerateListFromSlice([]int{1, 3, 4, 5})},
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5}), 5, GenerateListFromSlice([]int{2, 3, 4, 5})},
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5}), 0, GenerateListFromSlice([]int{1, 2, 3, 4, 5})},
		{GenerateListFromSlice([]int{1}), 1, nil},
		{GenerateListFromSlice([]int{1}), 2, nil},
		{GenerateListFromSlice([]int{1, 2}), 2, GenerateListFromSlice([]int{2})},
		{GenerateListFromSlice([]int{3, 3, 3}), 2, GenerateListFromSlice([]int{3, 3})},
		{nil, 2, nil},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := removeNthFromEnd(testCase.input, testCase.node)

			if !reflect.DeepEqual(output, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, removeNthFromEnd(%v, %v) = %v, want: %v.", i,
					testCase.input,
					testCase.node,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestReverseList(t *testing.T) {
	testCases := []struct {
		input          *ListNode
		expectedOutput *ListNode
	}{
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5}), GenerateListFromSlice([]int{5, 4, 3, 2, 1})},
		{GenerateListFromSlice([]int{1, 2, 3, 4}), GenerateListFromSlice([]int{4, 3, 2, 1})},
		{GenerateListFromSlice([]int{1, 2, 3}), GenerateListFromSlice([]int{3, 2, 1})},
		{GenerateListFromSlice([]int{1}), GenerateListFromSlice([]int{1})},
		{GenerateListFromSlice([]int{1, 1}), GenerateListFromSlice([]int{1, 1})},
		{GenerateListFromSlice([]int{1, 2}), GenerateListFromSlice([]int{2, 1})},
		{GenerateListFromSlice([]int{}), GenerateListFromSlice([]int{})},
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}),
			GenerateListFromSlice([]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1})},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := reverseList(testCase.input)

			if !reflect.DeepEqual(output, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, reverseList(%v) = %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		input          *ListNode
		input2         *ListNode
		expectedOutput *ListNode
	}{
		{
			GenerateListFromSlice([]int{1, 2, 4}),
			GenerateListFromSlice([]int{1, 3, 4}),
			GenerateListFromSlice([]int{1, 1, 2, 3, 4, 4}),
		},
		{
			GenerateListFromSlice([]int{1}),
			GenerateListFromSlice([]int{0}),
			GenerateListFromSlice([]int{0, 1}),
		},
		{
			GenerateListFromSlice([]int{0}),
			GenerateListFromSlice([]int{1}),
			GenerateListFromSlice([]int{0, 1}),
		},
		{
			GenerateListFromSlice([]int{1, 2, 3}),
			GenerateListFromSlice([]int{4, 5, 6}),
			GenerateListFromSlice([]int{1, 2, 3, 4, 5, 6}),
		},
		{
			GenerateListFromSlice([]int{6, 10, 43}),
			GenerateListFromSlice([]int{1, 2}),
			GenerateListFromSlice([]int{1, 2, 6, 10, 43}),
		},
		{
			GenerateListFromSlice([]int{6, 10, 43}),
			GenerateListFromSlice([]int{1, 2, 69, 121}),
			GenerateListFromSlice([]int{1, 2, 6, 10, 43, 69, 121}),
		},
		{
			GenerateListFromSlice([]int{}),
			GenerateListFromSlice([]int{}),
			GenerateListFromSlice([]int{}),
		},
		{
			GenerateListFromSlice([]int{7}),
			GenerateListFromSlice([]int{}),
			GenerateListFromSlice([]int{7}),
		},
		{
			GenerateListFromSlice([]int{}),
			GenerateListFromSlice([]int{8}),
			GenerateListFromSlice([]int{8}),
		},
		{
			GenerateListFromSlice([]int{4, 2, 1}),
			GenerateListFromSlice([]int{1, 3, 4}),
			GenerateListFromSlice([]int{1, 3, 4, 4, 2, 1}),
		},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := mergeTwoLists(testCase.input, testCase.input2)

			if !reflect.DeepEqual(output, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, mergeTwoLists(%v, %v) = %v, want: %v.", i,
					testCase.input,
					testCase.input2,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}