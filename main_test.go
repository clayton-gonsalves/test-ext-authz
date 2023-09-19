package main

import (
	"testing"
)

// Test function for the helper function.
func TestHelper(t *testing.T) {
	expected := "hello\n"
	got := helper()

	if got != expected {
		t.Errorf("Expected \"%s\", but got \"%s\"", expected, got)
	}
}
