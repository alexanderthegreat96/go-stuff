package main

import (
	"fmt"

	"github.com/alexanderthegreat96/go-stuff/hands-on-exercise-9/variables"
)

func main() {
	my_public_function := variables.PublicFunction()
	fmt.Println("Calling a method from an internal package  / folder: ", my_public_function)

	// short declaration operator
	username, password, weight, height := variables.ReturnVariables()

	fmt.Println("Calling variables via function.")
	fmt.Printf("username: %v of type:  %T \n", username, username)
	fmt.Printf("password: %v of type:  %T \n", password, password)
	fmt.Printf("height: %v of type:  %T \n", height, height)
	fmt.Printf("weight: %v of type:  %T \n", weight, weight)

	address := "Upper Farborn Road 12, VA, United States \n"

	fmt.Println(address)

	// Keep the window open
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
