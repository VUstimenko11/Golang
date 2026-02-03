package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type wordCount struct {
	word  string
	count int
}

func Top10(text string) []string {
	words := strings.Fields(text)

	counts := make(map[string]int)
	for _, word := range words {
		counts[word]++
	}

	items := make([]wordCount, 0, len(counts))
	for word, count := range counts {
		items = append(items, wordCount{word: word, count: count})
	}

	sort.Slice(items, func(i, j int) bool {
		if items[i].count != items[j].count {
			return items[i].count > items[j].count
		}
		return items[i].word < items[j].word
	})

	result := make([]string, 0, 10)
	for i := 0; i < len(items) && i < 10; i++ {
		result = append(result, items[i].word)
	}

	return result
}
