package ds

// NewSet creates a new Set
func NewSet[T comparable](data ...[]T) *Set[T] {
	s := &Set[T]{
		x: make(map[T]struct{}),
	}
	for _, d := range data {
		for _, item := range d {
			s.Add(item)
		}
	}
	return s
}

// Set is an unordered container which holds an unordered set of items with no duplicates
type Set[T comparable] struct {
	x map[T]struct{}
}

// Add an item to the set
func (s *Set[T]) Add(item T) {
	s.x[item] = struct{}{}
}

// Contains checks if the set contains an item
func (s *Set[T]) Contains(item T) bool {
	_, ok := s.x[item]
	return ok
}

// Slice returns the contents of the set as a slice
func (s *Set[T]) Slice() []T {
	ret := make([]T, len(s.x))
	j := 0
	for item := range s.x {
		ret[j] = item
		j += 1
	}
	return ret
}

// Iter allows the set to be an iface.Iterable
func (s *Set[T]) Iter() <-chan T {
	iterChan := make(chan T)
	go func(resultChan chan T, data map[T]struct{}) {
		for item := range data {
			resultChan <- item
		}
		close(iterChan)
	}(iterChan, s.x)
	return iterChan
}
