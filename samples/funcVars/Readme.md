## Pros
* Different, but doesn't change anything about the invocation! Just like passing dependency as parameter to functions.
* Very efficient; very close to "dependency as parameters" benchmark

## Cons
* Functions are less flexible than structures.
* Can't implement interfaces directly.

## Benchmark (Singleton)
`BenchmarkLibrary-4   	20000000	        99.0 ns/op	      16 B/op	       1 allocs/op
PASS
ok  	github.com/rdadbhawala/dutdip/samples/funcVars	2.281s
Success: Benchmarks passed.`