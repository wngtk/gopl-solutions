// exercise: 4.5
package main

import "fmt"

func main() {
	// Example usage
	strings := []string{"apple", "orange", "banana", "apple", "apple", }
	deduplicated := deduplicate(strings)
	fmt.Println(deduplicated) // Output: ["apple", "orange", "banana", "apple"]
}

func deduplicate(strings []string) []string {
	j := 0
	for i := 1; i < len(strings); i++ {
		if strings[i] != strings[j] {
			j++
			strings[j] = strings[i]
		}
	}
	return strings[:j+1]
}
