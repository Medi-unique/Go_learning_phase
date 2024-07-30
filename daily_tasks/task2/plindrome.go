package main

import "unicode"

func normalize(s string) string {
	var normalized []rune
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsNumber(r) {
			normalized = append(normalized, unicode.ToLower(r))
		}
	}
	return string(normalized)
}

func palindrome(s string) string {
	s = normalize(s)
	runes := []rune(s)
	n := len(runes) - 1
	l, r := 0, n

	for l <= r {
		if runes[l] != runes[r] {
			return "Not Palindrome"
		}
		l++
		r--
	}
	return "Palindrome"
}