package strings_challenges

import (
	"math"
)

func reverse(x int) int {
	var result = 0
	for x > 0 || x < 0 {
		result = result*10 + x%10
		x /= 10
	}
	if result >= math.MaxInt32 || result <= math.MinInt32 {
		return 0
	}
	return result
}
