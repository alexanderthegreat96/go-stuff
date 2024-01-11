package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// package bcrypt is used for generating hashes

func main() {
	pass := "alexande12312312312"
	bs, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.MinCost)

	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("For pass: %s we have: %s\n", pass, bs)

	// to compare hash with pass
	err = bcrypt.CompareHashAndPassword([]byte(bs), []byte(pass))
	if err != nil {
		fmt.Println(err.Error())
	}
}
