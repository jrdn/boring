package fn

import (
	"github.com/jrdn/boring/ds"
	"github.com/jrdn/boring/types"
)

// SliceIterator is provided as a convenience for making a List, since boring
// Lists wrap a slice and are iterable, but using this function can make for
// more readable code in some cases.
func SliceIterator[T any](s []T) types.Iterable[T] {
	return ds.NewList[T](s)
}
