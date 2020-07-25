package searching

func merge(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) != m+n {
		return
	}

	cursor := len(nums1) - 1
	// Convert length to cursors
	m--
	n--
	for cursor >= 0 {
		if n < 0 {
			// take the rest of m
			for i := m; i >= 0; i-- {
				nums1[cursor] = nums1[i]
				cursor--
			}
		} else if m < 0 {
			// take the rest of n
			for i := n; i >= 0; i-- {
				nums1[cursor] = nums2[i]
				cursor--
			}
		} else {
			if nums1[m] <= nums2[n] {
				nums1[cursor] = nums2[n]
				n--
			} else {
				nums1[cursor] = nums1[m]
				m--
			}
		}

		cursor--
	}

}

// Use extra memory
func mergeAlt(nums1 []int, m int, nums2 []int, n int) {
	newArr := make([]int, m+n)
	newIndex := 0
	index2 := 0
	for i := 0; i < m; i++ {
		for j := index2; j < n; j++ {
			if nums1[i] > nums2[j] {
				index2++
				newArr[newIndex] = nums2[j]
				newIndex++
			} else {
				break
			}
		}
		newArr[newIndex] = nums1[i]
		newIndex++
	}
	for i := index2; i < n; i++ {
		newArr[newIndex] = nums2[i]
		newIndex++
	}
	for index := range newArr {
		nums1[index] = newArr[index]
	}
}
