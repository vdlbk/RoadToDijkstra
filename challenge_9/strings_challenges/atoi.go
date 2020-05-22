package strings_challenges

import "math"

func myAtoi(str string) int {
	x := 0
	modifier := 1
	forcePositive := false
	i := 0
	for _, c := range str {
		if c < '0' || c > '9' {
			if x != 0 {
				break
			} else if (c != ' ' && c != '-' && c != '+') || modifier == -1 || forcePositive || i > 0 {
				return 0
			} else if c == ' ' {
				continue
			} else if c == '+' {
				forcePositive = true
				continue
			}
		}

		if modifier == 1 && c == '-' {
			modifier = -1
			continue
		}
		x *= 10
		x += int(c-48) * modifier
		i++

		if x > math.MaxInt32 {
			return math.MaxInt32
		} else if x < math.MinInt32 {
			return math.MinInt32
		}
	}

	return x
}
