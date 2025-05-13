// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"os"
	"strings"
)

// !+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

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

// !+main
func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func index(vec []string, v string) (int, error) {
	for i, s := range vec {
		if s == v {
			return i, nil
		}
	}
	return -1, fmt.Errorf("not found")
}

func topoSort(m map[string][]string) (order []string, err error) {
	resolved := make(map[string]bool)
	var visitAll func(items []string, path []string)

	visitAll = func(items []string, path []string) {
		for _, item := range items {
			if seen, ok := resolved[item]; !seen && ok {
				start, _ := index(path, item)
				err = fmt.Errorf("cycle: %s", strings.Join(append(path[start:], item), "->"))
			}
			if _, ok := resolved[item]; !ok {
				resolved[item] = false
				visitAll(m[item], append(path, item))
				resolved[item] = true
				order = append(order, item)
			}
		}

	}

	for k := range m {
		if err != nil {
			return nil, err
		}
		visitAll([]string{k}, nil)
	}

	return order, nil
}

//!-main
