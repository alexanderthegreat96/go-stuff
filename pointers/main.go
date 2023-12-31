package main

import "fmt"

// pointers
// related to memory
// a pointer is a variable that stored the memory address
// of another variable

// every location memory has an address
// a pointer is a memory address
// use & in front of whatever variable you have
// in order to get it's memory address
// lexical elements:
// & -> address operator -> gives you the address
// * -> dereferencing operator
// *int -> type of pointer

func main() {

	x := 42 // variable
	y := &x // y points to x's memory address

	fmt.Printf("x is %d and it's memory address is %d\n", x, y) // that's the pointer, we get the pointer using & followed by var name: &my_Var

	// types of pointers
	// *string
	// *int
	// *map[]string

	z := "Alexander"
	fmt.Printf("Value is %v, it's address is: %d, and the type is %T\n", z, &z, &z)

	v := map[string]string{
		"first_name": "alex",
		"last_name":  "popescu",
	}
	// %p prints the hexadecimal representation of a memory address
	fmt.Printf("Value is %s, it's address is: %p, and the type is %T\n", v, &v, &v)

	// another way of representing a pointer variable

	var my_int int = 67
	var my_int_ptr *int = &my_int // i've declade the type of pointer

	// you can also use %v for the memory address

	fmt.Printf("Value is %d, it's address is: %v, and the type is %T\n", my_int, my_int_ptr, my_int_ptr)

	// Dereferencing pointers
	// *var
	// *int

	r := 56
	t := &r // assigned the address of r to t

	fmt.Printf("Value is %d, it's address is: %v, and the type is %T\n", r, t, t)

	// when using dereferencing, you essentially get the value of an address
	// for instance, t points to r
	// t contains the memory address of r
	// by using * in front of t (since t is a pointer)
	// it gives the value of r

	fmt.Printf("t: %v is assigned the address of r: %d and by using * in front of t, since t contains the address of r, we get the value of r: %d \n", t, r, *t)
	printAdressValueRefference(r, "r", "t")

	// another clearer example
	my_name := "alex"
	my_name_ptr := &my_name             // memory address / pointer &my_name
	my_name_dereference := *my_name_ptr // using the pointer to dereference / get the value of the referenced value -> *my_name_ptr

	fmt.Printf("given: %s, it has the address of %v and getting its value from the address is: %s\n", my_name, my_name_ptr, my_name_dereference)

	*my_name_ptr = "seth" // what i did here was to modify the value of my_name, through dererefencing
	// in other words, i used it's address which was stored by my_name_ptr and set it a different value

	fmt.Printf("Modified the value of: %s, through it's refference %v and set it to: %s \n", "alex", my_name_ptr, "seth")

	// What we learned so far
	// & gives the memory address
	// * gives the value of the memory address
	// using * can get the value of the address or modify the value, which is really cool

	// pointer types and refference types

	a := 42

	fmt.Printf("Before dereferencing: %d\n", a)
	intDelta(&a) // &a is the address
	fmt.Printf("After dereferencing: %d\n", a)

	// about slices
	// slices are built on top of arrays
	// which means that under the hood
	// they point to an array

	xi := []int{1, 2, 3, 4}

	fmt.Println(xi)

	sliceDelta(xi)

	fmt.Println(xi)

	// for maps

	m := make(map[string]int)
	m["james"] = 42
	fmt.Println(m["james"])
	mapDelta(m, "james")
	fmt.Println(m["james"])

	// Why did we do the stuff from above?
	// to illustrate the fact that
	// we pass things by value to modify them,
	// same goes for poointers
	// we pass the address as a value
	// then based on that address we modify something

	sm := map[string]string{
		"alex": "developer",
	}

	fmt.Printf("Using %p to change the value of %v\n", &sm, sm)

	// set a pointer to sm
	smt_r := &sm // *map[string]string

	// set a value for smt_r
	*smt_r = map[string]string{
		"alex": "manager",
	}

	fmt.Printf("Changed the value of %p and set it to: %v\n", smt_r, sm)

	// so essentially, i used the same place in memory instead of creating a new spot

	// doing the same with a slice

	ms := []string{"oop", "is", "trash", "fuck", "it"}

	fmt.Printf("Using %p to change the value of %v\n", &ms, ms)

	ms_r := &ms

	*ms_r = []string{"just", "replaced", "the", "items", "above"}

	fmt.Printf("Changed the value of %p and set it to: %v\n", ms_r, ms)

	// what happened here is quite simple, i simply used the same spot in memory to store
	// different states, one before modification and one after, all using the same address

	// POINTER SEMANTICS AND VALUE SEMANTICS

	// the logic of how something works
	// when should you use pointer semantics
	// when should you use value semantics

	// Value semantics in Go refers to when the actual data of a variable is passed to a function
	// or assigned to another variable. This means that the new variable or function parameter
	// gets a completely independent copy of the data

	// example

	// value semantics
	thisint := 43
	fmt.Println(thisint)               // prints the value of the variable thisint
	fmt.Println(addNumber(thisint, 2)) // creates a new copy of the data

	// without actually using it's address
	// thisint = 33

	// pointer semantics, no new variable is created, instead, the data gets pased tghere
	addNumberToVariable(&thisint, 2) // grabbed the address via &thisint
	fmt.Println(thisint)

	// POINTER AND VALUE SEMANTIC HEURISTICS
	// as a general idea, semantics describes how you implement the code, code-wise, that's it
	// AKA when should we use which

	// Easiest way to display it:
	// value semantics: copied values (pass by value)
	// pointer semantics: shared memory

	// 1. When to use value semantics?
	// - they are simpler and safer, they do not invoke a shared state of memory
	// - as a rule of thumb, if a function does not need to modiify it's input state
	// or the data that you are working with, use value semantics
	// 2. When to use pointer semantics?
	// - when working with larger data, since copying large structs or arrays can be inefficient
	// - if the data you're working with is large, you might use pointer semantics to avoid the cost
	// of copying the data. A rule of thumb: 64 bytes or larger: use pointer semantics
	// 3. Use pointer semantics for mutability:
	//  - If a function or method needs to modify it's receiver or input paramter, you'll need to use
	// pointer semantics

	// 4. Consistency:
	// It's important to be consistent, therefore, if some functions on a type use pointer semantics and
	// others use value semantics, it can lead to confussion. Tipically, once a type has a method with
	// pointer semantics, all methods on that type should have pointer semantics.
	// 5. Pointer semantics when interfacing with other code
	// If you're interfacting weith other code (like library or system call), you might need to use pointer semantics
	// For example, the json.Unmarshall function in the Go standard library requires a pointer to a value to populate it
	// with unmashalled data

	// These are just guidelines. The specifics can depend on the situation and sometimes, you might have good reason to
	// breake them, but it's generally good to follow through with them

	// VALUE SEMANTICS
	// they facilitate a higher level of integrity
	// since, the majority of bugs are related to mutation of memory
	// everybody getting a copy of the data will keep the data bug-free
	// if everyone isolates their copy, then the bug is easier to isolate on
	// a specific copy, rather than globally at any given point
	// also, value semantics reduce the side-effects around concurrency
	// more likely to keep the values on the stack

	// POINTERS, VALUES, THE HEAP AND THE STACK
	/*
		   The stack and the heap are two distinct regions of a computer's memory.

		   Stack:
		   1. Memory Allocation:
		      - Operates in a last-in, first-out (LIFO) manner.
		      - Managed by the compiler or runtime system.
		   2. Usage:
		      - Stores local variables and function call information.
		   3. Size:
		      - Generally limited and pre-allocated.
		   4. Speed:
		      - Fast access due to adjusting the stack pointer.
		   5. Automatic Management:
		      - Memory is deallocated when a function exits.

		   Heap:
		   1. Memory Allocation:
		      - Used for dynamic memory allocation.
		      - Managed explicitly by the programmer.
		   2. Usage:
		      - Allocates memory for data structures like arrays, linked lists, and objects.
		   3. Size:
		      - Dynamically grows or shrinks during execution.
		      - Typically larger than the stack.
		   4. Speed:
		      - Slower access due to searching for free memory blocks.
		   5. Manual Management:
		      - Memory needs explicit allocation and deallocation.
		        Failure to deallocate can lead to memory leaks.

			In summary, the stack is typically used for managing local variables and function calls in a structured and efficient way,
			while the heap is used for dynamic memory allocation, allowing for flexibility in managing data structures whose size may not be
			known at compile time. Both have their advantages and trade-offs, and understanding their differences is crucial for writing efficient
			and bug-free programs.
	*/

	// when we call functions, the values are in the stack
	// they are pre-allocated
	// it's afficient
	// it's self-cleaning

	// the heap
	// a region of dynamic allocation
	// basically, when you store soemthing in a memory address
	// that data is moved from the stack
	// must be cleaned by the dev
	// slower due to having to search a new region in memory
	// in Go, the garbage collector will release the stuff

	// when running the program
	// you can use go run --gcflags -m main.go
	// it will show how these things are going in
	// and out of the stack to the heap

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}

// simple flow of using a memory address
// getting it's address
// getting the value of a variable
// through it's pointer
// called dereferencing
func printAdressValueRefference(value int, varname string, refferencevar string) {

	fmt.Println("Showcasing memory address and derefferencing for: " + varname)
	fmt.Printf("%s with value: %d\n", varname, value)
	fmt.Printf("%s points to the address of %s which is: %v\n", refferencevar, varname, &value)
	fmt.Printf("Value of %s through the refference of %s is: %d\n", varname, refferencevar, *&value)
}

// *int is a pointer type
// so it takes a memory adress as parameter

func intDelta(n *int) {
	*n = 43 // dereference the original value of the address and setting a new value
}

// same for slices
func sliceDelta(ii []int) {
	ii[0] = 43
}

// passing by value and modifying things
func mapDelta(m map[string]int, s string) {
	m[s] = 43
}

// returns the data instead of mutating it's previous state
func addNumber(this int, many int) int {
	return this + many
}

// this becomes a pointer refference
func addNumberToVariable(this *int, many int) {
	*this = *this + 2 // modified the value of the variable passed by refference
}
