package main

import "fmt"

func main() {
	// to delete from a slice, we use APPEND along with SLICING

	// start with this slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println("x contains:", x)

	// use APPEND and SLICING to get these values here
	x = append(x[:4], x[6:]...)
	fmt.Println("Modified the slice of x to contain values untill index 4 and after index 6:", x)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
