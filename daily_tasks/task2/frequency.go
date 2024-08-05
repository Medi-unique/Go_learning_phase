package main

import (
	"strings"
	"unicode"
)

func Frequency(input string) map[string]int {
	wordCount := make(map[string]int)

	words := strings.FieldsFunc(input, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}
