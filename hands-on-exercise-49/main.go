package main

import (
	"fmt"

	"github.com/alexanderthegreat96/go-stuff/package-helpers/helpers"
)

func main() {
	// working with maps

	// Create a map with a key of TYPE string which is a person's "last Name",
	// and a value of type []string which stores their favorite things
	// Store the records in your map
	// Print out all of the values, along with theis index position in a slice

	prefferences := map[string][]string{
		"bond_james":       {"shaken", "not stirred", "maritinis", "fast cars"},
		"moneypenny_penny": {"intelligence", "literature", "computer science"},
		"no_dr":            {"cats", "ice cream", "sunsets"},
		"no_values":        {},
	}

	fmt.Printf("Prefferences map: %#v\n", prefferences)

	fmt.Println("Starting ranging: ")

	for name, personal := range prefferences {
		fmt.Printf("%s prefferences:\n", name)

		if len(personal) > 0 {
			for key, value := range personal {
				fmt.Printf(" - %d - %s\n", key, value)
			}
		} else {
			fmt.Println(" - No values")
		}
	}

	// let's start generating a map that contains slices
	// using a new function I just made

	// 200 elements

	large_generated_map := make(map[string][]string)
	for i := 1; i <= 200; i++ {

		word, _, _ := helpers.ReturnRandomWords()
		_, words, _ := helpers.ReturnRandomWords(8)

		large_generated_map[word] = words
	}

	fmt.Printf("This is our map with a bunch of values we generated: %#v", large_generated_map)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
