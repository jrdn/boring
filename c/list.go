package c

// NewList creates a new List, which is an iterable analogous to a slice
func NewList[T any](data ...[]T) *List[T] {
	l := &List[T]{}
	for _, d := range data {
		l.x = append(l.x, d...)
	}
	return l
}

// List is an iterable version of a slice
type List[T any] struct {
	x []T
}

// Get the slice that this List is wrapping
func (l *List[T]) Get() []T {
	return l.x
}

// Iter returns a channel which can be used with range to loop through the iterator
func (l *List[T]) Iter() <-chan T {
	iterChan := make(chan T)
	go func() {
		for _, item := range l.x {
			iterChan <- item
		}
		close(iterChan)
	}()
	return iterChan
}
