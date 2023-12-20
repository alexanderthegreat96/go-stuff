package main

import "fmt"

// methods
// create a user defined struct with
// the identifier as person with
// fields: first, age
// attach a method to type person
// the identifier should be speak
// the method should have the person say their name
// and age

// create a value of type person
// call the method

func main() {

	// defined data for the custom type
	dude := person{
		first: "Alex",
		age:   27,
	}

	// invoked the function on the custom type
	dude.speak()

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}

// defined the custom type
type person struct {
	first string
	age   int
}

// provided the type to the receiver
func (p person) speak() {
	fmt.Printf("Hello, I am %s and I am %d years old.\n", p.first, p.age)
}
