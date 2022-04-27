package fn

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSliceIterator(t *testing.T) {
	expected := []string{"foo", "bar", "baz"}
	iterable := SliceIterator[string](expected)

	result := Collect(context.TODO(), iterable).GetSlice()
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)

	for i, r := range result {
		assert.Equal(t, expected[i], r)
	}
}
