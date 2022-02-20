package iterator

import (
	"context"
	"fmt"
	"runtime"
	"sync"

	"github.com/jrdn/boring/types"
)

func NewIterator[T any](gen types.Generator[T]) types.Iterator[T] {
	ctx, cancel := context.WithCancel(context.Background())
	iter := &iterator[T]{
		c:      gen(ctx),
		ctx:    ctx,
		cancel: cancel,
		m:      &sync.Mutex{},
	}
	runtime.SetFinalizer(iter, func(i *iterator[T]) {
		fmt.Println("FINALIZING") // TODO debugging
		i.Close()
	})

	return iter
}

type iterator[T any] struct {
	c      chan T
	cancel func()
	closed bool
	ctx    context.Context
	m      *sync.Mutex
}

func (i *iterator[T]) Derive(gen types.Generator[T]) types.Iterator[T] {
	ctx, cancel := context.WithCancel(i.ctx)
	iter := &iterator[T]{
		c:      gen(ctx),
		ctx:    ctx,
		cancel: cancel,
		m:      &sync.Mutex{},
	}
	runtime.SetFinalizer(iter, func(i *iterator[T]) {
		fmt.Println("FINALIZING DERIVED") // TODO
		i.Close()
	})
	return iter
}

func (i *iterator[T]) Range() <-chan T {
	i.m.Lock()
	defer i.m.Unlock()
	return i.c
}

func (i *iterator[T]) Close() {
	i.m.Lock()
	defer i.m.Unlock()

	if !i.closed {
		i.cancel()
		close(i.c)
		i.closed = true
	}
}
