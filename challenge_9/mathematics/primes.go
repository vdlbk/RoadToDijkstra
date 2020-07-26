package mathematics

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}

	primes := make([]bool, n)

	for i := 2; i*i < len(primes); i++ {
		if !primes[i] {
			for j := i; i*j < len(primes); j++ {
				primes[i*j] = true
			}
		}
	}

	count := 0
	for i := 2; i < len(primes); i++ {
		if !primes[i] {
			count++
		}
	}
	return count
}
