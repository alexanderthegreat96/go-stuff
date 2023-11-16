package main

import "fmt"

func main() {

	// below is the code to create a data structure called a slice of ints

	x := []int{42, 43, 44, 45, 46, 47}

	// loop through the list and print the index and value
	// range loop

	for index, number := range x {
		fmt.Printf("Value is %d and its index is: %d \n", number, index)
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
