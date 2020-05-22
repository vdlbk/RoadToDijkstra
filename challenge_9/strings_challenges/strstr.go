package strings_challenges

func strStr(haystack string, needle string) int {

	if haystack == needle || len(needle) == 0 {
		return 0
	}
	if len(needle) > len(haystack) {
		return -1
	}

	for i := range haystack {
		x := i
		y := 0
		for x < len(haystack) && haystack[x] == needle[y] {
			x++
			if y++; y == len(needle) {
				return i
			}
		}
	}
	return -1
}

// The difference is small
func strStrOpti(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	for idx := range haystack {
		if haystack[idx] == needle[0] && len(needle) <= len(haystack)-idx {
			found := true
			for j := 1; j < len(needle); j++ {
				if haystack[idx+j] != needle[j] {
					found = false
					break
				}
			}
			if found {
				return idx
			}
		}
	}
	return -1
}
