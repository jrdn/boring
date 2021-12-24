package c

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPair(t *testing.T) {
	pair := NewPair[string, string]("foo", "bar")
	assert.Equal(t, "foo", pair.First())
	assert.Equal(t, "bar", pair.Second())

}

func TestNewTriplet(t *testing.T) {
	triplet := NewTriplet[string, string, string]("foo", "bar", "baz")
	assert.Equal(t, "foo", triplet.First())
	assert.Equal(t, "bar", triplet.Second())
	assert.Equal(t, "baz", triplet.Third())
}

func TestNewQuadruplet(t *testing.T) {
	now := time.Now()
	quadruplet := NewQuadruplet[string, int, time.Time, []string]("foo", 10, now, []string{"a", "b"})

	assert.Equal(t, "foo", quadruplet.First())
	assert.Equal(t, 10, quadruplet.Second())
	assert.True(t, now.Equal(quadruplet.Third()))
	assert.Equal(t, []string{"a", "b"}, quadruplet.Fourth())
}
