package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestListCmd(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	listCmd.Run(listCmd, []string{})

	if err := w.Close(); err != nil {
		t.Errorf("Failed to close pipe: %v", err)
	}
	
	os.Stdout = old

	var buf bytes.Buffer

	if _, err := buf.ReadFrom(r); err != nil {
		t.Errorf("Failed to read buffer: %v", err)
	}

	expected := listNotes
	if buf.String() != expected {
		t.Errorf("Expected contents of list.md, instead got %q", buf.String())
	}
}