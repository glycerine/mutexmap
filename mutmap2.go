package mutexmap

import (
	"sync"
)

// Mutexmap2 is simple typesafe wrapper around sync.Map.
// The API mirrors that of Mutexmap for drop-in compatability,
// even if the methods are no longer efficient on the
// underlying sync.Map and require iterating the entire
// sync.Map (for example, Len() or GetN() are slow here,
// as they read every key in the map).
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

// Len returns the number of keys. See also GetN()
func (m *Mutexmap2[K, V]) Len() (n int) {
	n = 0
	m.m.Range(func(key, value any) bool {
		n++
		return true
	})
	return
}

// GetValSlice returns all the values in the map in slc.
func (m *Mutexmap2[K, V]) GetValSlice() (slc []V) {
	m.m.Range(func(key, value any) bool {
		slc = append(slc, value.(V))
		return true
	})
	return
}

// GetKeySlice returns all the keys in the map in slc.
func (m *Mutexmap2[K, V]) GetKeySlice() (slc []K) {
	m.m.Range(func(key, value any) bool {
		slc = append(slc, key.(K))
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
// The returned n gives the count of items left in map after deleting key.
func (m *Mutexmap2[K, V]) GetValNDel(key K) (val V, n int, ok bool) {
	v, exists := m.m.LoadAndDelete(key)
	if exists {
		val = v.(V)
		ok = true
	}
	// Count remaining items
	n = 0
	m.m.Range(func(key, value any) bool {
		n++
		return true
	})
	return
}

// GetN returns the number of keys in the map.
func (m *Mutexmap2[K, V]) GetN() (n int) {
	n = 0
	m.m.Range(func(key, value any) bool {
		n++
		return true
	})
	return
}

// Clear deletes all keys from the map.
func (m *Mutexmap2[K, V]) Clear() {
	m.m.Range(func(key, value any) bool {
		m.m.Delete(key)
		return true
	})
}

// Update runs updateFunc on the Mutexmap2. It is _not_ atomic.
// It is not efficient, or cheap.
func (m *Mutexmap2[K, V]) Update(updateFunc func(m map[K]V)) {
	// Create a temporary map to work with
	tempMap := make(map[K]V)
	m.m.Range(func(key, value any) bool {
		tempMap[key.(K)] = value.(V)
		return true
	})

	// Apply the update function
	updateFunc(tempMap)

	// Clear and restore from the temp map
	m.Clear()
	for k, v := range tempMap {
		m.m.Store(k, v)
	}
}

// GetMapReset returns the underlying map and
// resets the internal map by re-allocating it anew.
// This is useful when you want to discard
// synchronization going forward. It is not efficient.
func (m *Mutexmap2[K, V]) GetMapReset() (mm map[K]V) {
	mm = make(map[K]V)
	m.m.Range(func(key, value any) bool {
		mm[key.(K)] = value.(V)
		return true
	})
	m.m = &sync.Map{}
	return
}

// Reset discards map contents, allocating it anew.
func (m *Mutexmap2[K, V]) Reset() {
	m.m = &sync.Map{}
}
