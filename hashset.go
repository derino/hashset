package hashset

// Set elements need to implement this interface
type Hasher[U comparable] interface {
	Hash() U
}

// A Set implementation for types that have a Hash method
type HashSet[U comparable, T Hasher[U]] map[U]T

// Add element to the set
func (s HashSet[U, T]) Add(elem T) {
	s[elem.Hash()] = elem
}

// Remove element from the set
func (s HashSet[U, T]) Remove(elem T) {
	delete(s, elem.Hash())
}

// Check element is in the set
func (s HashSet[U, T]) Has(elem T) bool {
	_, ok := s[elem.Hash()]
	return ok
}

// Convert the set to a slice
func (s HashSet[U, T]) ToList() []T {
	list := []T{}
	for _, val := range s {
		list = append(list, val)
	}
	return list
}
