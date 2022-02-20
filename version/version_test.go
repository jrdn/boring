package version

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    *Version
		wantErr bool
	}{
		{"fix leading zero", "1.02.3", &Version{major: 1, minor: 2, patch: 3}, false},
		{"not a version", "foo", nil, true},
		{"prerelease", "1.2.3-foo.bar", &Version{
			major:      1,
			minor:      2,
			patch:      3,
			prerelease: []string{"foo", "bar"},
		}, false},
		{"prerelease and build", "1.2.3-foo.bar+baz.quux", &Version{
			major:      1,
			minor:      2,
			patch:      3,
			prerelease: []string{"foo", "bar"},
			build:      []string{"baz", "quux"},
		}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
