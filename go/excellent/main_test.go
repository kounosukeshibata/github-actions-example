package main

import (
	"testing"
)

func TestEvenOrOdd(t *testing.T) {
	result := EvenOrOdd(2)
	if result != "Even" {
		t.Errorf("EvenOrOdd(2) = %s; want Even", result)
	}
}