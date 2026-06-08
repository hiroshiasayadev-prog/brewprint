package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestRunServerModeStdio(t *testing.T) {
	root := filepath.FromSlash("../../../..")
	input := strings.Join([]string{
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`,
		`{"jsonrpc":"2.0","method":"notifications/initialized","params":{}}`,
		`{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}`,
		`{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"list_records","arguments":{"kind":"decision","limit":1}}}`,
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
	if len(lines) != 3 {
		t.Fatalf("response lines len = %d, want 3\n%s", len(lines), stdout.String())
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
		if i == 2 {
			result := response["result"].(map[string]any)
			content := result["content"].([]any)
			first := content[0].(map[string]any)
			if first["type"] != "text" || !json.Valid([]byte(first["text"].(string))) {
				t.Fatalf("tools/call content is not text JSON: %#v", first)
			}
			if result["isError"] != false {
				t.Fatalf("tools/call isError = %#v, want false", result["isError"])
			}
		}
	}
}

func TestRunServerModeEmptyInputDoesNotWriteSummary(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if err := run([]string{"--root", filepath.FromSlash("../../../../")}, strings.NewReader(""), &stdout, &stderr); err != nil {
		t.Fatalf("run server mode empty input: %v\nstderr=%s", err, stderr.String())
	}
	if got := stdout.String(); got != "" {
		t.Fatalf("server mode stdout = %q, want empty", got)
	}
}

func TestRunSummaryMode(t *testing.T) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	if err := run([]string{"--root", filepath.FromSlash("../../../../"), "--summary"}, strings.NewReader(""), &stdout, &stderr); err != nil {
		t.Fatalf("run summary mode: %v\nstderr=%s", err, stderr.String())
	}
	got := stdout.String()
	for _, want := range []string{"design-records-mcp ready", "root:", "records:"} {
		if !strings.Contains(got, want) {
			t.Fatalf("summary stdout missing %q:\n%s", want, got)
		}
	}
}

func TestProcessStdioSmoke(t *testing.T) {
	repoRoot, err := filepath.Abs(filepath.FromSlash("../../../../"))
	if err != nil {
		t.Fatalf("resolve repo root: %v", err)
	}
	cmdDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("resolve cmd dir: %v", err)
	}
	input := strings.Join([]string{
		`{"jsonrpc":"2.0","id":1,"method":"initialize","params":{}}`,
		`{"jsonrpc":"2.0","method":"notifications/initialized","params":{}}`,
		`{"jsonrpc":"2.0","id":2,"method":"tools/list","params":{}}`,
		`{"jsonrpc":"2.0","id":3,"method":"tools/call","params":{"name":"list_records","arguments":{"kind":"decision","limit":1,"order_by":"id","order":"desc"}}}`,
		`{"jsonrpc":"2.0","id":4,"method":"tools/call","params":{"name":"get_record","arguments":{"id":"V01-ADR-076","include_body":false}}}`,
		`{"jsonrpc":"2.0","id":5,"method":"tools/call","params":{"name":"validate_records","arguments":{}}}`,
		`{"jsonrpc":"2.0","id":6,"method":"tools/call","params":{"name":"suggest_next_record","arguments":{"kind":"decision","title":"Process Smoke Should Not Exist"}}}`,
		`{"jsonrpc":"2.0","id":7,"method":"tools/call","params":{"name":"get_record","arguments":{"id":"ADR-999"}}}`,
		`{"jsonrpc":"2.0","id":8,"method":"tools/call","params":{"name":"get_record","arguments":{}}}`,
	}, "\n") + "\n"

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	cmd := exec.CommandContext(ctx, "go", "run", ".", "--root", repoRoot)
	cmd.Dir = cmdDir
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		t.Fatalf("open stdin pipe: %v", err)
	}
	if err := cmd.Start(); err != nil {
		t.Fatalf("start design-records-mcp process: %v", err)
	}
	if _, err := io.WriteString(stdin, input); err != nil {
		t.Fatalf("write stdin: %v", err)
	}
	if err := stdin.Close(); err != nil {
		t.Fatalf("close stdin: %v", err)
	}
	if err := cmd.Wait(); err != nil {
		if ctx.Err() != nil {
			t.Fatalf("design-records-mcp process timed out; stdout=%s\nstderr=%s", stdout.String(), stderr.String())
		}
		t.Fatalf("design-records-mcp process failed: %v\nstdout=%s\nstderr=%s", err, stdout.String(), stderr.String())
	}
	if ctx.Err() != nil {
		t.Fatalf("design-records-mcp process timed out; stdout=%s\nstderr=%s", stdout.String(), stderr.String())
	}

	out := stdout.String()
	for _, forbidden := range []string{"design-records-mcp ready", "root:", "records:"} {
		if strings.Contains(out, forbidden) {
			t.Fatalf("server mode stdout contains summary text %q:\n%s", forbidden, out)
		}
	}
	lines := strings.Split(strings.TrimSpace(out), "\n")
	if len(lines) != 8 {
		t.Fatalf("response lines len = %d, want 8\nstdout=%s\nstderr=%s", len(lines), out, stderr.String())
	}

	responses := make(map[int]map[string]any, len(lines))
	for _, line := range lines {
		var response map[string]any
		if err := json.Unmarshal([]byte(line), &response); err != nil {
			t.Fatalf("stdout line is not JSON-RPC JSON: %v\nline=%s\nstdout=%s", err, line, out)
		}
		if response["jsonrpc"] != "2.0" {
			t.Fatalf("response jsonrpc = %#v, want 2.0; line=%s", response["jsonrpc"], line)
		}
		id, ok := response["id"].(float64)
		if !ok {
			t.Fatalf("response missing numeric id: %#v", response)
		}
		responses[int(id)] = response
	}
	assertInitializeSmokeResponse(t, responses[1])
	assertToolsListSmokeResponse(t, responses[2])
	assertListRecordsSmokeResponse(t, responses[3])
	assertGetRecordSmokeResponse(t, responses[4])
	assertValidateRecordsSmokeResponse(t, responses[5])
	assertSuggestNextRecordSmokeResponse(t, repoRoot, responses[6])
	assertToolErrorSmokeResponse(t, responses[7], "record_not_found")
	assertToolErrorSmokeResponse(t, responses[8], "invalid_request")
}

func assertInitializeSmokeResponse(t *testing.T, response map[string]any) {
	t.Helper()
	if response == nil || response["error"] != nil {
		t.Fatalf("initialize response error: %#v", response)
	}
	result := response["result"].(map[string]any)
	if result["protocolVersion"] == "" {
		t.Fatalf("initialize missing protocolVersion: %#v", result)
	}
	serverInfo := result["serverInfo"].(map[string]any)
	if serverInfo["name"] != "brewprint-design-records-mcp" {
		t.Fatalf("serverInfo.name = %#v", serverInfo["name"])
	}
	capabilities := result["capabilities"].(map[string]any)
	if _, ok := capabilities["tools"]; !ok {
		t.Fatalf("initialize missing tools capability: %#v", capabilities)
	}
}

func assertToolsListSmokeResponse(t *testing.T, response map[string]any) {
	t.Helper()
	if response == nil || response["error"] != nil {
		t.Fatalf("tools/list response error: %#v", response)
	}
	result := response["result"].(map[string]any)
	tools := result["tools"].([]any)
	for _, name := range []string{"list_records", "validate_records", "get_record", "suggest_next_record"} {
		found := false
		for _, item := range tools {
			tool := item.(map[string]any)
			if tool["name"] == name {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("tools/list missing %s: %#v", name, tools)
		}
	}
}

func assertListRecordsSmokeResponse(t *testing.T, response map[string]any) {
	t.Helper()
	text := assertToolResultText(t, response, false)
	var payload struct {
		Records []map[string]any `json:"records"`
	}
	unmarshalJSONText(t, text, &payload)
	if len(payload.Records) == 0 {
		t.Fatalf("list_records returned no records: %s", text)
	}
}

func assertGetRecordSmokeResponse(t *testing.T, response map[string]any) {
	t.Helper()
	text := assertToolResultText(t, response, false)
	var payload map[string]any
	unmarshalJSONText(t, text, &payload)
	record := payload["record"].(map[string]any)
	if record["id"] != "V01-ADR-076" {
		t.Fatalf("get_record id = %#v, want V01-ADR-076; text=%s", record["id"], text)
	}
	if _, ok := record["body"]; ok {
		t.Fatalf("get_record include_body=false returned body: %s", text)
	}
	headings, ok := record["headings"].([]any)
	if !ok || len(headings) == 0 {
		t.Fatalf("get_record headings missing: %s", text)
	}
}

func assertValidateRecordsSmokeResponse(t *testing.T, response map[string]any) {
	t.Helper()
	text := assertToolResultText(t, response, false)
	var payload map[string]any
	unmarshalJSONText(t, text, &payload)
	if _, ok := payload["ok"].(bool); !ok {
		t.Fatalf("validate_records missing ok bool: %s", text)
	}
	if _, ok := payload["diagnostics"]; !ok {
		t.Fatalf("validate_records missing diagnostics: %s", text)
	}
	if payload["diagnostics"] != nil {
		if _, ok := payload["diagnostics"].([]any); !ok {
			t.Fatalf("validate_records diagnostics is not an array: %s", text)
		}
	}
}

func assertSuggestNextRecordSmokeResponse(t *testing.T, repoRoot string, response map[string]any) {
	t.Helper()
	text := assertToolResultText(t, response, false)
	var payload map[string]any
	unmarshalJSONText(t, text, &payload)
	for _, key := range []string{"next_id", "next_number", "suggested_path"} {
		if payload[key] == nil {
			t.Fatalf("suggest_next_record missing %s: %s", key, text)
		}
	}
	suggested, ok := payload["suggested_path"].(string)
	if !ok || suggested == "" {
		t.Fatalf("suggest_next_record suggested_path invalid: %s", text)
	}
	if _, err := os.Stat(filepath.Join(repoRoot, filepath.FromSlash(suggested))); !os.IsNotExist(err) {
		t.Fatalf("suggest_next_record suggested path exists or stat failed: path=%s err=%v", suggested, err)
	}
}

func assertToolErrorSmokeResponse(t *testing.T, response map[string]any, wantCode string) {
	t.Helper()
	text := assertToolResultText(t, response, true)
	var payload map[string]map[string]any
	unmarshalJSONText(t, text, &payload)
	if payload["error"]["code"] != wantCode {
		t.Fatalf("tool error code = %#v, want %s; text=%s", payload["error"]["code"], wantCode, text)
	}
	if payload["error"]["message"] == "" {
		t.Fatalf("tool error message missing: %s", text)
	}
}

func assertToolResultText(t *testing.T, response map[string]any, wantError bool) string {
	t.Helper()
	if response == nil || response["error"] != nil {
		t.Fatalf("tools/call protocol error: %#v", response)
	}
	result := response["result"].(map[string]any)
	if result["isError"] != wantError {
		t.Fatalf("isError = %#v, want %v; response=%#v", result["isError"], wantError, response)
	}
	content := result["content"].([]any)
	if len(content) != 1 {
		t.Fatalf("content len = %d, want 1; response=%#v", len(content), response)
	}
	first := content[0].(map[string]any)
	if first["type"] != "text" {
		t.Fatalf("content[0].type = %#v, want text", first["type"])
	}
	text, ok := first["text"].(string)
	if !ok || !json.Valid([]byte(text)) {
		t.Fatalf("content[0].text is not valid JSON string: %#v", first["text"])
	}
	return text
}

func unmarshalJSONText(t *testing.T, text string, out any) {
	t.Helper()
	if err := json.Unmarshal([]byte(text), out); err != nil {
		t.Fatalf("unmarshal tool text: %v\n%s", err, text)
	}
}
