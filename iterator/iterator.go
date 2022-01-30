package iterator

import (
	"context"
	"runtime"
)

type Generator[T any] func() T

func (g Generator[T]) Iter() (chan T, func()) {
	ctx, cancel := context.WithCancel(context.Background())
	results := make(chan T)

	go func() {
		defer close(results)
		done := ctx.Done()
		for {
			val := g()
			select {
			case results <- val:
				// sent value
			case <-done:
				// cancelled
				return
			}
		}
	}()

	return results, cancel
}

func NewIterator[T any](gen Generator[T]) *Iterator[T] {
	i := &Iterator[T]{
		gen: gen,
	}
	runtime.SetFinalizer(i, func(iterator *Iterator[T]) {
		iterator.Close()
	})
	return i
}

type Iterable[T any] interface {
	Iter() chan T
	Close()
}

type Iterator[T any] struct {
	gen    Generator[T]
	cancel func()
}

func (i *Iterator[T]) Close() {
	i.cancel()
}

func (i *Iterator[T]) Iter() chan T {
	valueChan, cancel := i.gen.Iter()
	i.cancel = cancel
	return valueChan
}
