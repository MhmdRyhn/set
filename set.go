// Idea taken from here: https://bitfieldconsulting.com/golang/generic-set

package set

type void struct{}

type set[T comparable] struct {
	hashTable map[T]void
	len       int
}

//
// This constructor function never returns a nil pointer.
// If the set is initialized with no item(s), pointer to
// an empty/null set is returned.
//
func New[T comparable](items ...T) *set[T] {
	s := set[T]{
		hashTable: map[T]void{},
		len:       0,
	}
	for _, item := range items {
		if _, exist := s.hashTable[item]; !exist {
			s.hashTable[item] = void{}
			s.len++
		}
	}
	return &s
}

func (s *set[T]) Len() int {
	return s.len
}

func (s *set[T]) IsEmpty() bool {
	return s.Len() == 0
}

//
// Identical to `set.IsEmpty()` method
//
func (s *set[T]) IsNull() bool {
	return s.Len() == 0
}

func (s *set[T]) Contains(item T) bool {
	_, exist := s.hashTable[item]
	return exist
}

func (s *set[T]) Add(items ...T) {
	for _, item := range items {
		if !s.Contains(item) {
			s.hashTable[item] = void{}
			s.len++
		}
	}
}

func (s *set[T]) Remove(item T) {
	if s.Contains(item) {
		delete(s.hashTable, item)
		s.len--
	}
}

func (s *set[T]) Members() []T {
	var members []T
	for item := range s.hashTable {
		members = append(members, item)
	}
	return members
}

func (s *set[T]) Union(s2 set[T]) *set[T] {
	result := s
	result.Add(s2.Members()...)
	return result
}

func (s *set[T]) Intersection(s2 set[T]) *set[T] {
	result := New[T]()
	for item := range s.hashTable {
		if s2.Contains(item) {
			result.Add(item)
		}
	}
	return result
}

func (s *set[T]) Difference(s2 set[T]) *set[T] {
	result := s
	for item := range s2.hashTable {
		if result.Contains(item) {
			result.Remove(item)
		}
	}
	return result
}

//
// Checks if two sets are equal or not.
//
func (s *set[T]) IsEqual(s2 set[T]) bool {
	if s.Len() != s2.Len() {
		return false
	}
	for item := range s.hashTable {
		if !s2.Contains(item) {
			return false
		}
	}
	return true
}

//
// Checks whether the set "s" and "s2" are disjoint or not.
//
func (s *set[T]) IsDisjoint(s2 set[T]) bool {
	if s.IsEmpty() && s2.IsEmpty() {
		return false
	}
	for item := range s.hashTable {
		if s2.Contains(item) {
			return false
		}
	}
	return true
}

//
// Whether the set "s" is a subset of the set "s2". "proper" param
// indicates if the subset is proper or not.
//
func (s *set[T]) IsSubsetOf(s2 set[T], proper bool) bool {
	if s.Len() > s2.Len() {
		return false
	}
	if proper && s.Len() == s2.Len() {
		return false
	}
	for item := range s.hashTable {
		if !s2.Contains(item) {
			return false
		}
	}
	return true
}
