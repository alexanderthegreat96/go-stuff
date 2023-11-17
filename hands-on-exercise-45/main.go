package main

import "fmt"

func main() {
	// start with this slice
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	// append to that slice the value 52
	x = append(x, 52)

	// print out the slice
	fmt.Println("This is our slice with the value 52 appended: ", x)

	// in one statement append to that slice these values
	// 53, 54, 55

	x = append(x, 53, 54, 55)

	// print the slice
	fmt.Println("This is our slice with the values 53, 54, 55 appended: ", x)

	// append to the slice this slice
	// y:= []int{56,57,58,59,60}

	y := []int{56, 57, 58, 59, 60}

	// anoher method is simply unpacking the slice
	// and appending the values slice...
	// using the 3 dots

	x = append(x, y...)

	fmt.Println("Copied y's values into x which results in:", x)

	// for this we need to create a new slice
	master_slice := [][]int{}
	master_slice = append(master_slice, x, y)

	fmt.Printf("Created a new master slice containing both slices: %v and %v resulting in %v \n", x, y, master_slice)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
