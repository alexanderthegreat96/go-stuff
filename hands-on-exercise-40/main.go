package main

import "fmt"

func main() {

	// playing with some arrays

	// use an array literal to store elements
	// in an array and let the compiler decide
	// how many elements we are going to use in that array
	// finally, print out the length of the array

	my_large_array := [...]string{"hello", "i am", "the greatest", "of all", "time", "my", "name is", "Muhhamad Ali"}

	fmt.Printf("The array my_large_array's output is: %v \nand has a length of: %d elements \n", my_large_array, len(my_large_array))

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
