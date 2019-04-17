package challenge_1

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Hi !")
}

func LongestCommonSubsequence(s1, s2 string) string {
	result := ""
	longestCommonSubsequence := ""
	idx := 0
	cursorI := 0
	cursorJ := 0
	for idx < len(s1) {
		c := s1[cursorI]
		d := s2[cursorJ]

		if c == d {
			result += string(d)
			cursorJ++
			cursorI++
		} else {
			cursorJ++
		}

		if cursorJ >= len(s2) {
			cursorJ = len(s2) - 1
			cursorI++
		}

		if cursorI >= len(s1) && cursorJ < len(s2)-1 {
			cursorI = len(s1) - 1
		}

		if cursorI >= len(s1) {
			idx++
			cursorI = idx
			cursorJ = 0
			if len(result) > len(longestCommonSubsequence) && float64(len(result)) <= shortestWord(s1, s2) {
				longestCommonSubsequence = result
			}
			result = ""
		}
	}

	return longestCommonSubsequence
}

func shortestWord(s1, s2 string) float64 {
	return math.Min(float64(len(s1)), float64(len(s2)))
}
