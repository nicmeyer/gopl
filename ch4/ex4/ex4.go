package main

import (
	"fmt"
)

func rotateLeft(s []int, n int) {
	if len(s) == 0 || n%len(s) == 0 {
		return
	}
	n = n % len(s)
	t := make([]int, len(s))
	copy(t, s)
	copy(s, t[n:])
	copy(s[len(s)-n:], t[:n])
}

func main() {
	s := [...]int{0, 1, 2, 3, 4, 5}
	rotateLeft(s[:], 2)
	fmt.Println(s)
}
