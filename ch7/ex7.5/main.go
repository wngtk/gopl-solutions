package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type LimitedReader struct {
	reader io.Reader
	limit  int64
	count  int64
}

func LimitReader(r io.Reader, n int64) io.Reader {
	reader := LimitedReader{reader: r, limit: n, count: 0}
	return &reader
}

func (l *LimitedReader) Read(p []byte) (int, error) {
	if l.count == l.limit {
		return 0, io.EOF
	}
	remain := l.limit - l.count
	var buf []byte
	if len(p) > int(remain) {
		buf = p[:remain]
	}

	n, err := l.reader.Read(buf)
	if err != nil {
		return 0, err
	}
	l.count += int64(n)
	copy(p, buf)
	return n, nil
}

func main() {
	r := LimitReader(strings.NewReader("1234567"), 6)

	input := bufio.NewScanner(r)
	for input.Scan() {
		fmt.Print(input.Text())
	}

}
