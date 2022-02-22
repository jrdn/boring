package version

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCondition(t *testing.T) {
	testVersion := NewVersion("1.2.3")
	tests := []struct {
		name     string
		cond     condition
		expected bool
	}{
		{"equal", condition{"=", NewVersion("1.2.3")}, true},
		{"not equal", condition{"=", NewVersion("3.2.1")}, false},
		{"greater", condition{">", NewVersion("1")}, true},
		{"not greater", condition{">", NewVersion("2")}, false},
		{"less", condition{"<", NewVersion("2")}, true},
		{"not less", condition{"<", NewVersion("1")}, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.cond.test(testVersion) != test.expected {
				t.Fail()
			}
		})
	}
}

func TestFilter_test(t *testing.T) {
	version := NewVersion("1.2.3")

	tests := []struct {
		name     string
		filter   Filter
		expected bool
	}{
		{"equal", "=1.2.3", true},
		{"in range", ">1,<2", true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.filter.test(test.filter.parse(), version)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestFilter_MatchAll(t *testing.T) {
	options := []Version{
		NewVersion("0.5.0"),
		NewVersion("1.0.0"),
		NewVersion("1.0.1"),
		NewVersion("1.2.3"),
		NewVersion("2.0.0"),
	}

	tests := []struct {
		name     string
		filter   Filter
		expected []Version
	}{
		{"all", ">0,<999", options},
		{"v1", ">=1,<2", []Version{"v1.0.0", "v1.0.1", "v1.2.3"}},
		{"exact match", "=1.2.3", []Version{"v1.2.3"}},
		{"higher than 1.2", ">1.2", []Version{"v1.2.3", "v2.0.0"}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := test.filter.MatchAll(options)
			require.Equal(t, test.expected, results)
		})
	}
}

func TestFilter_Match(t *testing.T) {
	options := []Version{
		NewVersion("0.5.0"),
		NewVersion("1.0.0"),
		NewVersion("1.0.1"),
		NewVersion("1.2.3"),
		NewVersion("2.0.0"),
	}

	tests := []struct {
		name     string
		filter   Filter
		expected Version
	}{
		{"all", ">0,<999", "v2.0.0"},
		{"v1", ">=1,<2", "v1.2.3"},
		{"exact match", "=1.2.3", "v1.2.3"},
		{"higher than 1.2", ">1.2", "v2.0.0"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			results := test.filter.Match(options)
			require.Equal(t, test.expected, results)
		})
	}
}
