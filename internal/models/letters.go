package models

import (
	"encoding/json"
	"sort"
	"strings"
	"unicode"
)

type Letters map[rune]int

type LetterValue struct {
	Letter string `json:"name"`
	Value  int    `json:"value"`
}

func NewLetters(s string) Letters {
	l := make(Letters)
	for _, r := range strings.ToLower(s) {
		if unicode.IsLetter(r) {
			l[r]++
		}
	}

	return l
}

func (l Letters) MarshalJSON() ([]byte, error) {
	res := make([]LetterValue, 0, len(l))
	for key, val := range l {
		res = append(res, LetterValue{
			Letter: string(key),
			Value:  val,
		})
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i].Value > res[j].Value
	})

	return json.Marshal(res)
}
