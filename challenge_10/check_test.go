package challenge

import (
	"math"
	"math/rand"
	"testing"
)

func TestBinaryGap(t *testing.T) {
	testCases := []struct {
		input          int
		expectedOutput int
	}{
		{9, 2},
		{529, 4},
		{20, 1},
		{1041, 5},
		{math.MaxInt64, 0},
		{561892, 3},
		{74901729, 4},
		{1376796946, 5},

		{15, 0},
		{32, 0},
		{0, 0},
		{-5, 0},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := binaryGap(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, binaryGap(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput int
	}{
		{[]int{3, 1, 2, 4, 3}, 1},
		{[]int{1, 2}, 1},
		{[]int{5, 5, 5, 5, 5}, 5},
		{[]int{5, 5, 5, 5, 5, 5}, 0},
		{[]int{-3, -2, -1, -1, -2, -3}, 0},
		{[]int{-3, -2, -1, 1, 2, 3}, 6},
		{[]int{-1000, 1000}, 2000},
		{generateRandomSlice(100000), 2000},

		{[]int{}, 0},
		{[]int{42}, 0},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := diff(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, diff(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func generateRandomSlice(n int) []int {
	var slice []int
	for i := 0; i < n; i++ {
		slice = append(slice, rand.Intn(1000))
	}
	return slice
}
