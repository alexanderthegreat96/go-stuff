package main

import "fmt"

// func (receiver) identifier(parameters) (returns) {code}
// parameters & arguments
// - we define our func with parameters (if any)
// - we call our func and pass in arguments (if any)
// EVERYTHING in Go is PASS by VALUE
// sample functions:
// - no params, no returns
// - 1 param, no returns
// - 1 param, 1 return
// - 2 params, 2 returns

// func (r receiver) identifier(p parameters) (return(s)) {<code>}

// no params, no returns
func foo() {
	fmt.Println("bar")
}

// 1 param, no returns
func thisFunc(s string) {
	fmt.Println("Hello, " + s)
}

// 1 param 1 return
func myFunc(s string) string {
	return "Hello, " + s
}

// 2 params 2 returns
func hello(name string, age int) (string, int) {
	return name, age
}

// variatic parameter
// unlimitted number of params

// func (r receiver) identifier(p ...)

func sumNumbers(numbers ...int) int {

	// converts numbers to []int
	// creates a slice of all of them

	result := 0
	if len(numbers) > 0 {
		for _, number := range numbers {
			result += number
		}
	}

	fmt.Println(numbers)
	fmt.Printf("Type of: %T\n", numbers)

	return result

}

// the variatic parameter needs to be the FINAL parameter
// notice how the variatic parameter
// is the final parameter

func sumProducts(productName string, prices ...int) (string, int) {
	productTotal := 0

	if len(prices) > 0 {
		for _, price := range prices {
			productTotal += price
		}
	}

	return productName, productTotal
}

func main() {

	// sample function calls

	foo()
	thisFunc("Alex")

	myFunc := myFunc("Alex")
	fmt.Println(myFunc)

	name, age := hello("Alex", 28)

	fmt.Printf("Hello, %s, you are %d years old.\n", name, age)

	// variadic parameters
	// aka, pass unlimitted arguments to a function
	// we use it by specifying a param name followed by ...TYPE
	// ex: ages...int
	// ex: names...string
	// it's worth noting that the variatic parameter can be invoke with 0 arguments
	// in other words, you can pass nothing
	// it can be used for optional arguments

	fmt.Printf("We summed the numbers: %d\n", sumNumbers(2, 3, 4, 6))
	productName, productTotal := sumProducts("Macbook Pro 2023 M2 16in", 2344, 2344, 3100)
	fmt.Printf("The %s based on it's quantity, costs $%d\n", productName, productTotal)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
