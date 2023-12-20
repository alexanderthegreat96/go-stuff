package main

import "fmt"

// annonymous functions exercise
// build an use an annonymous func

func main() {

	func() {
		fmt.Println("This is an annonymous function.")
	}()

	func(name string) {
		fmt.Println("This is an annonymous function with the argument of name set to Alex")
	}("Alex")

	my_func := func(name string) {
		fmt.Println("This is my name: " + name)
	}

	my_func("Alex")

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
