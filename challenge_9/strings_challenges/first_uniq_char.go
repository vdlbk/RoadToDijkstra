package strings_challenges

func firstUniqChar(s string) int {
	m := make(map[rune][]int)
	for i, c := range s {
		if _, ok := m[c]; ok {
			m[c] = append(m[c], i)
		} else {
			m[c] = []int{i}
		}

	}
	max := len(s)
	for _, v := range m {
		if x := v[0]; len(v) == 1 && x < max {
			max = x
		}
	}
	if max == len(s) {
		return -1
	}
	return max
}

func firstUniqCharOpti(s string) int {
	sLen := len(s)
	if sLen == 0 {
		return -1
	}
	if sLen == 1 {
		return 0
	}

	chars := [26]int{}
	// setup init value
	for i := range chars {
		chars[i] = -1
	}

	// find unique chars
	for i, v := range s {
		if chars[v-'a'] == -1 {
			chars[v-'a'] = i
		} else {
			chars[v-'a'] = -2
		}
	}

	// find min index
	ret := sLen + 1
	for _, v := range chars {
		if v >= 0 && v < ret {
			ret = v
		}
	}

	if ret == sLen+1 {
		return -1
	}
	return ret
}
