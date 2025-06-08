package main

import (
	"fmt"
	"io"
	"os"
)

type bytecounter struct {
	w     io.Writer
	count int64
}

func (c *bytecounter) Write(p []byte) (int, error) {
	n, err := c.w.Write(p)
	c.count += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	c := bytecounter{w: w}
	return &c, &c.count
}

func main() {
	w, count := CountingWriter(os.Stdout)

	fmt.Fprintf(w, "Hahaha\n")

	fmt.Println(*count) // 7
}
