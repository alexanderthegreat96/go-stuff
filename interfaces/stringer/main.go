package main

import (
	"fmt"
	"log"
	"strconv"
)

type count int

type book struct {
	title string
}

// usage of String() method will modify the behaviour
// of the custom type using the receiver of the function
// in this case, b book

func (b book) String() string {
	return fmt.Sprint("Book title is: " + b.title)
}

// same as above, only for the c count function
// receiver
func (c count) String() string {
	return fmt.Sprint("The count is: ", strconv.Itoa(int(c)))
}

func main() {
	// showcasing the Stringer interface
	// in this segment of the code

	b := book{
		title: "Lord of the rings",
	}

	var n count = 42

	fmt.Println(b) // printing type book as a string
	fmt.Println(n) // printing type count as a string also

	// Long story short, what ends up happenind is
	// the stringer interface in go, has a type called String()
	// by using this as a signature function, we get to modify the
	// bahaviour of types

	// wrapper for the log function
	// log function will provide addional
	// information

	log.Println(b)
	log.Println(n)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
