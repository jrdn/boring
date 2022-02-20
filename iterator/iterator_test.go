package iterator

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestIterator(t *testing.T) {
	expected := []string{"foo", "bar", "baz"}
	iter := NewIterator[string](func(ctx context.Context) chan string {
		c := make(chan string)
		defer close(c)
		done := ctx.Done()
		go func() {
			for _, x := range expected {
				select {
				case <-done:
					return
				case c <- x:
					// sent

				}
			}
		}()
		return c
	})

	var results []string
	for x := range iter.Range() {
		results = append(results, x)
		fmt.Println(x)
	}

	require.NotEmpty(t, results)

}
