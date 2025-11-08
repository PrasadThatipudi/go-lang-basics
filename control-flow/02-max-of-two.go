package main

import "fmt"

// TASK 2: Given two integers a and b, print the larger one.
// If they are equal, print "equal".
// Example: a = 10, b = 7 -> Output: 10
// Example: a = 5, b = 5 -> Output: equal
func main() {
	a := 10
	b := 12
	
	if a == b {
		fmt.Println("equal");
	} else if a < b {
		fmt.Println(b)
	} else {
		fmt.Println(a)
	}
}
