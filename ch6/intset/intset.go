// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"
)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint64
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}

//!-string

func (s *IntSet) Len() int {
	var len int
	for _, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				len++
			}
		}
	}
	return len
}

func (s *IntSet) Remove(x int) {
	word, bit := x/64, uint(x%64)
	if word < len(s.words) {
		s.words[word] &^= (1 << bit)
	}
}

func (s *IntSet) Clear() {
	for i := range s.words {
		s.words[i] = 0
	}
}

func (s *IntSet) Copy() *IntSet {
	var ret IntSet
	ret.UnionWith(s)
	return &ret
}

// AddAll allows a list of values to be added
func (s *IntSet) AddAll(values ...int) {
	for _, val := range values {
		s.Add(val)
	}
}

func (s *IntSet) Elems() []int {
	var ret []int
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<j) != 0 {
				ret = append(ret, 64*i+j)
			}
		}
	}
	return ret
}

func (s *IntSet) IntersectWith(t *IntSet) {
	for _, val := range s.Elems() {
		if !t.Has(val) {
			s.Remove(val)
		}
	}
}

func (s *IntSet) DifferenceWith(t *IntSet) {
	for _, val := range s.Elems() {
		if t.Has(val) {
			s.Remove(val)
		}
	}
}

func (s *IntSet) SymmetricDifference(t *IntSet) {
	// Iterate over the words of both sets
	for i := 0; i < len(s.words) || i < len(t.words); i++ {
		var wordS, wordT uint64
		if i < len(s.words) {
			wordS = s.words[i]
		}
		if i < len(t.words) {
			wordT = t.words[i]
		}

		// Compute the symmetric difference for this word
		symmetricWord := wordS ^ wordT

		// Update the set
		if i < len(s.words) {
			s.words[i] = symmetricWord
		} else {
			s.words = append(s.words, symmetricWord)
		}
	}

	// Trim any trailing zero words
	for len(s.words) > 0 && s.words[len(s.words)-1] == 0 {
		s.words = s.words[:len(s.words)-1]
	}
}
