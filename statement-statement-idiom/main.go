package main

import (
	"fmt"
	"math/rand"
)

// statement statement idiom
// basically, have multiple statements on the same line

func main() {
	fmt.Println("Initial statement upon loading the program...")

	x := 44

	y := 5

	fmt.Printf("x=%v \n y=%v\n", x, y)

	// on the same line
	// what happens here
	// we first check if the z is assigned the value
	// by the rand.Intn function
	// then if that is true
	// we check the second condition
	// sometimes it's useful and easier to write certain
	// things like this
	// also, scope of z is included in that line
	// and only in that line
	// limits it's scope

	if z := 2 * rand.Intn(x); z >= x {
		fmt.Printf("z is %v and that is GREATER THAT OR EQUAL to x which is %v\n", z, x)
	} else {
		fmt.Printf("z is %v and that is LESS THAN x which is %v\n", z, x)
	}

	// comma ok idiom
	// In Go (often referred to as Golang),
	// the "OK idiom" typically refers to a common
	// pattern used when working with maps.
	// It is used to determine whether a value exists in a map
	// and whether that value is actually present or not.
	// The "OK idiom" involves using a second variable
	// to check if a value exists in a map and whether it is valid or not.

	arr := make(map[string]int)

	arr["age"] = 15

	age_value, ok := arr["age"]
	if ok {
		fmt.Printf("Found a value for age in the map above, using the OK statement and it is %v\n", age_value)
	} else {
		fmt.Println("No key found")
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
