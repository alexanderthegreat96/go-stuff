package main

import (
	"fmt"
	"math"
)

// interfaces
// create a type square with lenght and width
// create a type circle with radius float64
// attach a method to each that calculates the area and returns it

type square struct {
	width  float64
	length float64
}

type circle struct {
	radius float64
}

// length times width
func (s square) area() float64 {
	return s.length * s.width
}

// PI * radius squared
func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// in other words
// anything that implementds the
// area shape
// is of type shape

type shape interface {
	area() float64
}

// the type shape
func info(s shape) float64 {
	return s.area()
}

func main() {
	sq := square{
		width:  120,
		length: 110,
	}

	ci := circle{
		radius: 310,
	}

	fmt.Printf("The area of the square is: %f and it is of type %T\n", sq.area(), sq)
	fmt.Printf("The area of the circle is: %f and it is of type %T\n", ci.area(), ci)

	sq2 := square{
		width:  100,
		length: 100,
	}

	ci2 := circle{
		radius: 56.2,
	}

	fmt.Printf("The area of the square shape is %f\n", info(sq2))
	fmt.Printf("The area of the circle shape is %f\n", info(ci2))

	// the reason why the info function works
	// is because we attached the area function to it
	// the area function is "hooked" to the type square and circle
	// it is capable of identifying the different types, based on whatever
	// type we give it
	// so cirlce and square fall under the type shape

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
