package main

import (
	"bufio"
	"io"
	"strconv"

	"github.com/rsookram/hanzi-count/internal/runes"
)

type stat struct {
	ch    rune
	count uint
}

func computeStats(counts, excluded *runes.Count) []stat {
	stats := make([]stat, 0)

	for r := runes.Min; r <= runes.MaxInclusive; r++ {
		count := counts.Of(r)
		if count == 0 {
			continue
		}

		if excluded.Of(r) > 0 {
			continue
		}

		stat := stat{
			ch:    r,
			count: count,
		}

		stats = append(stats, stat)
	}

	return stats
}

func printStats(w io.Writer, stats []stat) error {
	buf := bufio.NewWriter(w)
	defer buf.Flush()

	for _, stat := range stats {
		if _, err := io.WriteString(buf, string(stat.ch)); err != nil {
			return err
		}

		if _, err := io.WriteString(buf, " "); err != nil {
			return err
		}

		if _, err := io.WriteString(buf, strconv.FormatUint(uint64(stat.count), 10)); err != nil {
			return err
		}

		if _, err := io.WriteString(buf, "\n"); err != nil {
			return err
		}
	}

	return nil
}
