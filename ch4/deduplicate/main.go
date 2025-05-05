// exercise: 4.5
package main

import (
	"fmt"
	"unicode"
)

func main() {
	// Example usage
	strings := []string{"apple", "orange", "banana", "apple", "apple", }
	deduplicated := deduplicate(strings)
	fmt.Println(deduplicated) // Output: ["apple", "orange", "banana", "apple"]

	// Example usage of deduplicateSpace
	s := []byte("hello   world  !  ")
	deduplicatedSpace := deduplicateSpace(s)
	fmt.Println(string(deduplicatedSpace)) // Output: "hello world !"
}

func deduplicate(strings []string) []string {
	j := 0
	for i := 1; i < len(strings); i++ {
		if strings[i] != strings[j] {
			j++
			strings[j] = strings[i]
		}
	}
	return strings[:j+1]
}

// 练习 4.6: 编写一个函数,原地将一个UTF-8编码的[]byte类型的slice中相邻的空格(参考  unicode.IsSpace)替换成一个空格返回
func deduplicateSpace(s []byte) []byte {
	out := s[:0]
	for _, b := range s {
		if unicode.IsSpace(rune(b)) && len(out) > 0 && unicode.IsSpace(rune(out[len(out)-1])) {
			continue
		}
		out = append(out, b)
	}
	return out
}
