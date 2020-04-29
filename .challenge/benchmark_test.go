package challenge

import (
	"fmt"
	"testing"
)

func BenchmarkFoobar(b *testing.B) {
	benchmarks := []struct {
		n    int
	}{
		{0},
	}

	for _, benchmark := range benchmarks {
		b.Run(fmt.Sprintf("Foobar %v",benchmark.n), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				Foobar(benchmark.n)
			}
		})
	}
}
