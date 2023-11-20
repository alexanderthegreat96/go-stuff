package main

import (
	"fmt"
)

type Person struct {
	FirstName                string
	LastName                 string
	FavoriteIceCreamFlavours []string
}

func main() {
	// map struct
	// taking the values from the
	// previous exercise
	// store them as the type person
	// in a map
	// where the key it's the person's last name

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

	peopleMap := make(map[string]Person)

	// we're gonna do it automatically cau'se we are badass

	for _, value := range people {
		peopleMap[value.LastName] = Person{
			FirstName:                value.FirstName,
			LastName:                 value.LastName,
			FavoriteIceCreamFlavours: value.FavoriteIceCreamFlavours,
		}
	}

	for lastName, values := range peopleMap {
		fmt.Println(lastName)

		fmt.Printf("%s %s likes:\n", values.FirstName, values.LastName)
		for _, flavor := range values.FavoriteIceCreamFlavours {
			fmt.Println("- " + flavor)
		}
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
