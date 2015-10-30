package prime

type PrimeGetter interface {
	GetPrimes(n uint) (primes []uint64)
}
