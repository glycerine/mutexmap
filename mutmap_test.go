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

// Mutexmap2: sync.Map with type-safe wrapper:

func BenchmarkMutexMap2StringStringWrite(b *testing.B) {
	m := NewMutexmap2[string, string]()
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

func BenchmarkMutexMap2StringStringRead(b *testing.B) {
	m := NewMutexmap2[string, string]()
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

func BenchmarkMutexMap2StringStringMixed(b *testing.B) {
	m := NewMutexmap2[string, string]()
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
