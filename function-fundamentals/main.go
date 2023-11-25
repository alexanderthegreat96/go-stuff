package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Function Fundamentals

	// Recursion is a programming concept where a function calls itself in order to solve a problem.
	// In Go, like in many other programming languages, you can use recursion to solve problems that
	// can be broken down into smaller instances of the same problem.

	recursiveFunc(1, 1000)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}

// in other words, this function
// calls itself until it reaches the limit
func recursiveFunc(number int, stopAt int) {
	if number < stopAt {
		recursiveFunc(number+1, stopAt)
		fmt.Println("Currently at: " + strconv.Itoa(number))
	} else {
		fmt.Println("Stopped at: " + strconv.Itoa(stopAt))
	}
}
