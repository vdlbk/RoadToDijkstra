package challenge

import (
	"fmt"
	"math"
	"strconv"
)

func Solution(N int) int {
	// write your code in Go 1.4
	if N < 10 {
		return 10
	}

	result := N * 10
	if result > math.MaxInt64 {
		return 1000000000
	}

	return result
}

func Solution2(A int, B int) int {
	// write your code in Go 1.4
	result := A * B

	binary := strconv.FormatInt(int64(result), 2)

	fmt.Println(math.MaxInt64)
	fmt.Println(100000000 * 100000000)

	count := 0
	for _, c := range binary {
		if c == '1' {
			count++
		}
	}

	return count
}

func Solution3(S string) int {
	// write your code in Go 1.4

	if len(S) < 3 {
		return 0
	}

	var previousChar rune
	count := 0 // number of consecutive letters
	swapOperations := 0
	for idx, c := range S {

		if c == previousChar {
			count++
		} else {
			count = 1
		}

		// swap
		if count == 3 {
			// If the next char is also the same char on which I'm rignt now, I'm at the limit, I need the swap this one.
			// Otherwise, if it's different, I might swap for a char that will create a substring with 3 identical consecutive
			// letters, which I don't want because it will increase the number of swap operations
			if idx+1 < len(S) && rune(S[idx+1]) == c {
				//S[idx] = swap(c)
				previousChar = rune(swap(c))
			} else {
				//S[idx-1] = swap(c)
				previousChar = c
			}
			swapOperations++
			count = 1
		} else {
			previousChar = c
		}

	}

	return swapOperations
}

func swap(letter rune) uint8 {
	if letter == 'a' {
		return 'b'
	}
	return 'a'
}
