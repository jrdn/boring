package ds

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

// Get an item from the List
func (l *List[T]) Get(index int) T {
	return l.x[index]
}

// GetSlice the slice that this List is wrapping
func (l *List[T]) GetSlice() []T {
	return l.x
}
