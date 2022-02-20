package types

import "context"

type Generator[T any] func(ctx context.Context) chan T

type Iterator[T any] interface {
	Range() chan T
	Close()
}

// Iterable is a type which implements the boring iterator protocol
type Iterable[T any] interface {
	Iter() Iterator[T]
}

type Lengthable interface {
	Len() int
}
