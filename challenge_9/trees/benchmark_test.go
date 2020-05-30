package trees

import (
	"fmt"
	"testing"
)

func BenchmarkMaxDepth(b *testing.B) {
	benchmarks := []struct {
		input *TreeNode
	}{
		{GenerateTreeNode([]int{3, 9, 20, null, null, 15, 7})},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("maxDepth %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				maxDepth(benchmark.input)
			}
		})
	}
}
