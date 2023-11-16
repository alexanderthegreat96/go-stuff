package main

import "fmt"

// modulus operator
// to print all even numbers
// use the continue statement
func main() {

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Printf("%d is an odd number \n", i)
			continue
		}
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
