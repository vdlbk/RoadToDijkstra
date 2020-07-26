package mathematics

import (
	"fmt"
	"testing"
)

func BenchmarkFizzBuzz(b *testing.B) {
	benchmarks := []struct {
		input int
	}{
		{0},
		{-1},
		{1},
		{5},
		{15},
	}

	for _, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("fizzBuzz %v", benchmark.input), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				fizzBuzz(benchmark.input)
			}
		})
	}
}
