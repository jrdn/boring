package fn

import (
	"context"
	"math/rand"
	"runtime"
	"testing"
	"time"

	"github.com/jrdn/boring/types"
	"github.com/stretchr/testify/assert"
)

func TestFuncIterator(t *testing.T) {
	newFuncIterator := func() types.Iterable[int] {
		counter := 0

		return FuncIterator[int](func() (int, bool) {
			for counter < 10 {
				val := counter
				counter++
				return val, false
			}
			return 10, true
		})
	}

	result := Collect[int](context.TODO(), newFuncIterator()).GetSlice()
	expected := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, expected, result)
}

func TestInfiniteFuncIterator(t *testing.T) {
	rand.Seed(13) // 13 is a good seed. A seed of the people.

	ctx, cancel := context.WithCancel(context.Background())

	startGoroutines := runtime.NumGoroutine()
	t.Log("start goroutines: ", runtime.NumGoroutine())

	result := CollectN(ctx, 5, FuncIterator[int](func() (int, bool) {
		return rand.Intn(10), false //nolint:gosec
	})).GetSlice()

	t.Log("iterator goroutines: ", runtime.NumGoroutine())

	cancel()
	time.Sleep(500 * time.Millisecond)

	endGoroutines := runtime.NumGoroutine()
	t.Log("end goroutines: ", endGoroutines)
	assert.Equal(t, startGoroutines, endGoroutines)

	expected := []int{2, 1, 1, 5, 4}
	assert.Equal(t, expected, result)
}
