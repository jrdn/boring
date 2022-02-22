package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewVersionSpec(t *testing.T) {
	v1 := NewVersion("1.2.3-foo")
	require.NotEmpty(t, v1)
	v2 := NewVersion("3.2.1-bar")
	require.NotEmpty(t, v2)

	assert.False(t, v1.Equal(v2))
}

func TestVersionString(t *testing.T) {
	v := NewVersion("1")
	require.NotEmpty(t, v)

	assert.Equal(t, "v1.0.0", v.String())
}
