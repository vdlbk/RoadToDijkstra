package challenge_8

import (
	"fmt"
	"testing"
)

func BenchmarkPivotIndex(b *testing.B) {
	benchmarks := []struct {
		n    []int
	}{
		{[]int{1, 7, 3, 6, 5, 6}},
		{[]int{1, 2, 6, 2, 1}},
		{[]int{1, 2, 6, 1, 2}},
		{[]int{1, -2, 6, -1, 2}},
		{[]int{-1, -2, -6, -1, -2}},
		{[]int{-1,-1,-1,-1,0,1}},
		{[]int{-1,-1,-1,-1,-1,0}},
		{[]int{-1,-1,-1,-1,0,-1}},
		{[]int{-1,-1,0,0,-1,-1}},
		{[]int{-1,-1,-1,-1,-1,-1}},
		{[]int{5, 0, 3, 2}},
		{[]int{1, 2, 1}},
		{[]int{1, 1, 1}},
		{[]int{1, 2, 3}},
		{[]int{}},
	}

	for _, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("PivotIndex %v",benchmark.n), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				PivotIndex(benchmark.n)
			}
		})

		b.Run(fmt.Sprintf("PivotIndexBis %v",benchmark.n), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				PivotIndexBis(benchmark.n)
			}
		})
	}
}
