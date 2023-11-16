package main

import "fmt"

func main() {

	x := map[string]int{
		"James":  43,
		"Mathew": 22,
		"Alex":   28,
	}

	fmt.Println(x["James"])

	if key_exists, ok := x["Q"]; ok {
		fmt.Printf("Found a value for %d in map", key_exists)
	} else {
		fmt.Println("The specified key does not exist.")
	}

	// second version, check if it is NOT OK

	if key_found, ok := x["Q"]; !ok {
		fmt.Println("No value found for this int. ", key_found)
	}

	fmt.Println("Press ENTER to exit the program...")
	fmt.Scanln()

}
