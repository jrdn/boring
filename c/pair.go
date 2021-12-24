package c

func NewPair[X, Y any](a X, b Y) Pair[X, Y] {
	return Pair[X, Y]{first: a, second: b}
}

type Pair[X, Y any] struct {
	first  X
	second Y
}

func (p *Pair[X, Y]) First() X {
	return p.first
}

func (p *Pair[X, Y]) Second() Y {
	return p.second
}
