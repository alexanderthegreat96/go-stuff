package main

import "fmt"

// In Go, a closure is a function value that references variables from outside its body.
// The function may close over (capture) and store references to the variables from its surrounding lexical scope,
// allowing it to access those variables even after they are no longer in scope.
// Closures are a powerful and flexible feature of the Go programming language.

func main() {
	incrementor := incrementor()

	fmt.Printf("Calling incrementor returned function: %d\n", incrementor()) // 1
	fmt.Printf("Calling incrementor returned function: %d\n", incrementor()) // 2
	fmt.Printf("Calling incrementor returned function: %d\n", incrementor()) // 3
	fmt.Printf("Calling incrementor returned function: %d\n", incrementor()) // 4
	fmt.Printf("Calling incrementor returned function: %d\n", incrementor()) // 5
	fmt.Printf("Calling incrementor returned function: %d\n", incrementor()) // 6

	// it increments
	// because we call the parent function onces
	// but then we call the inner or the child function
	// multiple times. since the "x" variable is declared
	// outside the inner function
	// the inner function can keep adding
	// stuff to it

	fmt.Println("PRESS ENTER to exit...")
	fmt.Scanln()
}

// incrementor is closing over another functiuon
// and the function inside is inclosed in it's
// "parent" function

// an outer func
// is enclosing over
// an inner func

func incrementor() func() int {
	// it's useful because
	// we can initialize a variable
	// in the outer func
	// and use it to do something
	// with the inner func
	x := 0

	return func() int {
		x++
		return x // we call x which is still in scope due to being encapsulated into a parent function
	}
}
