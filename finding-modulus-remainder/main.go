package main

import "fmt"

func main() {
	x := 83 / 40
	y := 83 % 40

	fmt.Printf("Modulus %v, remainder %v\n", x, y)

	// finding even numbers

	for i := 0; i <= 100; i++ {

		if i%2 == 0 {
			fmt.Printf("Number: %d is even \n", i)
		} else {
			fmt.Printf("Number: %d is not even \n", i)
		}

	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
