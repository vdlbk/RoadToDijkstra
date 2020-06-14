package trees

import (
	"fmt"
	"testing"
)

func BenchmarkMaxDepth(b *testing.B) {
	benchmarks := []struct {
		input *TreeNode
	}{
		{GenerateTreeNode([]int{1, 2, 4})},
		{GenerateTreeNode([]int{3, 9, 20, null, null, 15, 7})},
		{GenerateTreeNode([]int{3, 9, 20, null, 9, 2, null, 5, 6, null, null, 10})},
		{GenerateTreeNode([]int{0, 1, 2, 3, 5, 4, 6, 7, 9, 11, 13, 8, 10, 12, 14})},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("maxDepth %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				maxDepth(benchmark.input)
			}
		})

		b.Run(fmt.Sprintf("maxDepthOpti %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				maxDepthOpti(benchmark.input)
			}
		})
	}
}

func BenchmarkSortedArrayToBST(b *testing.B) {
	benchmarks := []struct {
		input []int
	}{
		{[]int{}},
		{[]int{1}},
		{[]int{1, 2}},
		{[]int{1, 2, 3}},
		{[]int{1, 2, 3, 4, 5}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26,
			27, 28, 29, 30, 31, 32, 33, 34}},
		{[]int{-10, -3, 0, 5, 9}},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("sortedArrayToBST %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				sortedArrayToBST(benchmark.input)
			}
		})

		b.Run(fmt.Sprintf("sortedArrayToBSTOpti %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				sortedArrayToBSTOpti(benchmark.input)
			}
		})
	}
}
