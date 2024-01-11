package main

import "fmt"

// type constraints

func addInteger(first, second int) int {
	return first + second
}

func addFloat(first, second float64) float64 {
	return first + second
}

// a generic type is a type that encapsulates multiple types
// so in this example, we can provide both a float64 or an int
// as opposted to having just one type constraint

func genericAdd[genericType int | float64](a, b genericType) genericType {
	return a + b
}

// type set interface

type Number interface {
	~int | ~float64
}

func genericMultiply[genericType Number](a, b genericType) genericType {
	return a * b
}

func main() {
	fmt.Println(addInteger(2, 3))
	fmt.Println(addFloat(2.34, 3.11))
	fmt.Println(genericAdd(12.2, 22.2))
	fmt.Println(genericAdd(12, 22))
	fmt.Println(genericMultiply(12, 22))
	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
