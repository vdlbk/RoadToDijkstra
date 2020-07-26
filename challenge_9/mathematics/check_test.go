package mathematics

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		input          int
		expectedOutput []string
	}{
		{0, nil},
		{-1, nil},
		{1, []string{"1"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			result := fizzBuzz(testCase.input)

			if !reflect.DeepEqual(result, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, fizzBuzz(%v)= %v, want: %v.", i,
					testCase.input,
					result,
					testCase.expectedOutput,
				)
			}
		})
	}
}
