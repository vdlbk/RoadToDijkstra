package arrays

func rotate(nums []int, k int) {
	m := make(map[int]int)
	newPosition := 0
	for i := 0; i < len(nums); i++ {
		newPosition = (i + k) % len(nums)
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
