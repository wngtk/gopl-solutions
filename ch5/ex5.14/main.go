package main

import "fmt"

// !+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

//!-table

func breadthFirst(m map[string][]string, worklist []string) []string {
	var courses []string
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				courses = append(courses, item)
				worklist = append(worklist, m[item]...)
			}
		}
	}
	return courses
}

func main() {
	var keys []string
	for k := range prereqs {
		keys = append(keys, k)
	}

	for i, item := range breadthFirst(prereqs, keys) {
		fmt.Printf("%d: %s\n", i, item)
	}

}
