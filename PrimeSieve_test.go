package prime

import "testing"

func BenchmarkPrimeSieve_1000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieve{}, 1000, b)
}

func BenchmarkPrimeSieve_5000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieve{}, 5000, b)
}

func BenchmarkPrimeSieve_10000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieve{}, 10000, b)
}

func TestPrimeSieve_10primes(t *testing.T) {
	PrimeGetterTest(PrimeSieve{}, 10, 29, t)
}

func TestPrimeSieve_100primes(t *testing.T) {
	PrimeGetterTest(PrimeSieve{}, 100, 541, t)
}

func TestPrimeSieve_1000primes(t *testing.T) {
	PrimeGetterTest(PrimeSieve{}, 1000, 7919, t)
}
