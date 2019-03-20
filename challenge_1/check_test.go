package main

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
		ouput := LongestCommonSubsequence(testCase.input.s1, testCase.input.s2)

		if ouput != testCase.expectedOutput {
			t.Errorf("Ouput was incorrect, LongestCommonSubsequence()= %v, want: %v.", ouput, testCase.expectedOutput)
		}
	}

}
