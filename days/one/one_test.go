package main

import "testing"

func TestPartTwoCalc(t *testing.T) {
	exp := 50346

	if res := partTwoCalc(100756); res != exp {
		t.Errorf("Expected %d, got %d\n", exp, res)
	}
}
