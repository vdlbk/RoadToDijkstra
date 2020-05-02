package arrays

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{[]int{1, 2}, 2},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1, 1, 1}, 1},
		{[]int{-3, -1, -1, 0, 0, 0, 0, 0, 2}, 4},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := removeDuplicates(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, removeDuplicates(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{0, 10}, 10},
		{[]int{1}, 0},
		{[]int{3, 6, 3, 6, 3, 6, 3, 6}, 12},
		{[]int{4, 1, 3, 10}, 9},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := maxProfit(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, maxProfit(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	testCases := []struct {
		input          []int
		k              int
		expectedOutput []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 2, []int{6, 7, 1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4}, 2, []int{3, 4, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6}, 2, []int{5, 6, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6}, 3, []int{4, 5, 6, 1, 2, 3}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{-1, -100, 3, 99}, 3, []int{-100, 3, 99, -1}},
		{[]int{1, 2}, 2, []int{1, 2}},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			raw := testCase.input
			rotate(testCase.input, testCase.k)

			if !reflect.DeepEqual(testCase.input, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, rotate(%v, %v)= %v, want: %v.", i,
					raw,
					testCase.k,
					testCase.input,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 13}, true},
		{[]int{}, false},
		{[]int{10000000}, false},
	}

	for i, testCase := range testCases {
		t.Run("containsDuplicate", func(t *testing.T) {
			output := containsDuplicate(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, containsDuplicate(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
		{[]int{-1, -1, -2, -4, -2, -7, -4}, -7},
		{[]int{-1, -10, -2, -9, -3, -8, -4, -7, -5, -6, -6, -5, -7, -4, -8, -3, -9, -2, -10, -1, 0}, 0},
	}

	for i, testCase := range testCases {
		t.Run("singleNumber", func(t *testing.T) {
			output := singleNumber(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, singleNumber(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	testCases := []struct {
		input1         []int
		input2         []int
		expectedOutput []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
		{[]int{1, 2}, []int{3, 4}, []int{}},
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{1}, []int{1}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{[]int{-1, -2, -3}, []int{-2, -3, -4}, []int{-2, -3}},
	}

	for i, testCase := range testCases {
		t.Run("intersect", func(t *testing.T) {
			output := intersect(testCase.input1, testCase.input2)

			if !reflect.DeepEqual(output, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, intersect(%v, %v)= %v, want: %v.", i,
					testCase.input1,
					testCase.input2,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{1, 0, 9}, []int{1, 1, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{1}, []int{2}},
		{[]int{9}, []int{1, 0}},
	}

	for i, testCase := range testCases {
		t.Run("plusOne", func(t *testing.T) {
			output := plusOne(testCase.input)

			if !reflect.DeepEqual(output, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, plusOne(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}
