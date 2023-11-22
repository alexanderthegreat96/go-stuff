package main

import "fmt"

// function name with an identifier
// func (r receiver) identifier (p parameter) (r returns) {<code>}

func foo() {
	fmt.Println("Hello world!")
}

func main() {
	foo()

	// annonymous function

	// func (p parameter) (r returns) {<code>} (arguments)
	// func () {code} ()
	func() {
		fmt.Println("Annonymous function")
	}()

	func(name string) {
		fmt.Println("This is my name: " + name)
	}("Alexander") // arguments for the parameters

	// essentially, an annonymous function is a
	// function you declare within your code
	// that doesnt have an identifier
	// it can have parameters and arguments

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
