package main

import (
	"context"
	"fmt"

	"github.com/jrdn/boring/fn"
)

var counter int = 0

func main() {
	// genFunc is an iterator which produces counting up forever
	// it would stop iterating by returning true fits second return value

	genFunc := func() (int, bool) {
		for {
			counter = counter + 1
			return counter, false // returning true stops iteration, but this should be infinite
		}
	}
	ctx, cancel := context.WithCancel(context.Background())
	_ = cancel // could use the cancel function to abort all the iterators & ensure their goroutines get shut down
	// when no longer required

	// wrap the generator function in an iterator
	iter := fn.FuncIterator(genFunc)

	// filter values from the stream so it's just even numbers
	filteredIter := fn.Filter(ctx, func(input int) bool {
		return input%2 == 0
	}, iter)

	// Map over the values to double them
	doubledIter := fn.Map[int, int](ctx, func(input int) int {
		return input * 2
	}, filteredIter)

	// since the stream is infinite we can't just iterate over it,
	// so collect items until they're > 40
	collected := fn.CollectUntil(ctx, func(item int) bool {
		return item > 40
	}, doubledIter)

	// reduce to the sum of the collection
	reduced := fn.Reduce[int](ctx, func(a, b int) int {
		return a + b
	}, collected)

	fmt.Println(reduced)
}
