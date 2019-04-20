package challenge_6

import (
	"fmt"
	"testing"
)

func Benchmark_CountDigitOne(b *testing.B) {
	benchmarks := []struct {
		name string
		n    int
	}{
		{"10", 10},
		{"20", 20},
		{"100", 100},
		{"1000", 1000},
		{"1000000", 1000000},
	}

	for _, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("CountDigitOne %s/%d", benchmark.name, benchmark.n), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				CountDigitOne(benchmark.n)
			}
		})
		b.Run(fmt.Sprintf("SillyCountDigitOne %s/%d", benchmark.name, benchmark.n), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				SillyCountDigitOne(benchmark.n)
			}
		})
		b.Run(fmt.Sprintf("SolutionCountDigitOne %s/%d", benchmark.name, benchmark.n), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				SolutionCountDigitOne(benchmark.n)
			}
		})
	}

}
