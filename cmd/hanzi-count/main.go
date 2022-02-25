package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rsookram/hanzi-count/internal/runes"
)

func main() {
	excludes := flag.String("excludes", "", "File containing characters to exclude from the output")
	flag.Parse()

	fileNames := flag.Args()

	if len(fileNames) == 0 {
		fmt.Fprintln(os.Stderr, "Must specify files")
		os.Exit(1)
	}

	counts := countCharacters(fileNames)

	var excludedCounts *runes.Count
	if *excludes != "" {
		excludedCounts = countCharacters([]string{*excludes})
	} else {
		excludedCounts = runes.NewCount()
	}

	stats := computeStats(counts, excludedCounts)

	err := printStats(os.Stdout, stats)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
