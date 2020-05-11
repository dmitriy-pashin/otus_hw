package hw03_frequency_analysis //nolint:golint,stylecheck
import (
	"regexp"
	"sort"
	"strings"
)

const topLength = 10

func Top10(text string) []string {
	if len(text) == 0 {
		return []string{}
	}

	var counters = make(map[string]int)

	words := getWords(text)

	for _, w := range words {
		counters[strings.ToLower(w)]++
	}

	topWords := getTop(counters)

	if len(topWords) <= topLength {
		return topWords
	}

	return topWords[0 : topLength-1]
}

func getWords(text string) []string {
	words := regexp.MustCompile(`[А-Яа-яA-Za-z\-]{2,}|[А-Яа-яA-Za-z]`)
	return words.FindAllString(text, -1)
}

func getTop(countersMap map[string]int) []string {
	keys := make([]string, 0, len(countersMap))

	for key := range countersMap {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		return countersMap[keys[i]] > countersMap[keys[j]]
	})

	return keys
}
