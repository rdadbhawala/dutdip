## Pros
* Simple to implement
* Aligned with "New" functions

## Cons
* Every call to New is tightly coupled with Dependency
* Dependency is initialized before Service

## Benchmarking:
BenchmarkService-4   	2000000000	         1.58 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/rdadbhawala/dutdip/samples/depParams	3.540s
Success: Benchmarks passed.