package functions

import "fmt"

//the function has to be capitalized in order
// to be exported aka callable outside this package
func Test() {
	fmt.Println("Hello world!")
	fmt.Printf("How are you, %s? \n", private_method("Alex"))
}

// private method
// since it beggins with
// a lowercase letter
// it cannot be called outside this package
func private_method(name string) string {
	return name
}
