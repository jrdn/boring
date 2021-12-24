package fn

import (
	"github.com/jrdn/boring/c"
	"github.com/jrdn/boring/iface"
)

// Collect gathers all the values from an iterable and converts it to a List
func Collect[T any](iter ...iface.Iterable[T]) *c.List[T] {
	var ret []T
	for item := range Chain(iter...).Iter() {
		ret = append(ret, item)
	}
	return c.NewList[T](ret)
}

// CollectN gathers the first N values from an iterable and adds it to a List until the list contains
func CollectN[T any](n int, iter ...iface.Iterable[T]) *c.List[T] {
	var ret []T

	counter := 0
	for item := range Chain(iter...).Iter() {
		ret = append(ret, item)
		counter += 1
		if counter == n {
			break
		}
	}
	return c.NewList[T](ret)
}

type CollectUntilDecider[T any] func(item T) (stopIteration bool)

// CollectUntil gathers values from iterables into a List until a provided function returns true
func CollectUntil[T any](fn CollectUntilDecider[T], iter ...iface.Iterable[T]) *c.List[T] {
	var ret []T
	for item := range Chain(iter...).Iter() {
		if stopIteration := fn(item); stopIteration {
			break
		}
		ret = append(ret, item)
	}
	return c.NewList[T](ret)
}

// Map applies a function to a number of iterators and produces an iterable of the function's return values
func Map[T, R any](fn func(T) R, iterators ...iface.Iterable[T]) iface.Iterable[R] {
	c := make(chan R)

	go func() {
		chain := Chain[T](iterators...)
		for item := range chain.Iter() {
			c <- fn(item)
		}
		close(c)
	}()
	return IterableChannel(c)
}

// Reduce applies a function cumulatively to values from the iterable so the return value of one call is passed to
// the next's first argument, and the next value from the iterator to the second
//
// for example for the iterable 1,2,3 and a function like:
// func(a, b int) { return a + b }
// would be result in ((1+2)+3)
//
// this implies that the iterable must return at least two items for the function to be called.
// reducing an iterable that only returns one value results in that value, regardless of the function
func Reduce[T any](fn func(a, b T) T, data ...iface.Iterable[T]) T {
	buf := make([]T, 2)
	i := 0
	for item := range Chain(data...).Iter() {
		// need to initialize buf with the first two elements
		switch i {
		case 0:
			i += 1
			buf[0] = item
		default:
			i += 1
			buf[1] = item
			buf[0] = fn(buf[0], buf[1])
		}
	}
	return buf[0]
}

// Filter applies a function to an iterator and returns an iterable that contains only the values where the function
// returns true
func Filter[T any](fn func(T) bool, data ...iface.Iterable[T]) iface.Iterable[T] {
	c := make(chan T)
	go func() {
		for item := range Chain(data...).Iter() {
			if fn(item) {
				c <- item
			}
		}
		close(c)
	}()
	return IterableChannel[T](c)
}

// Chain stitches together multiple iterators into one stream of values,
// where the first iterator is consumed entirely, then the second, and so on
func Chain[T any](iterables ...iface.Iterable[T]) iface.Iterable[T] {
	c := make(chan T)
	go func() {
		for _, iterable := range iterables {
			for item := range iterable.Iter() {
				c <- item
			}
		}
		close(c)
	}()
	return IterableChannel[T](c)
}
