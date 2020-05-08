package strings_challenges

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	m := make(map[uint8]int)
	n := make(map[uint8]int)

	for i := 0; i < len(s); i++ {
		m[s[i]]++
		n[t[i]]++
	}
	if len(m) != len(n) {
		return false
	}

	for k, v := range m {
		if n[k] != v {
			return false
		}
	}
	return true
}

func isAnagramOpti(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	var m [26]int

	n := 0

	for _, c := range s {
		i := byte(c - 'a')

		if m[i] == 0 {
			n++
		}

		m[i]++
	}

	for _, c := range t {
		i := byte(c - 'a')

		m[i]--

		if m[i] == 0 {
			n--
		}
	}

	return n == 0
}
