package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(text string) []string {
	type words struct {
		word  string
		count int
	}

	if text == "" {
		str := make([]string, 0)
		return str
	}

	mp := make(map[string]int)
	str := strings.Fields(text)

	for _, res := range str {
		mp[res]++
	}

	result := make([]words, 0, len(mp))
	for word, count := range mp {
		result = append(result, words{word: word, count: count})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].count != result[j].count {
			return result[i].count > result[j].count
		}
		return result[i].word < result[j].word
	})

	resSlice := make([]string, 0, len(result))
	for _, val := range result {
		resSlice = append(resSlice, val.word)
	}

	if len(resSlice) < 10 {
		lenSlice := len(resSlice)
		return resSlice[:lenSlice]
	}

	return resSlice[:10]
}
