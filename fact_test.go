package main

import (
	"testing"
)

func TestFactoriel(t *testing.T) {
	r := factoriel(5)
	if r != 120 {
		t.Fatalf("fact(5): want: 120, got: %d", r)
	}

	r = factoriel(0)
	if r != 1 {
		t.Fatalf("fact(0): want: 1, got: %d", r)
	}

	r = factoriel(-1)
	if r != -1 {
		t.Fatalf("fact(-1): want: -1, got: %d", r)
	}
}
