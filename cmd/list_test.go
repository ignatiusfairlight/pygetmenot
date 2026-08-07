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

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)

	expected := listNotes
	if buf.String() != expected {
		t.Errorf("Expected contents of list.md, instead got %q", buf.String())
	}
}