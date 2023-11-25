package main

import "fmt"

// defer functions
// delay the function or code exection
// untill the rest of the segment is finished

// basically it runs after the function containing it
// exists

func main() {
	// start using the defer

	defer first()  // runs last
	defer second() // runs second
	defer third()  // runs first

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}

func first() {
	fmt.Println("I am the first function.")
}

func second() {
	fmt.Println("I am the second function")
}

func third() {
	fmt.Println("I am the third function")
}
