package main

import (
	"fmt"

	"github.com/rsookram/hanzi-count/internal/runes"
)

type stat struct {
	ch    rune
	count int
}

func computeStats(counts runes.Counter) []stat {
	stats := make([]stat, 0)

	for r := runes.Min; r <= runes.Max; r++ {
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
	for _, stat := range stats {
		fmt.Println(string(stat.ch), stat.count)
	}
}
