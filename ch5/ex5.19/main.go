package main

import (
	"fmt"
	"math/rand"
)

func nonzero() (ret int) {
	defer func() {
		switch p := recover(); p {
		default:
			ret = rand.Intn(127) + 1
		}
	}()
	panic("?")
}

func main() {
	fmt.Println(nonzero())
}
