package mutexmap

import (
	"sync"
)

// Mutexmap2 is simple typesafe wrapper around sync.Map.
// Note that Len() is slow here, as it must read
// through every key in the map.
type Mutexmap2[K comparable, V any] struct {
	m *sync.Map
}

// NewMutexmap creates a new mutex-protected map.
func NewMutexmap2[K comparable, V any]() *Mutexmap2[K, V] {
	return &Mutexmap2[K, V]{
		m: &sync.Map{},
	}
}

// Get returns the value val for key.
func (m *Mutexmap2[K, V]) Get(key K) (val V, ok bool) {
	v, exists := m.m.Load(key)
	if !exists {
		return val, false
	}
	return v.(V), true
}

// Len returns the number of keys. It is not atomic,
// and there may not be a single right answer if there
// are concurrent modifications. It is also slow,
// as it must enumerate all the keys in the sync.Map.
func (m *Mutexmap2[K, V]) Len() (n int) {
	n = 0
	m.m.Range(func(key, value any) bool {
		n++
		return true
	})
	return
}

// Set a single key to value val.
func (m *Mutexmap2[K, V]) Set(key K, val V) {
	m.m.Store(key, val)
}

// Del deletes key from the map.
func (m *Mutexmap2[K, V]) Del(key K) {
	m.m.Delete(key)
}

// GetValNDel returns the val for key, and deletes it.
func (m *Mutexmap2[K, V]) GetValNDel(key K) (val V, ok bool) {
	v, exists := m.m.LoadAndDelete(key)
	if exists {
		val = v.(V)
		ok = true
	}
	return
}

// Clear deletes all keys from the map.
func (m *Mutexmap2[K, V]) Clear() {
	m.m.Clear()
}

// Reset discards map contents, allocating it anew.
func (m *Mutexmap2[K, V]) Reset() {
	m.m.Clear()
}
