package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

var index = `
<!DOCTYPE html>
<html>
	<head>Test Doc</head>
	<body>
		<a href="https://www.baidu.com">Baidu</a>
	</body>
</html>
`

type myreader struct {
	s      string
	offset int
}

func newMyReader(s string) *myreader {
	var reader = myreader{s: s}
	return &reader
}

func (r *myreader) Read(p []byte) (int, error) {
	if r.offset >= len(r.s) {
		return 0, io.EOF
	}
	n := copy(p, []byte(r.s))
	r.offset += n
	return n, nil
}

func main() {
	doc, err := html.Parse(newMyReader(index))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

//!-main

// !+visit
// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
