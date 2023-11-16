package main

import (
	"fmt"
	"math/rand"
)

// continuation of hands-on-exercise-23
// create 2 random ints between 0 inclusive and 10 exclusive
// in other words, the number picked will be from 0 to 9

func pickTwoRandomNumbers(first int, second int) (int, int) {
	return rand.Intn(first), rand.Intn(second)
}

func main() {
	x, y := pickTwoRandomNumbers(10, 10)
	fmt.Printf("Picked numbers: %d and %d \n", x, y)

	// using switch statements for this kind of stuff

	switch {
	case x < 4 && y < 4:
		fmt.Printf("%d and %d are smaller than 4 \n", x, y)
	case x > 6 && y > 6:
		fmt.Printf("%d and %d are larger than 6 \n", x, y)
	case x >= 4 && x <= 6:
		fmt.Printf("%d is greater or equal than 4 but smaller than 6 \n", x)
	case y != 5:
		fmt.Printf("%d is not equal to 5 \n", y)
	default:
		fmt.Println("None of the cases from above match.")
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
