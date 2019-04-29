package challenge_4

import (
	"reflect"
	"testing"
)

func Test_DailyTemperatures(t *testing.T) {
	testCases := []struct {
		id             int
		input          []int
		expectedOutput []int
	}{
		{0, []int{50, 60}, []int{1, 0}},
		{1, []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{2, []int{40, 39, 38, 37}, []int{0, 0, 0, 0}},
		{3, []int{60, 59, 58, 57, 57, 57, 57, 63}, []int{7, 6, 5, 4, 3, 2, 1, 0}},
		{4, []int{59,57,34,34,88,85,92,78,30,57,39,94,74,58,35,52,75,70,88,45,94,78,72,55,75,43,60,76,88,33,79,98,44,97,78,47,62,59,72,31,93,45,33,48,75,96,87,76,69,54,80,38,90,32,91,90,53,88,51,92,44,76,48,68,66,83,89,88,44,47,48,49,77,63,80,95,71,80,93,39,35,94,33,47,43,98,31,73,32,68,60,87,54,51,54,31,84,68,96,40,35,38,54,50,53,98,84,42,46,56,86,58,49,72,69,96,73,94,82,62,60,92,48,96,85,75,76,49,41,43,71,76,44,55,62,68,43,69,60,38,93,60,94,37,96,55,77,82,98,96,76,66,99,64,89,63,91,59,92,36,61,86,49,87,89,76,70,75,33,98,99,87,34,37,86,92,63,71,48,74,52,42,55,94,74,45,67,73,36,43,68,36,63,96,53,99,38,78,83,51,46,47,67,75,93,99,54,95,92,34,34,89,42,65,47,57,91,77,87,87,92,31,60,71,48,62,66,57,45,98,79,50,64,49,96,43,88,58,70,56,72,80,41,48,88,91,69,61,59,73,95,80,83,40,67,58,97,50,36,54,78,40,89,74,82,95,49,87,37,93,38,37,92,64,34,49,49,89,76,61,64,92,95,34,92,57,65,82,54,34,53,52,79,97,97,93,97,32,72,84,42,56,45,72,60,47,36,84,86,44,33,43,62,42,90,83,78,68,62,61,58,90,80,52,67,77,68,58,87,50,90,39,55,90,30,48,85,50,31,93,35,51,85,97,60,60,79,51,58,67,42,58,62,67,77,92,36,68,79,96,60,66,57,54,79,85,75,86,59,38,31,95,63,68,34,96,45,65,42,62,81,38,41,43,37,91,52,80,30,93,90,56,55,45,76,93,92,80,96,43,52,34,33,87,66,42,66,41,81,85,72,34,52,45,40,99,31,76,30,84,73,82,67,97,79,60,51,68,100,45,60,98,76,84,47,78,45,33,81,100,67,97,52,90,35,70,59,92,58,73,88,73,63,43,97,90,31,69,79,39,59,43,40,33,46,56,69,47,100,40,73,91,32,90,92,60,75,73,76,50,70,99,60,31,88,61,93,91,88,33,32,31,58,58,86,51,98,59,73,57,74,96,35,53,43,55,51,45,98,53,67,47,67,64,57,93,46,84,37,85,44,89,70,58,98,69,66,96,50,85,57,73,92,34,66,58,68,63,38,51,31,66,51,54,77,79,89,40,84,62,51,81,43,80,43,32,74,99,57,82,87,71,95,43,87,42,78,43,79,40,41,61,75,47,88,90,66,98,59,54,56,55,71,81,95,69,88,91,91,53,100,86,93,37,43,48,91,80,73,62,42,79,66,81,77,35,64,57,99,30,78,46,73,30,95,79,88,91,76,41,98,40,41,66,39,31,49,67,66,53,34,49,48,50,84,44,39,62,72,46,39,62,46,94,62,51,96,77,43,66,61,65,94,49,66,73,53,45,84,30,94,52,67,93,73,94,60,37,84,82,56,67,61,46,46,65,58,56,52,69,41,92,33,43,36,65,68,44,74,72,89,36,52,57,52,30,100,63,66,46,49,42,50,45,64,36,71,92,81,64,56,82,63,46,88,79,93,83,69,48,36,48,85,96,58,46,65,67,37,46,62,37,35,78,30,45,56,60,56,88,73,52,88,40,69,61,55,33,84,76,86,60,71,48,94,33,89,40,37,41,92,94,94,93,61,82,77,62,79,51,71,59,81,71,45,39,48,66,88,69,36,100,91,40,95,54,85,92,88,67,31,78,42,39,70,58,81,98,63,51,85,38,59,40,75,90,90,91,81,30,43,87,84,38,34,50,64,90,99,43,60,68,80,34,53,49,64,32,96,54,42,64,48,100,71,63,40,69,69,48,79,37,59,62,87,49,82,48,64,41,84,51,86,53,79,36,61,55,57,40,68,76,42,90,49,87,33,86,60,41,63,84,94,47,83,77,58,47,93,97,45,39,52,98,82,100,91,65,61,69,46,83,38,32,66,34,47,85,30,78,65,91,75,43,74,92,46,65,53,58,64,69,77,88,74,85,98,31,67,56,67,77,65,73,63,55,92,97,60,88,50,92,57,39,44,62,40,39,34,32,64,53,61,31,59,48,46,78,69,41,74,31,63,42,87,68,65,86,88,49,69,51,48,73,58,60,34,72,67,96,100,46,71,62,94,66,34,37,53,45,41,66,36,48,52,63,52,39,36,93,53,52,75,55,91,72,88,89,38,69},
			[]int{4,3,2,1,2,1,5,4,1,2,1,20,4,3,1,1,2,1,2,1,11,7,2,1,3,1,1,1,3,1,1,121,1,52,6,1,2,1,2,1,5,2,1,1,1,40,6,3,2,1,2,1,2,1,5,4,1,2,1,16,1,4,1,2,1,1,9,8,1,1,1,1,2,1,1,10,1,1,3,2,1,4,1,2,1,67,1,4,1,2,1,7,4,1,2,1,2,1,7,3,1,1,3,1,1,47,4,1,1,1,5,2,1,2,1,33,1,6,3,2,1,2,1,25,16,1,14,3,1,1,1,9,1,1,1,2,1,3,2,1,2,1,2,1,4,1,1,1,4,3,2,1,276,1,2,1,2,1,11,1,1,2,1,1,5,4,1,2,1,1,258,4,1,1,1,8,1,2,1,4,2,1,1,10,9,1,1,6,1,1,3,1,1,2,1,233,1,1,6,3,1,1,1,1,1,223,1,22,21,2,1,5,1,3,1,1,4,1,2,1,9,1,1,6,1,1,3,2,1,186,4,1,2,1,22,1,9,1,2,1,1,3,1,1,1,5,3,2,1,1,6,1,4,1,2,1,159,2,1,1,2,1,3,1,1,28,1,2,1,13,2,1,10,4,1,2,1,4,3,1,1,1,11,1,9,1,1,6,4,1,2,1,1,122,121,1,119,1,1,9,1,2,1,4,3,2,1,1,6,3,1,1,2,1,25,6,5,4,3,2,1,18,6,1,1,3,2,1,2,1,9,1,1,6,1,1,3,2,1,4,1,1,1,72,2,1,9,1,1,5,1,1,1,1,1,4,1,1,1,56,1,3,2,1,1,2,1,4,3,2,1,4,1,2,1,40,1,3,1,1,5,1,1,2,1,4,1,2,1,9,5,3,2,1,1,3,2,1,17,1,3,2,1,12,4,1,2,1,1,6,5,1,3,2,1,13,1,2,1,4,1,2,1,5,4,2,1,1,0,1,1,8,1,6,1,3,2,1,1,0,1,27,1,4,1,2,1,7,1,1,4,3,2,1,14,13,1,1,10,1,6,3,2,1,1,1,2,1,0,1,1,3,1,1,7,1,2,1,3,1,1,109,2,1,2,1,10,9,8,3,2,1,2,1,2,1,61,1,2,1,1,7,1,2,1,3,2,1,49,1,5,1,3,2,1,9,1,2,1,2,1,3,2,1,33,2,1,30,1,3,1,1,25,1,2,1,8,4,1,2,1,3,1,1,1,1,11,1,9,2,1,6,1,4,2,1,1,33,1,1,2,1,15,1,10,1,2,1,6,1,1,1,2,1,1,2,1,13,4,1,2,1,1,1,6,1,1,3,2,1,0,1,16,1,1,1,12,6,3,2,1,2,1,5,4,1,2,1,88,1,4,1,2,1,6,1,1,3,2,1,76,1,1,4,2,1,1,7,6,5,1,2,1,1,9,2,1,1,5,2,1,2,1,3,2,1,50,5,1,3,1,1,44,1,1,3,2,1,2,1,36,1,1,2,1,31,2,1,13,12,1,8,3,2,1,4,3,2,1,2,1,15,1,2,1,1,2,1,2,1,6,1,1,3,2,1,0,1,8,1,2,1,2,1,2,1,1,9,3,2,1,3,2,1,2,1,7,5,4,3,1,1,1,58,2,1,1,6,1,1,3,2,1,6,1,1,1,2,1,15,2,1,12,1,4,3,2,1,2,1,4,1,2,1,27,1,4,2,1,1,1,20,19,18,1,13,2,1,4,1,2,1,6,5,2,1,1,1,3,2,1,0,2,1,13,1,1,10,9,2,1,5,2,1,2,1,1,21,2,1,5,1,2,1,1,2,1,11,3,1,1,6,5,2,1,1,1,1,15,1,1,1,6,1,2,1,2,1,5,2,1,2,1,0,6,2,1,3,2,1,4,1,1,1,19,1,4,1,2,1,2,1,11,1,9,1,4,1,2,1,1,2,1,9,1,7,1,5,2,1,1,1,7,1,4,3,2,1,1,4,2,1,1,2,1,0,19,2,1,2,1,6,2,1,3,1,1,4,1,2,1,4,3,1,1,11,1,4,1,1,1,1,1,3,1,1,54,1,3,1,1,5,1,3,2,1,1,43,1,2,1,38,3,1,1,5,4,3,2,1,7,1,5,1,3,2,1,7,2,1,4,1,2,1,4,2,1,1,11,1,3,2,1,6,1,2,1,2,1,1,0,1,2,1,0,14,1,1,3,2,1,8,1,1,1,4,3,2,1,0,2,1,2,1,0,1,1,0,1,0}},
	}

	for _, testCase := range testCases {
		output := DailyTemperatures(testCase.input)

		if !reflect.DeepEqual(output, testCase.expectedOutput) {
			t.Errorf("[%d] Output was incorrect, DailyTemperatures()= %v, want: %v.", testCase.id, output, testCase.expectedOutput)
		}
	}
}

func Test_DailyTemperaturesBis(t *testing.T) {
	testCases := []struct {
		id             int
		input          []int
		expectedOutput []int
	}{
		{0, []int{50, 60}, []int{1, 0}},
		{1, []int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{2, []int{40, 39, 38, 37}, []int{0, 0, 0, 0}},
		{3, []int{60, 59, 58, 57, 57, 57, 57, 63}, []int{7, 6, 5, 4, 3, 2, 1, 0}},
		{4, []int{59,57,34,34,88,85,92,78,30,57,39,94,74,58,35,52,75,70,88,45,94,78,72,55,75,43,60,76,88,33,79,98,44,97,78,47,62,59,72,31,93,45,33,48,75,96,87,76,69,54,80,38,90,32,91,90,53,88,51,92,44,76,48,68,66,83,89,88,44,47,48,49,77,63,80,95,71,80,93,39,35,94,33,47,43,98,31,73,32,68,60,87,54,51,54,31,84,68,96,40,35,38,54,50,53,98,84,42,46,56,86,58,49,72,69,96,73,94,82,62,60,92,48,96,85,75,76,49,41,43,71,76,44,55,62,68,43,69,60,38,93,60,94,37,96,55,77,82,98,96,76,66,99,64,89,63,91,59,92,36,61,86,49,87,89,76,70,75,33,98,99,87,34,37,86,92,63,71,48,74,52,42,55,94,74,45,67,73,36,43,68,36,63,96,53,99,38,78,83,51,46,47,67,75,93,99,54,95,92,34,34,89,42,65,47,57,91,77,87,87,92,31,60,71,48,62,66,57,45,98,79,50,64,49,96,43,88,58,70,56,72,80,41,48,88,91,69,61,59,73,95,80,83,40,67,58,97,50,36,54,78,40,89,74,82,95,49,87,37,93,38,37,92,64,34,49,49,89,76,61,64,92,95,34,92,57,65,82,54,34,53,52,79,97,97,93,97,32,72,84,42,56,45,72,60,47,36,84,86,44,33,43,62,42,90,83,78,68,62,61,58,90,80,52,67,77,68,58,87,50,90,39,55,90,30,48,85,50,31,93,35,51,85,97,60,60,79,51,58,67,42,58,62,67,77,92,36,68,79,96,60,66,57,54,79,85,75,86,59,38,31,95,63,68,34,96,45,65,42,62,81,38,41,43,37,91,52,80,30,93,90,56,55,45,76,93,92,80,96,43,52,34,33,87,66,42,66,41,81,85,72,34,52,45,40,99,31,76,30,84,73,82,67,97,79,60,51,68,100,45,60,98,76,84,47,78,45,33,81,100,67,97,52,90,35,70,59,92,58,73,88,73,63,43,97,90,31,69,79,39,59,43,40,33,46,56,69,47,100,40,73,91,32,90,92,60,75,73,76,50,70,99,60,31,88,61,93,91,88,33,32,31,58,58,86,51,98,59,73,57,74,96,35,53,43,55,51,45,98,53,67,47,67,64,57,93,46,84,37,85,44,89,70,58,98,69,66,96,50,85,57,73,92,34,66,58,68,63,38,51,31,66,51,54,77,79,89,40,84,62,51,81,43,80,43,32,74,99,57,82,87,71,95,43,87,42,78,43,79,40,41,61,75,47,88,90,66,98,59,54,56,55,71,81,95,69,88,91,91,53,100,86,93,37,43,48,91,80,73,62,42,79,66,81,77,35,64,57,99,30,78,46,73,30,95,79,88,91,76,41,98,40,41,66,39,31,49,67,66,53,34,49,48,50,84,44,39,62,72,46,39,62,46,94,62,51,96,77,43,66,61,65,94,49,66,73,53,45,84,30,94,52,67,93,73,94,60,37,84,82,56,67,61,46,46,65,58,56,52,69,41,92,33,43,36,65,68,44,74,72,89,36,52,57,52,30,100,63,66,46,49,42,50,45,64,36,71,92,81,64,56,82,63,46,88,79,93,83,69,48,36,48,85,96,58,46,65,67,37,46,62,37,35,78,30,45,56,60,56,88,73,52,88,40,69,61,55,33,84,76,86,60,71,48,94,33,89,40,37,41,92,94,94,93,61,82,77,62,79,51,71,59,81,71,45,39,48,66,88,69,36,100,91,40,95,54,85,92,88,67,31,78,42,39,70,58,81,98,63,51,85,38,59,40,75,90,90,91,81,30,43,87,84,38,34,50,64,90,99,43,60,68,80,34,53,49,64,32,96,54,42,64,48,100,71,63,40,69,69,48,79,37,59,62,87,49,82,48,64,41,84,51,86,53,79,36,61,55,57,40,68,76,42,90,49,87,33,86,60,41,63,84,94,47,83,77,58,47,93,97,45,39,52,98,82,100,91,65,61,69,46,83,38,32,66,34,47,85,30,78,65,91,75,43,74,92,46,65,53,58,64,69,77,88,74,85,98,31,67,56,67,77,65,73,63,55,92,97,60,88,50,92,57,39,44,62,40,39,34,32,64,53,61,31,59,48,46,78,69,41,74,31,63,42,87,68,65,86,88,49,69,51,48,73,58,60,34,72,67,96,100,46,71,62,94,66,34,37,53,45,41,66,36,48,52,63,52,39,36,93,53,52,75,55,91,72,88,89,38,69},
			[]int{4,3,2,1,2,1,5,4,1,2,1,20,4,3,1,1,2,1,2,1,11,7,2,1,3,1,1,1,3,1,1,121,1,52,6,1,2,1,2,1,5,2,1,1,1,40,6,3,2,1,2,1,2,1,5,4,1,2,1,16,1,4,1,2,1,1,9,8,1,1,1,1,2,1,1,10,1,1,3,2,1,4,1,2,1,67,1,4,1,2,1,7,4,1,2,1,2,1,7,3,1,1,3,1,1,47,4,1,1,1,5,2,1,2,1,33,1,6,3,2,1,2,1,25,16,1,14,3,1,1,1,9,1,1,1,2,1,3,2,1,2,1,2,1,4,1,1,1,4,3,2,1,276,1,2,1,2,1,11,1,1,2,1,1,5,4,1,2,1,1,258,4,1,1,1,8,1,2,1,4,2,1,1,10,9,1,1,6,1,1,3,1,1,2,1,233,1,1,6,3,1,1,1,1,1,223,1,22,21,2,1,5,1,3,1,1,4,1,2,1,9,1,1,6,1,1,3,2,1,186,4,1,2,1,22,1,9,1,2,1,1,3,1,1,1,5,3,2,1,1,6,1,4,1,2,1,159,2,1,1,2,1,3,1,1,28,1,2,1,13,2,1,10,4,1,2,1,4,3,1,1,1,11,1,9,1,1,6,4,1,2,1,1,122,121,1,119,1,1,9,1,2,1,4,3,2,1,1,6,3,1,1,2,1,25,6,5,4,3,2,1,18,6,1,1,3,2,1,2,1,9,1,1,6,1,1,3,2,1,4,1,1,1,72,2,1,9,1,1,5,1,1,1,1,1,4,1,1,1,56,1,3,2,1,1,2,1,4,3,2,1,4,1,2,1,40,1,3,1,1,5,1,1,2,1,4,1,2,1,9,5,3,2,1,1,3,2,1,17,1,3,2,1,12,4,1,2,1,1,6,5,1,3,2,1,13,1,2,1,4,1,2,1,5,4,2,1,1,0,1,1,8,1,6,1,3,2,1,1,0,1,27,1,4,1,2,1,7,1,1,4,3,2,1,14,13,1,1,10,1,6,3,2,1,1,1,2,1,0,1,1,3,1,1,7,1,2,1,3,1,1,109,2,1,2,1,10,9,8,3,2,1,2,1,2,1,61,1,2,1,1,7,1,2,1,3,2,1,49,1,5,1,3,2,1,9,1,2,1,2,1,3,2,1,33,2,1,30,1,3,1,1,25,1,2,1,8,4,1,2,1,3,1,1,1,1,11,1,9,2,1,6,1,4,2,1,1,33,1,1,2,1,15,1,10,1,2,1,6,1,1,1,2,1,1,2,1,13,4,1,2,1,1,1,6,1,1,3,2,1,0,1,16,1,1,1,12,6,3,2,1,2,1,5,4,1,2,1,88,1,4,1,2,1,6,1,1,3,2,1,76,1,1,4,2,1,1,7,6,5,1,2,1,1,9,2,1,1,5,2,1,2,1,3,2,1,50,5,1,3,1,1,44,1,1,3,2,1,2,1,36,1,1,2,1,31,2,1,13,12,1,8,3,2,1,4,3,2,1,2,1,15,1,2,1,1,2,1,2,1,6,1,1,3,2,1,0,1,8,1,2,1,2,1,2,1,1,9,3,2,1,3,2,1,2,1,7,5,4,3,1,1,1,58,2,1,1,6,1,1,3,2,1,6,1,1,1,2,1,15,2,1,12,1,4,3,2,1,2,1,4,1,2,1,27,1,4,2,1,1,1,20,19,18,1,13,2,1,4,1,2,1,6,5,2,1,1,1,3,2,1,0,2,1,13,1,1,10,9,2,1,5,2,1,2,1,1,21,2,1,5,1,2,1,1,2,1,11,3,1,1,6,5,2,1,1,1,1,15,1,1,1,6,1,2,1,2,1,5,2,1,2,1,0,6,2,1,3,2,1,4,1,1,1,19,1,4,1,2,1,2,1,11,1,9,1,4,1,2,1,1,2,1,9,1,7,1,5,2,1,1,1,7,1,4,3,2,1,1,4,2,1,1,2,1,0,19,2,1,2,1,6,2,1,3,1,1,4,1,2,1,4,3,1,1,11,1,4,1,1,1,1,1,3,1,1,54,1,3,1,1,5,1,3,2,1,1,43,1,2,1,38,3,1,1,5,4,3,2,1,7,1,5,1,3,2,1,7,2,1,4,1,2,1,4,2,1,1,11,1,3,2,1,6,1,2,1,2,1,1,0,1,2,1,0,14,1,1,3,2,1,8,1,1,1,4,3,2,1,0,2,1,2,1,0,1,1,0,1,0}},
	}

	for _, testCase := range testCases {
		output := DailyTemperaturesBis(testCase.input)

		if !reflect.DeepEqual(output, testCase.expectedOutput) {
			t.Errorf("[%d] Output was incorrect, DailyTemperaturesBis()= %v, want: %v.", testCase.id, output, testCase.expectedOutput)
		}
	}
}
