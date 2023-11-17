package main

import "fmt"

func main() {
	// Use a slice literal to store these elements in
	// a slice. Also print out the slice and length
	// of the slice

	// try to get a for range loop on the slice

	names := []string{"Alice", "Bob", "Charlie", "David", "Eva", "Frank", "Grace", "Hank", "Ivy", "Jack",
		"Kate", "Leo", "Mia", "Nathan", "Olivia", "Peter", "Quinn", "Rachel", "Sam", "Tina",
		"Ulysses", "Vera", "Walter", "Xena", "Yvonne", "Zack", "Sophia", "Liam", "Emma", "Noah"}

	// print out the slice and it's length
	fmt.Printf("Names list is here: %v and it has %d elements \n", names, len(names))

	// let's loop over it

	fmt.Println("Looping over the names:")

	// i could have specified the key
	// but since i do not really need it
	// i skipped it using the blank identifier _

	for _, name := range names {
		fmt.Println(name)
	}

	// values can be accesed using their numeric index / key
	fmt.Println("Value for key: 3", names[3])

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()

}
