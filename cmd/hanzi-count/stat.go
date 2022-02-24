package main

import (
	"bufio"
	"os"
	"strconv"

	"github.com/rsookram/hanzi-count/internal/runes"
)

type stat struct {
	ch    rune
	count int
}

func computeStats(counts *runes.Count) []stat {
	stats := make([]stat, 0)

	for r := runes.Min; r <= runes.MaxInclusive; r++ {
		count := counts.Of(r)
		if count == 0 {
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

func printStats(stats []stat) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for _, stat := range stats {
		w.Write([]byte(string(stat.ch)))
		w.Write([]byte(" "))
		w.Write([]byte(strconv.Itoa(stat.count)))
		w.Write([]byte("\n"))
	}
}
