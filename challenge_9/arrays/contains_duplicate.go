package arrays

func containsDuplicate(nums []int) bool {
	if len(nums) < 2 {
		return false
	}
	m := make(map[int]bool)
	for _, x := range nums {
		if ok := m[x]; ok {
			return true
		} else {
			m[x] = true
		}
	}
	return false
}

func containsDuplicateAlt(nums []int) bool {
	moi:= make(map[int]bool,0)

	for i:= 0; i< len(nums); i++{

		num := nums[i]
		if _, ok:= moi[num]; ok{
			return  true
		}

		moi[num]= true
	}

	return false
}
