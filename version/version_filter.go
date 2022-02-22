package version

import (
	"sort"
	"strings"
)

// Filter is a way to narrow a selection of versions according to rules
// Multiple conditions can be specified by joining them with a comma. Each condition is ANDed-together
//
// The supported operations are:
// COMPARISON
// <, <=, =, >=, >
// Asserts that the version is greater than / less than / equal to another given version.
// The rule only requires the major version to be specified, omitted minor/patch versions are assumed to be 0
// Example: ">2.4, <3"
type Filter string

func (f Filter) parse() []*condition {
	filter := strings.ReplaceAll(string(f), " ", "")
	filterParts := strings.Split(filter, ",")

	conditions := make([]*condition, len(filterParts))

	operators := []string{">=", "<=", "<", ">", "="}

	for i, part := range filterParts {
	operatorLoop:
		for _, operator := range operators {
			if strings.HasPrefix(part, operator) {
				conditions[i] = &condition{
					operator: operator,
					version:  NewVersion(strings.TrimPrefix(part, operator)),
				}
				break operatorLoop
			}
		}
	}
	return conditions
}
func (f Filter) Match(options []Version) Version {
	matched := f.MatchAll(options)
	sort.Slice(matched, func(i, j int) bool {
		return matched[i].Compare(matched[j]) == 1
	})
	if len(matched) == 0 {
		return ""
	}
	return matched[0]
}

func (f Filter) MatchAll(options []Version) (result []Version) {
	conditions := f.parse()
	for _, o := range options {
		if f.test(conditions, o) {
			result = append(result, o)
		}
	}
	return
}

func (f Filter) test(conditions []*condition, v Version) bool {
	for _, c := range conditions {
		if !c.test(v) {
			return false
		}
	}
	return true
}

type condition struct {
	operator string
	version  Version
}

func (c condition) test(v Version) bool {
	comp := c.version.Compare(v)
	switch {
	case comp == 0:
		return c.operator == "=" || c.operator == "<=" || c.operator == ">="
	case comp > 0:
		return c.operator == "<" || c.operator == "<="
	case comp < 0:
		return c.operator == ">" || c.operator == ">="
	}
	return false
}
