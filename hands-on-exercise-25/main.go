package main

import (
	"fmt"
	"math/rand"
)

// create one random int between 0 inclusive and 5 exclusive
// assign the value to a variable with the identifier x
// use a switch statement to print a description of the variable
// and the value
// run the code 42 times to print the iteration number

func main() {

	for i := 1; i <= 42; i++ {

		fmt.Printf("Running iteration number %d\n", i)

		var x int = rand.Intn(5)
		fmt.Printf("Chose number %d \n", x)

		switch x {
		case 0:
			fmt.Printf("For Iteration of %v, x is %d \n", i, x)
		case 1:
			fmt.Printf("For Iteration of %v, x is %d \n", i, x)
		case 2:
			fmt.Printf("For Iteration of %v, x is %d \n", i, x)
		case 3:
			fmt.Printf("For Iteration of %v, x is %d \n", i, x)
		case 4:
			fmt.Printf("For Iteration of %v, x is %d \n", i, x)
		default:
			fmt.Println("Case is unknown")
		}

	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
