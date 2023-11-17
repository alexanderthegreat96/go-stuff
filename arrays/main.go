package main

import "fmt"

func main() {
	// arrays are inflexible objects
	// meaning that their length needs
	// to be declared first
	// it's not recommended to be used
	// you might see them sometimes
	// slices are a better option

	// one way of initializing them
	var my_array [6]int = [6]int{1, 3, 4, 5, 6, 7}

	fmt.Println("Got an array of 6 elements here: ", my_array)

	// another way of doing the same thing

	var example_array [3]int

	// when working with something like this
	// we append data via KEY

	example_array[0] = 14
	example_array[1] = 23
	example_array[2] = 25

	fmt.Println("We got another array of 3 elements: ", example_array)

	// the easiest way to initialize them

	easy_array := [4]int{9, 8, 7}

	fmt.Println("Easiest initialization: ", easy_array)

	// lets do a string variant
	// we can use 3 dots to basically
	// let the compiler figure out
	// how many elements we want

	string_array := [...]string{"Hello", "World"}

	fmt.Println("String array: ", string_array)

	string_array_manual := [3]string{"Whats", "going", "on?"}

	fmt.Println("String array manual: ", string_array_manual)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
