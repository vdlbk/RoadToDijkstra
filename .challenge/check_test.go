package challenge

import (
	"testing"
)

func TestFoobar(t *testing.T) {
	testCases := []struct {
		input          int
		expectedOutput int
	}{
		{0, -1},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := Foobar(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, Foobar(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}
