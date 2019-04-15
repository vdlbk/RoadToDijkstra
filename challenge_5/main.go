package challenge_5

func ValidateStackSequences(pushed []int, popped []int) bool {
	buffer := make([]int, 0, len(popped))
	pushedSize := len(pushed)

	j := 0
	for _, x := range pushed {
		buffer = append(buffer, x)
		for len(buffer) > 0 && j < pushedSize && buffer[len(buffer) -1] == popped[j] {
			buffer = buffer[:len(buffer) -1]
			j++
		}
	}
	return j == pushedSize
}
