package main

import (
	"fmt"
	"log"

	"github.com/alexanderthegreat96/go-stuff/dependency-management/functions"
	"github.com/raja/argon2pw"
)

func main() {

	// Importing a package involves using the package_name/folder
	// Packages are stored in folders
	// A module can have multiple packages
	// Also a module can use other modules as well

	fmt.Println("Dependency management in Go")

	// Imported package + function

	functions.Test()

	// Requiring external modules
	// go get module_name_here

	// Let's use a password hashing package
	// github.com/raja/argon2pw

	testPassString := "alexanderisthecoolest1231231"
	// assigning 2 vars because it has 2 returns, the hashed value and the error, if defined

	hashedPass, err := argon2pw.GenerateSaltedHash(testPassString)

	if err != nil {
		log.Panicf("Password hashes encountered an error. Err: %v", err)
	}

	fmt.Printf("Hashed password for %s is %s \n", testPassString, hashedPass)

	//Let's check if the passwords match now
	checkHashStatus, checkHashError := argon2pw.CompareHashWithPassword(hashedPass, testPassString)

	if checkHashError != nil {
		log.Panicf("Password hash check returned an error. Err: %v", err)
	}

	fmt.Printf("Password comparison for password: %s and it's hash: %s is %t \n", testPassString, hashedPass, checkHashStatus)

	// Call some variables from a package in a different file
	// The stuff is found in a file called variables
	// as you can see, file name is irrelevant for importing
	// as it's all done through package name

	fmt.Printf("Hello, %s %s \n", functions.FirsName, functions.LastName)

	// keep program running

	fmt.Println("Press Enter to exit...")
	fmt.Scanln()

}
