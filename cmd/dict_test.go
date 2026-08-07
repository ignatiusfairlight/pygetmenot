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

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	buf.ReadFrom(r)

	expected := "Python dictionary, yes\n"
	if buf.String() != expected {
		t.Errorf("expected %q got %q", expected, buf.String())
	}
}