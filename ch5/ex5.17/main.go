package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}

func ElementsByTagName(doc *html.Node, name ...string) []*html.Node {
	var m = make(map[string]bool)

	for _, v := range name {
		m[v] = true
	}

	var result []*html.Node
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && m[n.Data] {
			result = append(result, n)
		}
	}, nil)
	return result
}

func printNode(n *html.Node) {
	var content string
	var attribute string
	if n.Type == html.ElementNode {
		for _, v := range n.Attr {
			attribute += fmt.Sprintf("%s=\"%s\" ", v.Key, v.Val)
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			if c.Type == html.TextNode {
				content += c.Data
			}
		}
	}
	fmt.Printf("<%s %s>%s</%s>\n", n.Data, attribute, content, n.Data)
}

func main() {
	resp, err := http.Get(os.Args[1])
	if err != nil {
		resp.Body.Close()
		fmt.Fprintf(os.Stderr, "ElementByTagName: %v\n", err)
		os.Exit(1)
	}

	doc, _ := html.Parse(resp.Body)
	resp.Body.Close()

	images := ElementsByTagName(doc, "img")
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")

	for _, v := range images {
		printNode(v)
	}

	for _, v := range headings {
		printNode(v)
	}
}
