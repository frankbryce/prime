// credit goes to...
// https://golang.org/doc/play/sieve.go

// A concurrent prime sieve
package prime

type PrimeSieve struct{}

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func (p PrimeSieve) generate(ch chan<- uint64) {
	for i := uint64(2); ; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func (p PrimeSieve) filter(in <-chan uint64, out chan<- uint64, prime uint64) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}

// The prime sieve: Daisy-chain Filter processes.
func (p PrimeSieve) GetPrimes(n uint) []uint64 {
	ch := make(chan uint64) // Create a new channel.
	go p.generate(ch)       // Launch Generate goroutine.
	primes := make([]uint64, 0, n)
	for i := uint(0); i < n; i++ {
		prime := <-ch
		primes = append(primes, prime)
		ch1 := make(chan uint64)
		go p.filter(ch, ch1, prime)
		ch = ch1
	}
	return primes
}
