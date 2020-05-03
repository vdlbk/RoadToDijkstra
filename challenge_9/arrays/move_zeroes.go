package arrays

func moveZeroes(nums []int) {
	cursor := 0
	for i := 0; i < len(nums); i++ {
		if nums[cursor] == 0 {
			nums = append(nums[:cursor], append(nums[cursor+1:], 0)...)
		} else {
			cursor++
		}
	}
}

func moveZeroesAlt(nums []int) {
	if len(nums) <= 1 {
		return
	}

	zero := 0
	for ; zero < len(nums) && 0 != nums[zero]; zero++ {
	}

	for nonzero := zero + 1; nonzero < len(nums); nonzero++ {
		if 0 != nums[nonzero] {
			nums[zero] = nums[nonzero]
			zero++
		}
	}

	for ; zero < len(nums); zero++ {
		nums[zero] = 0
	}
}
