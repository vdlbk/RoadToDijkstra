package dynamic

import (
	"fmt"
	"testing"
)

func BenchmarkClimbStairs(b *testing.B) {
	benchmarks := []struct {
		input int
	}{
		{0},
		{1},
		{2},
		{3},
		{35},
	}

	for _, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("climbStairs %v", benchmark.input), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				climbStairs(benchmark.input)
			}
		})
	}
}
