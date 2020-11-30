package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	// fmt.Println(s)
}

func echo3() {
	strings.Join(os.Args[1:], " ")
	// fmt.Println(strings.Join(os.Args[1:], " "))
}

func timeEcho2() {
	start := time.Now()
	echo2()
	fmt.Printf("echo2: %.2fs elapsed\n", time.Since(start).Seconds())
}

func timeEcho3() {
	start := time.Now()
	echo3()
	fmt.Printf("echo3: %.2fs elapsed\n", time.Since(start).Seconds())
}

func main() {
	timeEcho2()
	timeEcho3()
}
