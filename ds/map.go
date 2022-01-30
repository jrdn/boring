package ds

import (
	"github.com/jrdn/boring/iterator"
)

// NewMap creates a new Map
func NewMap[K comparable, V any](data ...map[K]V) *Map[K, V] {
	m := &Map[K, V]{
		x: make(map[K]V),
	}

	for _, d := range data {
		for k, v := range d {
			m.x[k] = v
		}
	}

	return m
}

// Map is an iterable version of a go map
type Map[K comparable, V any] struct {
	x map[K]V
}

// Contains checks if the Map contains a value
func (m Map[K, V]) Contains(key K) bool {
	_, ok := m.x[key]
	return ok
}

// Iter allows the Map to be an iface.Iterable
func (m Map[K, V]) Iter() iterator.Iterable[Pair[K, V]] {
	data := make([]Pair[K, V], len(m.x))
	i := 0
	for key, val := range m.x {
		data[i] = NewPair[K, V](key, val)
		i++
	}
	return NewList(data).Iter()
}

// Get a value from the map
func (m Map[K, V]) Get(key K) (V, bool) {
	ret, ok := m.x[key]
	return ret, ok
}

// GetMap gets the underlying map
func (m Map[K, V]) GetMap() map[K]V {
	return m.x
}

func (m Map[K, V]) Len() int {
	return len(m.x)
}

// NewOrderedMap creates a new ordered map
func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		Map: NewMap[K, V](),
	}
}

type OrderedMap[K comparable, V any] struct {
	*Map[K, V]
	order []K
}

// Append returns true if the item was added, or false if it was not appended (because it already exists)
func (om *OrderedMap[K, V]) Append(key K, value V) bool {
	if om.Contains(key) {
		return false
	}

	om.order = append(om.order, key)
	om.x[key] = value
	return true
}

func (om *OrderedMap[K, V]) Iter() iterator.Iterable[Pair[K, V]] {
	data := make([]Pair[K, V], len(om.order))

	for i, key := range om.order {
		val, _ := om.Get(key)
		data[i] = NewPair[K, V](key, val)
	}
	return NewList(data).Iter()
}
