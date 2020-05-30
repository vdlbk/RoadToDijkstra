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
