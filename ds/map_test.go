package ds

import (
	"testing"

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
