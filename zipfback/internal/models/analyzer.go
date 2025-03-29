package models

import (
	"regexp"
	"strings"
	"unicode"
)

type LetterFrequency struct{}

func NewLetterFrequency() *LetterFrequency {
	return &LetterFrequency{}
}

func (l *LetterFrequency) Run(s string) Result {
	r := make(Result)
	for _, c := range strings.ToLower(s) {
		if unicode.IsLetter(c) {
			r[string(c)]++
		}
	}

	return r
}

type WordFrequency struct{}

func NewWordFrequency() *WordFrequency {
	return &WordFrequency{}
}

func (w *WordFrequency) Run(s string) Result {
	r := make(Result)

	// Compile the regex to match all punctuation characters
	re := regexp.MustCompile(`[^\w\s]`)

	// Replace all matches with an empty string
	sClean := re.ReplaceAllString(s, "")

	for _, word := range strings.Split(sClean, " ") {
		w := strings.ToLower(word)
		if w != "" {
			r[w]++
		}
	}

	return r
}
