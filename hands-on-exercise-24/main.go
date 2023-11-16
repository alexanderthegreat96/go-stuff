package main

import (
	"fmt"
	"math/rand"
)

// random numbers
func pickTwoRandomNumbers(first int, second int) (int, int) {
	return rand.Intn(first), rand.Intn(second)
}

// altered the structure of the program to make it cleaner

func computeNumbers() {
	x, y := pickTwoRandomNumbers(10, 10)
	fmt.Printf("Picked numbers: %d and %d \n", x, y)

	// using switch statements for this kind of stuff

	switch {
	case x < 4 && y < 4:
		fmt.Printf("%d and %d are smaller than 4 \n", x, y)
		fallthrough
	case x > 6 && y > 6:
		fmt.Printf("%d and %d are larger than 6 \n", x, y)
		fallthrough
	case x >= 4 && x <= 6:
		fmt.Printf("%d is greater or equal than 4 but smaller than 6 \n", x)
		fallthrough
	case y != 5:
		fmt.Printf("%d is not equal to 5 \n", y)
		fallthrough
	default:
		fmt.Println("None of the cases from above match.")
	}
}

// create a program that has a loop that prins every number from 0 to 99

func main() {

	// initial part of the program

	for i := 0; i < 100; i++ {
		fmt.Printf("This is number %d out of 100\n", i)
	}

	// modified to run 100 times
	// and added code from the previous exercise

	for e := 1; e <= 100; e++ {
		fmt.Printf("Executing iteration number: %d \n", e)

		// it will call the function 100 times
		// picking the random numbers between 0 and 9
		computeNumbers()
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
