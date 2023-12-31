package main

import "fmt"

// method sets
// a method set is a set of methods attached to a type.
// this concept is key to the go's interface mechanism,
// and it's associated with both value types and pointer types
// the method set of type T consists of all the methods with received T or T
// the method set of type *T consists of all the methods with received *T or T
// the idea of method set is integral to how interfaces are implemented and
// used in go
// An interface in Go defines a method set, and any type whose method is set
// is a superset of the interface's method set, is considered to implement that
// interface
// A crucial thing to remember in Go: if you define a method with a pointer recciever,
// the method is only in the method set of the pointer type. This is important in the context
// of interfaces because if an itnerface requires a method that's defined on the pointer (not on the value)
//, then oyu can only use a pointer to that type to satisfy the interface, not a value of the type

type dog struct {
	first string
}

func (d dog) walk() {
	fmt.Println("My name is: " + d.first + ", and I am walking...")
}

// we can use both pointers and normal variable for this
func (d *dog) run() {
	d.first = "Rover" // we mutated the value of the type dog through the refference
	fmt.Println("My name is: " + d.first + ", and I am running...")
}

type actions interface {
	walk()
	run()
}

func doAction(a actions) {
	a.run()
	a.walk()
}

func main() {
	d1 := dog{
		first: "henry",
	}

	d1.walk()
	d1.run()

	d2 := &dog{
		first: "padget",
	}

	d2.walk()
	d2.run()

	// d3 := dog{
	// 	first: "seth",
	// }

	// cannot use d3 (variable of type dog) as actions value in argument to doAction: dog does not implement actions (method run has pointer receiver)
	// doAction(d3)

	d4 := &dog{
		first: "Seth",
	}

	doAction(d4)

}
