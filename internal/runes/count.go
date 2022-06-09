// Package runes implements a counter for runes in the CJK Unified Ideographs
// Unicode block. It provides efficient read, increment, and merge operations
// at the expense of higher memory usage.
package runes

import (
	"errors"
	"math"
)

const (
	// Min is the first character in the CJK Unified Ideographs Unicode block.
	Min = '一'
	// MaxInclusive is the last character in the CJK Unified Ideographs Unicode
	// block.
	MaxInclusive = '\u9fff'

	total = MaxInclusive - Min + 1
)

var errOutOfBounds = errors.New("out of bounds")

// Count maintains counts of runes in the range specified by Min and
// MaxInclusive.
type Count struct {
	array [total]uint
}

// NewCount returns a new Count with the counts of all runes set to 0.
func NewCount() *Count {
	return &Count{}
}

// Of returns the count of the given rune.
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

// MergeWith adds the counts from the given Count to this Count.
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
