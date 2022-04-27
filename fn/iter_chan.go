package fn

import (
	"context"

	"github.com/jrdn/boring/types"
)

// ChanIterator wraps a channel in an iterator
func ChanIterator[T any](c <-chan T) types.Iterable[T] {
	return &chanIterator[T]{c: c}
}

type chanIterator[T any] struct {
	c <-chan T
}

// Iter iterates the channel
func (c *chanIterator[T]) Iter(ctx context.Context) <-chan T {
	retChan := make(chan T)
	go func() {
		for x := range c.c {
			select {
			case <-ctx.Done():
				return
			case retChan <- x:
			}
		}
		close(retChan)
	}()
	return retChan
}
