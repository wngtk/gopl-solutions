// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import "fmt"

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	fmt.Println(x.Len()) // "4"

	x.Remove(144)
	fmt.Println(x.Has(144), x.Len()) // "false 3"

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
	// 4
	// false 3
}

func Example_three() {
	var x IntSet
	x.AddAll(9, 6, 0, 4, 8, 5, 9, 0, 2)

	fmt.Printf("%v\n", &x)
	fmt.Println(x.Len())
	y := x.Copy()
	fmt.Println(y.String())
	y.Clear()
	fmt.Println(y.String())
	fmt.Printf("%v\n", &x)

	// Output:
	// {0 2 4 5 6 8 9}
	// 7
	// {0 2 4 5 6 8 9}
	// {}
	// {0 2 4 5 6 8 9}
}

func Example_four() {
	var x IntSet
	x.AddAll(0, 1, 2, 3, 4, 5, 6, 7, 8, 9)

	for i, val := range x.Elems() {
		if i != 0 {
			fmt.Printf(" ")
		}
		fmt.Printf("%d", val)
	}

	// Output:
	// 0 1 2 3 4 5 6 7 8 9
}

func Example_symmetricDifference() {
	var x, y IntSet
	x.AddAll(1, 2, 3, 4)
	y.AddAll(3, 4, 5, 6)

	x.SymmetricDifference(&y)
	fmt.Println(x.String()) // "{1 2 5 6}"

	// Output:
	// {1 2 5 6}
}

func Example_differenceWith() {
	var x, y IntSet
	x.AddAll(1, 2, 3, 4)
	y.AddAll(3, 4, 5, 6)

	x.DifferenceWith(&y)
	fmt.Println(x.String()) // "{1 2}"

	// Output:
	// {1 2}
}

func Example_intersectWith() {
	var x, y IntSet
	x.AddAll(1, 2, 3, 4)
	y.AddAll(3, 4, 5, 6)

	x.IntersectWith(&y)
	fmt.Println(x.String()) // "{3 4}"

	// Output:
	// {3 4}
}
