package main

import "testing"

func TestPartOneCalc(t *testing.T) {
	exp := 3500

	if res := partOneCalc([]int{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}); res != exp {
		t.Errorf("Expected %d, got %d\n", exp, res)
	}
}
