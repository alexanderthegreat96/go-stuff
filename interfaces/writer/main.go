package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
)

// In Go, the io.Writer interface is a fundamental interface used for writing data. It is defined as follows:
// The Write method writes len(p) bytes from p to the underlying data stream.
// It returns the number of bytes written and an error if any.

// type Writer interface {
// 	Write(p []byte) (n int, err error)
// }

type person struct {
	first string
}

// implements the writer interface
// takes a pointer as an argument
// and a receiver as type

func (p person) writeOut(w io.Writer) error {
	_, err := w.Write([]byte(p.first))
	return err
}

func main() {
	// let's create a new file
	// to see how we can deal with
	// the writter interface

	// returns back a pointer
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatalf("Something happened: %s", err)
	}

	defer file.Close() // close the file once everything else is done (write) or until error appears

	// create a slice of bytes
	// to be able to write to a file
	f := []byte("Hello gophers!")

	// handle the error
	_, err = file.Write(f)

	if err != nil {
		log.Fatalf("Something happened: %s", err)
	}

	// bonus
	// relationship between a string an a byte in go

	// In go, a string and a []byte are two different types, but they are closely
	// related and can often bne converted between each other.

	// A string in Go represents a sequence of characters. It is an immutable type,
	// which means, you cannot modify the individual characters within a string.
	// String values are always interpreted as UTF-8 encoded Unicode text.

	// A []byte is a slice of bytes, where each element represents a single byte
	// It is a mutable type, so you can modify individual bytes within a byte slice.
	// It can be used to represent binary data or text in various encodings.

	// Go provided built-in functions to convert between strings and byte slices

	// Writtin to either a file or a byte buffer -> see buffers for this

	// create a new file
	new_file, error := os.Create("example.txt")
	if error != nil {
		log.Fatalf("Something happened: %s", error)
	}

	defer new_file.Close()

	p := person{
		first: "alexander",
	}

	var bytesBuffer bytes.Buffer // bytes buffer

	// write to the file itself
	p.writeOut(new_file)

	// write to the bytes buffer

	p.writeOut(&bytesBuffer) // needs to be a pointer to the buffer

	// to see what's in the buffer buffer.String()

	fmt.Println("This is what got stored in the buffer: " + bytesBuffer.String())

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
