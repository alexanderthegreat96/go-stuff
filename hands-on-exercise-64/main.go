package main

import "fmt"

// unit tests in Go

// unit test is a type of software testing where individual components
// or units of a software are tested. The purpose is to validate that
// each unit of the software performs as designed. a unit test is the
// smallest tewstable part of any software. it usually has one or few
// inputs and usually a single output

// unit tests in go are tipically written using the built-in testing package,
// "testing". This package does not require any third party libraries, but
// it's somewhat limited compared to other languages' testing tools. despite
// it's simplicity, it has enough features to cosntruct effective unit tests

// when you ask if this differs from a "test" in go, it's worth clarifying that
// a "unit test" is a subset of "test". tests in software can take many forms
// such as: integration tests, functional tests, system tests, etc

// an "integration test", for example, in contrast to a "unit test", would test
// the interaction between multiple components of the system, to ensure they work
// together correctly.

// to summarize, in go, a unit test is just one kind of test you can conduct, focused
// on verifying the correct behavior of a small, isolated piece of functionality.
// other types of tests have different focuses and may require different tools and
// techniques

func main() {
	fmt.Printf("%T", add)
	fmt.Printf("%T", substract)
	fmt.Printf("%T", doMath)

	x := doMath(42, 16, add)
	fmt.Println(x)

	y := substract(42, 16)
	fmt.Println(y)

}

// created test for this function
func doMath(a int, b int, f func(int, int) int) int {
	return f(a, b)
}

// created test for this function
func add(a int, b int) int {
	return a + b
}

// created test for this function
func substract(a int, b int) int {
	return a - b
}
