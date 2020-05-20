package strings_challenges

import (
	"fmt"
)

// 48 -> 57  : 0 -> 9
// 65 -> 90  : A -> Z
// 97 -> 122 : a-> z
func isPalindrome(s string) bool {
	x := 0
	y := len(s) - 1

	ignore := func(r uint8) bool {
		return r < 48 || (r > 57 && r < 65) || (r > 90 && r < 97) || r > 122
	}
	for x < y {
		if ignore(s[x]) {
			x++
			continue
		} else if ignore(s[y]) {
			y--
			continue
		}
		a := s[x]
		if a > 90 {
			a -= 32
		}
		b := s[y]
		if b > 90 {
			b -= 32
		}
		if a != b {
			return false
		}
		fmt.Println(a, string(a), string(b), b)
		x++
		y--

	}
	return true
}

func isPalindromeOpti(s string) bool {
	if len(s) == 0 || len(s) == 1 {
		return true
	}
	i := 0
	j := len(s) - 1
	delta := byte('a') - byte('A')
	for i <= j {
		chi := s[i]
		chj := s[j]
		if s[i] >= 'A' && s[i] <= 'Z' {
			chi += delta
		}
		if s[j] >= 'A' && s[j] <= 'Z' {
			chj += delta
		}
		if !(chi >= 'a' && chi <= 'z') && !(chi >= '0' && chi <= '9') {
			i++
			continue
		}
		if !(chj >= 'a' && chj <= 'z') && !(chj >= '0' && chj <= '9') {
			j--
			continue
		}
		if chi != chj {
			return false
		}
		i++
		j--
	}
	return true
}
