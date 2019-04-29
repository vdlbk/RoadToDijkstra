package challenge_2

import "testing"

func Test_Reverse(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"ABBA", "ABBA"},
		{"", ""},
		{"aa", "aa"},
		{"foobar", "raboof"},
		{"0123456", "6543210"},
	}

	for _, testCase := range testCases {
		output := Reverse(testCase.input)

		if output != testCase.expectedOutput {
			t.Errorf("Output was incorrect, Reverse()= %v, want: %v.", output, testCase.expectedOutput)
		}
	}

}

func Test_AltReverse(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput string
	}{
		{"ABBA", "ABBA"},
		{"", ""},
		{"aa", "aa"},
		{"foobar", "raboof"},
		{"0123456", "6543210"},
	}

	for _, testCase := range testCases {
		output := AltReverse(testCase.input)

		if output != testCase.expectedOutput {
			t.Errorf("Output was incorrect, AltReverse()= %v, want: %v.", output, testCase.expectedOutput)
		}
	}

}

func Test_MissingItem(t *testing.T) {
	type input struct {
		a1 []int
		a2 []int
	}
	testCases := []struct {
		input          input
		expectedOutput int
	}{
		{input{
			[]int{4, 12, 9, 5, 6},
			[]int{4, 9, 12, 6},
		}, 5},
		{input{
			[]int{4},
			[]int{},
		}, 4},
		{input{
			[]int{1, 2, 3},
			[]int{1, 2},
		}, 3},
		{input{
			[]int{7, 0, 23, 14, 8, 2, 4},
			[]int{0, 23, 14, 8, 2, 4},
		}, 7},
		{input{
			[]int{3, 5, 7, 9},
			[]int{3, 5, 7, 9},
		}, -1},
	}

	for _, testCase := range testCases {
		output := MissingItem(testCase.input.a1, testCase.input.a2)

		if output != testCase.expectedOutput {
			t.Errorf("Output was incorrect, MissingItem()= %v, want: %v.", output, testCase.expectedOutput)
		}
	}

}
