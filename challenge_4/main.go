package main

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
