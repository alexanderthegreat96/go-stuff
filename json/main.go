package main

import (
	"encoding/json" // package that we want for handling json operations
	"fmt"
)

// using json
// in order to encode and decode data

// ussually, we need to define a struct for that

type Player struct {
	ID        int
	Username  string
	Platforms []string
}

// related th person
type Friend struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// created this for unmarshaling the json string
type Person struct {
	ID            string   `json:"_id"`
	Index         int      `json:"index"`
	GUID          string   `json:"guid"`
	IsActive      bool     `json:"isActive"`
	Balance       string   `json:"balance"`
	Picture       string   `json:"picture"`
	Age           int      `json:"age"`
	EyeColor      string   `json:"eyeColor"`
	Name          string   `json:"name"`
	Gender        string   `json:"gender"`
	Company       string   `json:"company"`
	Email         string   `json:"email"`
	Phone         string   `json:"phone"`
	Address       string   `json:"address"`
	About         string   `json:"about"`
	Registered    string   `json:"registered"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
	Tags          []string `json:"tags"`
	Friends       []Friend `json:"friends"`
	Greeting      string   `json:"greeting"`
	FavoriteFruit string   `json:"favoriteFruit"`
}

func main() {
	var player Player = Player{
		1,
		"alexander",
		[]string{"uplay", "psn", "xbl"},
	}

	fmt.Println(player)

	// encode data to json

	output, jsonError := json.Marshal(player)

	if jsonError != nil {
		fmt.Println(jsonError.Error())
	}

	// output will return []byte
	fmt.Printf("output is of type %T\n", output)

	// in order to see the output
	// we have to convert it to a string

	fmt.Println(string(output))

	// unmarshaling JSON

	var jsonString string = `
	{
		"_id": "6597cb9ebb5e29d4be1cf709",
		"index": 0,
		"guid": "030e8aeb-032b-4cc7-9603-6ca6aaf7edef",
		"isActive": false,
		"balance": "$2,964.29",
		"picture": "http://placehold.it/32x32",
		"age": 23,
		"eyeColor": "green",
		"name": "Wyatt Taylor",
		"gender": "male",
		"company": "MAGNEATO",
		"email": "wyatttaylor@magneato.com",
		"phone": "+1 (858) 416-2285",
		"address": "171 Sackett Street, Coldiron, Virgin Islands, 3376",
		"about": "Labore anim est incididunt eiusmod quis eu fugiat aute elit quis incididunt incididunt adipisicing cillum. Voluptate exercitation excepteur et consequat magna dolore. Quis anim qui exercitation consectetur. Qui minim tempor et aliqua aliquip sunt. Veniam occaecat ea proident eiusmod sunt sunt laborum exercitation nulla dolor.\r\n",
		"registered": "2018-08-06T07:07:17 -03:00",
		"latitude": -38.921139,
		"longitude": 156.778543,
		"tags": [
		  "cupidatat",
		  "incididunt",
		  "deserunt",
		  "proident",
		  "elit",
		  "qui",
		  "nisi"
		],
		"friends": [
		  {
			"id": 0,
			"name": "Mcguire Workman"
		  },
		  {
			"id": 1,
			"name": "Sharlene Humphrey"
		  },
		  {
			"id": 2,
			"name": "Mcdonald Morton"
		  }
		],
		"greeting": "Hello, Wyatt Taylor! You have 8 unread messages.",
		"favoriteFruit": "strawberry"
	  }
	`

	// so we essentially take the json string
	// convert it to a slice of byte
	// then  we need a struct to accomodate the output
	var jsonBytes []byte = []byte(jsonString)
	// used the struct for Perosn
	// it is important to know that the structure
	// must match the JSON

	var person Person = Person{}

	// provide the json bytes
	// and refference to the person variable
	decodeJsonErr := json.Unmarshal(jsonBytes, &person)

	if decodeJsonErr != nil {
		fmt.Println(decodeJsonErr.Error())
	} else {
		// person var that holds the struct
		// was populated by the json
		fmt.Println(person)
	}

}
