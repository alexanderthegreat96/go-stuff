package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var executionLimit int = 100
	var line int

	for {
		fmt.Printf("We are at loop %d \n", line)

		// what happens here
		// we define a random int for x, not larger than 5
		// we then check on the same line if the value
		// is equal to 3

		if x := rand.Intn(5); x == 3 {
			fmt.Println("x is equal to 3")
		}

		//rather different way to stop a loop

		if line == executionLimit {
			fmt.Println("Execution reached 100, stopping")
			break
		}

		line++
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
