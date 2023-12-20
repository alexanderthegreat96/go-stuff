package main

import "fmt"

// exercise callbacks
// a callback is when a fucntion receives a function
// and calls it within that function

func main() {
	fmt.Println("Showcasing providing a function to another function.")

	unnamed := func() {
		fmt.Println("Running...")
	}

	FuncWrapper(unnamed)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}

func FuncWrapper(f func()) {
	fmt.Println("Running before func...")

	f()

	fmt.Println("Running after func...")

}
