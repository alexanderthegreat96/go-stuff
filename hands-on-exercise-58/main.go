package main

import (
	"fmt"
	"math/rand"

	"github.com/alexanderthegreat96/go-stuff/package-helpers/helpers"
)

// basic functions
// tasks:
// - create a func with the identifier foo that returns an int
// - create a func with the identifier bar that returns an int and a string
// - call both funcs
// - print out their results

func main() {
	fmt.Printf("Returned random int: %d\n", giveInt())

	number, name := giveNameAndInt() // assiged the returns as variables

	fmt.Printf("Returned random int: %d and string %s\n", number, name)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}

func giveInt() int {
	return rand.Intn(99)
}

func giveNameAndInt() (number int, name string) {
	number = rand.Intn(99)
	name, _, _ = helpers.ReturnRandomWords()

	return number, name

}
