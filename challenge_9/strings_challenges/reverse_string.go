package strings_challenges

func reverseString(s []byte) {
	length := len(s)
	if length == 0 || length == 1 {
		return
	}
	middle := length / 2
	for i := 0; i < middle; i++ {
		s[i], s[length-1-i] = s[length-1-i], s[i]
	}
}

func reverseStringOptimal(s []byte) {
	left, right := 0, len(s)-1
	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}
}
