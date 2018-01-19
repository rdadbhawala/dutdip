## Pros
* Simple to implement
* Aligned with "New" functions

## Cons
* Every call to New is tightly coupled with Dependency
* Dependency is initialized before Service
* More orchestration requires "package" variables, structures, functions
* Singletons required a separate function to setup.

## Benchmarking:
```go
BenchmarkService-4      2000000000               1.55 ns/op            0 B/op          0 allocs/op
BenchmarkSintleton-4    2000000000               0.39 ns/op            0 B/op          0 allocs/op
PASS
ok      github.com/rdadbhawala/dutdip/samples/depParams 4.372s
```
