package designrecords

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestADRH1Parser(t *testing.T) {
	tests := []struct {
		name  string
		line  string
		valid bool
		num   string
		title string
	}{
		{name: "valid", line: "# 076: Design Records MCP", valid: true, num: "076", title: "Design Records MCP"},
		{name: "adr prefix invalid", line: "# ADR-076: Design Records MCP"},
		{name: "not zero padded invalid", line: "# 76: Design Records MCP"},
		{name: "non ascii colon invalid", line: "# 076： Design Records MCP"},
		{name: "missing whitespace after colon invalid", line: "# 076:Design Records MCP"},
		{name: "empty trimmed title invalid", line: "# 076:   "},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			num, title, valid := parseADRH1(tt.line)
			if valid != tt.valid {
				t.Fatalf("valid = %v, want %v", valid, tt.valid)
			}
			if num != tt.num || title != tt.title {
				t.Fatalf("num/title = %q/%q, want %q/%q", num, title, tt.num, tt.title)
			}
		})
	}
}

func TestADRRecordParserIssuesAndMetadata(t *testing.T) {
	raw := "# 076: Design Records MCP\n\n" +
		"- **status**: accepted\n" +
		"- **date**: 2026-05-11\n" +
		"- **depends_on**: ADR-050, ADR-068\n" +
		"- **supersedes**: ADR-001, ADR-002\n" +
		"- **migrated_to_spec**: 2026-05-12\n" +
		"- **unknown**: ignored\n" +
		"## Stop\n" +
		"- **status**: proposed\n"
	record, candidate, issues := parseADRRecord("docs/adr/075-wrong.md", raw)
	if record == nil {
		t.Fatal("record is nil")
	}
	if record.ID != "ADR-076" || record.Title != "Design Records MCP" || record.Status != RecordStatusAccepted {
		t.Fatalf("record = %#v", record)
	}
	assertStrings(t, record.DependsOn, []string{"ADR-050", "ADR-068"})
	assertStrings(t, record.Supersedes, []string{"ADR-001", "ADR-002"})
	if record.MigratedToSpec == nil || *record.MigratedToSpec != "2026-05-12" {
		t.Fatalf("MigratedToSpec = %#v", record.MigratedToSpec)
	}
	if !candidate.FilenameIDMismatch {
		t.Fatal("FilenameIDMismatch = false, want true")
	}
	if !hasIssue(issues, DiagnosticFilenameIDMismatch) {
		t.Fatalf("missing filename mismatch issue: %#v", issues)
	}
}

func TestADRInvalidH1RemainsCandidateWithoutFilenameDerivedID(t *testing.T) {
	record, candidate, issues := parseADRRecord("docs/adr/076-design-records-mcp.md", "# ADR-076: Design Records MCP\n")
	if record != nil {
		t.Fatalf("record = %#v, want nil", record)
	}
	if candidate.ID != "" || candidate.NormalizedID != "" {
		t.Fatalf("candidate derived ID from filename: %#v", candidate)
	}
	if candidate.Included || candidate.SkipReason != "invalid_adr_h1" {
		t.Fatalf("candidate inclusion = %#v", candidate)
	}
	if !hasIssue(issues, DiagnosticInvalidH1Title) {
		t.Fatalf("missing invalid H1 issue: %#v", issues)
	}
}

func TestADRMetadataEmptyValuesAndInvalidMigratedToSpec(t *testing.T) {
	raw := "# 076: Design Records MCP\n" +
		"- **status**: accepted\n" +
		"- **depends_on**: \n" +
		"- **supersedes**:   \n" +
		"- **migrated_to_spec**: tomorrow\n" +
		"> stop\n" +
		"- **depends_on**: ADR-999\n"
	record, _, issues := parseADRRecord("docs/adr/076-design-records-mcp.md", raw)
	if record == nil {
		t.Fatal("record is nil")
	}
	assertStrings(t, record.DependsOn, []string{})
	assertStrings(t, record.Supersedes, []string{})
	if record.MigratedToSpec == nil || *record.MigratedToSpec != "tomorrow" {
		t.Fatalf("MigratedToSpec = %#v", record.MigratedToSpec)
	}
	if !hasIssue(issues, DiagnosticInvalidMigratedToSpec) {
		t.Fatalf("missing invalid migrated_to_spec issue: %#v", issues)
	}

	record, _, issues = parseADRRecord("docs/adr/076-design-records-mcp.md", "# 076: Design Records MCP\n- **migrated_to_spec**: 2026-02-31\n")
	if record == nil {
		t.Fatal("record is nil")
	}
	if !hasIssue(issues, DiagnosticInvalidMigratedToSpec) {
		t.Fatalf("calendar-invalid migrated_to_spec should produce issue: %#v", issues)
	}

	record, _, issues = parseADRRecord("docs/adr/076-design-records-mcp.md", "# 076: Design Records MCP\n- **migrated_to_spec**: \n")
	if record.MigratedToSpec != nil {
		t.Fatalf("empty MigratedToSpec = %#v, want nil", record.MigratedToSpec)
	}
	if hasIssue(issues, DiagnosticInvalidMigratedToSpec) {
		t.Fatalf("empty migrated_to_spec should not produce issue: %#v", issues)
	}
}

func TestSpecRecordParser(t *testing.T) {
	raw := "---\n" +
		"status: draft\n" +
		"depends_on:\n" +
		"  - docs/adr/999-path-only.md\n" +
		"design_record:\n" +
		"  id: SPEC-design-records-mcp-schema\n" +
		"  kind: spec\n" +
		"  status: confirmed\n" +
		"  depends_on:\n" +
		"    - ADR-076\n" +
		"  supersedes:\n" +
		"    - ADR-001\n" +
		"  migrated_to_spec: 2026-05-12\n" +
		"---\n\n" +
		"# Design Records MCP schema\n"
	record, candidate, issues := parseSpecRecord("docs/spec/design-records-mcp/schema.md", raw)
	if record == nil {
		t.Fatal("record is nil")
	}
	if candidate.ID != record.ID || candidate.NormalizedID != "SPEC-DESIGN-RECORDS-MCP-SCHEMA" {
		t.Fatalf("candidate = %#v", candidate)
	}
	if record.Status != RecordStatusDraft {
		t.Fatalf("Status = %q, want draft", record.Status)
	}
	assertStrings(t, record.DependsOn, []string{"ADR-076"})
	assertStrings(t, record.Supersedes, []string{})
	if record.MigratedToSpec != nil {
		t.Fatalf("MigratedToSpec = %#v, want nil", record.MigratedToSpec)
	}
	if !hasIssue(issues, DiagnosticSpecStatusMismatch) {
		t.Fatalf("missing status mismatch issue: %#v", issues)
	}
}

func TestSpecSilentSkips(t *testing.T) {
	withoutDesignRecord := "---\nstatus: draft\n---\n# Existing spec\n"
	record, candidate, issues := parseSpecRecord("docs/spec/existing.md", withoutDesignRecord)
	if record != nil || candidate.Path != "" || len(issues) != 0 {
		t.Fatalf("design_record absent should silent skip: record=%#v candidate=%#v issues=%#v", record, candidate, issues)
	}

	invalidKind := "---\nstatus: draft\ndesign_record:\n  id: TASK-001\n  kind: task\n---\n# Task\n"
	record, candidate, issues = parseSpecRecord("docs/spec/task.md", invalidKind)
	if record != nil || len(issues) != 0 {
		t.Fatalf("invalid kind should silent skip record/issues: record=%#v issues=%#v", record, issues)
	}
	if candidate.Path == "" || candidate.Included || candidate.SkipReason != "unsupported_design_record_kind" {
		t.Fatalf("invalid kind candidate = %#v", candidate)
	}
}

func TestSpecInvalidH1Issue(t *testing.T) {
	raw := "---\nstatus: draft\ndesign_record:\n  id: SPEC-x\n  kind: spec\n---\n\n## Not H1\n"
	record, _, issues := parseSpecRecord("docs/spec/x.md", raw)
	if record == nil {
		t.Fatal("record is nil")
	}
	if record.Title != "" {
		t.Fatalf("Title = %q, want empty", record.Title)
	}
	if !hasIssue(issues, DiagnosticInvalidH1Title) {
		t.Fatalf("missing invalid H1 issue: %#v", issues)
	}
}

func TestHeadingsExtractionExcludesFrontMatterAndFences(t *testing.T) {
	raw := "---\nsummary: '# not a heading'\n---\n" +
		"# Title\n" +
		"```yaml\n" +
		"# not a heading\n" +
		"```\n" +
		"## Section\n" +
		"Setext\n---\n"
	headings := extractHeadings(raw)
	if len(headings) != 2 {
		t.Fatalf("headings = %#v, want 2", headings)
	}
	if headings[0] != (Heading{Level: 1, Text: "Title"}) || headings[1] != (Heading{Level: 2, Text: "Section"}) {
		t.Fatalf("headings = %#v", headings)
	}
}

func TestBuildIndexDiscoversRecordsAndPreservesRawBody(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/001-test.md", "# 001: Test ADR\n- **status**: accepted\n- **depends_on**:\n")
	writeTestFile(t, root, "docs/adr/nested/002-skip.md", "# 002: Skip\n- **status**: accepted\n")
	rawSpec := "---\nstatus: draft\ndesign_record:\n  id: SPEC-test\n  kind: spec\n  depends_on:\n    - ADR-001\n---\n# Test spec\n"
	writeTestFile(t, root, "docs/spec/test.md", rawSpec)
	writeTestFile(t, root, "docs/spec/existing.md", "---\nstatus: draft\n---\n# Existing\n")

	cfg, err := NewConfig(root)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	if len(idx.Records) != 2 {
		t.Fatalf("records = %#v, want 2", idx.Records)
	}
	if got := recordIDs(idx.Records); !sameStrings(got, []string{"ADR-001", "SPEC-test"}) {
		t.Fatalf("record IDs = %#v", got)
	}
	spec := findRecord(idx.Records, "SPEC-test")
	if spec == nil || spec.RawBody != rawSpec {
		t.Fatalf("raw body was not preserved: %#v", spec)
	}
}

func TestBuildIndexRepositoryBootstrapRecords(t *testing.T) {
	root := findRepoRoot(t)
	cfg, err := NewConfig(root)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	for _, id := range []string{"ADR-050", "ADR-067", "ADR-068", "ADR-069", "ADR-070", "ADR-071", "ADR-072", "ADR-073", "ADR-074", "ADR-075", "ADR-076", "ADR-077"} {
		record := findRecord(idx.Records, id)
		if record == nil {
			t.Fatalf("missing %s in repository index", id)
		}
		if record.Kind != RecordKindDecision {
			t.Fatalf("%s kind = %q, want decision", id, record.Kind)
		}
	}
	for _, id := range []string{"SPEC-design-records-mcp-overview", "SPEC-design-records-mcp-schema", "SPEC-design-records-mcp-tools"} {
		record := findRecord(idx.Records, id)
		if record == nil {
			t.Fatalf("missing %s in repository index", id)
		}
		if record.Kind != RecordKindSpec {
			t.Fatalf("%s kind = %q, want spec", id, record.Kind)
		}
	}
	if findRecord(idx.Records, "docs/spec/overview.md") != nil {
		t.Fatal("design_record-less existing spec was indexed")
	}
}

func writeTestFile(t *testing.T, root, rel, content string) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(rel))
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile: %v", err)
	}
}

func hasIssue(issues []ParseIssue, category DiagnosticCategory) bool {
	for _, issue := range issues {
		if issue.Category == category {
			return true
		}
	}
	return false
}

func assertStrings(t *testing.T, got, want []string) {
	t.Helper()
	if !sameStrings(got, want) {
		t.Fatalf("strings = %#v, want %#v", got, want)
	}
}

func sameStrings(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	for i := range got {
		if got[i] != want[i] {
			return false
		}
	}
	return true
}

func recordIDs(records []Record) []string {
	ids := make([]string, 0, len(records))
	for _, record := range records {
		ids = append(ids, record.ID)
	}
	sort.Strings(ids)
	return ids
}

func findRecord(records []Record, id string) *Record {
	for i := range records {
		if records[i].ID == id {
			return &records[i]
		}
	}
	return nil
}

func findRepoRoot(t *testing.T) string {
	t.Helper()
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd: %v", err)
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		next := filepath.Dir(dir)
		if next == dir {
			t.Fatal("repo root not found")
		}
		dir = next
	}
}
