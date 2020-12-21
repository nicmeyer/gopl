package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

var f = flag.String("f", "sha256", "hash function")

func main() {
	flag.Parse()
	for _, arg := range flag.Args() {
		switch *f {
		case "sha384":
			fmt.Printf("%s: %x\n", arg, sha512.Sum384([]byte(arg)))
		case "sha512":
			fmt.Printf("%s: %x\n", arg, sha512.Sum512([]byte(arg)))
		default:
			fmt.Printf("%s: %x\n", arg, sha256.Sum256([]byte(arg)))
		}
	}
}
