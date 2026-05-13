package main

import (
	"bytes"
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"
)

func TestRunServerModeStdio(t *testing.T) {
	root := filepath.FromSlash("../..")
	input := strings.Join([]string{
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`,
		`{"jsonrpc":"2.0","method":"notifications/initialized","params":{}}`,
		`{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}`,
	}, "\n") + "\n"

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if err := run([]string{"--root", root}, strings.NewReader(input), &stdout, &stderr); err != nil {
		t.Fatalf("run server mode: %v\nstderr=%s", err, stderr.String())
	}
	if strings.Contains(stdout.String(), "design-records-mcp ready") || strings.Contains(stdout.String(), "records:") {
		t.Fatalf("server mode wrote summary to stdout:\n%s", stdout.String())
	}
	lines := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	if len(lines) != 2 {
		t.Fatalf("response lines len = %d, want 2\n%s", len(lines), stdout.String())
	}
	for i, line := range lines {
		var response map[string]any
		if err := json.Unmarshal([]byte(line), &response); err != nil {
			t.Fatalf("line %d is not JSON: %v\n%s", i+1, err, line)
		}
		if response["jsonrpc"] != "2.0" {
			t.Fatalf("line %d jsonrpc = %#v, want 2.0", i+1, response["jsonrpc"])
		}
		if _, ok := response["error"]; ok {
			t.Fatalf("line %d unexpected error: %#v", i+1, response)
		}
	}
}

func TestRunServerModeEmptyInputDoesNotWriteSummary(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if err := run([]string{"--root", filepath.FromSlash("../..")}, strings.NewReader(""), &stdout, &stderr); err != nil {
		t.Fatalf("run server mode empty input: %v\nstderr=%s", err, stderr.String())
	}
	if got := stdout.String(); got != "" {
		t.Fatalf("server mode stdout = %q, want empty", got)
	}
}

func TestRunSummaryMode(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if err := run([]string{"--root", filepath.FromSlash("../.."), "--summary"}, strings.NewReader(""), &stdout, &stderr); err != nil {
		t.Fatalf("run summary mode: %v\nstderr=%s", err, stderr.String())
	}
	got := stdout.String()
	for _, want := range []string{"design-records-mcp ready", "root:", "records:"} {
		if !strings.Contains(got, want) {
			t.Fatalf("summary stdout missing %q:\n%s", want, got)
		}
	}
}
