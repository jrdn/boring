package ds

import (
	"context"

	"github.com/jrdn/boring/iterator"
	"github.com/jrdn/boring/types"
)

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

// Iter allows the List to be an iface.Iterable
func (l *List[T]) Iter() types.Iterator[T] {
	return iterator.NewIterator[T](func(ctx context.Context) <-chan T {
		iterChan := make(chan T)
		go func(ctx context.Context) {
			done := ctx.Done()
			defer close(iterChan)
			for _, item := range l.x {
				select {
				case <-done:
					return
				case iterChan <- item:
					// send successful
				}

			}
		}(ctx)
		return iterChan
	})
}
