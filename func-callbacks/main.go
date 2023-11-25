package main

import "fmt"

// callbacks
// providing a function as an argument
// to a function

// a callback is passing in a function as an argument

// adds to numbers, basic shit
func add(x int, y int) int {
	return x + y
}

// substracts 2 values
func subtract(x int, y int) int {
	return x - y
}

// as you can see
// the third parameter
// f uses the type func
// this func, uses 2 or could use more, or even none
// inputs int int, are arg types, then has a return type if int afterwatdfs
// func (arg1_type, arg2_type) return_type
// func (string, string) string
// func ([]string, int) map[string]int
// whatever the case may be

// important, the function signature argument
// MUST MATCH the function you are going to provide it
// ex: f func(int, int) int will not work with a function
// that is myfunc(name string, name string) string
// different types for both return and arguments

func doMath(x int, y int, f func(int, int) int) int {
	return f(x, y)
}

func main() {

	fmt.Printf("doMath type: %T\n", doMath)
	fmt.Printf("add type: %T\n", add)
	fmt.Printf("subtract type: %T\n", subtract)

	x := doMath(2, 4, add)      // called the function as a parameter
	y := doMath(4, 3, subtract) // called the function as a parameter

	fmt.Printf("%d\n%d\n", x, y)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
