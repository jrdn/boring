package c

import (
	"fmt"
	"testing"

	"github.com/jrdn/boring/iface"
)

func TestSet_Iter(t *testing.T) {
	s := NewSet[string]()
	s.Add("foo")
	s.Add("bar")

	var iter iface.Iterable[string] = s
	for item := range iter.Iter() {
		fmt.Println(item)
	}
}

func TestSet_Contains(t *testing.T) {

}
