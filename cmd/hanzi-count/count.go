package main

import (
	"io/ioutil"
	"sync"

	"github.com/rsookram/hanzi-count/internal/runes"
)

func countCharacters(paths []string) runes.Counter {
	in := gen(paths)

	out := merge(
		countWorker(in),
		countWorker(in),
		countWorker(in),
		countWorker(in),
	)

	total := runes.NewCount()
	for cs := range out {
		total.MergeWith(cs)
	}

	return total
}

func gen(paths []string) <-chan string {
	out := make(chan string, len(paths))

	for _, p := range paths {
		out <- p
	}
	close(out)

	return out
}

func countWorker(in <-chan string) <-chan runes.Counter {
	out := make(chan runes.Counter)

	go func() {
		counts := runes.NewCount()
		for path := range in {
			data, err := ioutil.ReadFile(path)
			if err != nil {
				continue
			}

			for _, r := range string(data) {
				counts.Increment(r)
			}
		}
		out <- counts

		close(out)
	}()

	return out
}

func merge(cs ...<-chan runes.Counter) <-chan runes.Counter {
	var wg sync.WaitGroup
	out := make(chan runes.Counter)

	output := func(c <-chan runes.Counter) {
		for n := range c {
			out <- n
		}
		wg.Done()
	}
	wg.Add(len(cs))
	for _, c := range cs {
		go output(c)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}
