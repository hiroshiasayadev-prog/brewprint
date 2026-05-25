package designrecords

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestSuggestNextRecordBasicDecisionResponse(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/076-design-records-mcp.md", "# 076: Design Records MCP\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/adr/077-boundary.md", "# 077: Boundary\n- **status**: accepted\n")
	idx := buildTestIndex(t, root)

	resp, err := SuggestNextRecord(context.Background(), idx, SuggestNextRecordRequest{
		Kind:  RecordKindDecision,
		Title: " Design Records MCP implementation package layout ",
	})
	if err != nil {
		t.Fatalf("SuggestNextRecord: %v", err)
	}
	if resp.Kind != RecordKindDecision {
		t.Fatalf("Kind = %q, want decision", resp.Kind)
	}
	if resp.Title != "Design Records MCP implementation package layout" {
		t.Fatalf("Title = %q, want trimmed title", resp.Title)
	}
	if resp.NextID != "ADR-078" || resp.NextNumber != 78 || resp.ExistingMaxID != "ADR-077" {
		t.Fatalf("numbering response = %#v", resp)
	}
	wantPath := "docs/adr/078-design-records-mcp-implementation-package-layout.md"
	if resp.SuggestedPath != wantPath {
		t.Fatalf("SuggestedPath = %q, want %q", resp.SuggestedPath, wantPath)
	}
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(resp.SuggestedPath))); !os.IsNotExist(err) {
		t.Fatalf("suggested path stat error = %v, want file to be absent", err)
	}

	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal response: %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(encoded, &raw); err != nil {
		t.Fatalf("Unmarshal response: %v", err)
	}
	for _, field := range []string{"kind", "title", "next_id", "next_number", "suggested_path", "existing_max_id"} {
		if _, ok := raw[field]; !ok {
			t.Fatalf("response missing %q: %s", field, encoded)
		}
	}
	for _, unexpected := range []string{"body", "headings", "diagnostics", "candidates", "records"} {
		if _, ok := raw[unexpected]; ok {
			t.Fatalf("response unexpectedly includes %q: %s", unexpected, encoded)
		}
	}
}

func TestSuggestNextRecordNumberingRules(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/001-one.md", "# 001: One\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/adr/077-seventy-seven.md", "# 077: Seventy seven\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/adr/099-ninety-nine.md", "# 099: Ninety nine\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/adr/100-one-hundred.md", "# 100: One hundred\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/adr/999-invalid-h1.md", "# ADR-999: Invalid H1\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/spec/design-records-mcp/schema.md", "---\nstatus: confirmed\ndesign_record:\n  id: SPEC-design-records-mcp-schema\n  kind: spec\n  status: confirmed\n---\n# Design Records MCP schema\n")
	idx := buildTestIndex(t, root)
	idx.Records = append(idx.Records,
		Record{ID: "ADR-999", Kind: RecordKindSpec, Title: "Spec-shaped ADR ID"},
		Record{ID: "ADR-101-extra", Kind: RecordKindDecision, Title: "Malformed decision ID"},
	)

	resp, err := SuggestNextRecord(context.Background(), idx, SuggestNextRecordRequest{
		Kind:  RecordKindDecision,
		Title: "Next",
	})
	if err != nil {
		t.Fatalf("SuggestNextRecord: %v", err)
	}
	if resp.NextID != "ADR-101" || resp.NextNumber != 101 || resp.ExistingMaxID != "ADR-100" {
		t.Fatalf("numbering response = %#v, want max ADR-100 plus one", resp)
	}
	if resp.SuggestedPath != "docs/adr/101-next.md" {
		t.Fatalf("SuggestedPath = %q", resp.SuggestedPath)
	}
}

func TestSuggestNextRecordNoExistingDecisionsStartsAtADR001(t *testing.T) {
	idx := &Index{Records: []Record{
		{ID: "SPEC-design-records-mcp-schema", Kind: RecordKindSpec},
	}}

	resp, err := SuggestNextRecord(context.Background(), idx, SuggestNextRecordRequest{
		Kind:  RecordKindDecision,
		Title: "First decision",
	})
	if err != nil {
		t.Fatalf("SuggestNextRecord: %v", err)
	}
	if resp.NextID != "ADR-001" || resp.NextNumber != 1 || resp.ExistingMaxID != "" {
		t.Fatalf("response = %#v, want first ADR with empty existing_max_id", resp)
	}
	if resp.SuggestedPath != "docs/adr/001-first-decision.md" {
		t.Fatalf("SuggestedPath = %q", resp.SuggestedPath)
	}

	encoded, err := json.Marshal(resp)
	if err != nil {
		t.Fatalf("Marshal response: %v", err)
	}
	var raw map[string]any
	if err := json.Unmarshal(encoded, &raw); err != nil {
		t.Fatalf("Unmarshal response: %v", err)
	}
	if _, ok := raw["existing_max_id"]; ok {
		t.Fatalf("existing_max_id should be omitted when empty: %s", encoded)
	}
}

func TestSuggestNextRecordSlugGeneration(t *testing.T) {
	idx := &Index{}
	tests := []struct {
		name  string
		title string
		want  string
	}{
		{
			name:  "ascii title",
			title: "Design Records MCP implementation package layout",
			want:  "docs/adr/001-design-records-mcp-implementation-package-layout.md",
		},
		{
			name:  "mixed punctuation",
			title: "ADR: Foo / Bar",
			want:  "docs/adr/001-adr-foo-bar.md",
		},
		{
			name:  "repeated separators",
			title: "Hello --- World___Again",
			want:  "docs/adr/001-hello-world-again.md",
		},
		{
			name:  "leading and trailing separators",
			title: " / Hello, World! / ",
			want:  "docs/adr/001-hello-world.md",
		},
		{
			name:  "uppercase lowercasing",
			title: "UPPER Case 123",
			want:  "docs/adr/001-upper-case-123.md",
		},
		{
			name:  "non ascii only",
			title: "日本語タイトル",
			want:  "docs/adr/001.md",
		},
		{
			name:  "mixed ascii and non ascii",
			title: "Hello 日本語 World",
			want:  "docs/adr/001-hello-world.md",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := SuggestNextRecord(context.Background(), idx, SuggestNextRecordRequest{
				Kind:  RecordKindDecision,
				Title: tt.title,
			})
			if err != nil {
				t.Fatalf("SuggestNextRecord: %v", err)
			}
			if resp.SuggestedPath != tt.want {
				t.Fatalf("SuggestedPath = %q, want %q", resp.SuggestedPath, tt.want)
			}
		})
	}
}

func TestSuggestNextRecordRequestErrors(t *testing.T) {
	idx := &Index{}
	tests := []struct {
		name string
		req  SuggestNextRecordRequest
		code ErrorCode
	}{
		{
			name: "missing kind",
			req:  SuggestNextRecordRequest{Title: "Title"},
			code: ErrorCodeInvalidRequest,
		},
		{
			name: "empty kind",
			req:  SuggestNextRecordRequest{Kind: "", Title: "Title"},
			code: ErrorCodeInvalidRequest,
		},
		{
			name: "whitespace kind",
			req:  SuggestNextRecordRequest{Kind: RecordKind(" \t"), Title: "Title"},
			code: ErrorCodeInvalidRequest,
		},
		{
			name: "spec kind",
			req:  SuggestNextRecordRequest{Kind: RecordKindSpec, Title: "Title"},
			code: ErrorCodeUnsupportedKind,
		},
		{
			name: "unknown kind",
			req:  SuggestNextRecordRequest{Kind: RecordKind("task"), Title: "Title"},
			code: ErrorCodeUnsupportedKind,
		},
		{
			name: "missing title",
			req:  SuggestNextRecordRequest{Kind: RecordKindDecision},
			code: ErrorCodeInvalidRequest,
		},
		{
			name: "empty title",
			req:  SuggestNextRecordRequest{Kind: RecordKindDecision, Title: ""},
			code: ErrorCodeInvalidRequest,
		},
		{
			name: "whitespace title",
			req:  SuggestNextRecordRequest{Kind: RecordKindDecision, Title: " \t"},
			code: ErrorCodeInvalidRequest,
		},
		{
			name: "nil index",
			req:  SuggestNextRecordRequest{Kind: RecordKindDecision, Title: "Title"},
			code: ErrorCodeInvalidRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index := idx
			if tt.name == "nil index" {
				index = nil
			}
			resp, err := SuggestNextRecord(context.Background(), index, tt.req)
			if err == nil {
				t.Fatal("SuggestNextRecord error = nil, want ToolError")
			}
			toolErr, ok := err.(*ToolError)
			if !ok {
				t.Fatalf("error = %T %v, want *ToolError", err, err)
			}
			if toolErr.Code != tt.code {
				t.Fatalf("error code = %q, want %q", toolErr.Code, tt.code)
			}
			if resp != (SuggestNextRecordResponse{}) {
				t.Fatalf("error response = %#v, want zero response", resp)
			}
		})
	}
}

func TestSuggestNextRecordNoSideEffects(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/077-boundary.md", "# 077: Boundary\n- **status**: accepted\n")
	idx := buildTestIndex(t, root)
	before := cloneIndexForSuggestTest(idx)

	resp, err := SuggestNextRecord(context.Background(), idx, SuggestNextRecordRequest{
		Kind:  RecordKindDecision,
		Title: "No side effect",
	})
	if err != nil {
		t.Fatalf("SuggestNextRecord: %v", err)
	}
	if !reflect.DeepEqual(idx, before) {
		t.Fatalf("index mutated:\nbefore=%#v\nafter=%#v", before, idx)
	}
	if _, err := os.Stat(filepath.Join(root, filepath.FromSlash(resp.SuggestedPath))); !os.IsNotExist(err) {
		t.Fatalf("suggested path stat error = %v, want file to be absent", err)
	}
}

func TestSuggestNextRecordRepositoryBootstrap(t *testing.T) {
	root := findRepoRoot(t)
	cfg, err := NewConfig(root)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}

	maxNum := 0
	maxID := ""
	for _, record := range idx.Records {
		if record.Kind != RecordKindDecision {
			continue
		}
		num, ok := decisionRecordNumber(record.ID)
		if !ok {
			continue
		}
		if num > maxNum {
			maxNum = num
			maxID = record.ID
		}
	}
	if maxNum == 0 {
		t.Fatal("repository index has no decision records")
	}

	resp, err := SuggestNextRecord(context.Background(), idx, SuggestNextRecordRequest{
		Kind:  RecordKindDecision,
		Title: "Repository bootstrap check",
	})
	if err != nil {
		t.Fatalf("SuggestNextRecord: %v", err)
	}
	if resp.NextNumber != maxNum+1 || resp.NextID != formatDecisionRecordID(maxNum+1) || resp.ExistingMaxID != maxID {
		t.Fatalf("response = %#v, want max %s plus one", resp, maxID)
	}
	wantPath := "docs/adr/" + formatDecisionRecordNumber(maxNum+1) + "-repository-bootstrap-check.md"
	if resp.SuggestedPath != wantPath {
		t.Fatalf("SuggestedPath = %q, want %q", resp.SuggestedPath, wantPath)
	}
}

func cloneIndexForSuggestTest(idx *Index) *Index {
	if idx == nil {
		return nil
	}
	clone := *idx
	clone.Records = append([]Record(nil), idx.Records...)
	for i := range clone.Records {
		if idx.Records[i].Decision != nil {
			clone.Records[i].Decision = &DecisionDetail{
				DependsOn:      cloneStringSliceForSuggestTest(idx.Records[i].Decision.DependsOn),
				Supersedes:     cloneStringSliceForSuggestTest(idx.Records[i].Decision.Supersedes),
				MigratedToSpec: idx.Records[i].Decision.MigratedToSpec,
			}
		}
		if idx.Records[i].Spec != nil {
			clone.Records[i].Spec = &SpecDetail{DependsOn: cloneStringSliceForSuggestTest(idx.Records[i].Spec.DependsOn)}
		}
		if idx.Records[i].Investigation != nil {
			clone.Records[i].Investigation = cloneInvestigationDetail(idx.Records[i].Investigation)
		}
		clone.Records[i].Headings = cloneHeadingsForSuggestTest(idx.Records[i].Headings)
	}
	clone.Diagnostics = cloneDiagnosticsForSuggestTest(idx.Diagnostics)
	clone.Candidates = cloneCandidatesForSuggestTest(idx.Candidates)
	clone.ParseIssues = cloneParseIssuesForSuggestTest(idx.ParseIssues)
	clone.PathIssues = clonePathIssuesForSuggestTest(idx.PathIssues)
	return &clone
}

func cloneStringSliceForSuggestTest(in []string) []string {
	if in == nil {
		return nil
	}
	out := make([]string, len(in))
	copy(out, in)
	return out
}

func cloneHeadingsForSuggestTest(in []Heading) []Heading {
	if in == nil {
		return nil
	}
	out := make([]Heading, len(in))
	copy(out, in)
	return out
}

func cloneDiagnosticsForSuggestTest(in []Diagnostic) []Diagnostic {
	if in == nil {
		return nil
	}
	out := make([]Diagnostic, len(in))
	copy(out, in)
	return out
}

func cloneCandidatesForSuggestTest(in []RecordCandidate) []RecordCandidate {
	if in == nil {
		return nil
	}
	out := make([]RecordCandidate, len(in))
	copy(out, in)
	return out
}

func cloneParseIssuesForSuggestTest(in []ParseIssue) []ParseIssue {
	if in == nil {
		return nil
	}
	out := make([]ParseIssue, len(in))
	copy(out, in)
	return out
}

func clonePathIssuesForSuggestTest(in []PathIssue) []PathIssue {
	if in == nil {
		return nil
	}
	out := make([]PathIssue, len(in))
	copy(out, in)
	return out
}

func formatDecisionRecordID(num int) string {
	return "ADR-" + formatDecisionRecordNumber(num)
}

func formatDecisionRecordNumber(num int) string {
	return fmt.Sprintf("%03d", num)
}
