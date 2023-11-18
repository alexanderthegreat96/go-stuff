package main

import "fmt"

func main() {

	// delete data from a map

	my_map := make(map[string]string)

	my_map["Alex"] = "Popescu"
	my_map["Andrei"] = "Stanescu"
	my_map["Paul"] = "Antonescu"

	fmt.Println("Intial Map: ", my_map)
	fmt.Println("Deleting: Andrei")

	delete(my_map, "Andrei")

	for key, value := range my_map {
		fmt.Printf("%s - %s\n", key, value)
	}

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
