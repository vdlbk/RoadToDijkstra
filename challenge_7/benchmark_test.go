package challenge_7

import (
	"fmt"
	"testing"
)

func Benchmark_CountDigitOne(b *testing.B) {
	benchmarks := []struct {
		name string
		input    []int
	}{
		{"3", []int{2,2,1,1,3}},
		{"1", []int{2,2,1}},
		{"6", []int{4, 1, 2, 6, 3, 5, 1, 2, 3, 5, 4}},
	}

	for _, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("SingleNumber %s/%d", benchmark.name, len(benchmark.input)), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				SingleNumber(benchmark.input)
			}
		})
		b.Run(fmt.Sprintf("SingleNumberBis %s/%d", benchmark.name, len(benchmark.input)), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				SingleNumberBis(benchmark.input)
			}
		})
	}

}
