package main

import "fmt"

func main() {

	// use the slice code from hands-on-exercise-43
	my_slice := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// slice and print 42, 43, 44, 45, 46

	fmt.Println("First slice: ", my_slice[:5])

	// slice and print 47, 48, 49, 50, 51

	fmt.Println("Second slice: ", my_slice[5:])

	// slice and print 43, 44, 45, 46, 47

	fmt.Println("Third slice: ", my_slice[1:6])

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
