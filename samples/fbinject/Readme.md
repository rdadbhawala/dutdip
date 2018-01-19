## Pros


## Cons
* Service itself is a struct. Not using an interface.
* Couldn't resolve multiple instances of dependency; was always a Singleton.
* May be an error in benchmark test, but current numbers are terrible.

## Benchmark
```go
BenchmarkLibrary-4   	   10000	   3716104 ns/op	  160497 B/op	   20008 allocs/op
PASS
ok  	github.com/rdadbhawala/dutdip/samples/fbinject	38.098s
Success: Benchmarks passed.
```
