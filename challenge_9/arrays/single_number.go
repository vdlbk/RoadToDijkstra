package arrays

func singleNumber(nums []int) int {
	m := make(map[int]int, len(nums))
	for _, x := range nums {
		m[x]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}

func singleNumberAlt(nums []int) int {
	var sum int
	for _, n := range nums {
		sum ^= n
	}
	return sum
}