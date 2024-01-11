package main

import "fmt"

// type alias
// giving a type a different name basically

type myAlias int

func main() {
	var age myAlias = 52
	fmt.Println(age)
}
