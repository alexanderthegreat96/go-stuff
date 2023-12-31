package main

import (
	"fmt"
	"math"
	"strconv"
)

// closure exercise
// enclosed the scope of the variable in some code block

func raiseToPower(providedNumber float64) func() float64 {
	// each time this function is run
	// the current will increase by 1
	var current float64

	return func() float64 {
		current++
		// raise the provided number to the power of current
		return math.Pow(providedNumber, current)
	}
}

// a more functional approach
// we take in the input_string
// and keep adding + 1 to each print
// notice how we did not provide an input to the main function
// as we do not need that, but we provided an input to the return function
// func numberedPrint() func(input_string string) string {}
// and replicated it into the return func(input_string string) string {}

func numberedPrint() func(input_string string) string {
	start := 0

	return func(input_string string) string {
		start++
		return strconv.Itoa(start) + ". " + input_string
	}
}

func main() {
	raise := raiseToPower(6)

	// invoked it multiple times
	fmt.Println(raise())
	fmt.Println(raise())
	fmt.Println(raise())
	fmt.Println(raise())
	fmt.Println(raise())
	fmt.Println(raise())

	// different approach
	// a more practical use case
	numbered_print := numberedPrint()

	fmt.Println(numbered_print("Alex"))
	fmt.Println(numbered_print("Joe"))
	fmt.Println(numbered_print("Michael"))
	fmt.Println(numbered_print("David"))
	fmt.Println(numbered_print("Seth"))

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
