package main

import "fmt"

func main() {

	// In Go, the make function is used for creating and initializing slices, maps, and channels.
	// It allocates and initializes the underlying data structure and returns
	// a value of the specified type (slice, map, or channel).
	// The make function has different signatures depending on the type of data structure you want to create.

	// composite slice -> slice literal
	slice_example := []string{"A", "B", "C", "D", "E", "F"}
	fmt.Printf("This is a regular slice: %v, which contains %d elements\n", slice_example, len(slice_example))

	//slice by using the make function

	// slice is built on top of an array
	// declare []slice_type, length, capacity
	// it's useful when dealing with performance boosts
	// in other words, instead of letting the runtime do the check
	// and to keep increasing the slice size
	// we simply state it beforehand
	// this way, the runtime doesnt need to do anything

	make_slice := make([]int, 0, 10)

	//appending values for the 10 spaces created
	for i := 0; i < 10; i++ {
		make_slice = append(make_slice, i)
	}

	// yes, we can add more than the limit
	// the runtime will create a new slice
	// and keep doing this on repeat
	// this has a performance cost
	// hence why we use make([]int, 0, 10) <- metered slice
	// and not []int{} <- unmetered slice

	fmt.Printf("Slice made using make: %v, which contains %d elements and a capacity of %d \n", make_slice, len(make_slice), cap(make_slice))

}
