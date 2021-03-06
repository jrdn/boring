package ds

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewLinkedList(t *testing.T) {
	expected := []string{"foo", "bar", "baz", "quux"}
	ll := NewLinkedList[string](expected)

	for i := 0; i < len(expected); i++ {
		x, ok := ll.Get(i)
		require.True(t, ok, "failed to get %d", i)
		require.Equal(t, expected[i], x)
	}
}
