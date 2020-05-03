package arrays

import (
	"fmt"
	"testing"
)

func BenchmarkRemoveDuplicates(b *testing.B) {
	benchmarks := []struct {
		n []int
	}{
		{[]int{}},
	}

	for _, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("removeDuplicates %v", benchmark.n), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				removeDuplicates(benchmark.n)
			}
		})
	}
}

func BenchmarkMaxProfit(b *testing.B) {
	benchmarks := []struct {
		n []int
	}{
		{[]int{7, 1, 5, 3, 6, 4}},
		{[]int{7, 1, 5, 3, 6, 4, 7, 1, 5, 3, 6, 4, 7, 1, 5, 3, 6, 4, 7, 1, 5, 3, 6, 4}},
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
		input []int
		k     int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 2},
		{[]int{1, 2, 3, 4}, 2},
		{[]int{1, 2, 3, 4, 5}, 2},
		{[]int{1, 2, 3, 4, 5, 6}, 2},
		{[]int{1, 2, 3, 4, 5, 6}, 3},
		{[]int{-1, -100, 3, 99}, 2},
		{[]int{-1, -100, 3, 99}, 3},
		{[]int{1, 2}, 2},
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

func BenchmarkContainsDuplicate(b *testing.B) {
	benchmarks := []struct {
		n []int
	}{
		{[]int{1, 2, 3, 1}},
		{[]int{1, 2, 3, 4}},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 13}},
		{[]int{}},
		{[]int{10000000}},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("containsDuplicate %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				containsDuplicate(benchmark.n)
			}
		})
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("containsDuplicateAlt %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				containsDuplicateAlt(benchmark.n)
			}
		})
	}
}

func BenchmarkSingleNumber(b *testing.B) {
	benchmarks := []struct {
		n []int
	}{
		{[]int{2, 2, 1}},
		{[]int{4, 1, 2, 1, 2}},
		{[]int{1}},
		{[]int{-1, -1, -2, -4, -2, -7, -4}},
		{[]int{-1, -10, -2, -9, -3, -8, -4, -7, -5, -6, -6, -5, -7, -4, -8, -3, -9, -2, -10, -1, 0}},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("singleNumber %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				singleNumber(benchmark.n)
			}
		})
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("singleNumberAlt %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				singleNumberAlt(benchmark.n)
			}
		})
	}
}

func BenchmarkIntersect(b *testing.B) {
	benchmarks := []struct {
		n []int
		m []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}},
		{[]int{1, 2}, []int{3, 4}},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{[]int{-1, -2, -3}, []int{-2, -3, -4}},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("intersect %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				intersect(benchmark.n, benchmark.m)
			}
		})
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("intersectAlt1 %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				intersectAlt1(benchmark.n, benchmark.m)
			}
		})
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("intersectOpti %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				intersectOpti(benchmark.n, benchmark.m)
			}
		})
	}
}

func BenchmarkPlusOne(b *testing.B) {
	benchmarks := []struct {
		n []int
	}{
		{[]int{1, 2, 3}},
		{[]int{4, 3, 2, 1}},
		{[]int{1, 0, 9}},
		{[]int{9, 9, 9}},
		{[]int{1}},
		{[]int{9}},
		{[]int{2, 4, 5, 6, 7, 8, 9, 0, 2, 4, 5, 6, 7, 8, 9, 0, 2, 4, 5, 6, 7, 8, 9, 0, 2, 4, 5, 6, 2, 4, 5, 6, 7, 8, 9}},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("plusOne %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				plusOne(benchmark.n)
			}
		})
	}
}

func BenchmarkMoveZeroes(b *testing.B) {
	benchmarks := []struct {
		n []int
	}{
		{[]int{0, 1, 0, 3, 12}},
		{[]int{0, 1}},
		{[]int{1, 2, 3}},
		{[]int{0, 0, 0, 0}},
		{[]int{1, 0, 0, 0}},
		{[]int{1}},
		{[]int{}},
		{[]int{
			2, 4, 5043, 4343, 0, 363, 345, 238, 0, 67, 33, 4, 0, 6, 2, 7, 7, 7, 0, 0, 0, 0, 0, 1, 0, 3, 0,
		}},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("moveZeroes %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				moveZeroes(benchmark.n)
			}
		})
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("moveZeroesAlt %d", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				moveZeroesAlt(benchmark.n)
			}
		})
	}
}
