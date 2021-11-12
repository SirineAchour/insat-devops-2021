package main

import (
	"testing"
)

func TestFactoriel(t *testing.T) {
	r := factorial(5)
	if r != 120 {
		t.Fatalf("fact(5): want: 120, got: %d", r)
	}

	r = factorial(0)
	if r != 1 {
		t.Fatalf("fact(0): want: 1, got: %d", r)
	}

	r = factorial(-1)
	if r != -1 {
		t.Fatalf("fact(-1): want: -1, got: %d", r)
	}
}
