package main

import "fmt"

// categories:
// built int types -> int, float, bool etc
// refference tyoes  -> slices, maps, channels, functions, interfaces
// structs

// concrete type -> something that can be created with this type
// interface type -> something that satisfies an interface. aka abstract type

type Friend struct {
	first string
	age   int
}

// this is an abstract type
// since implementing it
// must satisfy ot's constraints

// type Reader interface {
// 	Reader []byte
// }

func main() {
	// concrete type as we are creating values with them
	friend := Friend{
		"alex",
		22,
	}

	fmt.Println(friend)
}
