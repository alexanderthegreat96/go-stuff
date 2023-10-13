package main

import (
	"fmt"

	"github.com/alexanderthegreat96/puppy"
)

func main() {
	//calling function from package we already build

	fmt.Println(puppy.BarkAt("Alexander"))

	fmt.Println("Press ENTER to exist...")
	fmt.Scanln()
}
