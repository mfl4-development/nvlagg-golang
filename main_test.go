package main

import (
	"bytes"
	"os"
	"testing"
)

func TestMainOutput(t *testing.T) {
	// Save original stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Call main
	main()

	// Restore stdout
	w.Close()
	os.Stdout = old

	// Read output
	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Check output
	expected := "Hello, World!\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}
