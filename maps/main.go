package main

import (
	"fmt"

	"github.com/alexanderthegreat96/go-stuff/package-helpers/helpers"
)

// In the context of programming in Go (also known as Golang),
// "maps" refer to a built-in data type that represents an unordered collection
//  of key-value pairs. In Go, maps are somewhat similar to dictionaries or hash maps
// in other programming languages.

// composite data structures
// maps / dictionaries / associative arrays
// they store key => value pairs

func main() {

	// 2 ways of initializing a map
	// using make function map[key_type]value_type

	var first_map = make(map[string]int)
	first_map["alex"] = 12

	fmt.Printf("First map made with make: %#v\n", first_map)

	// using map literal

	test_map := map[string]int{
		"Bob":   15,
		"Randy": 56,
	}

	test_map["Mathew"] = 15

	fmt.Printf("Second map made with map literal: %#v\n", test_map)

	// let's generate a map with a lot of
	// key => value pairs

	large_map := map[string]int{
		"tatheus": 17,
		"brabus":  99,
		"X5":      53,
	}

	for i := 16; i <= 100; i++ {
		random, error := helpers.RandomStringGenerator(8)

		if error != nil {
			fmt.Println("Error: " + error.Error())
			break
		}

		large_map[random] = i * 2
	}

	fmt.Printf("We got a map with a lot of values. %#v\n", large_map)

	// let's range over a map using for range

	fmt.Println("Ranging over a map:")

	for key, value := range large_map {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}

	// DELETE an element from a map
	// we use the delete built-in function

	delete(large_map, "brabus")

	// in case we use the wrong key, the compiler
	// won't panic and nothing will happen ->
	// delete(large_map, "KeyThatDontExist")

	// !!!!!!!! important

	// in go, if you look for a key in a map and it
	// does not exist
	// the language will return 0
	// it won't crash or panic

	fmt.Printf("Looking for [drake] in the large_map. This will return: %d\n", large_map["drake"])

	// in order to check if a value
	// exists inside a map
	// based on a key
	// we use the commma ok idiom

	searched_value, ok := large_map["something"]

	if ok {
		fmt.Printf("Found value: %d\n", searched_value)
	} else {
		fmt.Println("Cannot find value when using key: something")
	}

	// we can also just run  a check without returning the value
	if _, ok := large_map["alex"]; !ok {
		fmt.Println("Could not find [alex] inside large map")
	}

	// another way, which is also shorter

	if found_value, ok := large_map["X5"]; ok {
		fmt.Printf("Found a value by key: X5 which is %d\n", found_value)
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
