// exercise 3.10

// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		// fmt.Printf("  %s\n", comma(os.Args[i]))
		// fmt.Printf("  %s\n", comma2(os.Args[i]))
		fmt.Printf("  %s\n", comma3(os.Args[i]))
	}
}

// !+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}

	// Use a bytes.Buffer to build the string efficiently.
	var buf bytes.Buffer

	start := n % 3
	if start == 0 {
		start = 3
	}

	// Write the first part of the string (before the first comma).
	fmt.Fprintf(&buf, "%s", s[:start])

	for i := start; i+3 <= n; i += 3 {
		fmt.Fprintf(&buf, ",%s", s[i:i+3])
	}

	return buf.String()
}

func comma3(s string) string {
	flag := ""
	if s[0] == '-' || s[0] == '+' {
		flag = s[:1]
		s = s[1:]
	}

	dot := strings.Index(s, ".")
	if dot == -1 {
		return flag + comma2(s)
	}
	return flag + comma2(s[:dot]) + "." + comma2(s[dot+1:])
}

//!-
