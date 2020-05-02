package arrays

func maxProfit(prices []int) int {
	max := 0
	for i, x := range prices {
		y := 0
		if i < len(prices)-1 {
			y = prices[i+1]
		}

		z := y - x
		if z > 0 {
			max += z
		}
	}
	return max
}
