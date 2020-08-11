package challenge

// O(n)
func Geography(input []string) bool {

	// O(n)
	cities := make(map[string]bool)

	// O(n)
	for _, city := range input {
		if cities[city] {
			return false
		} else {
			cities[city] = true
		}
	}

	// unique list of cities
	// O(n)
	var previousLetter rune = -1
	for _, city := range input {
		if len(city) == 0 {
			return false
		}
		firstLetter := rune(city[0])

		if previousLetter != -1 && firstLetter != previousLetter {
			return false
		}

		previousLetter = rune(city[len(city)-1])

	}

	return true
}

func CheckDuplicates(input []string) bool {
	cities := make(map[string]bool)

	// O(n)
	for _, city := range input {
		if cities[city] {
			return true
		} else {
			cities[city] = true
		}
	}

	return false
}
