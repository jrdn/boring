package version

// Implementation of Semantic Versioning
// https://semver.org/

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	coreRE       = regexp.MustCompile(`(?P<major>[[:digit:]]+)\.(?P<minor>[[:digit:]]+)\.(?P<patch>[[:digit:]])`)
	prereleaseRE = regexp.MustCompile(`-(?P<prerelease>[A-Za-z0-9\.]+)*(\z|\+)`)
	buildRE      = regexp.MustCompile(`\+(?P<build>[A-Za-z0-9\.]|)*\z`)
)

type Version struct {
	major int64
	minor int64
	patch int64

	prerelease []string
	build      []string
}

func (v Version) String() string {
	var prerelease, build string
	if len(v.prerelease) > 0 {
		prerelease = "-" + strings.Join(v.prerelease, ".")
	}
	if len(v.build) > 0 {
		build = "+" + strings.Join(v.build, ".")
	}

	return fmt.Sprintf("%d.%d.%d%s%s", v.major, v.minor, v.patch, prerelease, build)
}

func (v Version) Compare(x Version) int {
	if v.major > x.major {
		return 1
	} else if v.major < x.major {
		return -1
	}
	if v.minor > x.minor {
		return 1
	} else if v.minor < x.minor {
		return -1
	}
	if v.patch > x.patch {
		return 1
	} else if v.patch < x.patch {
		return -1
	}

	vPrereleaseLen := len(v.prerelease)
	xPrereleaseLen := len(x.prerelease)
	// if one has prerelease tags and the other doesn't, the one without resolves higher
	if vPrereleaseLen == 0 && xPrereleaseLen != 0 {
		return 1
	} else if xPrereleaseLen == 0 && vPrereleaseLen != 0 {
		return -1
	}

	// if they both have prerelease tags, compare them left to right until there is a difference
	for i, vt := range v.prerelease {
		// the version with more prerelease tags has precedence over the one with fewer, if all preceding tags are equal
		if i >= xPrereleaseLen {
			return 1
		}
		xt := x.prerelease[i]

		if vt == xt {
			continue
		}
		if vt > xt {
			return 1
		} else {
			return -1
		}
	}

	return 0
}

func Parse(v string) (*Version, error) {
	version := &Version{}

	match := coreRE.FindStringSubmatch(v)
	if match == nil {
		return nil, errors.New("failed to parse version core")
	}

	major, err := strconv.ParseInt(match[coreRE.SubexpIndex("major")], 10, 64)
	if err != nil {
		return nil, errors.New("failed to parse major version as int")
	}
	version.major = major

	minor, err := strconv.ParseInt(match[coreRE.SubexpIndex("minor")], 10, 64)
	if err != nil {
		return nil, errors.New("failed to parse minor version as int")
	}
	version.minor = minor

	patch, err := strconv.ParseInt(match[coreRE.SubexpIndex("patch")], 10, 64)
	if err != nil {
		return nil, errors.New("failed to parse patch version as int")
	}
	version.patch = patch

	prerelease := prereleaseRE.FindStringSubmatch(v)
	if prerelease == nil {
		return nil, errors.New("failed to parse prerelease")
	}
	version.prerelease = strings.Split(prerelease[prereleaseRE.SubexpIndex("prerelease")], ".")

	return version, nil
}
