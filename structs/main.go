package main

import "fmt"

// important to notice
// exporting  / making
// a type public, requires
// you to use a capital letter
// person -> private
// Person -> public

type Person struct {
	FirstName  string
	LastName   string
	Age        int
	Height     float32
	Salary     int
	Ocupations []string
}

// embeddable structs
// aka structs in structs

type People struct {
	Employees []Person
	Company   string
}

func main() {

	// structs are composite types
	// they serve as complex data structures
	// allowing us to define custom
	// types that contain multiple data types

	// it agregates multiple types together

	uncle := Person{
		FirstName:  "Marius",
		LastName:   "Popescu",
		Age:        54,
		Height:     6.3,
		Salary:     60000,
		Ocupations: []string{"welder", "mechanic", "electrician", "builder"},
	}

	fmt.Printf("I have an uncle, his name is %s %s, he is %d years old, with a height of %f feet. \n", uncle.FirstName, uncle.LastName, uncle.Age, uncle.Height)
	fmt.Println("He had various jobs accross time, some of them include:")

	for _, job := range uncle.Ocupations {
		fmt.Println("- ", job)
	}

	// embedded structs
	// struct within a struct

	luxoft := []Person{{
		FirstName:  "Popescu",
		LastName:   "Alexandru",
		Age:        33,
		Salary:     57000,
		Ocupations: []string{"engineer", "electronics specialist", "developer"},
	}, {
		FirstName:  "Toma",
		LastName:   "Ion",
		Age:        27,
		Salary:     120000,
		Ocupations: []string{"entrepreneur", "specialist", "business owner"},
	}}

	microsoft := []Person{
		{
			FirstName:  "Dontay",
			LastName:   "Wilder",
			Age:        35,
			Salary:     78000,
			Ocupations: []string{"boxer", "coach"},
		},
	}

	companyEmployees := []People{{
		Employees: luxoft,
		Company:   "Luxoft",
	}, {
		Employees: microsoft,
		Company:   "Microsoft",
	}}

	for _, company := range companyEmployees {
		fmt.Printf("%s has the following employees:\n", company.Company)

		if len(company.Employees) > 0 {
			for _, employee := range company.Employees {
				fmt.Printf("%s %s, age: %d gets paid: %d, for the roles:\n", employee.FirstName, employee.LastName, employee.Age, employee.Salary)

				if len(employee.Ocupations) > 0 {
					for _, ocupation := range employee.Ocupations {
						fmt.Println("- ", ocupation)
					}
				} else {
					fmt.Println("No ocupation found.")
				}
			}
		}
	}

	// annonymous structs
	// basically
	// we define a struct directly into the
	// variable

	annon := struct {
		name    string
		address string
	}{
		name:    "Marius",
		address: "Near the plaza 123",
	}

	fmt.Printf("This is: %s and he lives at %s\n", annon.name, annon.address)

	// This allows us to compose
	// data on the fly, if we do not need
	// the data model to be used multiple times

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
