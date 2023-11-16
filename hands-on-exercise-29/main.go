package main

import "fmt"

func main() {

	// create a loop that runs 5 times
	// nest a loop within the first loop  that also prins 5 times
	// print something in each loop to illustrate what is occuring

	for i := 1; i <= 30; i++ {
		fmt.Printf("Main loop is at %d \n", i)

		for x := 1; x <= 5; x++ {
			fmt.Printf("Nested loop is at %d, while the main loop is at %v \n", x, i)

			if x == 5 {
				fmt.Println("Nested loop done")
			}
		}

		if i == 30 {
			fmt.Println("Main loop done.")
		}
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
