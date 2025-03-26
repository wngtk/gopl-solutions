// Exercise 1.2
// Echo4 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, arg := range os.Args[1:] {
		fmt.Printf("%d: %s\n", i, arg)
	}
}
