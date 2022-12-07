package ysl

type Set[T comparable] interface {
	Cardinality() int
	Add(item T) bool
	Pop() (T, bool)
	Intersect(other Set[T]) Set[T]
}

func NewSet[T comparable](items ...T) Set[T] {
	s := newSet[T]()
	for _, i := range items {
		s.Add(i)
	}
	return &s
}

type set[T comparable] map[T]struct{}

func newSet[T comparable]() set[T] {
	return make(set[T])
}

func (s *set[T]) Cardinality() int {
	return len(*s)
}

func (s *set[T]) Add(item T) bool {
	if _, ok := (*s)[item]; ok {
		return false
	}
	(*s)[item] = struct{}{}
	return true
}

func (s *set[T]) Pop() (v T, ok bool) {
	for k := range *s {
		delete(*s, k)
		return k, true
	}
	return
}

func (s *set[T]) Intersect(other Set[T]) Set[T] {
	o := other.(*set[T])

	intersection := newSet[T]()
	// loop over smaller set
	if s.Cardinality() < other.Cardinality() {
		for el := range *s {
			if o.contains(el) {
				intersection.Add(el)
			}
		}
	} else {
		for elem := range *o {
			if s.contains(elem) {
				intersection.Add(elem)
			}
		}
	}
	return &intersection
}

func (s *set[T]) contains(v T) bool {
	_, ok := (*s)[v]
	return ok
}
