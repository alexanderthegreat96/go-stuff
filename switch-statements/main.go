package main

import "fmt"

func main() {
	// let's deal with switch statements

	x := 33
	y := 100
	z := 111

	// conditionals
	switch {
	case x > 33:
		fmt.Println("x is larger than 33")
	case x <= 33:
		fmt.Println("x is smaller or equal to 33")
	default:
		fmt.Println("x has a default value of 33")
	}

	// booleans
	// cased value

	switch y {
	case 100:
		fmt.Println("y is equal to 100")
	case 140:
		fmt.Println("y is equal to 140")
	default:
		fmt.Println("y has a default value of 100")
	}

	// example with fallthrough
	// meaning that whatever condition si true
	// move to the next one

	switch z {
	case 111:
		fmt.Println("z is equal to 111")
		fallthrough
	case 120:
		fmt.Println("z is euqla to 120 cause the first condition is true and fallsthrough")
	default:
		fmt.Println("z default value is 111")
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
