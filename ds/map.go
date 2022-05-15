package ds

import (
	"context"

	"github.com/jrdn/boring/types"
)

// NewMap creates a new Map.
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

// Map is an iterable version of a go map.
type Map[K comparable, V any] struct {
	x map[K]V
}

// Contains checks if the Map contains a value.
func (m Map[K, V]) Contains(key K) bool {
	_, ok := m.x[key]

	return ok
}

// Iter allows the Map to be an iface.Iterable.
func (m Map[K, V]) Iter(ctx context.Context) <-chan Pair[K, V] {
	iterChan := make(chan Pair[K, V])

	go func() {
		defer close(iterChan)

		for k, v := range m.x {
			select {
			case <-ctx.Done():
				return // break
			case iterChan <- NewPair[K, V](k, v):
				// sent the message
			} //nolint:wsl
		}
	}()

	return iterChan
}

// Get a value from the map.
func (m Map[K, V]) Get(key K) (V, bool) {
	ret, ok := m.x[key]

	return ret, ok
}

// GetMap gets the underlying map.
func (m Map[K, V]) GetMap() map[K]V {
	return m.x
}

func (m Map[K, V]) Len() int {
	return len(m.x)
}

var (
	_ types.Iterable[Pair[string, int]] = &Map[string, int]{}
	_ types.Lengthable                  = &Map[string, int]{}
)

//
// // NewOrderedMap creates a new ordered map
// func NewOrderedMap[K comparable, V any]() *OrderedMap[K, V] {
// 	return &OrderedMap[K, V]{
// 		Map: NewMap[K, V](),
// 	}
// }
//
// type OrderedMap[K comparable, V any] struct {
// 	*Map[K, V]
// 	order []K
// }
//
// // Append returns true if the item was added, or false if it was not appended (because it already exists)
// func (om *OrderedMap[K, V]) Append(key K, value V) bool {
// 	if om.Contains(key) {
// 		return false
// 	}
//
// 	om.order = append(om.order, key)
// 	om.x[key] = value
// 	return true
// }
//
// func (om *OrderedMap[K, V]) Iter( ) <-chan Pair[K, V] {
// 	iterChan := make(chan Pair[K, V])
// 	go func() {
// 		defer close(iterChan)
//
// 		for _, k := range om.order {
// 			v := om.x[k]
// 			iterChan <- NewPair[K, V](k, v)
// 		}
// 		close(iterChan)
// 	}()
// 	return iterChan
// }
//
// var _ types.Iterable[Pair[string, int]] = &OrderedMap[string, int]{}
// var _ types.Lengthable = &OrderedMap[string, int]{}
