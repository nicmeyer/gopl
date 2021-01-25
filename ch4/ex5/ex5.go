package main

import "fmt"

func cadup(s []string) []string {
	var i int = 0
	for _, v := range s {
		if s[i] == v {
			continue
		}
		i++
		s[i] = v
	}
	return s[:i+1]
}

func main() {
	s := []string{"g", "e", "e", "k", "s", "f", "o", "r", "g", "e", "e", "g"}
	s = cadup(s)
	fmt.Println(s)
}
