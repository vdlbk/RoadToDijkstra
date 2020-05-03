package arrays

func twoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))
	for i, x := range nums {
		m[x] = i
	}
	for i, x := range nums {
		if position, ok := m[target-x]; ok && i < position {
			return []int{i, position}
		}
	}
	return []int{}
}
