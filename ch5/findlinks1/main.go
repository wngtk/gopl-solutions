// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}

	for k, v := range countElements(make(map[string]int), doc) {
		fmt.Printf("%10s%5d\n", k, v)
	}

	allText(os.Stdout, doc)
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

	if n.FirstChild != nil {
		links = visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		links = visit(links, n.NextSibling)
	}

	return links
}

//!-visit

func countElements(count map[string]int, n *html.Node) map[string]int {
	if n.Type == html.ElementNode {
		count[n.Data]++
	}

	if n.FirstChild != nil {
		count = countElements(count, n.FirstChild)
	}

	if n.NextSibling != nil {
		count = countElements(count, n.NextSibling)

	}
	return count
}

func allText(out io.Writer, n *html.Node) {
	if n.Type == html.ElementNode {
		if n.Data == "style" || n.Data == "script" {
			return
		}
	}

	if n.Type == html.TextNode {
		if strings.TrimSpace(n.Data) != "" {
			fmt.Fprintf(out, "%s", n.Data)
		}
	}

	if n.FirstChild != nil {
		allText(out, n.FirstChild)
	}

	if n.NextSibling != nil {
		allText(out, n.NextSibling)
	}
}

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
