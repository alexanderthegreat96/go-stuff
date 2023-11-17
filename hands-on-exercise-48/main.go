package main

import "fmt"

func main() {
	// create a slice of a slice of a string

	first_names := []string{"Alex", "Mike"}
	second_names := []string{"Andrew", "Mathew"}

	fmt.Printf("First slice: %v, Second Slice: %v\n", first_names, second_names)

	names := [][]string{first_names, second_names}

	fmt.Printf("Both slices: %v\n", names)

	// iterate through the slices
	// iterate through all records of that slice

	fmt.Println("Iterating through both slices")

	for _, arr := range names {
		for _, name := range arr {
			fmt.Println(name)
		}
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
