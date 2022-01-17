package fn

import (
	"github.com/jrdn/boring/ds"
	"github.com/jrdn/boring/types"
)

// Collect gathers all the values from an iterable and converts it to a List
func Collect[T any](iter ...types.Iterable[T]) *ds.List[T] {
	var ret []T
	for item := range Chain(iter...).Iter() {
		ret = append(ret, item)
	}
	return ds.NewList[T](ret)
}

// CollectN gathers the first N values from an iterable and adds it to a List until the list contains
func CollectN[T any](n int, iter ...types.Iterable[T]) *ds.List[T] {
	var ret []T

	counter := 0
	for item := range Chain(iter...).Iter() {
		ret = append(ret, item)
		counter += 1
		if counter == n {
			break
		}
	}
	return ds.NewList[T](ret)
}

type CollectUntilDecider[T any] func(item T) (stopIteration bool)

// CollectUntil gathers values from iterables into a List until a provided function returns true
func CollectUntil[T any](fn CollectUntilDecider[T], iter ...types.Iterable[T]) *ds.List[T] {
	var ret []T
	for item := range Chain(iter...).Iter() {
		if stopIteration := fn(item); stopIteration {
			break
		}
		ret = append(ret, item)
	}
	return ds.NewList[T](ret)
}

// Map applies a function to a number of iterators and produces an iterable of the function's return values
func Map[T, R any](fn func(T) R, iterators ...types.Iterable[T]) types.Iterable[R] {
	c := make(chan R)

	go func() {
		defer close(c)
		chain := Chain[T](iterators...)
		for item := range chain.Iter() {
			c <- fn(item)
		}
	}()
	return ChanIterator(c)
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
func Reduce[T any](fn func(a, b T) T, data ...types.Iterable[T]) T {
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
func Filter[T any](fn func(T) bool, data ...types.Iterable[T]) types.Iterable[T] {
	c := make(chan T)
	go func() {
		defer close(c)
		for item := range Chain(data...).Iter() {
			if fn(item) {
				c <- item
			}
		}
	}()
	return ChanIterator[T](c)
}

// Chain stitches together multiple iterators into one stream of values,
// where the first iterator is consumed entirely, then the second, and so on
func Chain[T any](iterables ...types.Iterable[T]) types.Iterable[T] {
	c := make(chan T)
	go func() {
		defer close(c)
		for _, iterable := range iterables {
			for item := range iterable.Iter() {
				c <- item
			}
		}
	}()
	return ChanIterator[T](c)
}

// Tee makes two iterators which operate together, so an item must be read from both before either progresses to the
// next item
func Tee[T any](iter types.Iterable[T]) (types.Iterable[T], types.Iterable[T]) {
	chan1 := make(chan T)
	chan2 := make(chan T)

	go func() {
		defer close(chan1)
		defer close(chan2)
		for item := range iter.Iter() {
			chan1 <- item
			chan2 <- item
		}
	}()

	return ChanIterator[T](chan1), ChanIterator[T](chan2)
}

// Repeat a value infinitely
func Repeat[T any](element T) types.Iterable[T] {
	return FuncIterator(func() (T, bool) {
		return element, false
	})
}

// RepeatTimes repeats a value a given number of times
func RepeatTimes[T any](element T, times int) types.Iterable[T] {
	counter := 0
	return FuncIterator(func() (T, bool) {
		if counter < times-1 {
			counter++
			return element, false
		}
		counter++
		return element, true
	})
}

// Zip consumes two iterators and joins them into a Pair
// it continues iterating until the shortest iterator is exhausted.
func Zip[A, B any](a types.Iterable[A], b types.Iterable[B]) types.Iterable[ds.Pair[A, B]] {
	result := make(chan ds.Pair[A, B])
	aIter := a.Iter()
	bIter := b.Iter()
	// TODO well this will leak a goroutine every time the iterators are of unequal length
	go func() {
		defer close(result)
		for {
			x, ok := <-aIter
			if !ok {
				break
			}
			y, ok := <-bIter
			if !ok {
				break
			}
			result <- ds.NewPair[A, B](x, y)
		}
	}()
	return ChanIterator[ds.Pair[A, B]](result)
}

// Pairwise consumes an iterator and returns them two at a times
func Pairwise[T any](iter types.Iterable[T]) types.Iterable[ds.Pair[T, T]] {
	iterChan := make(chan ds.Pair[T, T])

	source := iter.Iter()

	go func() {
		defer close(iterChan)
		for {
			first, more := <-source
			if !more {
				break
			}

			second := <-source
			iterChan <- ds.NewPair[T, T](first, second)
		}
	}()

	return ChanIterator[ds.Pair[T, T]](iterChan)
}
