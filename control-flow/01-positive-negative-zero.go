package main

import "fmt"

// TASK 1: Determine if a number is positive, negative, or zero.
// Print one of: "positive", "negative", or "zero".
// Example: number = 5 -> Output: positive
// Example: number = -2 -> Output: negative
// Example: number = 0 -> Output: zero
func main() {
	number := -1

	if number > 0 {
		fmt.Println("positive")
	} else if number < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero");
	}
}
