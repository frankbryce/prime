package prime

import "testing"

func GetPrimesBenchmark(p PrimeGetter, n uint, b *testing.B) {
	for i := 0; i < b.N; i++ {
		p.GetPrimes(n)
	}
}

func PrimeGetterTest(p PrimeGetter, n uint, exp uint64, t *testing.T) {
	ans := p.GetPrimes(n)
	if ans[n-1] != exp {
		t.Error("Answer: ", ans[n-1], ", was unexpected: ", exp)
	}
}
