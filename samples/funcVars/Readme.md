## Pros
* Different, but doesn't change anything about the invocation! Just like passing dependency as parameter to functions.
* Very efficient; very close to "dependency as parameters" benchmark

## Cons
* Functions are less flexible than structures.
* Can't implement interfaces directly.

## Benchmark (Singleton)
```go
BenchmarkLibrary-4      30000000                49.9 ns/op            16 B/op          1 allocs/op
PASS
ok      github.com/rdadbhawala/dutdip/samples/funcVars  2.790s
```