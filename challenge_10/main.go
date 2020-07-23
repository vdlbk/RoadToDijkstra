package challenge

import (
	"math"
	"strconv"
)

func binaryGap(N int) int {
	if N <= 0 {
		return 0
	}
	longestGap := 0
	currentGap := 0
	binary := strconv.FormatInt(int64(N), 2)
	for _, c := range binary {
		if c == '1' {
			if currentGap > longestGap {
				longestGap = currentGap
			}
			currentGap = 0
		} else {
			currentGap++
		}
	}

	return longestGap
}

func diff(A []int) int {
	if len(A) < 2 {
		return 0
	}
	minDiff := math.MaxInt64
	sliceSum := sum(A)
	before := A[0]
	after := sliceSum - before

	for i := 1; i < len(A); i++ {
		diff := before - after
		if diff < 0 {
			diff *= -1
		}

		if diff < minDiff {
			minDiff = diff
		}

		before += A[i]
		after -= A[i]

	}

	return minDiff
}

func sum(slice []int) int {
	sum := 0
	for _, x := range slice {
		sum += x
	}
	return sum
}
