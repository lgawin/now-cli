package main

import (
	"bytes"
	"os"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestRunOutputMatrix(t *testing.T) {
	oldVersion := Version
	Version = "v1.2.3-test-tag"
	defer func() { Version = oldVersion }()

	baseUTC := time.Date(2026, time.July, 10, 10, 0, 0, 0, time.UTC)

	locWarsaw, err := time.LoadLocation("Europe/Warsaw")
	if err != nil {
		t.Fatalf("Failed to load timezone: %v", err)
	}
	locNewYork, err := time.LoadLocation("America/New_York")
	if err != nil {
		t.Fatalf("Failed to load timezone: %v", err)
	}

	expectedUsageOutput := `Usage of now:
  A fast CLI tool to generate the current time in various formats.

Options:
  -f string
    	Shorthand for format (default "iso")
  -format string
    	Time format (e.g., 'epoch' or 'iso') (default "iso")
  -v	Shorthand for version
  -version
    	Print version`
	tests := []struct {
		name           string
		args           []string
		expectedOutput string
		location       *time.Location
	}{
		// --- version
		{
			name:           "Version Flag Output",
			args:           []string{"now", "--version"},
			expectedOutput: "now-cli version: v1.2.3-test-tag",
		},
		{
			name:           "Long Version Flag Output",
			args:           []string{"now", "-v"},
			expectedOutput: "now-cli version: v1.2.3-test-tag",
		},
		// --- Formatting ---
		// --- Short Flags ---
		{
			name:           "Format Flag - ISO",
			args:           []string{"now", "--format", "iso"},
			expectedOutput: "2026-07-10T10:00:00Z",
			location:       time.UTC,
		},
		{
			name:           "Format Flag - Epoch",
			args:           []string{"now", "--format", "epoch"},
			expectedOutput: "1783677600",
			location:       time.UTC,
		},
		// --- Long Flags ---
		{
			name:           "Long Format Flag - ISO",
			args:           []string{"now", "-f", "iso"},
			expectedOutput: "2026-07-10T10:00:00Z",
			location:       time.UTC,
		},
		{
			name:           "Long Format Flag - Epoch",
			args:           []string{"now", "-f", "epoch"},
			expectedOutput: "1783677600",
			location:       time.UTC,
		},
		// --- Defaults ---
		{
			name:           "Format Flag - default is iso",
			args:           []string{"now"},
			expectedOutput: "2026-07-10T10:00:00Z",
			location:       time.UTC,
		},
		// --- Explicit Timezone Variations ---
		{
			name:           "Format Flag - ISO (Warsaw CEST / UTC+2)",
			args:           []string{"now", "--format", "iso"},
			expectedOutput: "2026-07-10T12:00:00+02:00",
			location:       locWarsaw,
		},
		{
			name:           "Format Flag - ISO (New York EDT / UTC-4)",
			args:           []string{"now", "--format", "iso"},
			expectedOutput: "2026-07-10T06:00:00-04:00",
			location:       locNewYork,
		},
		{
			name:           "Epoch Format is Timezone Independent",
			args:           []string{"now", "-f", "epoch"},
			expectedOutput: "1783677600",
			location:       locWarsaw, // Even in Warsaw, the absolute epoch match holds
		},
		// --- Help Flags ---
		{
			name:           "Short Help Flag",
			args:           []string{"now", "-h"},
			expectedOutput: expectedUsageOutput,
		},
		{
			name:           "Long Help Flag",
			args:           []string{"now", "--help"},
			expectedOutput: expectedUsageOutput,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer

			targetZone := time.UTC
			if tt.location != nil {
				targetZone = tt.location
			}

			fakeClock := func() time.Time {
				return baseUTC.In(targetZone)
			}

			run(&buf, tt.args, fakeClock)

			output := strings.TrimSpace(buf.String())
			if output != tt.expectedOutput {
				t.Errorf("%s failed:\nGot:      %q\nExpected: %q", tt.name, output, tt.expectedOutput)
			}
		})
	}
}

func TestHelpFlags(t *testing.T) {
	// Build the binary once for the tests
	buildCmd := exec.Command("go", "build", "-o", "test-now")
	if err := buildCmd.Run(); err != nil {
		t.Fatalf("Failed to build test binary: %v", err)
	}
	// Clean up the binary after tests complete
	defer os.Remove("test-now")

	tests := []struct {
		name     string
		flag     string
		expected string
	}{
		{name: "Short Help Flag", flag: "-h", expected: "Usage of now:"},
		{name: "Long Help Flag", flag: "--help", expected: "Usage of now:"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			cmd := exec.Command("./test-now", tc.flag)
			output, err := cmd.CombinedOutput()

			// Help flags should exit cleanly with status code 0
			if err != nil {
				t.Fatalf("Expected exit code 0, got error: %v (Output: %s)", err, string(output))
			}

			if !strings.Contains(string(output), tc.expected) {
				t.Errorf("Expected output to contain %q, got: %s", tc.expected, string(output))
			}
		})
	}
}
