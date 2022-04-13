package hashset

// Set elements need to implement this interface
type Hasher[U comparable] interface {
	Hash() U
}

// A Set implementation for types that have a Hash method
type Set[U comparable, T Hasher[U]] map[U]T

// Add element to the set
func NewSet[U comparable, T Hasher[U]](elems []T) Set[U, T] {
	s := Set[U, T]{}
	for _, elem := range elems {
		s.Add(elem)
	}
	return s
}

// Add element to the set
func (s Set[U, T]) Add(elem T) {
	s[elem.Hash()] = elem
}

// Remove element from the set
func (s Set[U, T]) Remove(elem T) {
	delete(s, elem.Hash())
}

// Check element is in the set
func (s Set[U, T]) Has(elem T) bool {
	_, ok := s[elem.Hash()]
	return ok
}

// Convert the set to a slice
func (s Set[U, T]) ToList() []T {
	list := []T{}
	for _, val := range s {
		list = append(list, val)
	}
	return list
}

// Creates a copy of the set
func (s Set[U, T]) Clone() Set[U, T] {
	clone := Set[U, T]{}
	for k, v := range s {
		clone[k] = v
	}
	return clone
}

// Update the set by taking the union with the other set
func (s Set[U, T]) Union(other Set[U, T]) {
	for _, v := range other {
		s.Add(v)
	}
}

// Compute the union of s1 and s2
func Union[U comparable, T Hasher[U]](s1, s2 Set[U, T]) Set[U, T] {
	union := s1.Clone()
	for _, v := range s2 {
		union.Add(v)
	}
	return union
}

// Update the set by taking the intersection with the other set
func (s Set[U, T]) Intersect(other Set[U, T]) {
	for k, v := range s {
		_, ok := other[k]
		if !ok {
			s.Remove(v)
		}
	}
}

// Compute the intersection of s1 and s2
func Intersect[U comparable, T Hasher[U]](s1, s2 Set[U, T]) Set[U, T] {
	intersection := Set[U, T]{}
	for k, v := range s1 {
		_, ok := s2[k]
		if ok {
			intersection.Add(v)
		}
	}
	return intersection
}

// Update the set by taking the difference with the other set
func (s Set[U, T]) Difference(other Set[U, T]) {
	for k, v := range s {
		_, ok := other[k]
		if ok {
			s.Remove(v)
		}
	}
}

// Compute the difference of s1 from s2
func Difference[U comparable, T Hasher[U]](s1, s2 Set[U, T]) Set[U, T] {
	diff := Set[U, T]{}
	for k, v := range s1 {
		_, ok := s2[k]
		if !ok {
			diff.Add(v)
		}
	}
	return diff
}

// Checks whether the set is a subset of the other set
func (s Set[U, T]) IsSubset(other Set[U, T]) bool {
	return len(Intersect(s, other)) == len(s)
}

// Checks whether s1 is a subset of s2
func IsSubset[U comparable, T Hasher[U]](s1, s2 Set[U, T]) bool {
	return s1.IsSubset(s2)
}

// Checks whether the set is a superset of the other set
func (s Set[U, T]) IsSuperset(other Set[U, T]) bool {
	return other.IsSubset(s)
}

// Checks whether s1 is a superset of s2
func IsSuperset[U comparable, T Hasher[U]](s1, s2 Set[U, T]) bool {
	return s1.IsSuperset(s2)
}

// Checks whether the set has no intersection with the other set
func (s Set[U, T]) IsDisjoint(other Set[U, T]) bool {
	return len(Intersect(s, other)) == 0
}

// Checks whether s1 and s2 have no intersection
func IsDisjoint[U comparable, T Hasher[U]](s1, s2 Set[U, T]) bool {
	return s1.IsDisjoint(s2)
}

// Checks whether s contains the same elements as the other
func (s Set[U, T]) Equal(other Set[U, T]) bool {
	if len(s) != len(other) {
		return false
	}
	for k := range s {
		_, ok := other[k]
		if !ok {
			return false
		}
	}

	return true
}

// Checks whether s1 and s2 contain the same elements
func Equal[U comparable, T Hasher[U]](s1, s2 Set[U, T]) bool {
	return s1.Equal(s2)
}
