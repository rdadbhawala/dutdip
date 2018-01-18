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
`BenchmarkLibrary-4   	20000000	        73.7 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/rdadbhawala/dutdip/samples/sarulabs	2.702s
Success: Benchmarks passed.`

### With Casting (Singleton)
`BenchmarkLibrary-4   	20000000	        83.7 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/rdadbhawala/dutdip/samples/sarulabs	2.077s
Success: Benchmarks passed.`

### With separate scopes
`BenchmarkLibrary-4   	 1000000	      1146 ns/op	     429 B/op	       5 allocs/op
PASS
ok  	github.com/rdadbhawala/dutdip/samples/sarulabs	1.926s
Success: Benchmarks passed.`

### Request Scope
`BenchmarkLibrary-4   	  500000	      3064 ns/op	    1118 B/op	      12 allocs/op
PASS
ok  	github.com/rdadbhawala/dutdip/samples/sarulabs	2.183s
Success: Benchmarks passed.`
