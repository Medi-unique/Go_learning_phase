package main

import (
	"strings"
	
)
func count(s string) map[string]int {
	mycount := make(map[string]int)
	for _, char := range strings.Split(s, ""){
		if (char >= "A" && char <= "Z") || (char >= "a" && char <= "z"){
			mycount[strings.ToLower(char)]+=1
		}

	}
		return mycount
	

}