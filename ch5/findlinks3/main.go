// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 139.

// Findlinks3 crawls the web, starting with the URLs on the command line.
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"gopl.io/ch5/links"
)

// !+breadthFirst
// breadthFirst calls f for each item in the worklist.
// Any items returned by f are added to the worklist.
// f is called at most once for each item.
func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

//!-breadthFirst

var targetHost string

func save(rawURL string) error {
	res, err := url.Parse(rawURL)
	if err != nil {
		return err
	}

	if res.Host != targetHost {
		targetHost = res.Host
		return nil
	}

	resp, err := http.Get(rawURL)
	defer resp.Body.Close()
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("get %s status: %d", rawURL, resp.StatusCode)
	}

	dir := "response"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("creating %s: %v", dir, err)
	}

	path := res.Path
	if path == "" {
		path = "index"
	} else if strings.HasSuffix(path, "/") {
		path = strings.TrimRight(path, "/")
	}
	filename := strings.TrimPrefix(strings.ReplaceAll(path, "/", "_"), "_")
	if filepath.Ext(filename) == "" {
		filename += ".html"
	}

	filePath := filepath.Join(dir, filename)
	file, err := os.Create(filePath)
	defer file.Close()
	if err != nil {
		return fmt.Errorf("creating file %s: %v", filePath, err)
	}

	io.Copy(file, resp.Body)

	return nil
}

// !+crawl
func crawl(url string) []string {
	fmt.Println(url)

	for uri := range cared {
		if url[:min(len(uri), len(url))] == uri {
			err := save(url)
			log.Print(err)
		}
	}

	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!-crawl

var cared = make(map[string]bool)

// !+main
func main() {
	// Crawl the web breadth-first,
	// starting from the command-line arguments.
	startUrls := os.Args[1:]

	if len(startUrls) > 0 {
		targetHost = startUrls[0]
	}

	for _, url := range startUrls {
		cared[url] = true
	}
	breadthFirst(crawl, startUrls)
}

//!-main
