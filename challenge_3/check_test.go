package challenge_3

import (
	"testing"

	"github.com/vdlbk/RoadToDijkstra/challenge_3/lib"
)

func Test_Permute(t *testing.T) {
	testCases := []struct {
		id             int
		input          []int
		expectedOutput [][]int
	}{
		{0, []int{1, 2, 3}, lib.CreateMatrix([]int{1, 2, 3}, []int{1, 3, 2}, []int{2, 1, 3}, []int{2, 3, 1}, []int{3, 1, 2}, []int{3, 2, 1})},
		{1, []int{1, 2}, lib.CreateMatrix([]int{1, 2}, []int{2, 1})},
		{2, []int{1}, lib.CreateMatrix([]int{1})},
		{3, []int{}, lib.CreateMatrix()},
		{4, []int{1, 2, 3, 4}, lib.CreateMatrix(
			[]int{1, 2, 3, 4}, []int{1, 2, 4, 3}, []int{1, 3, 2, 4}, []int{1, 3, 4, 2}, []int{1, 4, 2, 3}, []int{1, 4, 3, 2},
			[]int{2, 1, 3, 4}, []int{2, 1, 4, 3}, []int{2, 3, 1, 4}, []int{2, 3, 4, 1}, []int{2, 4, 1, 3}, []int{2, 4, 3, 1}, 
			[]int{3, 1, 2, 4}, []int{3, 1, 4, 2}, []int{3, 2, 1, 4}, []int{3, 2, 4, 1}, []int{3, 4, 1, 2}, []int{3, 4, 2, 1}, 
			[]int{4, 1, 2, 3}, []int{4, 1, 3, 2}, []int{4, 2, 1, 3}, []int{4, 2, 3, 1}, []int{4, 3, 1, 2}, []int{4, 3, 2, 1},
			)},
	}

	for _, testCase := range testCases {
		output := Permute(testCase.input)

		if !lib.IsMatrixEqual(output, testCase.expectedOutput) {
			t.Errorf("[%d] Output was incorrect, Permute()= %v, want: %v.", testCase.id, output, testCase.expectedOutput)
		}
	}
}