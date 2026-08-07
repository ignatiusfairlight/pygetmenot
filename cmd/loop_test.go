package cmd

import (
	"bytes"
	"os"
	"testing"
)

func TestLoopCmd(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	loopCmd.Run(loopCmd, []string{})

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)

	expected := loopNotes
	if buf.String() != expected {
		t.Errorf("Expected contents of loop.md, instead got %q", buf.String())
	}
}