package ds

import (
	"context"
	"fmt"
	"testing"

	"github.com/jrdn/boring/types"
	"github.com/stretchr/testify/require"
)

func TestNewMap(t *testing.T) {
	m := NewMap[string, string](map[string]string{"foo": "bar", "baz": "quux"}, map[string]string{"asdf": "qwer"})
	require.NotNil(t, m)
}

func TestMap_Contains(t *testing.T) {
	m := NewMap[string, string](map[string]string{"foo": "bar", "baz": "quux"})
	require.True(t, m.Contains("foo"))
	require.True(t, m.Contains("baz"))
	require.False(t, m.Contains("asdfj"))
}

func TestMap_Iter(t *testing.T) {
	var iter types.Iterable[Pair[string, string]] = NewMap[string, string](map[string]string{"foo": "bar", "baz": "quux"})

	for item := range iter.Iter(context.Background()) {
		fmt.Println(item)
	}
}

// type testComparable struct {
// 	val int
// }

// func TestNewOrderedMap(t *testing.T) {
// 	expected := []Pair[testComparable, int]{
// 		NewPair(testComparable{val: 5}, 5),
// 		NewPair(testComparable{val: 10}, 10),
// 		NewPair(testComparable{val: 1}, 1),
// 	}
//
// 	om := NewOrderedMap[testComparable, int]()
// 	for _, item := range expected {
// 		om.Append(item.First(), item.Second())
// 	}
//
// 	// should fail to add the item since it's already contained in the ordered map
// 	assert.False(t, om.Append(expected[0].First(), expected[0].Second()))
// 	index := 0
// 	for item := range om.Iter() {
// 		assert.Equal(t, expected[index].First(), item.First())
// 		assert.Equal(t, expected[index].Second(), item.Second())
// 		index += 1
// 	}
// 	assert.NotEqual(t, 0, index)
// }
