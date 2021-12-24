package iface

type Iterable[T any] interface {
	Iter() <-chan T
}

type Container[T any] interface {
	Contains(T) bool
}
