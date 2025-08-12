package generictype

type Pair[A any, B any] struct {
	a A
	b B
}

func NewPair[A any, B any](a A, b B) Pair[A, B] {
	return Pair[A, B]{
		a: a,
		b: b,
	}
}

func (p Pair[A, B]) A() A {
	return p.a
}

func (p Pair[A, B]) B() B {
	return p.b
}
