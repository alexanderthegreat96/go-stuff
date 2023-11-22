package main

import "fmt"

// In the Go programming language, the defer statement is used
// to ensure that a function call is performed later in a program's
// execution, usually for purposes of cleanup. The deferred function
// calls are executed in Last In, First Out (LIFO) order.
// The defer statement is often used to simplify functions that
// perform various setup and teardown tasks.

// another definition from go.dev

// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call is not executed
// until the surrounding function returns.

// func (r receiver) identifier (p paramter) (return (s)) { <code> }

func foo() {
	fmt.Println("I am foo")
}

func bar() {
	fmt.Println("I am bar")
}

func main() {

	// when using defer
	// we basically tell the
	// program DO NOT RUN THIS
	// untill the outer sorounding
	// function / code has executed

	defer foo()
	bar()

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
