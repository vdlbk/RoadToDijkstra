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
