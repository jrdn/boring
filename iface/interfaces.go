package iface

// Iterable is a type which implements the boring iterator protocol
type Iterable[T any] interface {
	Iter() <-chan T
}

type Lengthable interface {
	Len() int
}
