package fn

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIterator(t *testing.T) {
	counter := 0
	result := Collect[int](Iterator[int](func() (int, bool) {
		for counter < 10 {
			val := counter
			counter = counter + 1
			return val, false
		}
		return 10, true
	})).Get()

	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, expected, result)
}
