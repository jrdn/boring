package types

import "context"

// Iterable is a type which implements the boring iterator protocol
type Iterable[T any] interface {
	Iter(ctx context.Context) <-chan T
}

type Lengthable interface {
	Len() int
}
