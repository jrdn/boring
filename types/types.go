package types

type Container[T any] interface {
	Contains(other T) bool
}

type Lengthable interface {
	Len() int
}
