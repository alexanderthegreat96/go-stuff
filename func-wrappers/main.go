package main

import "fmt"

// In Go, a function wrapper is a higher-order function that takes one or more functions as arguments and returns a new function.
// Function wrappers are often used to add some behavior or modify the input/output of the wrapped function.
// This is a common pattern in functional programming.

// Common use casses
// - logging
// - timing and profiling
// - authentication and authorization
// - error handling

func main() {
	wrapper := modifyFunctions(example) // provided the function as a parameter
	fmt.Println(wrapper(15, 19))        // provided the function paramters
	// finally executed the function

}

func modifyFunctions(inputFunc func(int, int) int) func(int, int) int {
	fmt.Println("This function is called through the wrapper.")

	return func(x, y int) int { // then returned a function with the same signature
		return inputFunc(x, y)
	}
}

func example(x, y int) int {
	return x + y
}
