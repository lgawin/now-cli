package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunOutput(t *testing.T) {
	oldVersion := Version
	Version = "v1.2.3-test-tag"
	defer func() { Version = oldVersion }()

	var buf bytes.Buffer
	run(&buf)

	output := strings.TrimSpace(buf.String())
	expected := "now-cli version: v1.2.3-test-tag"

	if output != expected {
		t.Errorf("Unexpected output string.\nGot:      %q\nExpected: %q", output, expected)
	}
}
