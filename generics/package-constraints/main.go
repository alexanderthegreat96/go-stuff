package main

// package constraints

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// Package constraints defines a set of useful constraints to be used with type parameters.
// https://pkg.go.dev/golang.org/x/exp/constraints

type Number interface {
	constraints.Integer | constraints.Float
}

func genericMultiply[genericType Number](a, b genericType) genericType {
	return a * b
}

func main() {
	fmt.Println(genericMultiply(11, 2))
	fmt.Println(genericMultiply(12.3, 21.4))
}
