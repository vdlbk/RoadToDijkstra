package arrays

import "sort"

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	n := make(map[int]int)
	for _, x := range nums1 {
		m[x]++
	}

	for _, x := range nums2 {
		n[x]++
	}

	var result []int
	for k, v := range m {
		if y, ok := n[k]; ok {
			z := v
			if y < v {
				z = y
			}
			for i := 0; i < z; i++ {
				result = append(result, k)
			}
		}
	}
	return result
}

func intersectAlt1(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	result := []int{}
	if len(nums1) == 0 || len(nums2) == 0 {
		return result
	}
	for i := 0; i < len(nums1); i++ {
		m[nums1[i]] += 1
	}

	for i := 0; i < len(nums2); i++ {
		if m[nums2[i]] > 0 {
			result = append(result, nums2[i])
			m[nums2[i]] -= 1
		}
	}
	return result
}

func intersectOpti(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}
	sort.Ints(nums1)
	sort.Ints(nums2)

	i, j := 0, 0
	var results []int
	for i < len(nums1) && j < len(nums2) {
		m, n := nums1[i], nums2[j]
		if m == n {
			results = append(results, m)
			i += 1
			j += 1
		} else if m < n {
			i += 1
		} else {
			j += 1
		}
	}
	return results
}
