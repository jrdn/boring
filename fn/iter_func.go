package fn

import (
	"context"

	"github.com/jrdn/boring/types"
)

type IteratorFunc[T any] func() (retVal T, stopIteration bool)

// FuncIterator turns a function into an iterator When a new value is needed
// from the iterator, the function is called to produce it. will continue to
// iterate over the function results in this way while the 2nd return value
// stays false.
func FuncIterator[T any](fn IteratorFunc[T]) types.Iterable[T] {
	return &fnIter[T]{
		fn: fn,
	}
}

type fnIter[T any] struct {
	fn IteratorFunc[T]
}

// Iter iterates the function.
func (iter fnIter[T]) Iter(ctx context.Context) <-chan T {
	c := make(chan T)

	go func() {
		defer close(c)

		stopIteration := false
		ctxDone := ctx.Done()

		for !stopIteration {
			select {
			case <-ctxDone:
				return
			default:
				var ret T
				ret, stopIteration = iter.fn()
				c <- ret
			}
		}
	}()

	return c
}
