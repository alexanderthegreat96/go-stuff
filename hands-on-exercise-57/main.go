package main

import "fmt"

// play around with named returns
func main() {
	numbers := []int{45, 66, 55, 33, 223, 121, 443}

	fmt.Printf("The sum of the numbers %v is %d\n", numbers, sum(numbers))                         // given slice of numbers
	fmt.Printf("The sum of my given numbers 1,44,33,22,12 is %d\n", sumNumbers(1, 44, 33, 22, 12)) // unlimitted args
	fmt.Printf("The sum of the unfurled numbers %v is %d\n", numbers, sumNumbers(numbers...))      // given an unfurled slice

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}

// named returns -> define total as a named return
func sum(numbers []int) (total int) {
	total = 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// named returns with n ulimitted args
func sumNumbers(numbers ...int) (total int) {
	total = 0

	for _, number := range numbers {
		total += number
	}

	return total
}
