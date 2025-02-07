mutexmap: generics in action
==========

Wrap your map in a RWMutex to make it goroutine safe.

Generics make this usable.

---
Author: Jason E. Aten, Ph.D.

License: 3-clause BSD style, see the LICENSE file.

---

benchmarks
==========

~~~
/*
go version go1.23.5 darwin/amd64
$ go test -v -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/glycerine/mutexmap
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkMutexMapIntIntWrite
BenchmarkMutexMapIntIntWrite-8          	 9017359	       170.9 ns/op	       9 B/op	       0 allocs/op
BenchmarkSyncMapIntIntWrite
BenchmarkSyncMapIntIntWrite-8           	 4397559	       360.6 ns/op	      65 B/op	       3 allocs/op
BenchmarkMutexMapIntIntRead
BenchmarkMutexMapIntIntRead-8           	24207072	        49.18 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMapIntIntRead
BenchmarkSyncMapIntIntRead-8            	181523422	         7.027 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutexMapIntIntMixed
BenchmarkMutexMapIntIntMixed-8          	23658688	        46.01 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMapIntIntMixed
BenchmarkSyncMapIntIntMixed-8           	27765556	        54.49 ns/op	       5 B/op	       0 allocs/op
BenchmarkMutexMapStringStringWrite
BenchmarkMutexMapStringStringWrite-8    	 3426976	       395.1 ns/op	     166 B/op	       4 allocs/op
BenchmarkSyncMapStringStringWrite
BenchmarkSyncMapStringStringWrite-8     	 2652388	       544.2 ns/op	     206 B/op	       7 allocs/op
BenchmarkMutexMapStringStringRead
BenchmarkMutexMapStringStringRead-8     	19379794	        58.71 ns/op	      50 B/op	       1 allocs/op
BenchmarkSyncMapStringStringRead
BenchmarkSyncMapStringStringRead-8      	31712832	        39.72 ns/op	      50 B/op	       1 allocs/op
BenchmarkMutexMapStringStringMixed
BenchmarkMutexMapStringStringMixed-8    	 6075330	       191.8 ns/op	      61 B/op	       2 allocs/op
BenchmarkSyncMapStringStringMixed
BenchmarkSyncMapStringStringMixed-8     	15622306	        98.89 ns/op	      67 B/op	       2 allocs/op
BenchmarkMutexMap2StringStringWrite
BenchmarkMutexMap2StringStringWrite-8   	 2555968	       525.9 ns/op	     206 B/op	       7 allocs/op
BenchmarkMutexMap2StringStringRead
BenchmarkMutexMap2StringStringRead-8    	29895547	        40.78 ns/op	      50 B/op	       1 allocs/op
BenchmarkMutexMap2StringStringMixed
BenchmarkMutexMap2StringStringMixed-8   	15092859	        95.53 ns/op	      67 B/op	       2 allocs/op
PASS
ok  	github.com/glycerine/mutexmap	23.823s

Compilation finished at Fri Feb  7 16:46:15


With the latest release candidate go1.24rc3 (swiss-tables based map):

go version go1.24rc3 darwin/amd64
go test -v -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/glycerine/mutexmap
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkMutexMapIntIntWrite
BenchmarkMutexMapIntIntWrite-8          	 9030880	       166.8 ns/op	       8 B/op	       0 allocs/op
BenchmarkSyncMapIntIntWrite
BenchmarkSyncMapIntIntWrite-8           	21021043	        87.97 ns/op	      71 B/op	       3 allocs/op
BenchmarkMutexMapIntIntRead
BenchmarkMutexMapIntIntRead-8           	24077414	        48.41 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMapIntIntRead
BenchmarkSyncMapIntIntRead-8            	265517020	         4.445 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutexMapIntIntMixed
BenchmarkMutexMapIntIntMixed-8          	23059905	        50.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMapIntIntMixed
BenchmarkSyncMapIntIntMixed-8           	100000000	        16.28 ns/op	       7 B/op	       0 allocs/op
BenchmarkMutexMapStringStringWrite
BenchmarkMutexMapStringStringWrite-8    	 3201996	       415.0 ns/op	     155 B/op	       4 allocs/op
BenchmarkSyncMapStringStringWrite
BenchmarkSyncMapStringStringWrite-8     	 8965021	       164.2 ns/op	     230 B/op	       7 allocs/op
BenchmarkMutexMapStringStringRead
BenchmarkMutexMapStringStringRead-8     	20419197	        59.68 ns/op	      50 B/op	       1 allocs/op
BenchmarkSyncMapStringStringRead
BenchmarkSyncMapStringStringRead-8      	39455137	        32.21 ns/op	      50 B/op	       1 allocs/op
BenchmarkMutexMapStringStringMixed
BenchmarkMutexMapStringStringMixed-8    	 6571947	       207.3 ns/op	      61 B/op	       2 allocs/op
BenchmarkSyncMapStringStringMixed
BenchmarkSyncMapStringStringMixed-8     	23485053	        50.59 ns/op	      69 B/op	       2 allocs/op
BenchmarkMutexMap2StringStringWrite
BenchmarkMutexMap2StringStringWrite-8   	 8514170	       164.8 ns/op	     230 B/op	       7 allocs/op
BenchmarkMutexMap2StringStringRead
BenchmarkMutexMap2StringStringRead-8    	31341694	        35.44 ns/op	      50 B/op	       1 allocs/op
BenchmarkMutexMap2StringStringMixed
BenchmarkMutexMap2StringStringMixed-8   	23577817	        49.68 ns/op	      69 B/op	       2 allocs/op
PASS
ok  	github.com/glycerine/mutexmap	22.689s

Compilation finished at Fri Feb  7 16:47:57

*/
~~~
