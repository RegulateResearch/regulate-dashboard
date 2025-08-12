package generictype

type Triplet[A any, B any, C any] struct {
	Pair[A, B]
	c C
}

func NewTriplet[A any, B any, C any](a A, b B, c C) Triplet[A, B, C] {
	return Triplet[A, B, C]{
		Pair: NewPair(a, b),
		c:    c,
	}
}

func (t Triplet[A, B, C]) C() C {
	return t.c
}
