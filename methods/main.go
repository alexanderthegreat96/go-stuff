package main

import "fmt"

type person struct {
	firstname string
	lastname  string
	salary    int
}

// for this we will use the receivers
// func (r receiver) identifier (p params) (r return(s)) {<code>}
// receivers help attach a METHOD to a type
// aka define a type or use a defined  custom type
// then attach a funciton to it

func (p person) whoIsThis() {
	fmt.Printf("This is %s %s and their salary is %d\n", p.firstname, p.lastname, p.salary)
}

func main() {
	dude := person{
		firstname: "Alex",
		lastname:  "Smith",
		salary:    10000,
	}

	dude.whoIsThis()

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
