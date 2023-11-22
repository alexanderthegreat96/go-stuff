package main

import (
	"bytes"
	"fmt"
)

func main() {

	// A buffer is an area in memory where data can be stored temporarily.
	// It's like a holding space where information can be stored and retrieved.
	// Buffers are commonly used in programming to manage and transfer data efficiently.

	// initializing a buffer string
	buffer := bytes.NewBufferString("Hello ")

	// notice how buffer implements
	// the Stringer interface (String())
	fmt.Println(buffer.String())

	// adding more data to the string
	buffer.WriteString("Gophers! ")
	fmt.Println(buffer.String())

	// we can reset the buffer, wipping the in-memory data
	buffer.Reset()

	// adding data after it was reset
	buffer.WriteString("Adding data after the buffer was reset! ")

	fmt.Println(buffer.String())

	// we can also write bytes direcly
	buffer.Write([]byte("Wrote some bytes directly. "))

	fmt.Println(buffer.String())

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
