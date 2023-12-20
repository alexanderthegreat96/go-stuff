package main

import (
	"fmt"

	"github.com/alexanderthegreat96/go-stuff/package-helpers/helpers"
)

func main() {
	randomString, _ := helpers.RandomStringGenerator(90)
	fmt.Println(randomString)
	randomWord, _, _ := helpers.ReturnRandomWords(1)
	fmt.Println(randomWord)
	_, randomWords, _ := helpers.ReturnRandomWords(10)
	fmt.Println(randomWords)
}
