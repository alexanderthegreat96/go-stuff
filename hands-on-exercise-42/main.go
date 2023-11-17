package main

import "fmt"

func main() {

	// using a composite literal
	// - create an array which holds 5 values of type int
	// - assign values to each index position
	// - range over the array and print the values out

	var my_array [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("We just declared an array which holds %d elements, here it is: %v\n", len(my_array), my_array)

	my_second_array := [5]int{}

	my_second_array[0] = 1
	my_second_array[1] = 2
	my_second_array[2] = 3
	my_second_array[3] = 4
	my_second_array[4] = 5

	fmt.Printf("We just declared another array which holds %d elements, here it is: %v\n", len(my_second_array), my_second_array)

	for index, value := range my_second_array {
		fmt.Printf("We are at index %d with value %d which is of type %T\n", index, value, value)
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
