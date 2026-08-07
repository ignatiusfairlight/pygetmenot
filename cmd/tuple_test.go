package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestTupleCmd(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	tupleCmd.Run(tupleCmd, []string{})

	if err := w.Close(); err != nil {
		t.Errorf("Failed to close pipe: %v", err)
	}
	
	os.Stdout = old

	var buf bytes.Buffer

	if _, err := buf.ReadFrom(r); err != nil {
		t.Errorf("Failed to read buffer: %v", err)
	}

	expected := tupleNotes
	if buf.String() != expected {
		t.Errorf("Expected contents of tuple.md, instead got %q", buf.String())
	}
}