package main

import "fmt"

func main() {

	// if you wanna make this run indetinately
	// do not set a value for x
	// var x int
	// set x == x and it will run forever

	//var x int

	// for x == x {
	// 	fmt.Printf("Chose number %d \n", x)
	// 	x--
	// }

	x := 30

	// simply substract numbers
	for x > 0 {
		fmt.Printf("Chose number %d \n", x)
		x--
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
