package strings_challenges

import (
	"fmt"
	"testing"
)

func BenchmarkRemoveDuplicates(b *testing.B) {
	benchmarks := []struct {
		n []byte
	}{
		{[]byte{byte('H'), byte('a'), byte('n'), byte('n'), byte('a'), byte('h')}},
		{[]byte{byte('H'), byte('a'), byte('n'), byte('n'), byte('a'), byte('h'), byte('0'), byte('1'), byte(' ')}},
		{[]byte{byte('H'), byte('a'), byte('n'), byte('n'), byte('a'), byte('h'), byte('0'), byte('1'), byte(' '),
			byte('H'), byte('a'), byte('n'), byte('n'), byte('a'), byte('h'), byte('0'), byte('1'), byte(' '),
			byte('H'), byte('a'), byte('n'), byte('n'), byte('a'), byte('h'), byte('0'), byte('1'), byte(' ')}},
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("reverseString %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				reverseString(benchmark.n)
			}
		})
	}

	for i, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("reverseStringOptimal %v", i), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				reverseStringOptimal(benchmark.n)
			}
		})
	}
}
