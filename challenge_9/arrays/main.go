package arrays

/*
	removeDuplicates
 */
func removeDuplicates(nums []int) int {
	var last int
	n := 0
	for i, x := range nums {
		if i == 0 || x != last  {
			nums[n] = x
			n++
		}

		last = x
	}
	return n
}

/*
	maxProfit
 */
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

/*
	rotate
 */
func rotate(nums []int, k int) {
	m := make(map[int]int)
	newPosition := 0
	for i := 0; i < len(nums); i++{
		newPosition = (i + k)%len(nums)
		m[newPosition] = nums[newPosition]
		if x, exists := m[i]; exists {
			nums[newPosition] = x
		} else {
			nums[newPosition] = nums[i]
		}
	}

}

func rotateAlt(nums []int, k int) {
	l := len(nums)
	c := k % l
	copy(nums, append(nums[l-c:], nums[:l-c]...))
}

/*
	rotate
 */
func containsDuplicate(nums []int) bool {
	// TODO
	return false
}