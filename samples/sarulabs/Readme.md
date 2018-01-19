## Pros
* Definition is loosely coupled from instance variables.
* Definitions are named; allowing multiple definitions of same implementation
* Allows all types of injections.
* Compatible with "New" functions
* Supports "Scopes", nice relationships between Scopes
* "More" efficient when creating separate instances
* 

## Cons
* Requires casting
* Singletons within a context, means creating a separate scope/ context for each request which doesn't rely on singletons.

## Benchmark (Singleton)
```
BenchmarkLibrary-4        500000              2601 ns/op            1118 B/op         12 allocs/op
BenchmarkSingleton-4    20000000                84.5 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/rdadbhawala/dutdip/samples/sarulabs  3.962s
```
