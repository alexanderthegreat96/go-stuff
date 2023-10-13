package main

import "fmt"

func main() {

	// sequential

	var larger_number int = 44
	var smaller_number int = 22

	// conditional

	if larger_number > smaller_number {
		fmt.Printf("%v is larger than %v \n", larger_number, smaller_number)
	}

	var execute bool = true

	if execute {
		fmt.Println("Will execute!")
	}

	var x int = 55
	var y int = 13
	var z float64 = 67.4

	// performed a little type conversion
	// operators
	// can only be used when things have the same type
	// as them

	if x > y && x < int(z) {
		fmt.Printf("%v is larger than %v, but smaller than %v \n", x, y, z)
	}

	// conditionals with else if as well

	if x > int(z) {
		fmt.Printf("%v is larger than %v \n", x, z)
	} else if int(z) > y {
		fmt.Printf("%v is larger than %v \n", x, z)
	} else {
		fmt.Println("Nothing to display here")
	}

	fmt.Println("Please press ENTER to exist!")
	fmt.Scanln()
}
