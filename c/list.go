package c

func NewList[T any](data ...T) *List[T] {
	return &List[T]{data}
}

type List[T any] struct {
	x []T
}

func (l *List[T]) Iter() <-chan T {
	iterChan := make(chan T)
	go func() {
		for _, item := range l.x {
			iterChan <- item
		}
		close(iterChan)
	}()
	return iterChan
}
