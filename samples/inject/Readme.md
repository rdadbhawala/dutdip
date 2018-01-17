## Pros
* Can work with existing New functions

## Cons
* Dependencies become Singleton. "Provider" definition of the dependency couldn't be reused.
* Significantly inefficient than the "New" functions. AutoProvider is worse than normal Provider.
* Forces Singletons of Dependencies/ Services. Eliminates choice.
* If the primitives passed to a constructor are changing, new "Definition" instances are required.
* Do I need to create a new Definition for every instantiation of Service/ Dependency in code?

## Benchmark
`BenchmarkLibrary-4   	  300000	      6853 ns/op	     685 B/op	      14 allocs/op
PASS
ok  	github.com/rdadbhawala/dutdip/samples/inject	2.320s
Success: Benchmarks passed.`