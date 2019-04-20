package challenge_7

func SingleNumber(nums []int) int {
	m := make(map[int]bool)

	for i, x := range nums {
		if !m[x]{
			for j := i +1; j < len(nums); j++ {
				if nums[j] == x {
					m[x] = true
					break
				}
			}
			if !m[x] {
				return x
			}
		}
	}

	return -1
}

func SingleNumberBis(nums []int) int {
	m := make(map[int]int)

	for _, i := range nums {
		m[i]++
	}

	for k, v := range m {
		if v < 2 {
			return k
		}
	}

	return -1
}

func SingleNumberBit(nums []int) int {
	a := 0
	for _, x := range nums {
		a ^= x
	}

	return a
}
