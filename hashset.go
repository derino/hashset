package hashset

// Set elements need to implement this interface
type Hasher[U comparable] interface {
	Hash() U
}

// A Set implementation for types that have a Hash method
type Set[U comparable, T Hasher[U]] map[U]T

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

// Update the set by making a union with the given set
func (s Set[U, T]) Union(other Set[U, T]) {
	for _, v := range other {
		s.Add(v)
	}
}

// Compute a union of the given two sets
func Union[U comparable, T Hasher[U]](s1 Set[U, T], s2 Set[U, T]) Set[U, T] {
	union := s1.Clone()
	for _, v := range s2 {
		union.Add(v)
	}
	return union
}
