// Fetachall fetches URLs in parallel and reports their times and sizes as well as their body
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for _, url := range os.Args[1:] {
		replacer := strings.NewReplacer("://", "-", ".", "-")
		filename := replacer.Replace(url)
		file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		defer file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetchall: while opening %s: %v\n", filename, err)
		}
		msg := fmt.Sprintf("creation time: %s\nresult\n%v", start, <-ch)
		_, err = file.WriteString(msg) // write to file from channel ch
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetchall: while writing %s: %v\n", filename, err)
		}
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d\n%s\n", secs, len(b), b)
}
