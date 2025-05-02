package main

import (
	"fmt"
	"os"
)

func equal(x, y map[rune]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, v := range x {
		if y[k] != v {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Please provide a string to compare.")
		return
	}

	result1 := chCount(os.Args[1])
	result2 := chCount(os.Args[2])

	if equal(result1, result2) {
		fmt.Println("The strings are anagrams of each other.")
	} else {
		fmt.Println("The strings are not anagrams of each other.")
	}

}

func chCount(s string) map[rune]int {
	counts := make(map[rune]int)
	for _, ch := range s {
		counts[ch]++
	}
	return counts
}
