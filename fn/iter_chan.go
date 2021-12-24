package fn

import "github.com/jrdn/boring/iface"

// IterableChannel wraps a channel in an iterator
func IterableChannel[T any](c <-chan T) iface.Iterable[T] {
	return &chanIterator[T]{c: c}
}

type chanIterator[T any] struct {
	c <-chan T
}

// Iter iterates the channel
func (c *chanIterator[T]) Iter() <-chan T {
	retChan := make(chan T)
	go func() {
		for x := range c.c {
			retChan <- x
		}
		close(retChan)
	}()
	return retChan
}
