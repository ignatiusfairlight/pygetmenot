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

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)

	expected := tupleNotes
	if buf.String() != expected {
		t.Errorf("Expected contents of tuple.md, instead got %q", buf.String())
	}
}