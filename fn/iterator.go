package fn

import "github.com/jrdn/boring/iface"

type IteratorFunc[T any] func() (retVal T, stopIteration bool)

// Iterator turns a function into an iterator
// When a new value is needed from the iterator, the function is called to produce it
// It will continue to iterate over the function results in this way while the 2nd return value stays false.
func Iterator[T any](fn IteratorFunc[T]) iface.Iterable[T] {
	return &fnIter[T]{fn: fn}
}

type fnIter[T any] struct {
	fn func() (T, bool)
}

// Iter iterates the function
func (iter fnIter[T]) Iter() <-chan T {
	c := make(chan T)
	go func() {
		ret, stopIteration := iter.fn()
		c <- ret
		for !stopIteration {
			ret, stopIteration = iter.fn()
			c <- ret
		}
		close(c)
	}()

	return c
}
