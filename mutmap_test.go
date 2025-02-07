package mutexmap

import (
	"sync"
	"testing"
)

func BenchmarkMutexMapIntIntWrite(b *testing.B) {
	m := NewMutexmap[int, int]()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m.Set(i, i)
			i++
		}
	})
}

func BenchmarkSyncMapIntIntWrite(b *testing.B) {
	var m sync.Map
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m.Store(i, i)
			i++
		}
	})
}

func BenchmarkMutexMapIntIntRead(b *testing.B) {
	m := NewMutexmap[int, int]()
	for i := 0; i < 1000; i++ {
		m.Set(i, i)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m.Get(i % 1000)
			i++
		}
	})
}

func BenchmarkSyncMapIntIntRead(b *testing.B) {
	var m sync.Map
	for i := 0; i < 1000; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			m.Load(i % 1000)
			i++
		}
	})
}

func BenchmarkMutexMapIntIntMixed(b *testing.B) {
	m := NewMutexmap[int, int]()
	for i := 0; i < 1000; i++ {
		m.Set(i, i)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%10 == 0 { // 10% writes, 90% reads
				m.Set(i, i)
			} else {
				m.Get(i % 1000)
			}
			i++
		}
	})
}

func BenchmarkSyncMapIntIntMixed(b *testing.B) {
	var m sync.Map
	for i := 0; i < 1000; i++ {
		m.Store(i, i)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%10 == 0 { // 10% writes, 90% reads
				m.Store(i, i)
			} else {
				m.Load(i % 1000)
			}
			i++
		}
	})
}

/*
go version go1.23.5 darwin/amd64
go test -v -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/glycerine/mutexmap
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkMutexMapIntIntWrite
BenchmarkMutexMapIntIntWrite-8   	 9231715	       175.4 ns/op	       9 B/op	       0 allocs/op
BenchmarkSyncMapIntIntWrite
BenchmarkSyncMapIntIntWrite-8    	 4430647	       333.5 ns/op	      51 B/op	       3 allocs/op
BenchmarkMutexMapIntIntRead
BenchmarkMutexMapIntIntRead-8    	24757224	        47.64 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMapIntIntRead
BenchmarkSyncMapIntIntRead-8     	242166340	         6.652 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutexMapIntIntMixed
BenchmarkMutexMapIntIntMixed-8   	23693816	        43.65 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMapIntIntMixed
BenchmarkSyncMapIntIntMixed-8    	27424677	        53.16 ns/op	       5 B/op	       0 allocs/op
PASS
ok  	github.com/glycerine/mutexmap	9.645s


With the latest release candidate go1.24rc3 (swiss-tables based map):

go version go1.24rc3 darwin/amd64
jaten@Js-MacBook-Pro ~/go/src/github.com/glycerine/mutexmap (master) $ go test -v -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/glycerine/mutexmap
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkMutexMapIntIntWrite
BenchmarkMutexMapIntIntWrite-8   	 9120490	       159.9 ns/op	       8 B/op	       0 allocs/op
BenchmarkSyncMapIntIntWrite
BenchmarkSyncMapIntIntWrite-8    	19779579	        91.20 ns/op	      71 B/op	       3 allocs/op
BenchmarkMutexMapIntIntRead
BenchmarkMutexMapIntIntRead-8    	25259208	        46.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMapIntIntRead
BenchmarkSyncMapIntIntRead-8     	274623240	         4.508 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutexMapIntIntMixed
BenchmarkMutexMapIntIntMixed-8   	23732199	        49.48 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMapIntIntMixed
BenchmarkSyncMapIntIntMixed-8    	100000000	        15.60 ns/op	       7 B/op	       0 allocs/op
PASS
ok  	github.com/glycerine/mutexmap	9.470s

*/
