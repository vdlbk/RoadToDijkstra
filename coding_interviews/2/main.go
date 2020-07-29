package challenge

func CountRecurrentCharacter(input string) string {

	counters := make(map[rune]int)

	// O(n)
	for _, c := range input {
		counters[c]++
	}

	// O(n)
	for _, c := range input {
		if counters[c] > 1 {
			return string(rune(c))
		}
	}

	return ""
}

func IsAnagram(initial, anagram string) bool {

	if len(initial) != len(anagram) {
		return false
	}

	countersInitial := make(map[rune]int)
	countersAnagram := make(map[rune]int)

	for _, c := range initial {
		countersInitial[c]++
	}

	for _, c := range anagram {
		countersAnagram[c]++
	}

	for key, value := range countersInitial {
		if countersAnagram[key] != value {
			return false
		}
	}

	return true
}
