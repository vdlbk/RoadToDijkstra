package searching

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	testCases := []struct {
		input          []int
		m              int
		input2         []int
		n              int
		expectedOutput []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1, 1, 1, 0, 0, 0}, 3, []int{1, 1, 1}, 3, []int{1, 1, 1, 1, 1, 1}},
		{[]int{1, 0}, 1, []int{2}, 1, []int{1, 2}},
		{[]int{1, 2, 3}, 1, []int{2}, 1, []int{1, 2, 3}},
		{[]int{0, 2, 0, 0}, 2, []int{-1, 6}, 2, []int{-1, 0, 2, 6}},
		{[]int{0, 14, 0, 0}, 2, []int{-1, 6}, 2, []int{-1, 0, 6, 14}},
		{[]int{0, 14, 13, 0}, 2, []int{-1, 6}, 2, []int{-1, 0, 6, 14}},
		{[]int{}, 0, []int{}, 0, []int{}},
		{[]int{9, 18, 27}, 3, []int{}, 0, []int{9, 18, 27}},
		{[]int{0}, 0, []int{1}, 1, []int{1}},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			merge(testCase.input, testCase.m, testCase.input2, testCase.n)

			if !reflect.DeepEqual(testCase.input, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, merge()= %v, want: %v.", i,
					testCase.input,
					testCase.expectedOutput,
				)
			}
		})
	}
}
