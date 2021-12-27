package main

import (
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

	// wrap the generator function in an iterator
	iter := fn.FuncIterator(genFunc)

	// filter values from the stream so it's just even numbers
	filteredIter := fn.Filter(func(input int) bool {
		return input%2 == 0
	}, iter)

	// Map over the values to double them
	doubledIter := fn.Map[int, int](func(input int) int {
		return input * 2
	}, filteredIter)

	// since the stream is infinite we can't just iterate over it,
	// so collect items until they're > 40
	collected := fn.CollectUntil(func(item int) bool {
		return item > 40
	}, doubledIter)

	// reduce to the sum of the collection
	reduced := fn.Reduce[int](func(a, b int) int {
		return a + b
	}, collected)

	fmt.Println(reduced)
}
