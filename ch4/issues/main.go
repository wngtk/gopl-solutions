// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 112.
//!+

// Issues prints a table of GitHub issues matching the search terms.
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/wngtk/gopl-solutions/ch4/github"
)

// !+
func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)

	yearAgo, inAYear, inAMonth := make([]github.Issue, 0), make([]github.Issue, 0), make([]github.Issue, 0)

	for _, item := range result.Items {
		// a year ago
		if item.CreatedAt.Before(time.Now().AddDate(-1, 0, 0)) {
			yearAgo = append(yearAgo, *item)
		}
		// a month ago
		if item.CreatedAt.Before(time.Now().AddDate(0, -1, 0)) {
			inAYear = append(inAYear, *item)
		}
		// in a month
		if item.CreatedAt.After(time.Now().AddDate(0, -1, 0)) {
			inAMonth = append(inAMonth, *item)
		}
	}
	fmt.Printf("%d issues created in a year ago:\n", len(yearAgo))
	for _, item := range yearAgo {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("%d issues created in a year:\n", len(inAYear))
	for _, item := range inAYear {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Printf("%d issues created in a month:\n", len(inAMonth))
	for _, item := range inAMonth {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

//!-

/*
//!+textoutput
$ go build gopl.io/ch4/issues
$ ./issues repo:golang/go is:open json decoder
13 issues:
#5680    eaigner encoding/json: set key converter on en/decoder
#6050  gopherbot encoding/json: provide tokenizer
#8658  gopherbot encoding/json: use bufio
#8462  kortschak encoding/json: UnmarshalText confuses json.Unmarshal
#5901        rsc encoding/json: allow override type marshaling
#9812  klauspost encoding/json: string tag not symmetric
#7872  extempora encoding/json: Encoder internally buffers full output
#9650    cespare encoding/json: Decoding gives errPhase when unmarshalin
#6716  gopherbot encoding/json: include field name in unmarshal error me
#6901  lukescott encoding/json, encoding/xml: option to treat unknown fi
#6384    joeshaw encoding/json: encode precise floating point integers u
#6647    btracey x/tools/cmd/godoc: display type kind of each named type
#4237  gjemiller encoding/base64: URLEncoding padding is optional
//!-textoutput
*/
