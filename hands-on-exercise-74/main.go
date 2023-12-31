package main

import (
	"fmt"

	"github.com/alexanderthegreat96/go-stuff/package-helpers/helpers"
)

// create a value and assign it to a variable
// print that variable's address

func main() {
	x := "seth"
	fmt.Println(&x)

	helpers.Exit()
}
