// main_test.go
package main

import "testing"

// we test
// create a func that starts
// with Test followed by the name of the function
// TestAdd (Add is the function found in the main.go file)
// Then the rest is self explanatory
func TestAdd(t *testing.T) {
	total := Add(5, 5)

	if total != 10 {
		t.Errorf("Sum was incorrect, got %d, want: %d", total, 10)
	}
}
