// exercise: 4.2
package main

import (
	"crypto/sha256"
	"crypto/sha3"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strings"
)

var t = flag.String("t", "sha256", "hash type sha384, sha512, sha256")

func main() {
	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	data := strings.Join(flag.Args(), " ")

	switch hashType := *t; hashType {
	case "sha384":
		fmt.Printf("%x\n", sha3.Sum384([]byte(data)))
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(data)))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(data)))
	}

}
