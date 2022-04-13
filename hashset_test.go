package hashset

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type MyElement struct {
	id int
}

func (m *MyElement) Hash() int {
	return m.id
}

type YourElement struct {
	id string
}

func (y YourElement) Hash() string {
	return y.id
}

func TestHashSetWithPointerReceiver(t *testing.T) {
	s := Set[int, *MyElement]{}
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})
	assert.Equal(t, 3, len(s))
	assert.True(t, s.Has(&MyElement{id: 1}))
	assert.True(t, s.Has(&MyElement{id: 2}))
	assert.True(t, s.Has(&MyElement{id: 3}))

	s.Remove(&MyElement{id: 2})
	assert.Equal(t, 2, len(s))
	assert.True(t, s.Has(&MyElement{id: 1}))
	assert.False(t, s.Has(&MyElement{id: 2}))
	assert.True(t, s.Has(&MyElement{id: 3}))

	// Note that MyElement does not implement Hasher[int] (method Hash has pointer receiver)
	// s := Set[int, MyElement]{}  // compile error
}

func TestHashSetWithValueReceiverAndStringHashType(t *testing.T) {
	s := Set[string, YourElement]{}
	s.Add(YourElement{id: "1"})
	s.Add(YourElement{id: "2"})
	s.Add(YourElement{id: "2"})
	s.Add(YourElement{id: "3"})
	assert.Equal(t, 3, len(s))

	s2 := Set[string, *YourElement]{}
	s2.Add(&YourElement{id: "1"})
	s2.Add(&YourElement{id: "2"})
	s2.Add(&YourElement{id: "2"})
	s2.Add(&YourElement{id: "3"})
	assert.Equal(t, 3, len(s2))
}

func TestTypeAliasAndToList(t *testing.T) {
	type MyElementSet = Set[int, *MyElement]

	myElements := []*MyElement{{id: 1}, {id: 2}}
	s := MyElementSet{}
	for _, e := range myElements {
		s.Add(e)
	}

	assert.ElementsMatch(t, myElements, s.ToList())
}

func TestNewSet(t *testing.T) {
	elems := []*MyElement{{id: 1}, {id: 2}, {id: 2}, {id: 3}}
	s := NewSet[int, *MyElement](elems)
	assert.Equal(t, 3, len(s))
	assert.True(t, s.Has(&MyElement{id: 1}))
	assert.True(t, s.Has(&MyElement{id: 2}))
	assert.True(t, s.Has(&MyElement{id: 3}))
}

func TestClone(t *testing.T) {
	s := Set[int, *MyElement]{}
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})

	s2 := s.Clone()
	assert.Equal(t, 3, len(s2))
	assert.True(t, s2.Has(&MyElement{id: 1}))
	assert.True(t, s2.Has(&MyElement{id: 2}))
	assert.True(t, s2.Has(&MyElement{id: 3}))
}

func TestUnion(t *testing.T) {
	s := Set[int, *MyElement]{}
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})

	s2 := Set[int, *MyElement]{}
	s2.Add(&MyElement{id: 4})
	s2.Add(&MyElement{id: 3})

	union := Union(s, s2)
	assert.Equal(t, 4, len(union))
	assert.True(t, union.Has(&MyElement{id: 1}))
	assert.True(t, union.Has(&MyElement{id: 2}))
	assert.True(t, union.Has(&MyElement{id: 3}))
	assert.True(t, union.Has(&MyElement{id: 4}))

	s.Union(s2)
	assert.Equal(t, 4, len(s))
	assert.True(t, s.Has(&MyElement{id: 1}))
	assert.True(t, s.Has(&MyElement{id: 2}))
	assert.True(t, s.Has(&MyElement{id: 3}))
	assert.True(t, s.Has(&MyElement{id: 4}))
}

func TestIntersect(t *testing.T) {
	s := Set[int, *MyElement]{}
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})

	s2 := Set[int, *MyElement]{}
	s2.Add(&MyElement{id: 4})
	s2.Add(&MyElement{id: 3})

	isect := Intersect(s, s2)
	assert.Equal(t, 1, len(isect))
	assert.True(t, isect.Has(&MyElement{id: 3}))

	s.Intersect(s2)
	assert.Equal(t, 1, len(s))
	assert.True(t, s.Has(&MyElement{id: 3}))
}

func TestDifference(t *testing.T) {
	s := Set[int, *MyElement]{}
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})

	s2 := Set[int, *MyElement]{}
	s2.Add(&MyElement{id: 4})
	s2.Add(&MyElement{id: 3})

	diff := Difference(s, s2)
	assert.Equal(t, 2, len(diff))
	assert.True(t, diff.Has(&MyElement{id: 1}))
	assert.True(t, diff.Has(&MyElement{id: 2}))

	s.Difference(s2)
	assert.Equal(t, 2, len(s))
	assert.True(t, s.Has(&MyElement{id: 1}))
	assert.True(t, s.Has(&MyElement{id: 2}))
}

func TestIsSubset(t *testing.T) {
	s := Set[int, *MyElement]{}
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})

	s2 := Set[int, *MyElement]{}
	s2.Add(&MyElement{id: 4})
	s2.Add(&MyElement{id: 3})

	s3 := Set[int, *MyElement]{}
	s3.Add(&MyElement{id: 2})
	s3.Add(&MyElement{id: 3})

	empty := Set[int, *MyElement]{}

	assert.True(t, empty.IsSubset(s))
	assert.True(t, s.IsSubset(s))
	assert.False(t, s.IsSubset(s2))
	assert.False(t, s2.IsSubset(s))
	assert.True(t, s3.IsSubset(s))
	assert.False(t, s.IsSubset(s3))

	assert.True(t, IsSubset(s3, s))
	assert.False(t, IsSubset(s, s3))
}

func TestEqual(t *testing.T) {
	s := Set[int, *MyElement]{}
	s.Add(&MyElement{id: 1})
	s.Add(&MyElement{id: 2})
	s.Add(&MyElement{id: 3})

	s2 := Set[int, *MyElement]{}
	s2.Add(&MyElement{id: 2})
	s2.Add(&MyElement{id: 3})

	s3 := Set[int, *MyElement]{}
	s3.Add(&MyElement{id: 1})
	s3.Add(&MyElement{id: 2})
	s3.Add(&MyElement{id: 3})

	s4 := Set[int, *MyElement]{}
	s4.Add(&MyElement{id: 1})
	s4.Add(&MyElement{id: 2})
	s4.Add(&MyElement{id: 4})

	empty := Set[int, *MyElement]{}

	assert.False(t, empty.Equal(s))
	assert.False(t, s.Equal(empty))

	assert.False(t, s.Equal(s2))
	assert.False(t, s2.Equal(s))

	assert.False(t, s.Equal(s4))
	assert.False(t, s4.Equal(s))

	assert.True(t, s.Equal(s3))
	assert.True(t, s3.Equal(s))

	assert.True(t, Equal(s, s3))
	assert.False(t, Equal(s, s2))
}
