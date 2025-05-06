// 练习 4.9： 编写一个程序wordfreq程序，报告输入文本中每个单词出现的频率。在第一次调用Scan前先调用input.Split(bufio.ScanWords)函数，这样可以按单词而不是按行输入。
package main

import (
	"bufio"
	"os"
	"sort"
)

func main() {
	input := bufio.NewScanner(os.Stdin)
	freq := make(map[string]int) // Map to store word frequencies

	input.Split(bufio.ScanWords) // Split the input into words
	for input.Scan() {
		word := input.Text() // Get the current word
		freq[word]++         // Increment the frequency count for the word
	}

	keys := make([]string, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
	// Sort the words by their frequency
	sort.Slice(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})

	for _, word := range keys {
		println(word, freq[word]) // Print the word and its frequency
	}
}
