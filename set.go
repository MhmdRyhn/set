// Idea taken from here: https://bitfieldconsulting.com/golang/generic-set

package set

type void struct{}

type set[T comparable] map[T]void

func New[T comparable](items ...T) set[T] {
	s := set[T]{}
	for _, val := range items {
		s[val] = void{}
	}
	return s
}

func (s set[T]) Len() int {
	return len(s)
}

func (s set[T]) Contains(item T) bool {
	_, exist := s[item]
	return exist
}

func (s set[T]) Add(items ...T) {
	for _, val := range items {
		s[val] = void{}
	}
}

func (s set[T]) Remove(item T) {
	delete(s, item)
}

func (s set[T]) Members() []T {
	var members []T
	for v := range s {
		members = append(members, v)
	}
	return members
}

func (s set[T]) Union(s2 set[T]) set[T] {
	result := s
	result.Add(s2.Members()...)
	return result
}

func (s set[T]) Intersection(s2 set[T]) set[T] {
	result := New[T]()
	for item := range s {
		if s2.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

func (s set[T]) Difference(s2 set[T]) set[T] {
	result := s
	for item := range s2 {
		if result.Contains(item) {
			result.Remove(item)
		}
	}
	return result
}
