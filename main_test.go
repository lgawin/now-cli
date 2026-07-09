package main

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

func TestRunOutputMatrix(t *testing.T) {
	oldVersion := Version
	Version = "v1.2.3-test-tag"
	defer func() { Version = oldVersion }()

	mockTime := time.Date(2026, time.July, 10, 10, 0, 0, 0, time.UTC)
	fakeClock := func() time.Time {
		return mockTime
	}

	tests := []struct {
		name           string
		args           []string
		expectedOutput string
	}{
		{
			name:           "Default Version Output",
			args:           []string{"now", "--version"},
			expectedOutput: "now-cli version: v1.2.3-test-tag",
		},
		{
			name:           "Format Flag - ISO",
			args:           []string{"now", "-f", "iso"},
			expectedOutput: "2026-07-10T10:00:00Z", // Exact ISO string match
		},
		{
			name:           "Format Flag - Epoch",
			args:           []string{"now", "-f", "epoch"},
			expectedOutput: "1773136800", // Exact Unix epoch for July 10, 2026 10:00:00 UTC
		},
		{
			name:           "Format Flag - default is iso",
			args:           []string{"now"},
			expectedOutput: "2026-07-10T10:00:00Z", // Exact ISO string match
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			run(&buf, tt.args, fakeClock)

			output := strings.TrimSpace(buf.String())
			if output != tt.expectedOutput {
				t.Errorf("%s failed:\nGot:      %q\nExpected: %q", tt.name, output, tt.expectedOutput)
			}
		})
	}
}
