package main

import (
	"encoding/json"
	"fmt"
)

// marshaling or encoding data to json

// define a struct
// notice how everything beggins with
// uppercase lettering
// this is because they need to be "public / exporable"
// for json to work

type Person struct {
	First string
	Last  string
	Age   int
	Zip   int
}

func main() {

	person := Person{
		First: "Alexandru",
		Last:  "Popescu",
		Age:   28,
		Zip:   63663,
	}

	person2 := Person{
		First: "Matei",
		Last:  "Popescu",
		Age:   28,
		Zip:   63663,
	}

	// slice of type Person
	people := []Person{person, person2}

	// encoded will return a slice of bytes
	// so it will need to be converted to string
	// it is useful since, you can write bytes directly
	// to a file

	encoded, encodedErr := json.Marshal(people)

	if encodedErr != nil {
		fmt.Println(encodedErr.Error())
	} else {
		encodedString := string(encoded)
		fmt.Println(encodedString)
	}

}
