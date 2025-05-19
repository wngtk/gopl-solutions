package main

import (
	"fmt"
	"strings"
)

func join(seq string, elements ...string) string {
	return strings.Join(elements, seq)
}

func main() {
	fmt.Println(join(" ", "a", "b", "c"))
	fmt.Println(join(",", "Hello", "World"))
}
