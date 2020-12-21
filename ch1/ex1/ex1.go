// Ex1 prints its command-line arguments including the name of the command that invoked it.
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args, " "))
}
