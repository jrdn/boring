package c

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSet_Iter(t *testing.T) {
	results := make(map[string]struct{})
	for item := range NewSet[string]([]string{"asdf", "quux"}).Iter() {
		results[item] = struct{}{}
	}

	assert.Contains(t, results, "asdf")
	assert.Contains(t, results, "quux")
}

func TestSet_Contains(t *testing.T) {
	s := NewSet[string]([]string{"asdf", "quux"})
	assert.True(t, s.Contains("asdf"))
	assert.True(t, s.Contains("quux"))
	assert.False(t, s.Contains("hello world"))
}
