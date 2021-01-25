package main

import (
	"unicode"
	"fmt"
)

func squash(s []rune) []rune {
	var j int = 0
	for i := 0; i < len(s); i++ {
		if unicode.IsSpace(s[i]) == true {
			continue
		}
		s[j] = s[i]
		j++
	}
	return s[:j+1]
}

func main() {
}
