package ds

// Pair, Triplet, and Quadruplet exist as an easy way to return an object with multiple items of (potentially) different
// types. This is useful for things like Map iterators, where the the key and value are returned in a Pair
// Triplets and Quadruplets are less obviously useful but might come in handy sometimes.
// But it's unusual enough we haven't gone up to Quintuplets or Sextuplets or etc (yet...)

// NewPair creates a new Pair.
func NewPair[A, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{first: a, second: b}
}

// Pair holds two values.
type Pair[A, B any] struct {
	first  A
	second B
}

// First gets the first value.
func (tuple Pair[A, B]) First() A {
	return tuple.first
}

// Second gets the Second value.
func (tuple Pair[A, B]) Second() B {
	return tuple.second
}

// NewTriplet creates a new Triplet.
func NewTriplet[A, B, C any](a A, b B, c C) Triplet[A, B, C] {
	return Triplet[A, B, C]{
		Pair: Pair[A, B]{
			first:  a,
			second: b,
		},
		third: c,
	}
}

// Triplet holds three values.
type Triplet[A, B, C any] struct {
	Pair[A, B]
	third C
}

// Third returns the third value.
func (tuple Triplet[A, B, C]) Third() C {
	return tuple.third
}

// NewQuadruplet creates a new Quadruplet.
func NewQuadruplet[A, B, C, D any](a A, b B, c C, d D) Quadruplet[A, B, C, D] {
	return Quadruplet[A, B, C, D]{
		Triplet: Triplet[A, B, C]{
			Pair: Pair[A, B]{
				first:  a,
				second: b,
			},
			third: c,
		},
		fourth: d,
	}
}

// Quadruplet holds four values.
type Quadruplet[A, B, C, D any] struct {
	Triplet[A, B, C]
	fourth D
}

// Fourth gets the fourth value.
func (tuple Quadruplet[A, B, C, D]) Fourth() D {
	return tuple.fourth
}
