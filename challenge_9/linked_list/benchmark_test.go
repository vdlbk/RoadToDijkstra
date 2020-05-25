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
		{GenerateListFromSlice([]int{4, 5, 1, 9}), 5},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("deleteNode %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				deleteNode(benchmark.input, benchmark.node)
			}
		})
	}
}

func BenchmarkRemoveNthFromEnd(b *testing.B) {
	benchmarks := []struct {
		input *ListNode
		n     int
	}{
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5}), 1},
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5}), 2},
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5}), 3},
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5}), 4},
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5}), 5},
		{GenerateListFromSlice([]int{1}), 1},
		{GenerateListFromSlice([]int{3, 3, 3}), 2},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("removeNthFromEnd %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				removeNthFromEnd(benchmark.input, benchmark.n)
			}
		})

		b.Run(fmt.Sprintf("removeNthFromEndAlternative %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				removeNthFromEndAlternative(benchmark.input, benchmark.n)
			}
		})
	}
}
