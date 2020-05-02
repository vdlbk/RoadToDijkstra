package arrays

func removeDuplicates(nums []int) int {
	var last int
	n := 0
	for i, x := range nums {
		if i == 0 || x != last {
			nums[n] = x
			n++
		}

		last = x
	}
	return n
}
