package main

// An expression is a combination of values, variables, operators and function calls
// that are evaluated to produce a single value. It can be as simple as a literal value
// or a variable, or it can invole complex operations and function invocations

import "fmt"

func foo() {
	fmt.Println("Hello world!")
}

// just a regular function
func regularFunction() {
	fmt.Println("I am a regular function")
}

// will take 1 input of type func and run it
func runFuncInThisOne(f func()) {
	fmt.Println("I am about to run the f function")
	f()
}

// will take 2 inputs
// of type func
// and run them
func runTwoFunctions(f func(), e func()) {
	fmt.Println("Running the f and the e functions")
	defer f() // run after e
	e()
}

func main() {
	foo()

	// Expressions are the building blocks in Go programes, and they
	// are used to perform computations, manipulate values, and make
	// decissions based on conditions. The are assigned to variables,
	// passed as arguments to functions, or used in control flow statements
	// like if statements and loops.

	// The term "first-class citizen" reffers to the status of certain
	// entitires, such as values, types and functions, that are treated
	// equally and have the same capabilitiesas the other entities in the
	// language. It means that these entities can be assigned to variables,
	// passed as arguments to functions, and returned as values from functions,
	// just like any other data type in the language.

	// In Go, functions are considered "first-class citizens". They can be assigned
	// to varialbles, passed as arguments to other functions, and returned as vlaues from
	// functions. This allows Go to support higher-order functions, where functions can
	// operate on other functions. For example, you can defined a function that takes another
	// function as an argument and then call that functiuon withing the body of the function,

	// function expression
	ex := func() {
		fmt.Println("This is a function assigned to a variable")
	}

	ed := func() {
		fmt.Println("Another function assigned to a variable ")
	}

	ex() // notice how the variable becomes a function, this is because the variable inherit's the function's type
	ed()

	// running a function within a function
	runFuncInThisOne(ex)

	// running 2 functions within a function
	// called the function directly without ()
	runTwoFunctions(regularFunction, ed)

	// Long story short, a function expression is basically
	// creating a variable and assigning the func keyword to it

	alex := func() {
		fmt.Println("This is the Alex function expression")
	}

	alex()

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
