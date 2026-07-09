package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunOutput(t *testing.T) {
	var buf bytes.Buffer

	run(&buf)

	output := strings.TrimSpace(buf.String())
	expected := "now-cli skeleton v0.0.1"

	if output != expected {
		t.Errorf("Unexpected output string.\nGot:      %q\nExpected: %q", output, expected)
	}
}
