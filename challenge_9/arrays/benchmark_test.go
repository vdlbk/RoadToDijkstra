package arrays

import (
	"fmt"
	"testing"
)

func BenchmarkRemoveDuplicates(b *testing.B) {
	benchmarks := []struct {
		n    []int
	}{
		{[]int{}},
	}

	for _, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("removeDuplicates %v",benchmark.n), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				removeDuplicates(benchmark.n)
			}
		})
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	benchmarks := []struct {
		n    []int
	}{
		{[]int{7,1,5,3,6,4}},
		{[]int{7,1,5,3,6,4, 7,1,5,3,6,4, 7,1,5,3,6,4, 7,1,5,3,6,4}},
		{[]int{3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5}},
		{[]int{
			3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5,
			3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5,
			3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5,
			3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5,
			3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5,
			3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5,
			3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5,
			3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5, 3, 6, 4, 9, 2, 8, 1, 5,
		}},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("maxProfit %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				maxProfit(benchmark.n)
			}
		})
	}
}

func BenchmarkRotate(b *testing.B) {
	benchmarks := []struct {
		input          []int
		k int
	}{
		{[]int{1,2,3,4,5,6,7}, 3},
		{[]int{1,2,3,4,5,6,7},2},
		{[]int{1,2,3,4}, 2},
		{[]int{1,2,3,4,5}, 2},
		{[]int{1,2,3,4,5,6}, 2},
		{[]int{1,2,3,4,5,6}, 3},
		{[]int{-1,-100,3,99}, 2},
		{[]int{-1,-100,3,99}, 3,},
		{[]int{1,2}, 2},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("rotate %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				rotate(benchmark.input, benchmark.k)
			}
		})
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("rotateAlt %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				rotateAlt(benchmark.input, benchmark.k)
			}
		})
	}
}