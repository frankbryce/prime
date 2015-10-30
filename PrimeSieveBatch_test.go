package prime

import "testing"

func BenchmarkPrimeSieveBatch_1batch_1000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 1}, 1000, b)
}

func BenchmarkPrimeSieveBatch_2batch_1000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 2}, 1000, b)
}

func BenchmarkPrimeSieveBatch_5batch_1000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 5}, 1000, b)
}

func BenchmarkPrimeSieveBatch_10batch_1000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 10}, 1000, b)
}

func BenchmarkPrimeSieveBatch_20batch_1000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 20}, 1000, b)
}

func BenchmarkPrimeSieveBatch_1batch_10000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 1}, 10000, b)
}

func BenchmarkPrimeSieveBatch_2batch_10000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 2}, 10000, b)
}

func BenchmarkPrimeSieveBatch_5batch_10000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 5}, 10000, b)
}

func BenchmarkPrimeSieveBatch_10batch_10000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 10}, 10000, b)
}

func BenchmarkPrimeSieveBatch_20batch_10000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 20}, 10000, b)
}

func BenchmarkPrimeSieveBatch_50batch_10000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 50}, 10000, b)
}

func BenchmarkPrimeSieveBatch_100batch_10000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 100}, 10000, b)
}

func BenchmarkPrimeSieveBatch_500batch_10000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 500}, 10000, b)
}

func BenchmarkPrimeSieveBatch_1000batch_10000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 1000}, 10000, b)
}

func BenchmarkPrimeSieveBatch_10000batch_10000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 10000}, 10000, b)
}

func BenchmarkPrimeSieveBatch_1batch_5000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 1}, 5000, b)
}

func BenchmarkPrimeSieveBatch_3batch_5000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 3}, 5000, b)
}

func BenchmarkPrimeSieveBatch_10batch_5000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 10}, 5000, b)
}

func BenchmarkPrimeSieveBatch_30batch_5000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 30}, 5000, b)
}

func BenchmarkPrimeSieveBatch_100batch_5000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 100}, 5000, b)
}

func BenchmarkPrimeSieveBatch_300batch_5000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 300}, 5000, b)
}

func BenchmarkPrimeSieveBatch_1000batch_5000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 1000}, 5000, b)
}

func BenchmarkPrimeSieveBatch_3000batch_5000primes(b *testing.B) {
	GetPrimesBenchmark(PrimeSieveBatch{batchSize: 3000}, 5000, b)
}

func TestPrimeSieveBatch_1batch_10primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 1}, 10, 29, t)
}

func TestPrimeSieveBatch_1batch_100primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 1}, 100, 541, t)
}

func TestPrimeSieveBatch_1batch_1000primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 1}, 1000, 7919, t)
}

func TestPrimeSieveBatch_2batch_10primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 2}, 10, 29, t)
}

func TestPrimeSieveBatch_2batch_100primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 2}, 100, 541, t)
}

func TestPrimeSieveBatch_2batch_1000primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 2}, 1000, 7919, t)
}

func TestPrimeSieveBatch_5batch_10primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 5}, 10, 29, t)
}

func TestPrimeSieveBatch_5batch_100primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 5}, 100, 541, t)
}

func TestPrimeSieveBatch_5batch_1000primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 5}, 1000, 7919, t)
}

func TestPrimeSieveBatch_13batch_10primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 13}, 10, 29, t)
}

func TestPrimeSieveBatch_13batch_100primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 13}, 100, 541, t)
}

func TestPrimeSieveBatch_13batch_1000primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 13}, 1000, 7919, t)
}

func TestPrimeSieveBatch_37batch_10primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 13}, 10, 29, t)
}

func TestPrimeSieveBatch_37batch_100primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 13}, 100, 541, t)
}

func TestPrimeSieveBatch_37batch_1000primes(t *testing.T) {
	PrimeGetterTest(PrimeSieveBatch{batchSize: 13}, 1000, 7919, t)
}
