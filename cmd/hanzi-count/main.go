package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Parse()

	fileNames := flag.Args()

	if len(fileNames) == 0 {
		fmt.Fprintln(os.Stderr, "Must specify files")
		os.Exit(1)
	}

	counts := countCharacters(fileNames)

	stats := computeStats(counts)

	printStats(stats)
}
