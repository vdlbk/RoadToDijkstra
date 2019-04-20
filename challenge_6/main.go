package challenge_6

import (
	"strconv"
)

func CountDigitOne(n int) int {
	// short-circuit for negative integers
	if n <= 0 {
		return 0
	}

	r := 1
	for i := 2; i <= n; i++ {
		y := i
		for {
			if mod := y % 10; mod == 1 {
				r++
			}
			x := y / 10
			if x == 0 {
				break
			}
			y = x
		}
	}

	return r
}

func SillyCountDigitOne(n int) int {
	// short-circuit for negative integers
	if n <= 0 {
		return 0
	}

	r := 1
	for i := 2; i <= n; i++ {
		x := strconv.Itoa(i)
		for _, c := range x {
			if c == '1' {
				r++
			}
		}

	}

	return r
}

func SolutionCountDigitOne(n int) int {
	r := 0
	for i := 1; i <= n; i *=10 {
		y := i * 10
		r+=  (n / y) * i + min(max(n % y - i + 1, 0), i)
	}
	return r
}

func min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
