package linked_list

import (
	"fmt"
	"testing"
)

func BenchmarkDeleteNode(b *testing.B) {
	benchmarks := []struct {
		input *ListNode
		node  int
	}{
		{generateListFromSlice([]int{4, 5, 1, 9}), 5},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("deleteNode %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				deleteNode(benchmark.input, benchmark.node)
			}
		})
	}
}
