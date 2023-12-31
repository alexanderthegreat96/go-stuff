package main

import (
	"fmt"

	"github.com/alexanderthegreat96/go-stuff/package-helpers/helpers"
)

// implementing interfaces
// implement an interface and some methods
// maintain consistency with the code
// receivers should either stick to POINTER or VALUE semantics

type Person struct {
	title     string
	name      string
	age       int
	gender    string
	ocupation string
	skills    []string
}

func (p *Person) hire() {
	fmt.Println("We are going to hire: ")
	fmt.Printf("%s %s that is %d years old.\n", p.title, p.name, p.age)

	if p.gender == "female" || p.gender == "f" {
		fmt.Printf("She has worked as %s using various skills, including:\n", p.ocupation)
	} else if p.gender == "male" || p.gender == "m" {
		fmt.Printf("He has worked as %s using various skills, including:\n", p.ocupation)
	} else {
		fmt.Printf("They have worked as %s using various skills, including:\n", p.ocupation)
	}

	if len(p.skills) > 0 {
		for _, skill := range p.skills {
			fmt.Printf("- %s\n", skill)
		}
	}

}

func (p *Person) fire() {
	fmt.Println("We are going to fire: ")
	fmt.Printf("%s %s that is %d years old.\n", p.title, p.name, p.age)

	if p.gender == "female" || p.gender == "f" {
		fmt.Printf("She has worked as %s using various skills, including:\n", p.ocupation)
	} else if p.gender == "male" || p.gender == "m" {
		fmt.Printf("He has worked as %s using various skills, including:\n", p.ocupation)
	} else {
		fmt.Printf("They have worked as %s using various skills, including:\n", p.ocupation)
	}

	if len(p.skills) > 0 {
		for _, skill := range p.skills {
			fmt.Printf("- %s\n", skill)
		}
	}

}

// means that whatever function uses this type
// it must implement these 2 methods
type People interface {
	fire()
	hire()
}

func companyAction(action string, p People) {
	if action == "fire" {
		p.fire()
	} else if action == "hire" {
		p.hire()
	} else {
		fmt.Println("Action not defined. Must be either: fire or hire")
	}
}

func main() {

	// used a pinter
	// the reason for that is because
	// instead of copy-ing the data in 2 memory addresses
	// for both function
	// it will use the same one for both

	person := &Person{
		title:     "Mr",
		name:      "Popescu Adrian",
		age:       30,
		gender:    "male",
		ocupation: "dev-ops engineer",
		skills:    []string{"docker", "kubernetes", "aws", "gc"},
	}

	companyAction("hire", person)
	companyAction("fire", person)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	// using a different approach

	var employee People = &Person{
		title:     "Mr",
		name:      "Popescu Adrian",
		age:       30,
		gender:    "male",
		ocupation: "dev-ops engineer",
		skills:    []string{"docker", "kubernetes", "aws", "gc"},
	}
	fmt.Printf("Employee is of type %T\n", employee)

	employee.fire()
	employee.fire()

	helpers.Exit()

}
