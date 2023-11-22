package main

import (
	"fmt"
	"time"
)

// tellTheTime is a simple function that prints the current time to the console
func tellTheTime() {
	fmt.Println("The time is: " + time.Now().Format("2006-01-02 15:04:05"))
}

// logWrapper is a higher-order function that takes a function f as a parameter
// and returns a new function. The returned function wraps the original function
// by logging the duration it takes to run.
func logWrapper(f func()) func() {
	return func() {
		start := time.Now()                               // Record the start time
		f()                                               // Call the original function here
		duration := time.Since(start)                     // Calculate the duration
		fmt.Printf("Function took %v to run\n", duration) // Log the duration
	}
}

func main() {
	// Create a wrapper function by calling logWrapper with tellTheTime
	wrapper := logWrapper(tellTheTime)

	// Call the wrapped function
	wrapper()

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
