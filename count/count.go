// Package count is for counting words of a given inputstream
package count

import (
	"bufio"
	"io"
	"sort"
	"strings"
)

// words count the words of the io.Reader and stores everything
// inside a map.
func words(r io.Reader) map[string]int {
	counter := make(map[string]int)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		w := strings.Trim(scanner.Text(), "#[]()&{}/'´`\",.!;:")
		if skipWord(w) {
			continue
		}
		counter[w]++
	}
	return counter
}

// skipWord defines words, which should not be counted
func skipWord(w string) bool {
	switch w {
	case "":
		return true
	}
	return false
}

// Word representate the result of counting the occurences
// of a word in a text.
type Word struct {
	Name  string
	Count int
}

// Words scanns the data from the io.Reader and creates
// a slice of Words. The slice is sorted by the word
// count.
func Words(r io.Reader) []Word {
	wordsMap := words(r)
	words := []Word{}
	for w, c := range wordsMap {
		words = append(words, Word{w, c})
	}
	sort.Slice(words, func(i, j int) bool {
		return words[i].Count > words[j].Count
	})
	return words
}
