package n3

import (
	"regexp"
	"sort"
	"strings"
)

var re = regexp.MustCompile("[^a-zA-Zа-яА-Я-]+")

func Top10(text string) []string {
	parsed := re.ReplaceAllString(text, " ")
	words := strings.Fields(parsed)
	allWords := make(map[string]int)

	for _, word := range words {
		if word == "-" {
			continue
		}
		allWords[word]++
	}

	unsortedWords := make([]string, len(allWords))
	var id int
	for word := range allWords {
		unsortedWords[id] = word
		id++
	}

	sort.Slice(unsortedWords, func(i, j int) bool {
		if allWords[unsortedWords[i]] == allWords[unsortedWords[j]] {
			switch strings.Compare(unsortedWords[i], unsortedWords[j]) {
			case 1:
				return false
			default:
				return true
			}
		}

		return allWords[unsortedWords[i]] > allWords[unsortedWords[j]]
	})

	if len(unsortedWords) < 10 {
		return unsortedWords
	}

	return unsortedWords[:10]
}
