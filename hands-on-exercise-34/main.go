package main

import "fmt"

func main() {
	//messing arround with bools

	fmt.Println("true && true is:", (true && true))   // true -> because both are true
	fmt.Println("true && false is:", (true && false)) // false // because true and false -> that means false
	fmt.Println("true || true is:", (true || true))   // true // true or true -> means true
	fmt.Println("true || false is:", (true || false)) // true // true or false -> means true, as the first thing is evaluated
	fmt.Println("!true is:", !true)                   // false  -> means false because ! negates the value
	fmt.Println("!false is:", !false)                 // true

	fmt.Println("Press ENTER to exit...")
	fmt.Scanln()
}
