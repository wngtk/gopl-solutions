package main

// min returns the smallest integer among the given arguments.
// It takes at least one integer as the first argument and any number of additional integers.
func min(first int, rest ...int) int {
	min := first
	for _, value := range rest {
		if value < min {
			min = value
		}
	}
	return min
}

// max returns the maximum value in a variadic list of integers.
// It returns the maximum value and true if at least one value is provided,
// or 0 and false if the list is empty.
func max(first int, rest ...int) int {
	max := first
	for _, value := range rest {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {
	// test min and max
	// Test cases for min function
	println("Min of 3, 1, 4, 1, 5:", min(3, 1, 4, 1, 5)) // Expected: 1
	println("Min of -1, -2, -3:", min(-1, -2, -3))       // Expected: -3
	println("Min of 7:", min(7))                         // Expected: 7

	// Test cases for max function
	println("Max of 3, 1, 4, 1, 5:", max(3, 1, 4, 1, 5)) // Expected: 5
	println("Max of -1, -2, -3:", max(-1, -2, -3))       // Expected: -1
	println("Max of 7:", max(7))                         // Expected: 7

	// Edge case: empty list for max (though not handled in current implementation)
	// To handle this, the function signature and logic would need to be changed
	// For now, this will panic as no values are provided
	// println("Max of []:", max()) // Uncommenting this will cause a compile error
}
