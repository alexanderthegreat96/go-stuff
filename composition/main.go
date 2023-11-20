package main

import (
	"fmt"
	"strconv"
)

// composition
// OOP term
// composite data types
// aka
// agregates
// complex data types
// you put together multiple
// data types
// it's how we do OOP in to
// basically

type Engine struct {
	Cylinders  int
	Inline     bool
	Type       string
	Horsepower int
}

type CarModel struct {
	Brand   string
	Model   string
	Chassis string
}

// Takes the data for the engine type
func (e *Engine) Start() {

	engine_type := "v" + strconv.Itoa(e.Cylinders)
	if e.Inline {
		engine_type = "inline" + strconv.Itoa(e.Cylinders)
	}

	fmt.Printf("Start the %s engine, which has %d cylinders and it is a %s rocking %d horepower\n", e.Type, e.Cylinders, engine_type, e.Horsepower)
}

// takes the data for the CarModel Type
func (m *CarModel) Kind() {
	fmt.Printf("This is a %s %s which is a %s\n", m.Brand, m.Model, m.Chassis)
}

type Car struct {
	Engine
	CarModel
}

type alex string

func main() {

	car := Car{}

	car.Brand = "Mercedes"
	car.Model = "AMG E63"
	car.Chassis = "Sedan"

	car.Cylinders = 8
	car.Inline = false
	car.Type = "gas"
	car.Horsepower = 650

	car.Kind()
	car.Start()

	var dude alex = "This is my custom type"

	fmt.Printf("Created a custom type which is: %T and has the value %s", dude, dude)
}
