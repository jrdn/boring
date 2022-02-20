package ds

type Optional[T any] struct {
	x   T
	set bool
}

func (o *Optional[T]) Set(x T) {
	o.x = x
	o.set = true
}

func (o *Optional[T]) Get() (T, bool) {
	return o.x, o.set
}

func (o *Optional[T]) Check() bool {
	return o.set
}
