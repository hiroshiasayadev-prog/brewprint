package designrecords

import (
	"os"
	"path/filepath"
	"testing"
)

// fixtureBase is the path to the read-baseline fixture tree, relative to the package.
const fixtureBase = "testdata/read-baseline"

// writeTestFile creates a file at filepath.Join(root, relPath) with the given content.
// Preserved for use by authoring_test.go and authoring_guidance_test.go.
func writeTestFile(t *testing.T, root, relPath, content string) {
	t.Helper()
	full := filepath.Join(root, filepath.FromSlash(relPath))
	if err := os.MkdirAll(filepath.Dir(full), 0o755); err != nil {
		t.Fatalf("MkdirAll %q: %v", filepath.Dir(full), err)
	}
	if err := os.WriteFile(full, []byte(content), 0o644); err != nil {
		t.Fatalf("WriteFile %q: %v", full, err)
	}
}

// findRepoRoot walks up from the current working directory until it finds go.mod.
// Preserved for use by integration tests in this package.
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
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatal("could not find repo root (go.mod not found in any parent directory)")
		}
		dir = parent
	}
}

// findRecord returns a pointer to the first record in records whose ID equals id, or nil.
// Preserved for use by integration tests in this package.
func findRecord(records []Record, id string) *Record {
	for i := range records {
		if records[i].ID == id {
			return &records[i]
		}
	}
	return nil
}

func readFixtureFile(t *testing.T, relPath string) string {
	t.Helper()
	data, err := os.ReadFile(filepath.FromSlash(relPath))
	if err != nil {
		t.Fatalf("ReadFile %q: %v", relPath, err)
	}
	return string(data)
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

// ── C01: PRODUCT-ADR-SPEC-901 ─────────────────────────────────────────────────

func TestCurrentADRRecordParser_C01(t *testing.T) {
	path := fixtureBase + "/current/product/records/adr/spec/PRODUCT-ADR-SPEC-901-current-read-fixture-baseline.md"
	raw := readFixtureFile(t, path)
	record, candidate, issues := parseCurrentADRRecord(path, raw, "PRODUCT-")

	if record == nil {
		t.Fatalf("record is nil; issues = %#v", issues)
	}
	if record.ID != "PRODUCT-ADR-SPEC-901" {
		t.Fatalf("ID = %q, want PRODUCT-ADR-SPEC-901", record.ID)
	}
	if record.Kind != RecordKindDecision {
		t.Fatalf("Kind = %q, want decision", record.Kind)
	}
	if record.Title != "Use shared current read fixtures" {
		t.Fatalf("Title = %q", record.Title)
	}
	if record.Status != RecordStatusAccepted {
		t.Fatalf("Status = %q, want accepted", record.Status)
	}
	if candidate.FilenameIDMismatch {
		t.Fatalf("FilenameIDMismatch = true, want false; candidate = %#v", candidate)
	}
	if len(issues) != 0 {
		t.Fatalf("unexpected issues: %#v", issues)
	}
	if record.Decision == nil {
		t.Fatal("Decision detail is nil")
	}
	assertStrings(t, record.Decision.DependsOn, []string{})
	assertStrings(t, record.Decision.Supersedes, []string{})
	if record.Decision.MigratedToSpec != nil {
		t.Fatalf("MigratedToSpec = %#v, want nil", record.Decision.MigratedToSpec)
	}
}

func TestCurrentADRH1NoRepair(t *testing.T) {
	// Bare ADR ID without app namespace prefix must be rejected (no repair).
	for _, tc := range []struct {
		name string
		line string
		ns   string
	}{
		{"missing ns prefix", "# ADR-SPEC-901: title", "PRODUCT-"},
		{"wrong ns prefix", "# DRMCP-ADR-SPEC-901: title", "PRODUCT-"},
		{"legacy bare number form", "# ADR-076: title", "PRODUCT-"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, _, ok := parseCurrentADRH1(tc.line, tc.ns)
			if ok {
				t.Fatalf("parseCurrentADRH1 accepted %q with ns %q, want rejection", tc.line, tc.ns)
			}
		})
	}
}

func TestCurrentADRH1ValidForms(t *testing.T) {
	for _, tc := range []struct {
		line   string
		ns     string
		bareID string
		title  string
	}{
		{"# PRODUCT-ADR-SPEC-901: Use shared current read fixtures", "PRODUCT-", "ADR-SPEC-901", "Use shared current read fixtures"},
		{"# DRMCP-ADR-MCP-001: Something", "DRMCP-", "ADR-MCP-001", "Something"},
	} {
		bareID, title, ok := parseCurrentADRH1(tc.line, tc.ns)
		if !ok {
			t.Fatalf("parseCurrentADRH1(%q, %q) = false, want true", tc.line, tc.ns)
		}
		if bareID != tc.bareID || title != tc.title {
			t.Fatalf("bareID/title = %q/%q, want %q/%q", bareID, title, tc.bareID, tc.title)
		}
	}
}

// ── C02: DRMCP-INV-MCP-901 ───────────────────────────────────────────────────

func TestCurrentInvestigationRecordParser_C02(t *testing.T) {
	path := fixtureBase + "/current/drmcp/records/investigations/mcp/DRMCP-INV-MCP-901-current-read-fixture-observations.md"
	raw := readFixtureFile(t, path)
	record, candidate, issues := parseCurrentInvestigationRecord(path, raw, "DRMCP-")

	if record == nil {
		t.Fatalf("record is nil; issues = %#v", issues)
	}
	if record.ID != "DRMCP-INV-MCP-901" {
		t.Fatalf("ID = %q, want DRMCP-INV-MCP-901", record.ID)
	}
	if record.Kind != RecordKindInvestigation {
		t.Fatalf("Kind = %q, want investigation", record.Kind)
	}
	if record.Title != "Current read fixture observations" {
		t.Fatalf("Title = %q", record.Title)
	}
	if record.Status != RecordStatusConcluded {
		t.Fatalf("Status = %q, want concluded", record.Status)
	}
	if candidate.FilenameIDMismatch {
		t.Fatalf("FilenameIDMismatch = true; candidate = %#v", candidate)
	}
	if len(issues) != 0 {
		t.Fatalf("unexpected issues: %#v", issues)
	}
	if record.Investigation == nil {
		t.Fatal("Investigation detail is nil")
	}
	assertStrings(t, record.Investigation.SourceRefs, []string{"PRODUCT-ADR-SPEC-901"})
}

func TestCurrentInvestigationH1NoRepair(t *testing.T) {
	// Bare INV ID without app namespace prefix must be rejected.
	_, _, ok := parseCurrentInvestigationH1("# INV-MCP-901: title", "DRMCP-")
	if ok {
		t.Fatal("parseCurrentInvestigationH1 accepted bare ID without ns prefix")
	}
}

// ── C03: DRMCP-REQ-MCP-901 ───────────────────────────────────────────────────

func TestCurrentRequirementRecordParser_C03(t *testing.T) {
	path := fixtureBase + "/current/drmcp/records/requirements/mcp/DRMCP-REQ-MCP-901-current-read-fixture-baseline.md"
	raw := readFixtureFile(t, path)
	record, candidate, issues := parseCurrentRequirementRecord(path, raw, "DRMCP-")

	if record == nil {
		t.Fatalf("record is nil; issues = %#v", issues)
	}
	if record.ID != "DRMCP-REQ-MCP-901" {
		t.Fatalf("ID = %q, want DRMCP-REQ-MCP-901", record.ID)
	}
	if record.Kind != RecordKindRequirement {
		t.Fatalf("Kind = %q, want requirement", record.Kind)
	}
	if record.Title != "Current read fixture baseline" {
		t.Fatalf("Title = %q", record.Title)
	}
	if record.Status != RecordStatusAccepted {
		t.Fatalf("Status = %q, want accepted", record.Status)
	}
	if candidate.FilenameIDMismatch {
		t.Fatalf("FilenameIDMismatch = true; candidate = %#v", candidate)
	}
	if len(issues) != 0 {
		t.Fatalf("unexpected issues: %#v", issues)
	}
	if record.Requirement == nil {
		t.Fatal("Requirement detail is nil")
	}
	assertStrings(t, record.Requirement.SourceRefs, []string{"PRODUCT-ADR-SPEC-901"})
	assertStrings(t, record.Requirement.WorkItems, []string{"DRMCP-WORK-MCP-901"})
}

// ── C04: DRMCP-WORK-MCP-901 ──────────────────────────────────────────────────

func TestCurrentWorkItemRecordParser_C04(t *testing.T) {
	path := fixtureBase + "/current/drmcp/records/work-items/mcp/DRMCP-WORK-MCP-901-current-read-fixture-baseline.md"
	raw := readFixtureFile(t, path)
	record, candidate, issues := parseCurrentWorkItemRecord(path, raw, "DRMCP-")

	if record == nil {
		t.Fatalf("record is nil; issues = %#v", issues)
	}
	if record.ID != "DRMCP-WORK-MCP-901" {
		t.Fatalf("ID = %q, want DRMCP-WORK-MCP-901", record.ID)
	}
	if record.Kind != RecordKindWorkItem {
		t.Fatalf("Kind = %q, want work_item", record.Kind)
	}
	if record.Status != RecordStatusInProgress {
		t.Fatalf("Status = %q, want in_progress", record.Status)
	}
	if candidate.FilenameIDMismatch {
		t.Fatalf("FilenameIDMismatch = true; candidate = %#v", candidate)
	}
	if len(issues) != 0 {
		t.Fatalf("unexpected issues: %#v", issues)
	}
	if record.WorkItem == nil {
		t.Fatal("WorkItem detail is nil")
	}
	if record.WorkItem.SourceRequirement != "DRMCP-REQ-MCP-901" {
		t.Fatalf("SourceRequirement = %q", record.WorkItem.SourceRequirement)
	}
	assertStrings(t, record.WorkItem.Tasks, []string{"DRMCP-TASK-MCP-901-01"})
}

// ── C05: DRMCP-TASK-MCP-901-01 ───────────────────────────────────────────────

func TestCurrentTaskRecordParser_C05(t *testing.T) {
	path := fixtureBase + "/current/drmcp/records/tasks/mcp/DRMCP-TASK-MCP-901-01-current-read-fixture.md"
	raw := readFixtureFile(t, path)
	record, candidate, issues := parseCurrentTaskRecord(path, raw, "DRMCP-")

	if record == nil {
		t.Fatalf("record is nil; issues = %#v", issues)
	}
	if record.ID != "DRMCP-TASK-MCP-901-01" {
		t.Fatalf("ID = %q, want DRMCP-TASK-MCP-901-01", record.ID)
	}
	if record.Kind != RecordKindTask {
		t.Fatalf("Kind = %q, want task", record.Kind)
	}
	if record.Title != "Current read fixture" {
		t.Fatalf("Title = %q", record.Title)
	}
	if record.Status != RecordStatusInProgress {
		t.Fatalf("Status = %q, want in_progress", record.Status)
	}
	if candidate.FilenameIDMismatch {
		t.Fatalf("FilenameIDMismatch = true; candidate = %#v", candidate)
	}
	if len(issues) != 0 {
		t.Fatalf("unexpected issues: %#v", issues)
	}
	if record.Task == nil {
		t.Fatal("Task detail is nil")
	}
	if record.Task.WorkItem != "DRMCP-WORK-MCP-901" {
		t.Fatalf("WorkItem = %q", record.Task.WorkItem)
	}
	if record.Task.SourceRequirement != "DRMCP-REQ-MCP-901" {
		t.Fatalf("SourceRequirement = %q", record.Task.SourceRequirement)
	}
	if record.Task.Estimate != "0.5d" {
		t.Fatalf("Estimate = %q", record.Task.Estimate)
	}
	assertStrings(t, record.Task.DependsOn, []string{})
	assertStrings(t, record.Task.Outputs, []string{})
}

// ── C06: spec:product.fixture_baseline.overview ───────────────────────────────

func TestCurrentSpecRecordParser_C06(t *testing.T) {
	recordsRoot := fixtureBase + "/current/product/records"
	path := recordsRoot + "/spec/fixture-baseline/overview.md"
	raw := readFixtureFile(t, path)
	record, candidate, issues := parseCurrentSpecRecord(path, raw, recordsRoot, "product")

	if record == nil {
		t.Fatalf("record is nil; issues = %#v", issues)
	}
	if record.ID != "spec:product.fixture_baseline.overview" {
		t.Fatalf("ID = %q, want spec:product.fixture_baseline.overview", record.ID)
	}
	if record.Kind != RecordKindSpec {
		t.Fatalf("Kind = %q, want spec", record.Kind)
	}
	if record.Title != "Overview: Fixture baseline overview" {
		t.Fatalf("Title = %q", record.Title)
	}
	if record.Status != RecordStatusAccepted {
		t.Fatalf("Status = %q, want accepted", record.Status)
	}
	if candidate.ID != "spec:product.fixture_baseline.overview" {
		t.Fatalf("candidate.ID = %q", candidate.ID)
	}
	if !candidate.Included {
		t.Fatalf("candidate.Included = false; candidate = %#v", candidate)
	}
	if len(issues) != 0 {
		t.Fatalf("unexpected issues: %#v", issues)
	}
}

// ── C07: spec:product.fixture_baseline (index.md) ────────────────────────────

func TestCurrentSpecRecordParser_C07_Index(t *testing.T) {
	recordsRoot := fixtureBase + "/current/product/records"
	path := recordsRoot + "/spec/fixture-baseline/index.md"
	raw := readFixtureFile(t, path)
	record, candidate, issues := parseCurrentSpecRecord(path, raw, recordsRoot, "product")

	if record == nil {
		t.Fatalf("record is nil; issues = %#v", issues)
	}
	if record.ID != "spec:product.fixture_baseline" {
		t.Fatalf("ID = %q, want spec:product.fixture_baseline", record.ID)
	}
	if record.Status != RecordStatusAccepted {
		t.Fatalf("Status = %q, want accepted", record.Status)
	}
	if candidate.ID != "spec:product.fixture_baseline" {
		t.Fatalf("candidate.ID = %q", candidate.ID)
	}
	if !candidate.Included {
		t.Fatalf("candidate.Included = false")
	}
	if len(issues) != 0 {
		t.Fatalf("unexpected issues: %#v", issues)
	}
}

// ── C11: mixed sequential + spec (no repair) ──────────────────────────────────

func TestCurrentMixedInputsNoRepair_C11(t *testing.T) {
	adrPath := fixtureBase + "/current/product/records/adr/spec/PRODUCT-ADR-SPEC-901-current-read-fixture-baseline.md"
	specPath := fixtureBase + "/current/product/records/spec/fixture-baseline/overview.md"
	recordsRoot := fixtureBase + "/current/product/records"

	adrRaw := readFixtureFile(t, adrPath)
	specRaw := readFixtureFile(t, specPath)

	adrRecord, _, adrIssues := parseCurrentADRRecord(adrPath, adrRaw, "PRODUCT-")
	specRecord, _, specIssues := parseCurrentSpecRecord(specPath, specRaw, recordsRoot, "product")

	if adrRecord == nil {
		t.Fatalf("ADR record is nil; issues = %#v", adrIssues)
	}
	if specRecord == nil {
		t.Fatalf("spec record is nil; issues = %#v", specIssues)
	}
	if adrRecord.ID != "PRODUCT-ADR-SPEC-901" {
		t.Fatalf("ADR ID = %q", adrRecord.ID)
	}
	if specRecord.ID != "spec:product.fixture_baseline.overview" {
		t.Fatalf("spec ID = %q", specRecord.ID)
	}
	if len(adrIssues) != 0 {
		t.Fatalf("unexpected ADR issues: %#v", adrIssues)
	}
	if len(specIssues) != 0 {
		t.Fatalf("unexpected spec issues: %#v", specIssues)
	}
}

// ── C15: invalid source (missing parent) ─────────────────────────────────────

func TestCurrentSpecRecordParser_C15_MissingParent(t *testing.T) {
	arrangement := "arrangements/invalid-current-source"
	recordsRoot := fixtureBase + "/" + arrangement + "/current/product/records"
	path := recordsRoot + "/spec/invalid-source/missing-parent.md"
	raw := readFixtureFile(t, path)
	record, candidate, issues := parseCurrentSpecRecord(path, raw, recordsRoot, "product")

	// Source is invalid but path-addressable — record must be returned.
	if record == nil {
		t.Fatal("record is nil; invalid-but-addressable source must be retained")
	}
	if record.ID != "spec:product.invalid_source.missing_parent" {
		t.Fatalf("ID = %q, want spec:product.invalid_source.missing_parent", record.ID)
	}
	if !candidate.Included {
		t.Fatalf("candidate.Included = false; invalid-but-addressable source must be retained")
	}
	if !hasIssue(issues, DiagnosticMissingRequiredMetadata) {
		t.Fatalf("expected MissingRequiredMetadata issue for missing parent; issues = %#v", issues)
	}
}

// ── R06: YAML front matter spec rejected ──────────────────────────────────────

func TestCurrentSpecRecordParser_R06_YAMLFrontMatter(t *testing.T) {
	arrangement := "arrangements/invalid-spec-format"
	recordsRoot := fixtureBase + "/" + arrangement + "/current/product/records"
	path := recordsRoot + "/spec/invalid-format/yaml-current-spec.md"
	raw := readFixtureFile(t, path)
	record, candidate, issues := parseCurrentSpecRecord(path, raw, recordsRoot, "product")

	if record != nil {
		t.Fatalf("record must be nil for YAML front matter spec; got %#v", record)
	}
	if candidate.Included {
		t.Fatalf("candidate.Included must be false for YAML front matter spec")
	}
	if candidate.SkipReason != "yaml_front_matter_current_spec" {
		t.Fatalf("SkipReason = %q, want yaml_front_matter_current_spec", candidate.SkipReason)
	}
	if len(issues) == 0 {
		t.Fatal("expected at least one issue for YAML front matter rejection")
	}
}

// ── No-repair behavior (R02/R04/R05 at parser level) ─────────────────────────

func TestCurrentWorkflowH1NoRepair(t *testing.T) {
	tests := []struct {
		name string
		line string
		ns   string
	}{
		// R02: bare REQ ID without app namespace prefix — rejected (not repaired)
		{"REQ bare no ns", "# REQ-MCP-901: title", "DRMCP-"},
		// R05: bare TASK ID without app namespace prefix — rejected (not repaired)
		{"TASK bare no ns", "# TASK-MCP-901-01: title", "DRMCP-"},
		// wrong ns prefix — rejected (not silently rewritten to correct ns)
		{"wrong ns prefix", "# PRODUCT-REQ-MCP-901: title", "DRMCP-"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			_, _, ok := parseCurrentWorkflowH1(tc.line, tc.ns)
			if ok {
				t.Fatalf("parseCurrentWorkflowH1 accepted %q with ns %q, want rejection", tc.line, tc.ns)
			}
		})
	}
}

func TestCurrentWorkflowH1ValidForms(t *testing.T) {
	tests := []struct {
		line   string
		ns     string
		bareID string
		title  string
	}{
		{"# DRMCP-REQ-MCP-901: title", "DRMCP-", "REQ-MCP-901", "title"},
		{"# DRMCP-WORK-MCP-901: title", "DRMCP-", "WORK-MCP-901", "title"},
		{"# DRMCP-TASK-MCP-901-01: title", "DRMCP-", "TASK-MCP-901-01", "title"},
	}
	for _, tc := range tests {
		bareID, title, ok := parseCurrentWorkflowH1(tc.line, tc.ns)
		if !ok {
			t.Fatalf("parseCurrentWorkflowH1(%q, %q) = false, want true", tc.line, tc.ns)
		}
		if bareID != tc.bareID || title != tc.title {
			t.Fatalf("bareID/title = %q/%q, want %q/%q", bareID, title, tc.bareID, tc.title)
		}
	}
}

// ── deriveSpecRef ─────────────────────────────────────────────────────────────

func TestDeriveSpecRef(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		recordsRoot string
		appNS       string
		want        string
	}{
		{
			name:        "leaf spec",
			path:        "current/product/records/spec/fixture-baseline/overview.md",
			recordsRoot: "current/product/records",
			appNS:       "product",
			want:        "spec:product.fixture_baseline.overview",
		},
		{
			name:        "index collapses to parent",
			path:        "current/product/records/spec/fixture-baseline/index.md",
			recordsRoot: "current/product/records",
			appNS:       "product",
			want:        "spec:product.fixture_baseline",
		},
		{
			name:        "invalid-source path",
			path:        "arrangements/invalid-current-source/current/product/records/spec/invalid-source/missing-parent.md",
			recordsRoot: "arrangements/invalid-current-source/current/product/records",
			appNS:       "product",
			want:        "spec:product.invalid_source.missing_parent",
		},
		{
			name:        "invalid-format path",
			path:        "arrangements/invalid-spec-format/current/product/records/spec/invalid-format/yaml-current-spec.md",
			recordsRoot: "arrangements/invalid-spec-format/current/product/records",
			appNS:       "product",
			want:        "spec:product.invalid_format.yaml_current_spec",
		},
		{
			name:        "path outside records root",
			path:        "other/path/spec/foo/bar.md",
			recordsRoot: "current/product/records",
			appNS:       "product",
			want:        "",
		},
		{
			name:        "non-spec path",
			path:        "current/product/records/adr/spec/PRODUCT-ADR-SPEC-901.md",
			recordsRoot: "current/product/records",
			appNS:       "product",
			want:        "",
		},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := deriveSpecRef(tc.path, tc.recordsRoot, tc.appNS)
			if got != tc.want {
				t.Fatalf("deriveSpecRef = %q, want %q", got, tc.want)
			}
		})
	}
}

// ── stripBackticks ────────────────────────────────────────────────────────────

func TestStripBackticks(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{"`spec:product.fixture_baseline.overview`", "spec:product.fixture_baseline.overview"},
		{"`spec:product.fixture_baseline`", "spec:product.fixture_baseline"},
		{"spec:product.fixture_baseline", "spec:product.fixture_baseline"},
		{"", ""},
		{"`single`", "single"},
		{"` spaced `", " spaced "},
	}
	for _, tc := range tests {
		got := stripBackticks(tc.input)
		if got != tc.want {
			t.Fatalf("stripBackticks(%q) = %q, want %q", tc.input, got, tc.want)
		}
	}
}

// ── parseCurrentListField ─────────────────────────────────────────────────────

func TestParseCurrentListField(t *testing.T) {
	assertStrings(t, parseCurrentListField("[]"), []string{})
	assertStrings(t, parseCurrentListField(""), []string{})
	assertStrings(t, parseCurrentListField("PRODUCT-ADR-SPEC-901"), []string{"PRODUCT-ADR-SPEC-901"})
	assertStrings(t, parseCurrentListField("ADR-050, ADR-068"), []string{"ADR-050", "ADR-068"})
}

// ── HeadingsExtraction ────────────────────────────────────────────────────────

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

// ── F-MIN-02: spec date retention (C06, C07) ──────────────────────────────────

func TestCurrentSpecRecordDate_C06(t *testing.T) {
	recordsRoot := fixtureBase + "/current/product/records"
	path := recordsRoot + "/spec/fixture-baseline/overview.md"
	raw := readFixtureFile(t, path)
	record, _, issues := parseCurrentSpecRecord(path, raw, recordsRoot, "product")

	if record == nil {
		t.Fatalf("record is nil; issues = %#v", issues)
	}
	if record.Date != "2026-06-28" {
		t.Fatalf("Date = %q, want 2026-06-28", record.Date)
	}
}

func TestCurrentSpecRecordDate_C07(t *testing.T) {
	recordsRoot := fixtureBase + "/current/product/records"
	path := recordsRoot + "/spec/fixture-baseline/index.md"
	raw := readFixtureFile(t, path)
	record, _, issues := parseCurrentSpecRecord(path, raw, recordsRoot, "product")

	if record == nil {
		t.Fatalf("record is nil; issues = %#v", issues)
	}
	if record.Date != "2026-06-28" {
		t.Fatalf("Date = %q, want 2026-06-28", record.Date)
	}
}

// ── F-MIN-03: full-record parser rejection for R02 and R05 ───────────────────
//
// R02 (manifest): exact_input "REQ-MCP-901" — app_prefixless_id.
// At parser level: H1 "# REQ-MCP-901: title" with ns "DRMCP-" must be rejected.
// R05 (manifest): exact_input "TASK-MCP-901-01" — missing_app_prefix.
// At parser level: H1 "# TASK-MCP-901-01: title" with ns "DRMCP-" must be rejected.

func TestCurrentRequirementRecordParser_R02_BareIDRejected(t *testing.T) {
	raw := "# REQ-MCP-901: Bare requirement\n\n- **id**: REQ-MCP-901\n- **status**: captured\n- **date**: 2026-06-28\n- **source_refs**: PRODUCT-ADR-SPEC-901\n- **work_items**: DRMCP-WORK-MCP-901\n"
	record, candidate, _ := parseCurrentRequirementRecord("records/requirements/mcp/DRMCP-REQ-MCP-901-fixture.md", raw, "DRMCP-")
	if record != nil {
		t.Fatalf("parser must reject bare ID in H1 (no ns prefix); got record.ID = %q", record.ID)
	}
	if candidate.Included {
		t.Fatalf("candidate.Included must be false for rejected source; candidate = %#v", candidate)
	}
}

func TestCurrentTaskRecordParser_R05_BareIDRejected(t *testing.T) {
	raw := "# TASK-MCP-901-01: Bare task\n\n- **id**: TASK-MCP-901-01\n- **status**: in_progress\n- **date**: 2026-06-28\n- **work_item**: DRMCP-WORK-MCP-901\n- **source_requirement**: DRMCP-REQ-MCP-901\n- **estimate**: 0.5d\n- **depends_on**: []\n- **outputs**: []\n"
	record, candidate, _ := parseCurrentTaskRecord("records/tasks/mcp/DRMCP-TASK-MCP-901-01-fixture.md", raw, "DRMCP-")
	if record != nil {
		t.Fatalf("parser must reject bare TASK suffix (no ns prefix); got record.ID = %q", record.ID)
	}
	if candidate.Included {
		t.Fatalf("candidate.Included must be false for rejected source; candidate = %#v", candidate)
	}
}

// ── F-MIN-04: splitCommaListWithEmptyItems dedicated test ────────────────────

func TestSplitCommaListWithEmptyItems(t *testing.T) {
	t.Run("empty string", func(t *testing.T) {
		items, emptyItems := splitCommaListWithEmptyItems("")
		if len(items) != 0 {
			t.Fatalf("items = %v, want empty", items)
		}
		if len(emptyItems) != 0 {
			t.Fatalf("emptyItems = %v, want empty", emptyItems)
		}
	})

	t.Run("bracket notation []", func(t *testing.T) {
		items, emptyItems := splitCommaListWithEmptyItems("[]")
		if len(items) != 0 {
			t.Fatalf("items = %v, want empty for []", items)
		}
		if len(emptyItems) != 0 {
			t.Fatalf("emptyItems = %v, want empty for []", emptyItems)
		}
	})

	t.Run("one value", func(t *testing.T) {
		items, emptyItems := splitCommaListWithEmptyItems("PRODUCT-ADR-SPEC-901")
		assertStrings(t, items, []string{"PRODUCT-ADR-SPEC-901"})
		if len(emptyItems) != 0 {
			t.Fatalf("emptyItems = %v, want empty", emptyItems)
		}
	})

	t.Run("multiple values", func(t *testing.T) {
		items, emptyItems := splitCommaListWithEmptyItems("ADR-050, ADR-068")
		assertStrings(t, items, []string{"ADR-050", "ADR-068"})
		if len(emptyItems) != 0 {
			t.Fatalf("emptyItems = %v, want empty", emptyItems)
		}
	})

	t.Run("empty member", func(t *testing.T) {
		items, emptyItems := splitCommaListWithEmptyItems("ADR-050,,ADR-068")
		assertStrings(t, items, []string{"ADR-050", "ADR-068"})
		if len(emptyItems) != 1 {
			t.Fatalf("emptyItems = %v, want one empty marker", emptyItems)
		}
	})
}
