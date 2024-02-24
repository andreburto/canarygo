package canarygo

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestCanary(t *testing.T) {
	expected := "Black Canary, GO!"

	// Capture the stdout from Canary().
	originalStdout := os.Stdout
	readIn, writeOut, _ := os.Pipe()
	os.Stdout = writeOut

	// Call the function and capture the output.
	Canary()

	// Revert to the original stdout.
	writeOut.Close()
	stdOutput, _ := io.ReadAll(readIn)
	os.Stdout = originalStdout

	// Convert the bytes to a string.
	actual := string(stdOutput[:])

	if strings.Contains(actual, expected) != true {
		t.Error("Canary function did not contain extected output.")
	}
}
