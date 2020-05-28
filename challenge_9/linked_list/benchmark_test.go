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

func BenchmarkReverseList(b *testing.B) {
	benchmarks := []struct {
		input *ListNode
	}{
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5})},
		{GenerateListFromSlice([]int{1, 2, 3, 4})},
		{GenerateListFromSlice([]int{1, 2})},
		{GenerateListFromSlice([]int{})},
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 18, 19, 20, 21})},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("reverseList %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				reverseList(benchmark.input)
			}
		})

		b.Run(fmt.Sprintf("reverseListRecursive %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				reverseListRecursive(benchmark.input)
			}
		})
	}
}

func BenchmarkIsPalindrome(b *testing.B) {
	benchmarks := []struct {
		input *ListNode
	}{
		{GenerateListFromSlice([]int{1, 2, 3, 2, 1})},
		{GenerateListFromSlice([]int{20, 43, 2, 79, 2, 43, 20})},
		{GenerateListFromSlice([]int{1, 2, 3, 4, 5, 5, 4, 3, 2, 1})},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("isPalindrome %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				isPalindrome(benchmark.input)
			}
		})

		b.Run(fmt.Sprintf("isPalindromeOpti %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				isPalindromeOpti(benchmark.input)
			}
		})
	}
}

func BenchmarkHasCycle(b *testing.B) {
	benchmarks := []struct {
		input *ListNode
	}{
		{GenerateListFromSliceWithCycle([]int{3, 2, 0, -4}, 0)},
		{GenerateListFromSliceWithCycle([]int{3, 2, 0, -4}, 1)},
		{GenerateListFromSliceWithCycle([]int{3, 2, 0, -4}, 2)},
		{GenerateListFromSliceWithCycle([]int{3, 2, 0, -4}, 3)},
		//{GenerateListFromSliceWithCycle([]int{3, 2, 0, -4}, -2)},
		{GenerateListFromSliceWithCycle([]int{1, 2}, 0)},
		{GenerateListFromSliceWithCycle([]int{1, 2}, 1)},
		//{GenerateListFromSliceWithCycle([]int{1, 2}, 2)},
		{GenerateListFromSliceWithCycle([]int{1}, -1)},
		{GenerateListFromSliceWithCycle([]int{3, 2, 0, -4, 2, 0, -4, 2, 0, -4, 2, 0, -4, 2, 0, -4, 2, 0, -4, 2, 0, -4}, -1)},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("hasCycle %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				hasCycle(benchmark.input)
			}
		})

		b.Run(fmt.Sprintf("hasCycleOpti %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				hasCycleOpti(benchmark.input)
			}
		})
	}
}
