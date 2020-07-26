package dynamic

func climbStairs(n int) int {
	if n < 1 {
		return 0
	}

	if n == 1 {
		return 1
	}

	results := make([]int, n+1)
	results[1] = 1
	results[2] = 2
	for i := 3; i <= n; i++ {
		results[i] = results[i-1] + results[i-2]
	}

	return results[n]
}
