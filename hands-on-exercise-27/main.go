package main

import "fmt"

// create for loop that uses a break statement
// increment or decrement the variable being checked
// as a condition in the loop curriculum item

func main() {

	// for the purpose of this, i will create an infinite loop first
	// then once a condition is met
	// i will stop it

	var x int
	for {
		fmt.Printf("Iteration is at number: %d \n", x)

		if x == 40 {
			fmt.Printf("Iteration reached number: %d and stopped \n", x)
			break
		}
		x++
	}

	// second iteration
	// basically we keep substracting from 70
	// until it reaches 0 and stops

	fmt.Println("Second iteration here")

	y := 70

	for {
		fmt.Printf("Iteration is at %d \n", y)

		if y <= 0 {
			fmt.Println("Reached number 0 and stopped.")
			break
		}
		y--
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
