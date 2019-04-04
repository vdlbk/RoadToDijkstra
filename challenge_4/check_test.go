package main

import (
	"reflect"
	"testing"
)

func Test_DailyTemperatures(t *testing.T) {
	testCases := []struct {
		id             int
		input          []int
		expectedOutput []int
	}{
		{0, []int{50, 60}, []int{1, 0}},
		{1, []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{2, []int{40, 39, 38, 37}, []int{0, 0, 0, 0}},
		{3, []int{60, 59, 58, 57, 57, 57, 57, 63}, []int{7, 6, 5, 4, 3, 2, 1, 0}},
	}

	for _, testCase := range testCases {
		output := DailyTemperatures(testCase.input)

		if !reflect.DeepEqual(output, testCase.expectedOutput) {
			t.Errorf("[%d] Ouput was incorrect, DailyTemperatures()= %v, want: %v.", testCase.id, output, testCase.expectedOutput)
		}
	}
}
