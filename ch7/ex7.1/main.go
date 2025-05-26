// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"bufio"
	"fmt"
	"strings"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

//!-bytecounter

//!+linecounter

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	ret := 0
	input := bufio.NewScanner(strings.NewReader(string(p)))
	for input.Scan() {
		*c += LineCounter(1)
		ret++
	}
	return ret, nil
}

//!-linecounter

//!+wordcounter

type WordCount int

func (c *WordCount) Write(p []byte) (int, error) {
	count := 0
	scanner := bufio.NewScanner(strings.NewReader(string(p)))
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		advance, token, err = bufio.ScanWords(data, atEOF)
		if err == nil && token != nil {
			count++
		}
		return
	})
	for scanner.Scan() {
		*c += 1
	}
	return count, nil
}

//!-wordcounter

func main() {
	//!+main
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main

	var lc LineCounter
	lc.Write([]byte(`Go: Go is better than Rust!
Rust: Rust is better than Go!
`))
	fmt.Println(lc)
	lc = 0
	fmt.Fprintf(&lc, "hello, %s", name)
	fmt.Println(lc)

	var wc WordCount
	wc.Write([]byte(`Usage: wc [OPTION]... [FILE]...
  or:  wc [OPTION]... --files0-from=F
Print newline, word, and byte counts for each FILE, and a total line if
more than one FILE is specified.  A word is a non-zero-length sequence of
printable characters delimited by white space.

With no FILE, or when FILE is -, read standard input.

The options below may be used to select which counts are printed, always in
the following order: newline, word, character, byte, maximum line length.
  -c, --bytes            print the byte counts
  -m, --chars            print the character counts
  -l, --lines            print the newline counts
      --files0-from=F    read input from the files specified by
                           NUL-terminated names in file F;
                           If F is - then read names from standard input
  -L, --max-line-length  print the maximum display width
  -w, --words            print the word counts
      --total=WHEN       when to print a line with total counts;
                           WHEN can be: auto, always, only, never
      --help        display this help and exit
      --version     output version information and exit

GNU coreutils online help: <https://www.gnu.org/software/coreutils/>
Report any translation bugs to <https://translationproject.org/team/>
Full documentation <https://www.gnu.org/software/coreutils/wc>
or available locally via: info '(coreutils) wc invocation'`))
	fmt.Println(wc)
}
