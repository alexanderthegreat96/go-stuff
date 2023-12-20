package main

import (
	"reflect"
	"testing"

	"github.com/alexanderthegreat96/go-stuff/package-helpers/helpers"
)

// The code contains two test functions for the helpers package, one for testing the
// RandomStringGenerator function and another for testing the ReturnRandomWords function.
func TestRandomStringGenerator(t *testing.T) {
	value, err := helpers.RandomStringGenerator(80)

	if len(value) == 0 {
		t.Errorf("Something happened. Wanted a string but got an error: %v", err.Error())
	}
}

func TestReturnRandomWords(t *testing.T) {
	word, _, errWord := helpers.ReturnRandomWords()

	if len(word) == 0 {
		t.Errorf("Something happened. Wanted a string but got an error: %v", errWord.Error())
	}

	_, words, _ := helpers.ReturnRandomWords(10)

	// check returned type
	reflectValue := reflect.ValueOf(words)
	if reflectValue.Kind() != reflect.Slice {
		t.Errorf("Wanted slice but got %v", words)
	}

	// check if there are any values

	if len(words) == 0 {
		t.Errorf("Wanted a slice of words but got nothing")
	}

}
