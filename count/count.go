package count

import (
	"bufio"
	"io"
	"sort"
	"strings"
)

func words(r io.Reader) map[string]int {
	counter := make(map[string]int)
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		w := strings.Trim(scanner.Text(), "'\",.!;:")
		if skipWord(w) {
			continue
		}
		counter[w]++
	}
	return counter
}

func skipWord(w string) bool {
	switch w {
	case "":
		return true
	}
	return false
}

type Word struct {
	Name  string
	Count int
}

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
