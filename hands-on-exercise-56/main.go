package main

import "fmt"

func main() {

	// annonymous structs

	myself := struct {
		first     string
		last      string
		friends   map[string]int
		favDrinks []string
	}{
		first: "Alex",
		last:  "Popescu",
		friends: map[string]int{
			"marius": 33,
			"florin": 65,
			"tom":    12,
		},
		favDrinks: []string{"cola", "pepsi"},
	}

	fmt.Printf("This is myself: %s, %s and I like: %v\nThese are my friends: %v\n", myself.first, myself.last, myself.favDrinks, myself.friends)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
