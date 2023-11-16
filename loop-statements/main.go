package main

import "fmt"

func main() {
	// loop statements
	// you know what this is

	//basic loop

	for i := 0; i < 40; i++ {
		fmt.Printf("Printing message on line %v\n", i)
	}

	// looping over key value pairs

	my_map := make(map[string]string)

	my_map["username"] = "alexanderdth"
	my_map["password"] = "randompass12312"
	my_map["first_name"] = "Alexandru"

	for key, value := range my_map {
		fmt.Printf("%v:%v\n", key, value)
	}

	// similar to while

	y := 17

	for y < 15 {
		fmt.Printf("Counting from %v untill 15\n", y)
	}

	// takes you out of the loop

	for {
		fmt.Printf("%v \t\t third for the loop\n", y)
		if y > 20 {
			break
		}
		y++
	}

	//continue till the next iteration
	// even numbers only

	for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}

		fmt.Printf("Counting even numbers: %v\n", i)
	}

	// nested loops
	// long story short
	// first loop runs once
	// second loop runs 5 times
	// and this keeps going

	for i := 0; i < 20; i++ {
		fmt.Println("--")

		for j := 0; j < 5; j++ {
			fmt.Printf("oute loop %v \t inner loop %v\n", i, j)
		}
	}

	// for range loops

	// iterates over data structures
	// arrays, maps / dicts , etc

	numbers := []int{2, 3, 4, 5, 6, 7, 8}

	for i, v := range numbers {
		fmt.Printf("Lopping over slike at key %v and value %v\n", i, v)
	}

	// maps are the equivalent of dictionaries in python
	// Dictionaries are often also called maps, hashmaps, lookup tables, or associative arrays.

	map_data := map[string]int{
		"age":    23,
		"height": 185,
	}

	for k, v := range map_data {
		fmt.Printf("We are at key: %v with value: %v\n", k, v)
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
