package main

import "fmt"

func main() {
	// create a slice to store all the nbames of the states in the United States of America

	// ! use make and append to do this
	// goal: DO NOT HAVE THE ARRAY that underlies the slice created more than once

	states := make([]string, 0, 57)

	states = append(states, "AL|Alabama",
		"AK|Alaska",
		"AZ|Arizona",
		"AR|Arkansas",
		"CA|California",
		"CO|Colorado",
		"CT|Connecticut",
		"DE|Delaware",
		"FL|Florida",
		"GA|Georgia",
		"HI|Hawaii",
		"ID|Idaho",
		"IL|Illinois",
		"IN|Indiana",
		"IA|Iowa",
		"KS|Kansas",
		"KY|Kentucky",
		"LA|Louisiana",
		"ME|Maine",
		"MD|Maryland",
		"MA|Massachusetts",
		"MI|Michigan",
		"MN|Minnesota",
		"MS|Mississippi",
		"MO|Missouri",
		"MT|Montana",
		"NE|Nebraska",
		"NV|Nevada",
		"NH|New Hampshire",
		"NJ|New Jersey",
		"NM|New Mexico",
		"NY|New York",
		"NC|North Carolina",
		"ND|North Dakota",
		"OH|Ohio",
		"OK|Oklahoma",
		"OR|Oregon",
		"PA|Pennsylvania",
		"RI|Rhode Island",
		"SC|South Carolina",
		"SD|South Dakota",
		"TN|Tennessee",
		"TX|Texas",
		"UT|Utah",
		"VT|Vermont",
		"VA|Virginia",
		"WA|Washington",
		"WV|West Virginia",
		"WI|Wisconsin",
		"WY|Wyoming",
		"DC|District of Columbia",
		"AS|American Samoa",
		"GU|Guam",
		"MP|Northern Mariana Islands",
		"PR|Puerto Rico",
		"UM|United States Minor Outlying Islands",
		"VI|Virgin Islands, U.S.")

	// Print out len, cap and the values along with their index position without using range clauses
	fmt.Printf("The states slice has a length of %d and a capacity of %d\n", len(states), cap(states))

	// print the values without using range
	// so we basically use a loop

	for i := 0; i < len(states); i++ {
		fmt.Printf("State: %s is at index: %d\n", states[i], i)
	}

	fmt.Println("PRESS ENTER to exit...")
	fmt.Scanln()

}
