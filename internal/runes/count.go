package runes

import (
	"errors"
	"math"
)

const (
	// Min and MaxInclusive form the range for reasonably common Chinese characters
	Min          = '一'
	MaxInclusive = '龥'

	total = MaxInclusive - Min + 1
)

var errOutOfBounds = errors.New("out of bounds")

type Count struct {
	array [total]uint
}

func NewCount() *Count {
	return &Count{}
}

func (c *Count) Of(r rune) uint {
	i, err := index(r)
	if err != nil {
		return 0
	}

	return c.array[i]
}

// Increment increments the count for the given rune by one.
func (c *Count) Increment(r rune) {
	i, err := index(r)
	if err != nil {
		return
	}

	c.array[i]++
}

func (c *Count) MergeWith(other *Count) {
	for i, cnt := range other.array {
		c.array[i] += cnt
	}
}

func index(r rune) (uint, error) {
	if r < Min || r > MaxInclusive {
		return math.MaxUint, errOutOfBounds
	}

	return uint(r) - Min, nil
}
