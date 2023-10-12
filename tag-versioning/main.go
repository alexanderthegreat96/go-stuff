package main

import (
	"fmt"
	"strings"

	"github.com/alexanderthegreat96/puppy"
)

// initialize an array of 5 string elements
var commands [5]string

// prints all strings to lowercase
// this is a private method
// due to the fact that we are using
// lowercase

func to_lowercase(arr [5]string) {

	// logic
	// for key, value in range of arr
	// using blank identifier
	// we are not actually using the key,
	// so it would be irrelevant

	for _, value := range arr {
		fmt.Println(strings.ToLower(value))
	}
}

func main() {

	// populate that array
	commands = [5]string{"GIT TAG", "git tag vN.N.N", "GIT PUSH ORIGIN --TAGS", "GIT TAG", "GIT SHOW V1.0.0"}

	fmt.Println("Git versioning through tags.")
	fmt.Println("The commands we are going to use are: ")

	to_lowercase(commands)

	barkAt := puppy.BarkAt("Alex")

	// using function from package that was tagged
	fmt.Println(barkAt)

	// when requiring packages
	// a certain version can be specified in the go.mod file
	// see the example
	// found in the go.mod file
	// not just versions can be specified, but commit as well

}
