package challenge

import (
	"testing"
)

func TestFoobar(t *testing.T) {
	testCases := []struct {
		input          []string
		expectedOutput bool
	}{
		{[]string{"atlanta", "anchorage", "el paso", "oakland", "denver", "ridgewood"}, true},
		{[]string{"barcelona", "madrid", "roma"}, false},
		{[]string{"barcelona", "madrid", "barcelona"}, false},
		{[]string{"barcelona", "adrib", "barcelona"}, false},
		{[]string{"barcelona", "", "arcelona"}, false},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := Geography(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, Geography(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestBot(t *testing.T) {
	testCases := []struct {
		input          []string
		botInput       []string
		expectedOutput bool
	}{
		// player 1 is out of word
		{[]string{"atlanta"},
			[]string{"adrid"}, true},

		// player 1 does not answer with a correct word
		{[]string{"atlanta", "andorra"},
			[]string{"adrid"}, true},

		// player 1 is out of word
		{[]string{"atlanta", "denver"},
			[]string{"adrid", "ridgewood"}, true},

		// player 1 is trying to use multiple times the same word
		{[]string{"atlanta", "denver", "denver"},
			[]string{"adrid", "ridgewood"}, true},

		// bot does not answer with a correct city
		{[]string{"atlanta"},
			[]string{"barcelona", "madrid"}, false},

		// bot is out of word
		{[]string{"atlanta", "denver", "dallas"},
			[]string{"adrid", "ridgewood"}, false},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := Game(testCase.input, testCase.botInput)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, Geography(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}
