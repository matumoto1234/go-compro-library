package datastructure

import (
	"sort"

	"github.com/matumoto1234/go-compro-library/algorithm"
	"github.com/matumoto1234/go-compro-library/assert"
	"golang.org/x/exp/constraints"
)

type Compressor[T constraints.Ordered] struct {
	Xs []T
}

func NewCompressor[T constraints.Ordered](vs []T) *Compressor[T] {
	xs := make([]T, len(vs))
	copy(xs, vs)

	sort.Slice(xs, func(i, j int) bool {
		return xs[i] < xs[j]
	})

	return &Compressor[T]{
		Xs: algorithm.Uniq(xs),
	}
}

func (c *Compressor[T]) Do(x T) int {
	i := algorithm.LowerBound(c.Xs, x)
	assert.Do(
		c.Xs[i] == x,
		assert.Msg("Compressor.Do() : x is not in the original array"),
	)
	return i
}
