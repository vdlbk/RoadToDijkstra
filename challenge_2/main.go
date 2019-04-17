package challenge_2

import (
	"strings"

	"github.com/vdlbk/RoadToDijkstra/challenge_2/lib/qsort"
)

func Reverse(s string) string {
	reversed := ""
	for i := len(s) - 1; i >= 0; i-- {
		reversed += string(s[i])
	}

	return reversed
}

func AltReverse(s string) string {
	reversed := make([]string, len(s))
	for i, c := range s {
		reversed[len(s)-1-i] = string(c)
	}
	return strings.Join(reversed, "")
}

func MissingItem(a1, a2 []int) int {
	a2 = qsort.QuickSort(a2)

	for _, x := range a1 {
		var y int
		for _, y = range a2 {
			if x == y {
				break
			}
			if x < y {
				return x
			}
		}
		if len(a2) == 0 || x > y{
			return x
		}
	}
	return -1
}
