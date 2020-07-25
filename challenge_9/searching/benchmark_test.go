package searching

import (
	"fmt"
	"testing"
)

func BenchmarkMerge(b *testing.B) {
	benchmarks := []struct {
		input  []int
		m      int
		input2 []int
		n      int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3},
		{[]int{1, 1, 1, 0, 0, 0}, 3, []int{1, 1, 1}, 3},
		{[]int{1, 0}, 1, []int{2}, 1},
		{[]int{1, 2, 3}, 1, []int{2}, 1},
		{[]int{0, 2, 0, 0}, 2, []int{-1, 6}, 2},
		{[]int{0, 14, 0, 0}, 2, []int{-1, 6}, 2},
		{[]int{0, 14, 13, 0}, 2, []int{-1, 6}, 2},
		{[]int{}, 0, []int{}, 0},
		{[]int{9, 18, 27}, 3, []int{}, 0},
		{[]int{0}, 0, []int{1}, 1},
	}

	for _, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("merge %v", benchmark.n), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				merge(benchmark.input, benchmark.m, benchmark.input2, benchmark.n)
			}
		})
	}

	for _, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("mergeAlt %v", benchmark.n), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				mergeAlt(benchmark.input, benchmark.m, benchmark.input2, benchmark.n)
			}
		})
	}
}
