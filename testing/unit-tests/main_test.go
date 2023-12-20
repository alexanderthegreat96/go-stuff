package main

import "testing"

// testing doMath following the naming convention
func TestDoMath(t *testing.T) {
	gotSubstract := doMath(46, 16, substract)
	wantSubstract := 30

	if gotSubstract != wantSubstract {
		t.Errorf("Error. Want: %v and got %v", wantSubstract, gotSubstract)
	}

	gotAdd := doMath(46, 16, add)
	wantAdd := 62

	if gotAdd != wantAdd {
		t.Errorf("Error. Want: %v and got %v", wantAdd, gotAdd)
	}
}

// testing add following the naming convention
func TestAdd(t *testing.T) {
	got := add(46, 16)
	want := 62

	if want != got {
		t.Errorf("Error. Want: %v and got %v", want, got)
	}
}

// testing substract following the naming convention
func TestSubtract(t *testing.T) {
	got := substract(46, 16)
	want := 30
	if want != got {
		t.Errorf("Error. Want: %v and got %v", want, got)
	}
}
