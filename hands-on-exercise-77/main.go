package main

import (
	"fmt"

	"github.com/alexanderthegreat96/go-stuff/package-helpers/helpers"
)

// create 2 functions to change a field ina  struct called "first" of TYPE string
// one function will use value semantics -> this function will return some VALUE of some TYPE
// another function will use pointer semantics -> nothing will be returned

type dude struct {
	first string
}

// this one produces a copy of dude
// hence why it returns dude
func changeFirst(d dude, s string) dude {
	d.first = s
	return d
}

// this one modifies an instance
// hence why it returns nothing
func changeFirstPtr(d *dude, s string) {
	d.first = s
}

func main() {

	// initialize dude

	d := dude{
		first: "seth",
	}

	// print

	fmt.Println(d.first)

	// print the changed value
	// notice, the returned value has a different address than d
	// as the function creates a new copy of dude
	e := changeFirst(d, "alex")

	fmt.Println(e.first)
	fmt.Printf("e's addr: %p, d's addr: %p\n", &e, &d)

	// modify the state of d
	// by setting it's first to some value

	changeFirstPtr(&d, "roger")

	// the reason why we are printint d is
	// because the function changes it's state
	// through the pointer

	fmt.Println(d.first)

	helpers.Exit()
}
