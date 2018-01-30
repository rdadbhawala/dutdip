## Pros
* Different, but doesn't change anything about the invocation! Just like passing dependency as parameter to functions.
* Very efficient; very close to "dependency as parameters" benchmark

## Cons
* Functions are less flexible than structures.
* Can't implement interfaces directly.

## Benchmark (Singleton)
```go
BenchmarkLibrary-4      30000000                47.1 ns/op            16 B/op          1 allocs/op
BenchmarkSingleton-4    1000000000               2.71 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/rdadbhawala/dutdip/samples/funcVars  5.320s
```
