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
