package challenge_6

import (
	"fmt"
	"testing"
)

func Test_CountDigitOne(t *testing.T) {
	testCases := []struct {
		id             int
		input          int
		expectedOutput int
	}{
		{0, 13, 6},
		{1, 1, 1},
		{2, 2, 1},
		{3, 0, 0},
		{4, -1, 0},
		{5, 10000000, 7000001},
		{6, 824883294, 767944060},
	}

	for _, testCase := range testCases {
		output := SolutionCountDigitOne(testCase.input)
		//output := CountDigitOne(testCase.input)
		//output := SillyCountDigitOne(testCase.input)

		if output != testCase.expectedOutput {
			t.Errorf("[%d] Ouput was incorrect, CountDigitOne()= %v, want: %v.", testCase.id, output, testCase.expectedOutput)
		}
	}
}

func Test_main(t *testing.T) {
	x := 10
	y := 10

	fmt.Println(x / y, x % y)
}