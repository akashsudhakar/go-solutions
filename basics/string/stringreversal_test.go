package main

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	input := "Akash"
	rev := "hsakA"
	output, error := reverseString(input)
	if error != nil {
		t.Fatalf("Error during execution %v", error)
	}
	if !(output == rev) {
		t.Errorf("Output not matching %v ", output)
	}
}
