package main

import "fmt"

func main() {
	// logical operators showcase
	// && AND
	// || OR
	// ! NOT

	// true && true -> true
	// true && false -> false
	// true || true -> true -> true or true
	// true || false -> true -> true or false
	// !true -> false -> not true

	var isHired bool = true
	var isNoHired bool = false

	var isTall = true
	var isShort = false

	var isMale = true
	var isFemale = false

	// true && true

	if isHired && isMale {
		fmt.Println("The subject is both male and is hired")
	} else {
		fmt.Println("The subject is not male and is not hired")
	}

	// true && false

	if isHired && isFemale {
		fmt.Println("Subject is hired but is a female")
	}

	// true || false

	if isHired || isFemale {

		// evaluates to true
		fmt.Println("Subject is hired and is a female")
	}

	if isHired || isMale {
		// evaluates to true
		fmt.Println("Subject is hired and is a male")
	}

	// !true

	if isHired && !isTall {
		fmt.Println("Subject is hired and is not tall.")
	}

	if isNoHired || isShort {
		fmt.Println("Subject is not hired or is short.")
	}

	fmt.Println("Press ENTER to exist...")
	fmt.Scanln()

}
