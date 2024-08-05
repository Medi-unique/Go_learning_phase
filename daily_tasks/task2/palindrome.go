package main

import (
	"strings"
	"unicode"
)

func Palindrome(input string) bool {
	cleanedInput := strings.ToLower(strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			return r
		}
		return -1
	}, input))

	for i := 0; i < len(cleanedInput)/2; i++ {
		if cleanedInput[i] != cleanedInput[len(cleanedInput)-1-i] {
			return false
		}
	}
	return true
}
