package runes

const (
	// Min and Max form the range for reasonably common Chinese characters
	Min = '一'
	Max = '龥'

	total = Max - Min + 1
)

type Count struct {
	array [total]int
}

func NewCount() *Count {
	return &Count{}
}

func (c *Count) Of(r rune) int {
	i := index(r)
	if i < 0 || i >= total {
		return 0
	}

	return c.array[i]
}

func (c *Count) Increment(r rune) {
	i := index(r)
	if i < 0 || i >= total {
		return
	}

	c.array[i]++
}

func (c *Count) MergeWith(other *Count) {
	for i := 0; i < total; i++ {
		c.array[i] += other.array[i]
	}
}

func index(r rune) int {
	return int(r) - Min
}
