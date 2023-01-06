package main

import "testing"

func TestExample(t *testing.T) {
	expected := 2
	got := 1 + 1
	if expected != got {
		t.Fatalf("expected: %d, got: %d\n", expected, got)
	}
}
