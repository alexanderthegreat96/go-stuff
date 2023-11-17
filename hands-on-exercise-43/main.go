package main

import "fmt"

func main() {
	// Using composite literal
	// - create  SLICE of type INT
	// - assign these 10 values:
	// 42, 43, 44, 45, 46, 47, 48, 49, 50, 51
	// Range over the slice and print the values
	// - print out the VALUE and the TYPE

	my_slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	for index, value := range my_slice {
		fmt.Printf("Index %d with value %d is type %T \n", index, value, value)
	}

	// another way using make which is better
	// for performance

	fmt.Println("Printing a slice made with make")

	my_second_slice := make([]int, 0, 10)
	my_second_slice = append(my_second_slice, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51)

	for kk, vv := range my_second_slice {
		fmt.Printf("Index %d with value %d is type %T \n", kk, vv, vv)
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
