package fn

import (
	"github.com/jrdn/boring/iface"
)

// Collect gathers all the values from an iterable and converts it to a list
func Collect[T any](iter iface.Iterable[T]) []T {
	var ret []T
	for item := range iter.Iter() {
		ret = append(ret, item)
	}
	return ret
}

func Map[T, R any](fn func(T) R, iterators ...iface.Iterable[T]) iface.Iterable[R] {
	c := make(chan R)
	go func() {
		for _, iterator := range iterators {
			for item := range iterator.Iter() {
				c <- fn(item)
			}
		}
		close(c)
	}()
	return IterableChannel(c)
}

func Reduce[T any](fn func(a, b T) T, data ...iface.Iterable[T]) T {
	buf := make([]T, 2)
	i := 0
	// need to initialize buf with the first two elements
	// TODO rewrite to use a channel and run async
	for _, d := range data {
		for item := range d.Iter() {
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
	}

	return buf[0]
}

func Filter[T any](fn func(T) bool, iterable iface.Iterable[T]) iface.Iterable[T] {
	c := make(chan T)
	go func() {
		for item := range iterable.Iter() {
			if fn(item) {
				c <- item
			}
		}
		close(c)
	}()
	return IterableChannel[T](c)
}

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
