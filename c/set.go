package c

func NewSet[T comparable](data ...T) *Set[T] {
	s := &Set[T]{
		x: make(map[T]struct{}),
	}
	for _, item := range data {
		s.Add(item)
	}
	return s
}

type Set[T comparable] struct {
	x map[T]struct{}
}

func (s *Set[T]) Add(item T) {
	s.x[item] = struct{}{}
}

func (s *Set[T]) Has(item T) bool {
	_, ok := s.x[item]
	return ok
}

func (s *Set[T]) List() []T {
	ret := make([]T, len(s.x))
	j := 0
	for item := range s.x {
		ret[j] = item
		j += 1
	}
	return ret
}

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
