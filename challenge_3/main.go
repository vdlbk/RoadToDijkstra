package challenge_3

func Permute(nums []int) [][]int {
	var result [][]int
	if len(nums) == 0 {
		result = make([][]int, 0)
	} else if len(nums) < 2 {
		result = [][]int{nums}
	} else if len(nums) == 2 {
		result = [][]int{nums, []int{nums[1], nums[0]}}
	} else {
		for i := range nums { // O(n)
			x := nums[i]
			startBuffer := nums[:i]
			endBuffer := nums[i+1:]
			start := make([]int, len(startBuffer))
			copy(start, startBuffer)

			end := make([]int, len(endBuffer))
			copy(end, endBuffer)

			y := append(start, end...)
			m := Permute(y) // O(n-1!)
			for j := range m { 
				m[j] = append([]int{x}, m[j]...)
			}
			result = append(result, m...)
		}
	}

	return result
}
