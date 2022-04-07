package hashset

import (
	"reflect"
	"sort"
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

type MyElementSet = HashSet[int, *MyElement]

func TestHashSetWithPointerReceiver(t *testing.T) {
	s := HashSet[int, *MyElement]{}
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
	// s := HashSet[int, MyElement]{}  // compile error
}

func TestHashSetWithValueReceiverAndStringHashType(t *testing.T) {
	s := HashSet[string, YourElement]{}
	s.Add(YourElement{id: "1"})
	s.Add(YourElement{id: "2"})
	s.Add(YourElement{id: "2"})
	s.Add(YourElement{id: "3"})
	assert.Equal(t, 3, len(s))

	s2 := HashSet[string, *YourElement]{}
	s2.Add(&YourElement{id: "1"})
	s2.Add(&YourElement{id: "2"})
	s2.Add(&YourElement{id: "2"})
	s2.Add(&YourElement{id: "3"})
	assert.Equal(t, 3, len(s2))
}

func TestTypeAliasAndToList(t *testing.T) {
	myElements := []*MyElement{{id: 1}, {id: 2}}
	s := MyElementSet{}
	for _, e := range myElements {
		s.Add(e)
	}

	sList := s.ToList()
	sort.Slice(sList, func(i, j int) bool {
		return sList[i].id < sList[j].id
	})
	assert.True(t, reflect.DeepEqual(sList, myElements))
}
