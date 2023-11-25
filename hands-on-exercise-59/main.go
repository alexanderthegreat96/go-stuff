package main

import "fmt"

// variadic functions
// tasks:
// - create a fuynction with the identifier foo that:
// 	- takes a variadict parameter of type int // variadic -> unlimitted params
//  - pass a value of type []int into the func -> unfurl the slice
//  - return the sum of all values of type INT passed in
// - create a fuynct with the identifier bar that
//  - takes a parameter of type []int
//  - returns the sum of all values of type int passed in

func main() {
	numbers := []int{12, 34, 56, 78, 56, 345, 23}

	fmt.Printf("The sum of the unfurled slice %v is %d\n", numbers, foo(numbers...))
	fmt.Printf("The sum of the  slice %v is %d\n", numbers, bar(numbers))

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}

// variadic will convert to slice
func foo(ints ...int) (result int) {
	result = 0

	for _, number := range ints {
		result += number
	}

	return result
}

// non variadic, expects a slice
func bar(numbers []int) int {
	result := 0

	for _, number := range numbers {
		result += number
	}

	return result
}
