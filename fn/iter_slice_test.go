package fn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIterator(t *testing.T) {
	expected := []string{"foo", "bar", "baz"}
	iterable := SliceIterator[string](expected)

	result := Collect(iterable).GetSlice()
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)

	for i, r := range result {
		assert.Equal(t, expected[i], r)
	}
}
