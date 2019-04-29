package challenge_1

import "testing"

func Test_LongestCommonSubsequence(t *testing.T) {
	type input struct {
		s1 string
		s2 string
	}
	testCases := []struct {
		input          input
		expectedOutput string
	}{
		{input{"ABAZDC", "BACBAD"}, "ABAD"},
		{input{"AGGTAB", "GXTXAYB"}, "GTAB"},
		{input{"aaaa", "aa"}, "aa"},
		{input{"", "foobar"}, ""},
		{input{"ABBA", "ABCABA"}, "ABBA"},
	}

	for _, testCase := range testCases {
		output := LongestCommonSubsequence(testCase.input.s1, testCase.input.s2)

		if output != testCase.expectedOutput {
			t.Errorf("Output was incorrect, LongestCommonSubsequence()= %v, want: %v.", output, testCase.expectedOutput)
		}
	}

}
