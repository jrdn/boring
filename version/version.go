package version

import (
	"strings"

	"golang.org/x/mod/semver"
)

func NewVersion(version string) Version {
	if !strings.HasPrefix(version, "v") {
		version = "v" + version
	}
	if !semver.IsValid(version) {
		return ""
	}

	return Version(semver.Canonical(version))
}

// Version wraps the Go semver library which mostly matches SemVer 2.0
// https://pkg.go.dev/golang.org/x/mod/semver
type Version string

func (v Version) Equal(other Version) bool {
	return v.Compare(other) == 0
}

// Compare returns -1 if this version is < other, 1 if it's greater, and 0 if they're equal
func (v Version) Compare(other Version) int {
	return semver.Compare(string(v), string(other))
}

// String returns the canonical representation of this version
func (v Version) String() string {
	return semver.Canonical(string(v))
}
