package main

import (
	"fmt"
	"math/rand"
)

func randomInteger(n int) int {
	return rand.Intn(n)
}

func main() {
	// took the code from hands-on-exercise-19
	// didn't use the short declaration operator
	// as i wanted to be a little more classic

	var choseNumber int = 500                          // could be written as choseNumber := 60
	var randomInteger int = randomInteger(choseNumber) // could be written as randomInteger := randomInteger(choseNumber)

	fmt.Printf("%d is the nubmer picked up to %d \n", randomInteger, choseNumber)

	fmt.Println("Executed Conditional Ifs:")

	if randomInteger > 0 && randomInteger < 100 {
		fmt.Printf("The value %d is between 0 and 100 \n", randomInteger)
	}
	if randomInteger > 101 && randomInteger < 200 {
		fmt.Printf("The value %d is between 101 and 200 \n", randomInteger)
	}
	if randomInteger > 201 && randomInteger < 250 {
		fmt.Printf("The value %d is between 201 and 250 \n", randomInteger)
	}

	fmt.Println("Executed switch statements: ")

	switch {
	case randomInteger <= 100:
		fmt.Printf("The value %d is between 0 and 100 \n", randomInteger)
	case randomInteger > 101 && randomInteger < 200:
		fmt.Printf("The value %d is between 101 and 200 \n", randomInteger)
	case randomInteger > 201 && randomInteger < 250:
		fmt.Printf("The value %d is between 201 and 250 \n", randomInteger)
	default:
		fmt.Printf("The value %d is highter than 250 \n", randomInteger)
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
