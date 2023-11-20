package main

import "fmt"

type Person struct {
	FirstName                string
	LastName                 string
	FavoriteIceCreamFlavours []string
}

func main() {
	// struct with slices
	// create a new type called person
	// create 2 values
	// store them in a slice
	// loop over it

	var people []Person

	people = append(people, Person{
		FirstName:                "Mathew",
		LastName:                 "Simson",
		FavoriteIceCreamFlavours: []string{"vanilla", "cococa", "caramel"},
	})

	people = append(people, Person{
		FirstName:                "Andew",
		LastName:                 "Dumas",
		FavoriteIceCreamFlavours: []string{"cocoa", "vanilla"},
	})

	for _, person := range people {
		fmt.Printf("%s %s likes:\n", person.FirstName, person.LastName)
		for _, flavor := range person.FavoriteIceCreamFlavours {
			fmt.Println("- " + flavor)
		}
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
