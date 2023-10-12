package main

import (
	"fmt"

	"github.com/alexanderthegreat96/puppy"
)

// to updae packages
// go get -u package_name
// go get -i package_name@commit/branch
// this will force update to the latest commit

// different method

// go get package_name@latest
// probably easier to use

func main() {
	fmt.Println("Hello")

	a := puppy.Bark()
	b := puppy.Barks()

	c := puppy.BigBark()
	d := puppy.BigBarks()

	fmt.Println("Using Bark function from puppy package: ", a)
	fmt.Println("Using Barks function from puppy package: ", b)
	fmt.Println("Using BigBark function from puppy package: ", c)
	fmt.Println("Using BigBarks function from puppy package: ", d)

	// keep program running
	fmt.Println("Press Enter to exit...")
	fmt.Scanln()
}
