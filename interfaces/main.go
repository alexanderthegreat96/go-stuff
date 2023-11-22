package main

import "fmt"

// interfaces and polymorphism
// an interface in go defines a set of method signatures
// polymorphism is the ability of a value of a CERTAIN TYPE to also be another TYPE

type Engine struct {
	Name       string
	Horsepower int
	Cylinders  int
}

// function that implement the receiver e of type Engine
func (e Engine) startEngine() {
	fmt.Printf("Starting engine: %s with %d that has %d cylinders\n", e.Name, e.Horsepower, e.Cylinders)
}

// function that implement the receiver e of type Engine
func (e Engine) stopEngine() {
	fmt.Printf("Stopping engine: %s with %d that has %d cylinders\n", e.Name, e.Horsepower, e.Cylinders)
}

// interface that lays out the template of
// what functions will be available if implement
// Todd McLeod -> "Hey baby, if you got these methods, then you are my type"
type CarMustImplement interface {
	startEngine()
	stopEngine()
}

// funtion that implements interface
// this illustrates polymorphism
// any VALUE that uses this method is of type CarMustImplement
// even if in this case, it is using the type Engine

func turnKey(c CarMustImplement) {
	c.startEngine()
	c.stopEngine()
}

func main() {
	// showcasing usage of interfaces

	engine := Engine{
		Name:       "Hemi V8",
		Horsepower: 300,
		Cylinders:  8,
	}

	turnKey(engine)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
