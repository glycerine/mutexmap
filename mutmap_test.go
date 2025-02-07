package mutexmap

import (
	"strconv"
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

var testKeyString = "my somewhat long string key for benchmarking "
var testValString = "my somewhat long string value for benchmarking "

func BenchmarkMutexMapStringStringWrite(b *testing.B) {
	m := NewMutexmap[string, string]()
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := testKeyString + strconv.Itoa(i)
			val := testValString + strconv.Itoa(i)
			m.Set(key, val)
			i++
		}
	})
}

func BenchmarkSyncMapStringStringWrite(b *testing.B) {
	var m sync.Map
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := testKeyString + strconv.Itoa(i)
			val := testValString + strconv.Itoa(i)
			m.Store(key, val)
			i++
		}
	})
}

func BenchmarkMutexMapStringStringRead(b *testing.B) {
	m := NewMutexmap[string, string]()
	for i := 0; i < 1000; i++ {
		key := testKeyString + strconv.Itoa(i)
		val := testValString + strconv.Itoa(i)
		m.Set(key, val)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := testKeyString + strconv.Itoa(i%1000)
			m.Get(key)
			i++
		}
	})
}

func BenchmarkSyncMapStringStringRead(b *testing.B) {
	var m sync.Map
	for i := 0; i < 1000; i++ {
		key := testKeyString + strconv.Itoa(i)
		val := testValString + strconv.Itoa(i)
		m.Store(key, val)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			key := testKeyString + strconv.Itoa(i%1000)
			m.Load(key)
			i++
		}
	})
}

func BenchmarkMutexMapStringStringMixed(b *testing.B) {
	m := NewMutexmap[string, string]()
	for i := 0; i < 1000; i++ {
		key := testKeyString + strconv.Itoa(i)
		val := testValString + strconv.Itoa(i)
		m.Set(key, val)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%10 == 0 { // 10% writes, 90% reads
				key := testKeyString + strconv.Itoa(i)
				val := testValString + strconv.Itoa(i)
				m.Set(key, val)
			} else {
				key := testKeyString + strconv.Itoa(i%1000)
				m.Get(key)
			}
			i++
		}
	})
}

func BenchmarkSyncMapStringStringMixed(b *testing.B) {
	var m sync.Map
	for i := 0; i < 1000; i++ {
		key := testKeyString + strconv.Itoa(i)
		val := testValString + strconv.Itoa(i)
		m.Store(key, val)
	}
	b.ResetTimer()

	b.RunParallel(func(pb *testing.PB) {
		i := 0
		for pb.Next() {
			if i%10 == 0 { // 10% writes, 90% reads
				key := testKeyString + strconv.Itoa(i)
				val := testValString + strconv.Itoa(i)
				m.Store(key, val)
			} else {
				key := testKeyString + strconv.Itoa(i%1000)
				m.Load(key)
			}
			i++
		}
	})
}

/*
go version go1.23.5 darwin/amd64
$ go test -v -bench=. -benchmem
goos: darwin
goarch: amd64
pkg: github.com/glycerine/mutexmap
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkMutexMapIntIntWrite
BenchmarkMutexMapIntIntWrite-8         	 9267024	       168.4 ns/op	       9 B/op	       0 allocs/op
BenchmarkSyncMapIntIntWrite
BenchmarkSyncMapIntIntWrite-8          	 4195042	       332.9 ns/op	      52 B/op	       3 allocs/op
BenchmarkMutexMapIntIntRead
BenchmarkMutexMapIntIntRead-8          	25707594	        45.97 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMapIntIntRead
BenchmarkSyncMapIntIntRead-8           	240654564	         6.665 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutexMapIntIntMixed
BenchmarkMutexMapIntIntMixed-8         	24486024	        43.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMapIntIntMixed
BenchmarkSyncMapIntIntMixed-8          	28197913	        53.17 ns/op	       5 B/op	       0 allocs/op
BenchmarkMutexMapStringStringWrite
BenchmarkMutexMapStringStringWrite-8   	 3398826	       400.0 ns/op	     166 B/op	       4 allocs/op
BenchmarkSyncMapStringStringWrite
BenchmarkSyncMapStringStringWrite-8    	 2379595	       517.3 ns/op	     207 B/op	       7 allocs/op
BenchmarkMutexMapStringStringRead
BenchmarkMutexMapStringStringRead-8    	20131230	        58.07 ns/op	      50 B/op	       1 allocs/op
BenchmarkSyncMapStringStringRead
BenchmarkSyncMapStringStringRead-8     	34885687	        39.21 ns/op	      50 B/op	       1 allocs/op
BenchmarkMutexMapStringStringMixed
BenchmarkMutexMapStringStringMixed-8   	 6050472	       191.7 ns/op	      61 B/op	       2 allocs/op
BenchmarkSyncMapStringStringMixed
BenchmarkSyncMapStringStringMixed-8    	15723597	        96.59 ns/op	      67 B/op	       2 allocs/op
PASS
ok  	github.com/glycerine/mutexmap	19.178s


With the latest release candidate go1.24rc3 (swiss-tables based map):

go version go1.24rc3 darwin/amd64
goos: darwin
goarch: amd64
pkg: github.com/glycerine/mutexmap
cpu: Intel(R) Core(TM) i7-1068NG7 CPU @ 2.30GHz
BenchmarkMutexMapIntIntWrite
BenchmarkMutexMapIntIntWrite-8         	 8961474	       182.2 ns/op	       8 B/op	       0 allocs/op
BenchmarkSyncMapIntIntWrite
BenchmarkSyncMapIntIntWrite-8          	18723195	        87.84 ns/op	      71 B/op	       3 allocs/op
BenchmarkMutexMapIntIntRead
BenchmarkMutexMapIntIntRead-8          	23856789	        48.31 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMapIntIntRead
BenchmarkSyncMapIntIntRead-8           	270451916	         4.447 ns/op	       0 B/op	       0 allocs/op
BenchmarkMutexMapIntIntMixed
BenchmarkMutexMapIntIntMixed-8         	22815286	        51.06 ns/op	       0 B/op	       0 allocs/op
BenchmarkSyncMapIntIntMixed
BenchmarkSyncMapIntIntMixed-8          	100000000	        16.38 ns/op	       7 B/op	       0 allocs/op
BenchmarkMutexMapStringStringWrite
BenchmarkMutexMapStringStringWrite-8   	 3173473	       409.4 ns/op	     155 B/op	       4 allocs/op
BenchmarkSyncMapStringStringWrite
BenchmarkSyncMapStringStringWrite-8    	 8589674	       170.3 ns/op	     230 B/op	       7 allocs/op
BenchmarkMutexMapStringStringRead
BenchmarkMutexMapStringStringRead-8    	19470552	        61.15 ns/op	      50 B/op	       1 allocs/op
BenchmarkSyncMapStringStringRead
BenchmarkSyncMapStringStringRead-8     	40280904	        32.13 ns/op	      50 B/op	       1 allocs/op
BenchmarkMutexMapStringStringMixed
BenchmarkMutexMapStringStringMixed-8   	 6546421	       187.3 ns/op	      61 B/op	       2 allocs/op
BenchmarkSyncMapStringStringMixed
BenchmarkSyncMapStringStringMixed-8    	24407060	        50.42 ns/op	      69 B/op	       2 allocs/op
PASS
ok  	github.com/glycerine/mutexmap	18.943s
*/
