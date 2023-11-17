package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

// generates a random string
func randomStringGenerator(len int) (string, error) {
	// get some random bytes

	randomBytes := make([]byte, len)
	_, err := rand.Read(randomBytes)

	if err != nil {
		return "", err
	}

	// encode the random bytes
	// to string
	randomString := base64.URLEncoding.EncodeToString(randomBytes)
	return randomString[:len], nil
}

func main() {
	// slices are kind of like arrays
	// but they get initialized without
	// specifying the number of elements

	// slices are built on top of an array
	// it points towards it

	slices := []string{}

	// in order to add data to a slice
	// use append keyword

	slices = append(slices, "Alexander")
	slices = append(slices, "Andrew")
	slices = append(slices, "Mathew")

	for i := 1; i < 15; i++ {
		randomString, err := randomStringGenerator(7)

		if err != nil {
			fmt.Println("Error generating random string. Error: ", err.Error())
			break
		}

		slices = append(slices, randomString)
	}

	fmt.Printf("Here's the slice with the values appended to it: %v\n", slices)

	// slice with values already pre-added

	sliceWithValuesPreAdded := []string{"Hopkinks", "Anthony", "Joshua"}

	fmt.Printf("Here's another slice with the values pre-added: %v\n", sliceWithValuesPreAdded)

	// let's make a slice with a bunch of numbers in it
	numeric_slice := []int{23, 123, 445, 342, 121, 434, 323}

	for n := 321; n < 500; n++ {
		numeric_slice = append(numeric_slice, n)
	}

	fmt.Printf("We are appending some values starting from 321 within the slice and this is the result: %v \n", numeric_slice)

	// slicing a slice aka grabbing some parts from the slice

	// similar to python's lists
	// inclusive [starting_point:ending_point]
	// get everything from start to end
	// sort of like between

	fmt.Printf("Sliced the slice and took values from index 0 to index 5: %v \n", numeric_slice[0:5])

	// exlusive
	// [:end] -> get up until the end
	fmt.Printf("Sliced the slice and took values UNTILL index 3: %v\n", numeric_slice[:3])

	// inclusive
	// [start:] -> gets everything starting from that starting point

	fmt.Printf("Sliced the slice and took values FROM index 3: %v \n", numeric_slice[3:])

	// deleting from a slice

	// deleting is done using the append function and the inclusive or exclusive operators
	// so i took the elements from 0-8 from that slice and then made sure it won't go beyond 10
	// this way, the slice has been reduced

	numeric_slice = append(numeric_slice[:8], numeric_slice[:10]...)

	fmt.Printf("Showcasing deleting slice elements: %v \n", numeric_slice)

	// Multidimensional slices

	// basically, holds multiple slices
	// sort of like python's list contains other lists
	// or dictionaries

	first_slice := []string{"Alex", "Maria", "Car", "House", "Games"}
	second_slice := []string{"Automatic", "Transmission", "BMW", "Volvo"}

	main_slice := [][]string{first_slice, second_slice}

	fmt.Printf("First slice : %v\nSecond Slice: %v\nMain Slice : %v \n", first_slice, second_slice, main_slice)

	// ways to copy data from slice to slice

	// created an empty slice with the capacity
	// equal to the length of the copied slice

	new_slice := make([]string, len(first_slice))
	copy(new_slice, first_slice)

	fmt.Printf("Copied %v to copy_slice and now copy_slice contains: %v \n", first_slice, new_slice)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
