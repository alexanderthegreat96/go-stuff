package main

import (
	"fmt"
	"time"
)

// function wrappers exercise
// equivallent of decorators in python
// or thereabouts

func benchmark(inputFunc func(int), to int) { // gave it to parameter in other to have the function access to that
	fmt.Println("Starting benchmark...")
	start := time.Now()

	inputFunc(to)

	elapsed := time.Since(start)
	fmt.Printf("Execution time took %s\n", elapsed)
}

func countTo(to int) {
	for i := 0; i < to; i++ {
		fmt.Println(i)
	}
}

func main() {
	to := 100000
	benchmark(countTo, to)

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
