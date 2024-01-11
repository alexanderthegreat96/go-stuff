package main

import (
	"encoding/json"
	"fmt"
)

// unmarshaling a json string
// aka decoding a json string
// and linking it to a struct

type Person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
	Zip   int    `json:"Zip"`
}

func main() {
	// declared the intial string
	jsonString := `[{"First":"Alexandru","Last":"Popescu","Age":28,"Zip":63663},{"First":"Matei","Last":"Popescu","Age":28,"Zip":63663}]`

	// convert the string to a slice of bytes
	jsonBytes := []byte(jsonString)

	// declare the variable that will hold the decoded data
	// slice of type Person
	// since the json struture is
	// an array of people
	// holding mutliple People

	people := []Person{}

	decodeErr := json.Unmarshal(jsonBytes, &people)

	if decodeErr != nil {
		fmt.Println("We got a problem: " + decodeErr.Error())
	} else {
		for _, person := range people {
			fmt.Printf("Hello, %s %s, you are %d years old\n", person.First, person.Last, person.Age)
		}
	}

}
