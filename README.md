I saw a cool piece of code which showed how go programs have a life of their own.

https://golang.org/doc/play/sieve.go

I also saw that this was probably not as efficient as it could be, since the channels
_have_ to be the bottleneck for large number of prime numbers!  I extended the
daisy chain filtering idea, but instead each layer is responsible for filtering then
passing on 1 value, I allow it to pass on up to N values.

I created some benchmarks (run with -bench=5000 to see cool results) to see that batching
is actually helping!
