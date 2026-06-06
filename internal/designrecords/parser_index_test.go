package designrecords

import (
	"context"
	"os"
	"path/filepath"
	"sort"
	"strings"
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
	assertStrings(t, record.Decision.DependsOn, []string{"ADR-050", "ADR-068"})
	assertStrings(t, record.Decision.Supersedes, []string{"ADR-001", "ADR-002"})
	if record.Decision.MigratedToSpec == nil || *record.Decision.MigratedToSpec != "2026-05-12" {
		t.Fatalf("MigratedToSpec = %#v", record.Decision.MigratedToSpec)
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
	assertStrings(t, record.Decision.DependsOn, []string{})
	assertStrings(t, record.Decision.Supersedes, []string{})
	if record.Decision.MigratedToSpec == nil || *record.Decision.MigratedToSpec != "tomorrow" {
		t.Fatalf("MigratedToSpec = %#v", record.Decision.MigratedToSpec)
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
	if record.Decision.MigratedToSpec != nil {
		t.Fatalf("empty MigratedToSpec = %#v, want nil", record.Decision.MigratedToSpec)
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
	assertStrings(t, record.Spec.DependsOn, []string{"ADR-076"})
	if record.Decision != nil {
		t.Fatalf("Decision = %#v, want nil", record.Decision)
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

func TestInvestigationRecordParserIssuesAndMetadata(t *testing.T) {
	raw := "# INV-DOCS-001: investigation artifact format and lifecycle\n\n" +
		"- **status**: concluded\n" +
		"- **date**: 2026-05-19\n" +
		"- **trigger**: ADR-085\n" +
		"- **scope**: investigation format\n" +
		"- **non_scope**: writer tools\n" +
		"- **source_refs**:\n" +
		"  - ADR-085\n" +
		"  - spec:trace.resolve-and-validation\n" +
		"- **follow_up_candidates**:\n" +
		"  - ADR-086\n" +
		"- **follow_up_results**:\n" +
		"  - ADR-087\n" +
		"## Stop\n" +
		"  - ADR-999\n"
	record, candidate, issues := parseInvestigationRecord("docs/investigations/docs/INV-DOCS-001-investigation-artifact-format-and-lifecycle.md", raw)
	if record == nil {
		t.Fatal("record is nil")
	}
	if record.ID != "INV-DOCS-001" || record.Kind != RecordKindInvestigation || record.Title != "investigation artifact format and lifecycle" || record.Status != RecordStatusConcluded {
		t.Fatalf("record = %#v", record)
	}
	if record.Investigation == nil {
		t.Fatalf("Investigation detail missing: %#v", record)
	}
	if record.Investigation.Trigger != "ADR-085" || record.Investigation.Scope != "investigation format" || record.Investigation.NonScope != "writer tools" {
		t.Fatalf("investigation scalar metadata = %#v", record.Investigation)
	}
	assertStrings(t, record.Investigation.SourceRefs, []string{"ADR-085", "spec:trace.resolve-and-validation"})
	assertStrings(t, record.Investigation.FollowUpCandidates, []string{"ADR-086"})
	assertStrings(t, record.Investigation.FollowUpResults, []string{"ADR-087"})
	if candidate.FilenameIDMismatch || len(issues) != 0 {
		t.Fatalf("candidate/issues = %#v %#v", candidate, issues)
	}
}

func TestInvestigationInvalidH1AndFilenameMismatch(t *testing.T) {
	record, candidate, issues := parseInvestigationRecord("docs/investigations/docs/INV-DOCS-001-valid.md", "# INV-docs-001: invalid\n")
	if record != nil {
		t.Fatalf("record = %#v, want nil", record)
	}
	if candidate.Included || candidate.SkipReason != "invalid_investigation_h1" {
		t.Fatalf("candidate = %#v", candidate)
	}
	if !hasIssue(issues, DiagnosticInvalidH1Title) {
		t.Fatalf("missing invalid H1 issue: %#v", issues)
	}

	record, candidate, issues = parseInvestigationRecord("docs/investigations/docs/INV-DOCS-002-mismatch.md", "# INV-DOCS-001: mismatch\n- **status**: concluded\n")
	if record == nil {
		t.Fatal("record is nil")
	}
	if !candidate.FilenameIDMismatch || !hasIssue(issues, DiagnosticFilenameIDMismatch) {
		t.Fatalf("missing filename mismatch: candidate=%#v issues=%#v", candidate, issues)
	}
}

func TestWorkflowRecordParsersMetadataAndDiagnostics(t *testing.T) {
	reqRaw := "# REQ-MCP-003: Workflow support\n\n" +
		"- **id**: REQ-MCP-003\n" +
		"- **status**: accepted\n" +
		"- **date**: 2026-05-25\n" +
		"- **source_refs**:\n" +
		"  - ADR-091\n" +
		"- **work_items**:\n" +
		"  - WORK-MCP-003\n"
	req, candidate, issues := parseRequirementRecord("docs/requirements/mcp/REQ-MCP-003-workflow-support.md", reqRaw)
	if req == nil || candidate.FilenameIDMismatch || len(issues) != 0 {
		t.Fatalf("requirement parse = %#v candidate=%#v issues=%#v", req, candidate, issues)
	}
	if req.Kind != RecordKindRequirement || req.ID != "REQ-MCP-003" || req.Title != "Workflow support" || req.Status != RecordStatusAccepted {
		t.Fatalf("requirement record = %#v", req)
	}
	assertStrings(t, req.Requirement.SourceRefs, []string{"ADR-091"})
	assertStrings(t, req.Requirement.WorkItems, []string{"WORK-MCP-003"})

	workRaw := "# WORK-MCP-003: Workflow implementation\n\n" +
		"- **id**: WORK-MCP-003\n" +
		"- **status**: in_progress\n" +
		"- **date**: 2026-05-26\n" +
		"- **source_requirement**: REQ-MCP-003\n" +
		"- **impact_refs**:\n" +
		"  - ADR-092\n" +
		"- **tasks**:\n" +
		"  - TASK-MCP-003-04\n"
	work, candidate, issues := parseWorkItemRecord("docs/work-items/mcp/WORK-MCP-003-workflow-implementation.md", workRaw)
	if work == nil || candidate.FilenameIDMismatch || len(issues) != 0 {
		t.Fatalf("work item parse = %#v candidate=%#v issues=%#v", work, candidate, issues)
	}
	if work.Kind != RecordKindWorkItem || work.ID != "WORK-MCP-003" || work.Status != RecordStatusInProgress {
		t.Fatalf("work item record = %#v", work)
	}
	if work.WorkItem.SourceRequirement != "REQ-MCP-003" {
		t.Fatalf("source requirement = %q", work.WorkItem.SourceRequirement)
	}
	assertStrings(t, work.WorkItem.ImpactRefs, []string{"ADR-092"})
	assertStrings(t, work.WorkItem.Tasks, []string{"TASK-MCP-003-04"})

	taskRaw := "# TASK-MCP-003-04: Implement workflow records\n\n" +
		"- **id**: TASK-MCP-003-04\n" +
		"- **status**: in_progress\n" +
		"- **date**: 2026-05-26\n" +
		"- **work_item**: WORK-MCP-003\n" +
		"- **source_requirement**: REQ-MCP-003\n" +
		"- **estimate**: 1.5d\n" +
		"- **depends_on**:\n" +
		"- **outputs**:\n" +
		"  - implementation\n"
	task, candidate, issues := parseTaskRecord("docs/tasks/mcp/TASK-MCP-003-04-implement-workflow-records.md", taskRaw)
	if task == nil || candidate.FilenameIDMismatch || len(issues) != 0 {
		t.Fatalf("task parse = %#v candidate=%#v issues=%#v", task, candidate, issues)
	}
	if task.Kind != RecordKindTask || task.ID != "TASK-MCP-003-04" || task.Status != RecordStatusInProgress {
		t.Fatalf("task record = %#v", task)
	}
	if task.Task.WorkItem != "WORK-MCP-003" || task.Task.SourceRequirement != "REQ-MCP-003" || task.Task.Estimate != "1.5d" {
		t.Fatalf("task scalar metadata = %#v", task.Task)
	}
	assertStrings(t, task.Task.DependsOn, []string{})
	assertStrings(t, task.Task.Outputs, []string{"implementation"})
}

func TestWorkflowRecordParserIssues(t *testing.T) {
	record, candidate, issues := parseRequirementRecord("docs/requirements/mcp/REQ-MCP-003-valid.md", "# REQ-MCP-003 missing colon\n- **id**: REQ-MCP-003\n")
	if record != nil {
		t.Fatalf("record = %#v, want nil", record)
	}
	if candidate.Included || candidate.SkipReason != "invalid_workflow_h1" || !hasIssue(issues, DiagnosticInvalidH1Title) {
		t.Fatalf("invalid H1 candidate/issues = %#v %#v", candidate, issues)
	}

	record, candidate, issues = parseRequirementRecord("docs/requirements/mcp/REQ-MCP-003-mismatch.md", "# REQ-MCP-003: Mismatch\n- **id**: REQ-MCP-004\n- **status**: accepted\n")
	if record == nil {
		t.Fatal("record is nil")
	}
	if !candidate.FilenameIDMismatch || !hasIssue(issues, DiagnosticFilenameIDMismatch) {
		t.Fatalf("metadata mismatch candidate/issues = %#v %#v", candidate, issues)
	}

	task, candidate, issues := parseTaskRecord("docs/tasks/mcp/TASK-mcp-003-04-invalid.md", "# TASK-mcp-003-04: Invalid\n- **id**: TASK-mcp-003-04\n- **status**: todo\n")
	if task != nil {
		t.Fatalf("task = %#v, want nil", task)
	}
	if candidate.Included || candidate.SkipReason != "invalid_workflow_id" || !hasIssue(issues, DiagnosticInvalidWorkflowID) {
		t.Fatalf("invalid workflow ID candidate/issues = %#v %#v", candidate, issues)
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
	writeTestFile(t, root, "docs/investigations/docs/INV-DOCS-001-test.md", "# INV-DOCS-001: Test investigation\n- **status**: concluded\n- **date**: 2026-05-19\n- **trigger**: ADR-001\n- **scope**: test\n- **non_scope**: none\n- **source_refs**:\n  - ADR-001\n- **follow_up_candidates**:\n  - SPEC-test\n")
	writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-003-test.md", "# REQ-MCP-003: Test requirement\n- **id**: REQ-MCP-003\n- **status**: accepted\n- **date**: 2026-05-25\n- **source_refs**:\n  - ADR-001\n- **work_items**:\n  - WORK-MCP-003\n")
	writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-003-test.md", "# WORK-MCP-003: Test work item\n- **id**: WORK-MCP-003\n- **status**: implementation_pending\n- **date**: 2026-05-26\n- **source_requirement**: REQ-MCP-003\n- **impact_refs**:\n  - ADR-001\n- **tasks**:\n  - TASK-MCP-003-01\n")
	writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-003-01-test.md", "# TASK-MCP-003-01: Test task\n- **id**: TASK-MCP-003-01\n- **status**: todo\n- **date**: 2026-05-26\n- **work_item**: WORK-MCP-003\n- **source_requirement**: REQ-MCP-003\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n  - test\n")
	writeTestFile(t, root, "docs/tasks/m17-legacy.md", "# Legacy task\n")

	cfg, err := NewConfig(root)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	if len(idx.Records) != 6 {
		t.Fatalf("records = %#v, want 6", idx.Records)
	}
	if got := recordIDs(idx.Records); !sameStrings(got, []string{"ADR-001", "INV-DOCS-001", "REQ-MCP-003", "SPEC-test", "TASK-MCP-003-01", "WORK-MCP-003"}) {
		t.Fatalf("record IDs = %#v", got)
	}
	spec := findRecord(idx.Records, "SPEC-test")
	if spec == nil || spec.RawBody != rawSpec {
		t.Fatalf("raw body was not preserved: %#v", spec)
	}
	investigation := findRecord(idx.Records, "INV-DOCS-001")
	if investigation == nil || investigation.Kind != RecordKindInvestigation || investigation.Investigation == nil {
		t.Fatalf("investigation not indexed: %#v", investigation)
	}
	if legacy := findRecord(idx.Records, "m17"); legacy != nil {
		t.Fatalf("legacy task was indexed as workflow task: %#v", legacy)
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
	for _, id := range []string{"INV-DOCS-001", "INV-DOCS-002", "INV-DOCS-003"} {
		record := findRecord(idx.Records, id)
		if record == nil {
			t.Fatalf("missing %s in repository index", id)
		}
		if record.Kind != RecordKindInvestigation {
			t.Fatalf("%s kind = %q, want investigation", id, record.Kind)
		}
	}
	for _, tt := range []struct {
		id   string
		kind RecordKind
	}{
		{"REQ-MCP-003", RecordKindRequirement},
		{"WORK-MCP-003", RecordKindWorkItem},
		{"TASK-MCP-003-01", RecordKindTask},
		{"TASK-MCP-003-03", RecordKindTask},
		{"TASK-MCP-003-04", RecordKindTask},
	} {
		record := findRecord(idx.Records, tt.id)
		if record == nil {
			t.Fatalf("missing %s in repository index", tt.id)
		}
		if record.Kind != tt.kind {
			t.Fatalf("%s kind = %q, want %q", tt.id, record.Kind, tt.kind)
		}
	}
	for _, record := range idx.Records {
		if record.Kind == RecordKindTask && isLegacyMSeriesTaskPath(record.Path) {
			t.Fatalf("legacy M-series task was indexed: %#v", record)
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

func isLegacyMSeriesTaskPath(path string) bool {
	prefix := "docs/tasks/m"
	if !strings.HasPrefix(path, prefix) || len(path) <= len(prefix) {
		return false
	}
	next := path[len(prefix)]
	return next >= '0' && next <= '9'
}
