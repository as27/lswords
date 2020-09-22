package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/as27/lswords/count"
)

var (
	flagTop    = flag.Int("t", 25, "print top t words")
	flagMinLen = flag.Int("min", 3, "min length for the word")
)

var in io.Reader

func main() {
	flag.Parse()
	run(os.Stdin, os.Stdout, os.Stderr)
}

func run(in io.Reader, out, err io.Writer) {
	words := count.Words(in)
	total := 0
	for _, w := range words {
		total += w.Count
	}
	i := 1
	for _, w := range words {
		if i > *flagTop {
			break
		}
		if len(w.Name) < *flagMinLen {
			continue
		}
		prct := float64(w.Count) / float64(total) * 100
		fmt.Fprintf(out, "% 6d % 3.2f%%  %s\n",
			w.Count, prct, w.Name)
		i++
	}
}
