package main

import (
	"bufio"
	"io"
	"os"
	"runtime"
	"sync"

	"github.com/rsookram/hanzi-count/internal/runes"
)

// countCharacters counts the Chinese runes in the files at the given paths.
func countCharacters(paths []string) *runes.Count {
	in := gen(paths)

	ws := make([]<-chan *runes.Count, 0, runtime.NumCPU())
	for i := 0; i < cap(ws); i++ {
		ws = append(ws, countWorker(in))
	}

	out := merge(ws...)

	total := runes.NewCount()
	for cs := range out {
		total.MergeWith(cs)
	}

	return total
}

// gen returns a channel which emits the given paths.
func gen(paths []string) <-chan string {
	out := make(chan string, len(paths))

	for _, p := range paths {
		out <- p
	}
	close(out)

	return out
}

// countWorker reads filepaths off the given channel, and emits its results
// (counts) on the returned channel.
func countWorker(in <-chan string) <-chan *runes.Count {
	out := make(chan *runes.Count)

	go func() {
		counts := runes.NewCount()
		for path := range in {
			err := count(counts, path)
			if err != nil {
				continue
			}
		}
		out <- counts

		close(out)
	}()

	return out
}

// count computes the frequencies of each Chinese rune in the file with the
// given path.
func count(c *runes.Count, path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	for {
		if r, _, err := reader.ReadRune(); err != nil {
			if err == io.EOF {
				break
			} else {
				return err
			}
		} else {
			c.Increment(r)
		}
	}

	return nil
}

// merge reads results off the given channels, and emits all of them through
// the single returned channel.
func merge(cs ...<-chan *runes.Count) <-chan *runes.Count {
	var wg sync.WaitGroup
	out := make(chan *runes.Count)

	output := func(c <-chan *runes.Count) {
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
