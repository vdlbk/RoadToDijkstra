package arrays

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput int
	}{
		{[]int{1, 1, 2}, 2},
		{[]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5},
		{[]int{1, 2}, 2},
		{[]int{}, 0},
		{[]int{1}, 1},
		{[]int{1, 1, 1, 1}, 1},
		{[]int{-3, -1, -1, 0, 0, 0, 0, 0, 2}, 4},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := removeDuplicates(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, removeDuplicates(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{0, 10}, 10},
		{[]int{1}, 0},
		{[]int{3, 6, 3, 6, 3, 6, 3, 6}, 12},
		{[]int{4, 1, 3, 10}, 9},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := maxProfit(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, maxProfit(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestRotate(t *testing.T) {
	testCases := []struct {
		input          []int
		k              int
		expectedOutput []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 2, []int{6, 7, 1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4}, 2, []int{3, 4, 1, 2}},
		{[]int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{[]int{1, 2, 3, 4, 5, 6}, 2, []int{5, 6, 1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5, 6}, 3, []int{4, 5, 6, 1, 2, 3}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		{[]int{-1, -100, 3, 99}, 3, []int{-100, 3, 99, -1}},
		{[]int{1, 2}, 2, []int{1, 2}},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			raw := testCase.input
			rotate(testCase.input, testCase.k)

			if !reflect.DeepEqual(testCase.input, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, rotate(%v, %v)= %v, want: %v.", i,
					raw,
					testCase.k,
					testCase.input,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestContainsDuplicate(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 13}, true},
		{[]int{}, false},
		{[]int{10000000}, false},
	}

	for i, testCase := range testCases {
		t.Run("containsDuplicate", func(t *testing.T) {
			output := containsDuplicate(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, containsDuplicate(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestSingleNumber(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
		{[]int{-1, -1, -2, -4, -2, -7, -4}, -7},
		{[]int{-1, -10, -2, -9, -3, -8, -4, -7, -5, -6, -6, -5, -7, -4, -8, -3, -9, -2, -10, -1, 0}, 0},
	}

	for i, testCase := range testCases {
		t.Run("singleNumber", func(t *testing.T) {
			output := singleNumber(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, singleNumber(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	testCases := []struct {
		input1         []int
		input2         []int
		expectedOutput []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
		{[]int{1, 2}, []int{3, 4}, []int{}},
		{[]int{}, []int{}, []int{}},
		{[]int{1}, []int{1}, []int{1}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{[]int{-1, -2, -3}, []int{-2, -3, -4}, []int{-2, -3}},
	}

	for i, testCase := range testCases {
		t.Run("intersect", func(t *testing.T) {
			output := intersect(testCase.input1, testCase.input2)

			if !reflect.DeepEqual(output, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, intersect(%v, %v)= %v, want: %v.", i,
					testCase.input1,
					testCase.input2,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestPlusOne(t *testing.T) {
	testCases := []struct {
		input          []int
		expectedOutput []int
	}{
		{[]int{1, 2, 3}, []int{1, 2, 4}},
		{[]int{4, 3, 2, 1}, []int{4, 3, 2, 2}},
		{[]int{1, 0, 9}, []int{1, 1, 0}},
		{[]int{9, 9, 9}, []int{1, 0, 0, 0}},
		{[]int{1}, []int{2}},
		{[]int{9}, []int{1, 0}},
	}

	for i, testCase := range testCases {
		t.Run("plusOne", func(t *testing.T) {
			output := plusOne(testCase.input)

			if !reflect.DeepEqual(output, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, plusOne(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestMoveZero(t *testing.T) {
	testCases := []struct {
		raw            []int
		input          []int
		expectedOutput []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 1}, []int{0, 1}, []int{1, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 0, 0, 0}, []int{0, 0, 0, 0}, []int{0, 0, 0, 0}},
		{[]int{1, 0, 0, 0}, []int{1, 0, 0, 0}, []int{1, 0, 0, 0}},
		{[]int{1}, []int{1}, []int{1}},
		{[]int{}, []int{}, []int{}},
		{[]int{
			2, 4, 5043, 4343, 0, 363, 345, 238, 0, 67, 33, 4, 0, 6, 2, 7, 7, 7, 0, 0, 0, 0, 0, 1, 0, 3, 0,
		}, []int{
			2, 4, 5043, 4343, 0, 363, 345, 238, 0, 67, 33, 4, 0, 6, 2, 7, 7, 7, 0, 0, 0, 0, 0, 1, 0, 3, 0,
		}, []int{
			2, 4, 5043, 4343, 363, 345, 238, 67, 33, 4, 6, 2, 7, 7, 7, 1, 3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		}},
	}

	for i, testCase := range testCases {
		t.Run("moveZeroes", func(t *testing.T) {
			moveZeroes(testCase.input)

			if !reflect.DeepEqual(testCase.input, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, moveZeroes(%v)= %v, want: %v.", i,
					testCase.raw,
					testCase.input,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestTwoSum(t *testing.T) {
	testCases := []struct {
		input          []int
		k              int
		expectedOutput []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{1, 2, 0, 3}, 2, []int{1, 2}},
		{[]int{1, 2, 3, 4, 5}, 8, []int{2, 4}},
		{[]int{1, 2, 3, 4, 5, 3}, 8, []int{2, 4}},
		{[]int{1, 2}, 3, []int{0, 1}},
		{[]int{1, 2, 4, 4, 2, 1}, 8, []int{2, 3}},
	}

	for i, testCase := range testCases {
		t.Run("", func(t *testing.T) {
			output := twoSum(testCase.input, testCase.k)

			if !reflect.DeepEqual(output, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, rotate(%v, %v)= %v, want: %v.", i,
					testCase.input,
					testCase.k,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestIsValidSudoku(t *testing.T) {
	testCases := []struct {
		input          [][]byte
		expectedOutput bool
	}{
		{[][]byte{
			{byte('5'), byte('3'), byte('.'), byte('.'), byte('7'), byte('.'), byte('.'), byte('.'), byte('.')},
			{byte('6'), byte('.'), byte('.'), byte('1'), byte('9'), byte('5'), byte('.'), byte('.'), byte('.')},
			{byte('.'), byte('9'), byte('8'), byte('.'), byte('.'), byte('.'), byte('.'), byte('6'), byte('.')},
			{byte('8'), byte('.'), byte('.'), byte('.'), byte('6'), byte('.'), byte('.'), byte('.'), byte('3')},
			{byte('4'), byte('.'), byte('.'), byte('8'), byte('.'), byte('3'), byte('.'), byte('.'), byte('1')},
			{byte('7'), byte('.'), byte('.'), byte('.'), byte('2'), byte('.'), byte('.'), byte('.'), byte('6')},
			{byte('.'), byte('6'), byte('.'), byte('.'), byte('.'), byte('.'), byte('2'), byte('8'), byte('.')},
			{byte('.'), byte('.'), byte('.'), byte('4'), byte('1'), byte('9'), byte('.'), byte('.'), byte('5')},
			{byte('.'), byte('.'), byte('.'), byte('.'), byte('8'), byte('.'), byte('.'), byte('7'), byte('9')},
		},
			true},
		{[][]byte{
			{byte('8'), byte('3'), byte('.'), byte('.'), byte('7'), byte('.'), byte('.'), byte('.'), byte('.')},
			{byte('6'), byte('.'), byte('.'), byte('1'), byte('9'), byte('5'), byte('.'), byte('.'), byte('.')},
			{byte('.'), byte('9'), byte('8'), byte('.'), byte('.'), byte('.'), byte('.'), byte('6'), byte('.')},
			{byte('8'), byte('.'), byte('.'), byte('.'), byte('6'), byte('.'), byte('.'), byte('.'), byte('3')},
			{byte('4'), byte('.'), byte('.'), byte('8'), byte('.'), byte('3'), byte('.'), byte('.'), byte('1')},
			{byte('7'), byte('.'), byte('.'), byte('.'), byte('2'), byte('.'), byte('.'), byte('.'), byte('6')},
			{byte('.'), byte('6'), byte('.'), byte('.'), byte('.'), byte('.'), byte('2'), byte('8'), byte('.')},
			{byte('.'), byte('.'), byte('.'), byte('4'), byte('1'), byte('9'), byte('.'), byte('.'), byte('5')},
			{byte('.'), byte('.'), byte('.'), byte('.'), byte('8'), byte('.'), byte('.'), byte('7'), byte('9')},
		},
			false},
		{[][]byte{
			{byte('5'), byte('3'), byte('.'), byte('.'), byte('5'), byte('.'), byte('.'), byte('.'), byte('.')},
			{byte('6'), byte('.'), byte('.'), byte('1'), byte('9'), byte('5'), byte('.'), byte('.'), byte('.')},
			{byte('.'), byte('9'), byte('8'), byte('.'), byte('.'), byte('.'), byte('.'), byte('6'), byte('.')},
			{byte('8'), byte('.'), byte('.'), byte('.'), byte('6'), byte('.'), byte('.'), byte('.'), byte('3')},
			{byte('4'), byte('.'), byte('.'), byte('8'), byte('.'), byte('3'), byte('.'), byte('.'), byte('1')},
			{byte('7'), byte('.'), byte('.'), byte('.'), byte('2'), byte('.'), byte('.'), byte('.'), byte('6')},
			{byte('.'), byte('6'), byte('.'), byte('.'), byte('.'), byte('.'), byte('2'), byte('8'), byte('.')},
			{byte('.'), byte('.'), byte('.'), byte('4'), byte('1'), byte('9'), byte('.'), byte('.'), byte('5')},
			{byte('.'), byte('.'), byte('.'), byte('.'), byte('8'), byte('.'), byte('.'), byte('7'), byte('9')},
		},
			false},
		{[][]byte{
			{byte('5'), byte('3'), byte('.'), byte('.'), byte('7'), byte('.'), byte('.'), byte('.'), byte('.')},
			{byte('6'), byte('.'), byte('.'), byte('1'), byte('9'), byte('5'), byte('.'), byte('.'), byte('.')},
			{byte('.'), byte('5'), byte('8'), byte('.'), byte('.'), byte('.'), byte('.'), byte('6'), byte('.')},
			{byte('8'), byte('.'), byte('.'), byte('.'), byte('6'), byte('.'), byte('.'), byte('.'), byte('3')},
			{byte('4'), byte('.'), byte('.'), byte('8'), byte('.'), byte('3'), byte('.'), byte('.'), byte('1')},
			{byte('7'), byte('.'), byte('.'), byte('.'), byte('2'), byte('.'), byte('.'), byte('.'), byte('6')},
			{byte('.'), byte('6'), byte('.'), byte('.'), byte('.'), byte('.'), byte('2'), byte('8'), byte('.')},
			{byte('.'), byte('.'), byte('.'), byte('4'), byte('1'), byte('9'), byte('.'), byte('.'), byte('5')},
			{byte('.'), byte('.'), byte('.'), byte('.'), byte('8'), byte('.'), byte('.'), byte('7'), byte('9')},
		},
			false},
	}

	for i, testCase := range testCases {
		t.Run("isValidSudoku", func(t *testing.T) {
			output := isValidSudoku(testCase.input)

			if output != testCase.expectedOutput {
				t.Errorf("[%d] Output was incorrect, isValidSudoku(%v)= %v, want: %v.", i,
					testCase.input,
					output,
					testCase.expectedOutput,
				)
			}
		})
	}
}

func TestRotateMatrix(t *testing.T) {
	testCases := []struct {
		input          [][]int
		expectedOutput [][]int
	}{
		{[][]int{
			{1, 2, 3},
			{4, 5, 6},
			{7, 8, 9},
		}, [][]int{
			{7, 4, 1},
			{8, 5, 2},
			{9, 6, 3},
		}},
		{[][]int{
			{5, 1, 9, 11},
			{2, 4, 8, 10},
			{13, 3, 6, 7},
			{15, 14, 12, 16},
		}, [][]int{
			{15, 13, 2, 5},
			{14, 3, 4, 1},
			{12, 6, 8, 9},
			{16, 7, 10, 11},
		}},
		{[][]int{
			{5, 6},
			{5, 6},
		}, [][]int{
			{5, 5},
			{6, 6},
		}},
		{[][]int{
			{3},
		}, [][]int{
			{3},
		}},
		{[][]int{
			{},
		}, [][]int{
			{},
		}},
	}

	for i, testCase := range testCases {
		t.Run("rotateMatrix", func(t *testing.T) {
			rotateMatrix(testCase.input)

			if !reflect.DeepEqual(testCase.input, testCase.expectedOutput) {
				t.Errorf("[%d] Output was incorrect, rotateMatrix()= %v, want: %v.", i,
					testCase.input,
					testCase.expectedOutput,
				)
			}
		})
	}
}
