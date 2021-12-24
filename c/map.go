package c

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

type Map[K comparable, V any] struct {
	x map[K]V
}

func (m Map[K, V]) Contains(key K) bool {
	_, ok := m.x[key]
	return ok
}

func (m Map[K, V]) Iter() <-chan Pair[K, V] {
	iterChan := make(chan Pair[K, V])
	go func() {
		for k, v := range m.x {
			iterChan <- NewPair[K, V](k, v)
		}
		close(iterChan)
	}()
	return iterChan
}
