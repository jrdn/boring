package c

import (
	"fmt"
	"testing"

	"github.com/jrdn/boring/iface"
	"github.com/stretchr/testify/require"
)

func TestNewMap(t *testing.T) {
	m := NewMap[string, string](map[string]string{"foo": "bar", "baz": "quux"})
	require.NotNil(t, m)
}

func TestMap_Contains(t *testing.T) {
	m := NewMap[string, string](map[string]string{"foo": "bar", "baz": "quux"})
	require.True(t, m.Contains("foo"))
	require.True(t, m.Contains("baz"))
	require.False(t, m.Contains("asdfj"))
}

func TestMap_Iter(t *testing.T) {
	m := NewMap[string, string](map[string]string{"foo": "bar", "baz": "quux"})
	var iter iface.Iterable[Pair[string, string]] = m
	for item := range iter.Iter() {
		fmt.Println(item)
	}
}
