// Dup2 prints the cout and text of lines that appear more than once
// in the input. It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string][]string)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, files := range counts {
		if len(files) > 1 {
			fmt.Printf("%d\t%s: %s\n", len(files), line, strings.Join(files, ","))
		}
	}
}

func countLines(f *os.File, count map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		count[input.Text()] = append(count[input.Text()], f.Name())
	}
	// NOTE: ignoring potential error from input.Err()
}
