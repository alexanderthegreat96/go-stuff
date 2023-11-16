package main

import "fmt"

func main() {
	// use a for range loop to iterate
	// through the map below and print it's key and value

	x := map[string]int{
		"James":  43,
		"Mathew": 22,
		"Alex":   28,
	}

	for key, value := range x {
		fmt.Printf("%s is %d \n", key, value)
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
