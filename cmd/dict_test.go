package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestDictCmd(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	dictCmd.Run(dictCmd, []string{})

	if err := w.Close(); err != nil {
		t.Errorf("Failed to close pipe: %v", err)
	}
	
	os.Stdout = old

	var buf bytes.Buffer

	if _, err := buf.ReadFrom(r); err != nil {
		t.Errorf("Failed to read buffer: %v", err)
	}

	expected := dictNotes
	if buf.String() != expected {
		t.Errorf("Expected contents of dict.md, instead got %q", buf.String())
	}
}