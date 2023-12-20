package main

import "fmt"

// function expressions && func returns
func main() {
	fmt.Println("Showcasing function expressions, aka assigning a function to a variable and calling it through the variable.")
	my_func := func(name string) {
		fmt.Println("This is my name: " + name)
	}

	my_func("Alex")

	fmt.Println("Showcasing a return function. Similar to closures.")

	ThisReturnsZero := generateFunction() // call the function that returns a function

	var function_returned int = ThisReturnsZero() // the assigned variable will be called as a function

	fmt.Printf("This function was returned by generateFunction() and it is: %d \n", function_returned)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}

func generateFunction() func() int {
	return func() int {
		return 0
	}
}
