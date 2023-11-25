package main

import "fmt"

func foo() int {
	return 28
}

// In Go, you can return functions from functions.
// This is possible because functions are first-class citizens in Go,
// meaning you can treat them like any other value, such as integers or strings.

// looks a bit weird
// but it returns a function
// and that function needs to be
// invoked after assigning bar
// to a variable

func bar() func() int {
	// func keyword + return type
	// it can also have no return type
	// same philosophy as with normal funcs
	return func() int {
		return 28
	}
}

// Function that returns a function
func generateMultiplier(factor int) func(int) int {
	// The returned function takes an int and multiplies it by the factor
	return func(x int) int {
		return x * factor
	}
}

// puts together firstName, lastName and middleName
func namesTogether(firstName string, lastName string) func(string) string {
	return func(middleName string) string {
		return firstName + " " + middleName + " " + lastName
	}
}

func main() {

	// normal function
	// nothing fancy

	x := foo()
	fmt.Println(x)

	// running a function that returns a function
	returnFunction := bar()                 // returns a function
	runReturnedFunction := returnFunction() // calling the returned function
	fmt.Println(runReturnedFunction)        // printing the result

	y := generateMultiplier(6) // generate a function that multiplies by 6

	fmt.Println(y(5)) //Use the generated function and multiply 5 => 30

	z := namesTogether("Smit", "Weston")

	w := z("and")

	fmt.Println(w)

	fmt.Printf("%T \n", y)
	fmt.Printf("%T \n", z)
	fmt.Printf("%T \n", w)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
