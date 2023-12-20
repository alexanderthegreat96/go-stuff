package main

// testing in go
// go prtovides a built-in testing framework
// that simplifies the process of testing Go code.
// here's an overview of the file structure, naming
// convensions, and how testing works in Go

// 1. Test files
// - test files live in the same package as the code they test
// - they are named like this: "filename_test.go" where the filename is the name
// of the file that contains the code you want to test
// 2. Test Functions
// - must start with the word "Test", followed by a word that starts with a capital letter
// -  the signature of the test function is "func TestXxx(*testing.T)", where Xxx does not
// start with a lower letter

func main() {
	// nothing to see here
	// just testing
}

// let's do it

func Add(x, y int) int {
	return x + y
}
