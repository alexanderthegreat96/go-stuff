package main

// this example is aimed at showcasing
// the use of initializers

// initializer or intializers are functions
// that use the keyword init() -> func init() {}
// they run before the main() -> func main() {}
// meaning that normally, the program starts in
// package main and runs in the func main() {}
// func init() {} runs before func main() {}

import (
	"fmt"
	"time"

	"github.com/alexanderthegreat96/go-stuff/package-helpers/donut"
)

func init() {
	currentTime := time.Now()
	formattedTime := currentTime.Format("Jan 02, 2006 15:03")
	fmt.Println("Program started at: " + formattedTime)
}

func main() {
	fmt.Println("Drawning a rotating sphere...")
	donut.RenderDonut()
}
