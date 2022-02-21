package runes

import "testing"

func TestEmptyCounts(t *testing.T) {
	c := NewCount()

	tables := []struct {
		r rune
		n int
	}{
		{'a', 0},
		{'一', 0},
		{'个', 0},
		{'开', 0},
		{'龥', 0},
	}

	for _, table := range tables {
		count := c.Of(table.r)
		if count != table.n {
			t.Errorf("Count of %s is incorrect. Got %d, expected %d", string(table.r), count, table.n)
		}
	}
}

func TestASCIINotCounted(t *testing.T) {
	c := NewCount()

	s := "Go is..."
	for _, r := range s {
		c.Increment(r)
	}

	tables := []struct {
		r rune
		n int
	}{
		{'G', 0},
		{'o', 0},
		{' ', 0},
		{'i', 0},
		{'s', 0},
		{'.', 0},
	}

	for _, table := range tables {
		count := c.Of(table.r)
		if count != table.n {
			t.Errorf("Count of %s is incorrect. Got %d, expected %d", string(table.r), count, table.n)
		}
	}
}

func TestSingleIncrement(t *testing.T) {
	c := NewCount()

	s := "一个开源的编程语言"
	for _, r := range s {
		c.Increment(r)
	}

	tables := []struct {
		r rune
		n int
	}{
		{'一', 1},
		{'个', 1},
		{'开', 1},
		{'源', 1},
		{'的', 1},
		{'编', 1},
		{'程', 1},
		{'语', 1},
		{'言', 1},
	}

	for _, table := range tables {
		count := c.Of(table.r)
		if count != table.n {
			t.Errorf("Count of %s is incorrect. Got %d, expected %d", string(table.r), count, table.n)
		}
	}
}

func TestMultiIncrement(t *testing.T) {
	c := NewCount()

	s := "你可以编辑这里的代码！点击这里然后开始输入。"
	for _, r := range s {
		c.Increment(r)
	}

	tables := []struct {
		r rune
		n int
	}{
		{'你', 1},
		{'可', 1},
		{'以', 1},
		{'编', 1},
		{'辑', 1},
		{'这', 2},
		{'里', 2},
		{'的', 1},
		{'代', 1},
		{'码', 1},
		{'！', 0},
		{'点', 1},
		{'击', 1},
		{'然', 1},
		{'后', 1},
		{'开', 1},
		{'始', 1},
		{'输', 1},
		{'入', 1},
		{'。', 0},
	}

	for _, table := range tables {
		count := c.Of(table.r)
		if count != table.n {
			t.Errorf("Count of %s is incorrect. Got %d, expected %d", string(table.r), count, table.n)
		}
	}
}

func TestMerge(t *testing.T) {
	c := NewCount()
	c2 := NewCount()

	s := "你可以编辑这里的代码！"
	for _, r := range s {
		c.Increment(r)
	}

	s2 := "点击这里然后开始输入。"
	for _, r := range s2 {
		c2.Increment(r)
	}

	c.MergeWith(c2)

	tables := []struct {
		r rune
		n int
	}{
		{'你', 1},
		{'可', 1},
		{'以', 1},
		{'编', 1},
		{'辑', 1},
		{'这', 2},
		{'里', 2},
		{'的', 1},
		{'代', 1},
		{'码', 1},
		{'！', 0},
		{'点', 1},
		{'击', 1},
		{'然', 1},
		{'后', 1},
		{'开', 1},
		{'始', 1},
		{'输', 1},
		{'入', 1},
		{'。', 0},
	}

	for _, table := range tables {
		count := c.Of(table.r)
		if count != table.n {
			t.Errorf("Count of %s is incorrect. Got %d, expected %d", string(table.r), count, table.n)
		}
	}

	// c2 is unchanged
	tables = []struct {
		r rune
		n int
	}{
		{'点', 1},
		{'击', 1},
		{'这', 1},
		{'里', 1},
		{'然', 1},
		{'后', 1},
		{'开', 1},
		{'始', 1},
		{'输', 1},
		{'入', 1},
		{'。', 0},
	}

	for _, table := range tables {
		count := c2.Of(table.r)
		if count != table.n {
			t.Errorf("Count of %s is incorrect. Got %d, expected %d", string(table.r), count, table.n)
		}
	}
}
