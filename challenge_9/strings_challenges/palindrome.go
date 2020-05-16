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
	for x < y {
		if s[x] < 48 && s[x] > 57 && s[x] < 65 && s[x] > 90 && s[x] < 97 && s[x] > 122 {
			fmt.Println("lol", string(s[x]), s[x])
		}
		fmt.Println(s[x], string(s[x]), string(s[y]), s[y])
		// to lower with ascii <<
		// A = 65; a = 97  (32)
		x++
		y--

	}
	return false
}
