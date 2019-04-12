package challenge_4

import "math"

func DailyTemperatures(T []int) []int {
	nextWarmerTemperature := make([]int, len(T))
	for i, x := range T {
		next := 0
		for j := i + 1; j < len(T); j++ {
			if x < T[j] {
				next = j - i
				break
			}
		}
		nextWarmerTemperature[i] = next
	}
	return nextWarmerTemperature
}

func DailyTemperaturesBis(T []int) []int {
	nextWarmerTemperature := make([]int, len(T))
	// This could be an array of 71(101-30) item only but we will have to remove 30 each time we want to access this array so we won't increase perf, just a bit of space.
	next := make([]int, 101)
	for i := range next {
		next[i] = math.MaxInt64
	}
	for i := len(T)-1; i >= 0; i-- {
		warmerIdx := math.MaxInt64
		for t := T[i] + 1; t <= 100; t++ {
			if next[t] < warmerIdx {
				warmerIdx = next[t]
			}
		}
		if warmerIdx < math.MaxInt64 {
			nextWarmerTemperature[i] = warmerIdx - i
		}

		next[T[i]] = i
	}
	return nextWarmerTemperature
}
