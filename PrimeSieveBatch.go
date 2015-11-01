// A concurrent prime sieve
package prime

type PrimeSieveBatch struct {
	BatchSize uint
}

// Send the sequence 2, 3, 4, ... to channel 'ch'.
func (p PrimeSieveBatch) generate(ch chan<- []uint64) {
	outs := make([]uint64, 0, p.BatchSize)
	i := uint64(2)
	for ; ; i++ {
		addToOuts := true
		for j := 0; j < len(outs); j++ {
			if i%outs[j] == 0 {
				addToOuts = false
				break
			}
		}

		if addToOuts {
			outs = append(outs, i)
			if len(outs) == cap(outs) {
				ch <- outs
				outs = make([]uint64, 0, p.BatchSize)
				break
			}
		}
	}

	for ; ; i++ {
		outs = append(outs, i)
		if len(outs) == cap(outs) {
			ch <- outs
			outs = make([]uint64, 0, p.BatchSize)
		}
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func (p PrimeSieveBatch) filter(out chan<- []uint64, primes []uint64, in <-chan []uint64) {
	outs := make([]uint64, 0, p.BatchSize)
	for {
		ins := <-in // Receive value from 'in'.
		for i := 0; i < len(ins); i++ {
			addToOuts := true
			for j := 0; j < len(primes); j++ {
				if ins[i]%primes[j] == 0 {
					addToOuts = false
					break
				}
			}

			if addToOuts {
				outs = append(outs, ins[i])
				if len(outs) == cap(outs) {
					out <- outs
					outs = make([]uint64, 0, p.BatchSize)
				}
			}
		}
		ins = nil // let the garbage collector get this stuff as soon as possible
	}
}

// The prime sieve: Daisy-chain Filter processes.
func (p PrimeSieveBatch) GetPrimes(n uint) []uint64 {
	primes := make([]uint64, 0, n+p.BatchSize)
	ch := make(chan []uint64)
	go p.generate(ch)
	for len(primes) < int(n) {
		newPrimes := <-ch
		primes = append(primes, newPrimes...)
		ch1 := make(chan []uint64)
		go p.filter(ch1, newPrimes, ch)
		ch = ch1
	}
	return primes
}
