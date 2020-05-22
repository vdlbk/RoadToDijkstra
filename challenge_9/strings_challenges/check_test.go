package strings_challenges

import (
	"reflect"
	"testing"
)

func TestReverseString(t *testing.T) {
	testCases := []struct {
		input          []byte
		expectedOutput []byte
	}{
		{[]byte{byte('h'), byte('e'), byte('l'), byte('l'), byte('o')},
			[]byte{byte('o'), byte('l'), byte('l'), byte('e'), byte('h')}},
		{[]byte{byte('H'), byte('a'), byte('n'), byte('n'), byte('a'), byte('h')},
			[]byte{byte('h'), byte('a'), byte('n'), byte('n'), byte('a'), byte('H')}},
		{[]byte{},
			[]byte{}},
		{[]byte{byte('f')},
			[]byte{byte('f')}},
		{[]byte{byte('f'), byte('o'), byte('o')},
			[]byte{byte('o'), byte('o'), byte('f')}},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			reverseString(testCase.input)

			if !reflect.DeepEqual(testCase.input, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, reverseString() = %v, want: %v.", i,
					testCase.input,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestReverse(t *testing.T) {
	testCases := []struct {
		input          int
		expectedOutput int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{-21, -12},
		{12, 21},
		{10, 1},
		{0, 0},
		{-0, 0},
		{-1, -1},
		{1534236469, 0},
		{2147483648, 0},
		{-2147483648, 0},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := reverse(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, reverse(%v) = %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestFirstUniqChar(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput int
	}{
		{"leetcode", 0},     // l
		{"loveleetcode", 2}, // v
		{"foooof", -1},      // none
		{"x", 0},            // x
		{"", -1},            // none
		{"polop", 2},        // l
		{"aabbz", 4},        // z
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := firstUniqChar(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, firstUniqChar(%v) = %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestIsAnagram(t *testing.T) {
	testCases := []struct {
		input1         string
		input2         string
		expectedOutput bool
	}{
		{"anagram", "nagaram", true},
		{"ra", "ar", true},
		{"z", "z", true},
		{"", "", true},
		{"ofo", "ofo", true},
		{"poloolop", "poloolop", true},
		{"velo", "love", true},
		{"rat", "car", false},
		{"foo", "bar", false},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := isAnagram(testCase.input1, testCase.input2)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, isAnagram(%v, %v) = %v, want: %v.", i,
					testCase.input1,
					testCase.input2,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput bool
	}{
		{"A man, a plan, a canal: Panama", true},
		{"Hannah", true},
		{"", true},
		{"a", true},
		{"aa", true},

		{"race a car", false},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := isPalindrome(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, isPalindrome(%v) = %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestMyAtoi(t *testing.T) {
	testCases := []struct {
		input          string
		expectedOutput int
	}{
		{"42", 42},
		{"   -42", -42},
		{"-42", -42},
		{"--3", -0},
		{"42-", 42},
		{"- 42", 0},
		{"+3", 3},
		{"++3", 0},
		{"+-3", 0},
		{"4193 with words", 4193},
		{" -4193 with words", -4193},
		{"words and 987", 0},
		{"-91283472332", -2147483648},
		{"", 0},
		{"900001283472332", 2147483647},
		{"2147483648", 2147483647},
		{"-2147483648", -2147483648},
		{"0-1", 0},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := myAtoi(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, myAtoi(%v) = %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestStrStr(t *testing.T) {
	testCases := []struct {
		input          string
		needle         string
		expectedOutput int
	}{
		{"hello", "ll", 2},
		{"     hello", "ll", 7},
		{"     hello", "lz", -1},
		{"aaaaa", "bba", -1},
		{"aaaaa", "aa", 0},
		{"a", "aa", -1},
		{"a", "", 0},
		{"foo", "bar", -1},
		{"foo", "foo", 0},
		{"", "", 0},
		{"", "z", -1},
		{"mississippi", "issipi", -1},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := strStr(testCase.input, testCase.needle)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, strStr(%v) = %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}
