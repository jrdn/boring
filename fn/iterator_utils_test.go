package fn

import (
	"context"
	"testing"

	"github.com/jrdn/boring/ds"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCollect(t *testing.T) {
	input := []string{"foo", "bar", "baz", "quux"}

	result := Collect[string](context.TODO(), ds.NewList(input))

	require.NotNil(t, result)
	require.NotEmpty(t, result)

	for i, r := range result.GetSlice() {
		assert.Equal(t, input[i], r)
	}
}

func TestMap(t *testing.T) {
	data := []string{"foo", "bar", "baz", "quux"}
	lst := ds.NewList(data)
	results := Collect[int](context.TODO(), Map[string, int](context.TODO(), func(x string) int {
		return len(x)
	}, lst)).GetSlice()
	require.NotNil(t, results)
	require.NotEmpty(t, results)

	for i, v := range data {
		assert.Equal(t, len(v), results[i])
	}
}

func TestReduce(t *testing.T) {
	expected := 1 + 2 + 3 + 4 + 5
	lst := ds.NewList([]int{1, 2, 3, 4, 5})
	result := Reduce[int](context.TODO(), func(a, b int) int {
		return a + b
	}, lst)

	assert.Equal(t, expected, result)
}

func TestFilter(t *testing.T) {
	input := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	expected := []int{0, 2, 4, 6, 8, 10}

	result := Collect(context.TODO(), Filter[int](context.TODO(), func(x int) bool {
		return x%2 == 0
	}, ds.NewList(input))).GetSlice()

	assert.Equal(t, expected, result)
}

func TestChain(t *testing.T) {
	expected := []int{0, 1, 2, 3, 4}
	ctx := context.TODO()

	result := Collect(ctx, Chain[int](ctx,
		ds.NewList([]int{0, 1, 2}),
		ds.NewList([]int{3, 4}),
	)).GetSlice()

	require.NotNil(t, result)
	require.NotEmpty(t, result)

	assert.Equal(t, expected, result)
}

func TestTee(t *testing.T) {
	ctx := context.TODO()
	first, second := Tee[string](ctx, ds.NewList[string]([]string{"foo", "bar", "baz"}))

	firstIter := first.Iter(ctx)
	secondIter := second.Iter(ctx)

	for firstItem := range firstIter {
		secondItem := <-secondIter
		assert.Equal(t, firstItem, secondItem)
	}
}

func TestRepeat(t *testing.T) {
	ctx := context.TODO()
	result := CollectN(ctx, 3, Repeat("foo")).GetSlice()
	assert.Equal(t, []string{"foo", "foo", "foo"}, result)
}

func TestRepeatTimes(t *testing.T) {
	ctx := context.TODO()
	result := Collect(ctx, RepeatTimes("foo", 5)).GetSlice()
	assert.Len(t, result, 5)
	assert.Equal(t, "foo", result[0])
}

func TestPairwise(t *testing.T) {
	ctx := context.TODO()
	expected := [][]string{
		{"foo", "bar"},
		{"baz", "quux"},
	}
	pairwiseIter := Pairwise[string](ctx, ds.NewList[string]([]string{"foo", "bar", "baz", "quux"}))

	i := 0
	for pair := range pairwiseIter.Iter(ctx) {
		assert.Equal(t, expected[i][0], pair.First())
		assert.Equal(t, expected[i][1], pair.Second())
		i++
	}
}

func TestZip(t *testing.T) {
	expected := [][]string{
		{"foo", "bar"},
		{"baz", "quux"},
	}

	ctx := context.TODO()

	first := ds.NewList[string]([]string{"foo", "baz"})
	second := ds.NewList[string]([]string{"bar", "quux", "hello world"})
	i := 0

	for pair := range Zip[string, string](ctx, first, second).Iter(ctx) {
		assert.Equal(t, expected[i][0], pair.First())
		assert.Equal(t, expected[i][1], pair.Second())
		i++
	}
}
