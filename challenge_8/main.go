package challenge_8

func PivotIndex(nums []int) int {
	sum := 0
	leftsum := 0
	for _, x := range nums {
		sum += x
	}
	for i, x := range nums {
		if leftsum == sum - leftsum - x {
			return i
		}
		leftsum += x
	}
	return -1
}

func PivotIndexBis(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	sum := 0
	leftsum := 0
	for _, x := range nums {
		sum += x
	}
	for i, x := range nums {
		if leftsum == sum - x {
			return i
		}
		leftsum += x
		sum -= x
	}
	return -1
}
