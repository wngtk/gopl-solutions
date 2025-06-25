package main

import (
	"fmt"
	"sort"
)

func IsPalindrome(s sort.Interface) bool {
	for i := range s.Len() {
		j := s.Len() - i - 1
		if !s.Less(i, j) && !s.Less(j, i) {
			continue
		} else {
			return false
		}
	}
	return true
}

type String string

func (p String) Len() int { return len(p) }
func (p String) Less(i, j int) bool { return p[i] < p[j] }
func (p String) Swap(i, j int) { }

func main() {
	fmt.Println(IsPalindrome(String("HaaH")))
	fmt.Println(IsPalindrome(String("123")))
}
