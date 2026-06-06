package designrecords

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"testing"
)

const diagnosticSectionHeadingCaseMismatch DiagnosticCategory = "section_heading_case_mismatch"

func TestValidateRecordsDiagnosticCategories(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/001-valid.md", "# 001: Valid\n- **status**: accepted\n- **depends_on**:\n- **supersedes**:\n")
	writeTestFile(t, root, "docs/adr/002-bad-status-and-links.md", "# 002: Bad status and links\n- **status**: draft\n- **depends_on**: ADR-999\n- **supersedes**: ADR-998\n- **migrated_to_spec**: tomorrow\n")
	writeTestFile(t, root, "docs/adr/003-invalid-h1.md", "# ADR-003: Invalid H1\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/adr/004-mismatch.md", "# 005: Mismatch\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/adr/006-duplicate-a.md", "# 006: Duplicate A\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/adr/106-duplicate-b.md", "# 006: Duplicate B\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/spec/invalid-h1.md", "---\nstatus: draft\ndesign_record:\n  id: SPEC-invalid-h1\n  kind: spec\n---\n## Missing H1\n")
	writeTestFile(t, root, "docs/spec/bad-status.md", "---\nstatus: accepted\ndesign_record:\n  id: SPEC-bad-status\n  kind: spec\n---\n# Bad status\n")
	writeTestFile(t, root, "docs/spec/status-mismatch.md", "---\nstatus: draft\ndesign_record:\n  id: SPEC-status-mismatch\n  kind: spec\n  status: confirmed\n---\n# Status mismatch\n")
	writeTestFile(t, root, "docs/spec/missing-dep.md", "---\nstatus: draft\ndepends_on:\n  - docs/adr/997-policy-path-only.md\ndesign_record:\n  id: SPEC-missing-dep\n  kind: spec\n  depends_on:\n    - ADR-996\n---\n# Missing dep\n")

	idx := buildTestIndex(t, root)
	idx.PathIssues = append(idx.PathIssues, PathIssue{Path: "docs/adr/404-missing.md", Operation: "read", Err: errors.New("file does not exist")})

	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if resp.OK {
		t.Fatal("OK = true, want false")
	}
	for _, category := range []DiagnosticCategory{
		DiagnosticDuplicateID,
		DiagnosticFilenameIDMismatch,
		DiagnosticInvalidH1Title,
		DiagnosticInvalidStatusForKind,
		DiagnosticSpecStatusMismatch,
		DiagnosticMissingDependsOnTarget,
		DiagnosticMissingSupersedesTarget,
		DiagnosticInvalidMigratedToSpec,
		DiagnosticMissingRecordPath,
	} {
		if !hasDiagnostic(resp.Diagnostics, category) {
			t.Fatalf("missing diagnostic %s in %#v", category, resp.Diagnostics)
		}
	}
	assertDiagnostic(t, resp.Diagnostics, DiagnosticMissingDependsOnTarget, "ADR-002", "ADR-999")
	assertDiagnostic(t, resp.Diagnostics, DiagnosticMissingSupersedesTarget, "ADR-002", "ADR-998")
	assertDiagnostic(t, resp.Diagnostics, DiagnosticMissingDependsOnTarget, "SPEC-missing-dep", "ADR-996")
	assertDiagnosticPath(t, resp.Diagnostics, DiagnosticFilenameIDMismatch, "docs/adr/004-mismatch.md")
	assertDiagnosticPath(t, resp.Diagnostics, DiagnosticInvalidH1Title, "docs/adr/003-invalid-h1.md")
	assertDiagnosticPath(t, resp.Diagnostics, DiagnosticInvalidH1Title, "docs/spec/invalid-h1.md")
	assertDiagnosticPath(t, resp.Diagnostics, DiagnosticInvalidStatusForKind, "docs/adr/002-bad-status-and-links.md")
	assertDiagnosticPath(t, resp.Diagnostics, DiagnosticInvalidStatusForKind, "docs/spec/bad-status.md")
	assertDiagnosticPath(t, resp.Diagnostics, DiagnosticSpecStatusMismatch, "docs/spec/status-mismatch.md")
	assertDiagnosticPath(t, resp.Diagnostics, DiagnosticInvalidMigratedToSpec, "docs/adr/002-bad-status-and-links.md")
	assertDiagnosticPath(t, resp.Diagnostics, DiagnosticMissingRecordPath, "docs/adr/404-missing.md")
	if hasTarget(resp.Diagnostics, "ADR-997") {
		t.Fatal("top-level spec depends_on path list was treated as record dependency")
	}
	if countDiagnostics(resp.Diagnostics, DiagnosticDuplicateID) != 2 {
		t.Fatalf("duplicate diagnostics = %d, want 2", countDiagnostics(resp.Diagnostics, DiagnosticDuplicateID))
	}
}

func TestValidateRecordsAllowsMultipleDiagnosticsForOneRecord(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/001-bad.md", "# 001: Bad\n- **status**: draft\n- **depends_on**: ADR-999\n- **supersedes**: ADR-998\n- **migrated_to_spec**: tomorrow\n")

	idx := buildTestIndex(t, root)
	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	got := diagnosticsForRecord(resp.Diagnostics, "ADR-001")
	for _, category := range []DiagnosticCategory{
		DiagnosticInvalidStatusForKind,
		DiagnosticMissingDependsOnTarget,
		DiagnosticMissingSupersedesTarget,
		DiagnosticInvalidMigratedToSpec,
	} {
		if !hasDiagnostic(got, category) {
			t.Fatalf("record ADR-001 missing %s in %#v", category, got)
		}
	}
}

func TestValidateRecordsOKWhenNoErrorDiagnostics(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/001-valid.md", "# 001: Valid\n- **status**: accepted\n- **depends_on**:\n- **supersedes**:\n")
	writeTestFile(t, root, "docs/spec/valid.md", "---\nstatus: draft\ndesign_record:\n  id: SPEC-valid\n  kind: spec\n  depends_on:\n    - ADR-001\n---\n# Valid spec\n")
	writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-003-valid.md", "# REQ-MCP-003: Valid requirement\n- **id**: REQ-MCP-003\n- **status**: accepted\n- **date**: 2026-05-25\n- **source_refs**:\n  - ADR-001\n- **work_items**:\n  - WORK-MCP-003\n\n## Requirement\n\nRequirement content.\n\n## Required Outcome\n\nOutcome content.\n")
	writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-003-valid.md", "# WORK-MCP-003: Valid work item\n- **id**: WORK-MCP-003\n- **status**: not_started\n- **date**: 2026-05-26\n- **source_requirement**: REQ-MCP-003\n- **impact_refs**:\n  - ADR-001\n- **tasks**:\n  - TASK-MCP-003-01\n")
	writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-003-01-valid.md", "# TASK-MCP-003-01: Valid task\n- **id**: TASK-MCP-003-01\n- **status**: not_started\n- **date**: 2026-05-26\n- **work_item**: WORK-MCP-003\n- **source_requirement**: REQ-MCP-003\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n  - test\n")

	idx := buildTestIndex(t, root)
	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if !resp.OK || len(resp.Diagnostics) != 0 {
		t.Fatalf("response = %#v, want ok with no diagnostics", resp)
	}
}

func TestValidateRecordsWorkflowStatusAndParseDiagnostics(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-001-bad-status.md", "# REQ-MCP-001: Bad requirement status\n- **id**: REQ-MCP-001\n- **status**: not_started\n- **date**: 2026-05-25\n- **source_refs**:\n- **work_items**:\n")
	writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-001-bad-status.md", "# WORK-MCP-001: Bad work status\n- **id**: WORK-MCP-001\n- **status**: accepted\n- **date**: 2026-05-26\n- **source_requirement**: REQ-MCP-001\n- **impact_refs**:\n- **tasks**:\n")
	writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-001-01-bad-status.md", "# TASK-MCP-001-01: Bad task status\n- **id**: TASK-MCP-001-01\n- **status**: accepted\n- **date**: 2026-05-26\n- **work_item**: WORK-MCP-001\n- **source_requirement**: REQ-MCP-001\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n")
	writeTestFile(t, root, "docs/tasks/mcp/TASK-mcp-001-02-invalid.md", "# TASK-mcp-001-02: Invalid task ID\n- **id**: TASK-mcp-001-02\n- **status**: not_started\n")
	writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-002-mismatch.md", "# REQ-MCP-002: Mismatch\n- **id**: REQ-MCP-003\n- **status**: accepted\n")

	idx := buildTestIndex(t, root)
	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	for _, id := range []string{"REQ-MCP-001", "WORK-MCP-001", "TASK-MCP-001-01"} {
		if !hasDiagnosticForRecord(resp.Diagnostics, DiagnosticInvalidStatusForKind, id) {
			t.Fatalf("missing invalid_status_for_kind for %s in %#v", id, resp.Diagnostics)
		}
	}
	if !hasDiagnostic(resp.Diagnostics, DiagnosticInvalidWorkflowID) {
		t.Fatalf("missing invalid_workflow_id in %#v", resp.Diagnostics)
	}
	if !hasDiagnosticForRecord(resp.Diagnostics, DiagnosticFilenameIDMismatch, "REQ-MCP-002") {
		t.Fatalf("missing filename_id_mismatch in %#v", resp.Diagnostics)
	}
}

func TestValidateRecordsWorkflowMetadataStrictness(t *testing.T) {
	tests := []struct {
		name         string
		kind         RecordKind
		path         string
		content      string
		recordID     string
		category     DiagnosticCategory
		field        string
		value        string
		valuePresent bool
		validField   string
	}{
		{
			name:     "requirement missing date",
			kind:     RecordKindRequirement,
			path:     "docs/requirements/mcp/REQ-MCP-006-test.md",
			content:  "# REQ-MCP-006: Test\n- **id**: REQ-MCP-006\n- **status**: accepted\n- **source_refs**:\n- **work_items**:\n",
			recordID: "REQ-MCP-006",
			category: DiagnosticMissingRequiredMetadata,
			field:    "date",
		},
		{
			name:         "requirement invalid date",
			kind:         RecordKindRequirement,
			path:         "docs/requirements/mcp/REQ-MCP-006-test.md",
			content:      "# REQ-MCP-006: Test\n- **id**: REQ-MCP-006\n- **status**: accepted\n- **date**: 2026/06/01\n- **source_refs**:\n- **work_items**:\n",
			recordID:     "REQ-MCP-006",
			category:     DiagnosticInvalidMetadataValue,
			field:        "date",
			value:        "2026/06/01",
			valuePresent: true,
		},
		{
			name:     "requirement missing source_refs",
			kind:     RecordKindRequirement,
			path:     "docs/requirements/mcp/REQ-MCP-006-test.md",
			content:  "# REQ-MCP-006: Test\n- **id**: REQ-MCP-006\n- **status**: accepted\n- **date**: 2026-06-01\n- **work_items**:\n",
			recordID: "REQ-MCP-006",
			category: DiagnosticMissingRequiredMetadata,
			field:    "source_refs",
		},
		{
			name:       "requirement empty source_refs list is valid metadata",
			kind:       RecordKindRequirement,
			path:       "docs/requirements/mcp/REQ-MCP-006-test.md",
			content:    "# REQ-MCP-006: Test\n- **id**: REQ-MCP-006\n- **status**: accepted\n- **date**: 2026-06-01\n- **source_refs**:\n- **work_items**:\n",
			recordID:   "REQ-MCP-006",
			validField: "source_refs",
		},
		{
			name:         "requirement source_refs empty item",
			kind:         RecordKindRequirement,
			path:         "docs/requirements/mcp/REQ-MCP-006-test.md",
			content:      "# REQ-MCP-006: Test\n- **id**: REQ-MCP-006\n- **status**: accepted\n- **date**: 2026-06-01\n- **source_refs**:\n  -\n- **work_items**:\n",
			recordID:     "REQ-MCP-006",
			category:     DiagnosticEmptyRequiredMetadata,
			field:        "source_refs",
			valuePresent: true,
		},
		{
			name:     "requirement missing work_items",
			kind:     RecordKindRequirement,
			path:     "docs/requirements/mcp/REQ-MCP-006-test.md",
			content:  "# REQ-MCP-006: Test\n- **id**: REQ-MCP-006\n- **status**: accepted\n- **date**: 2026-06-01\n- **source_refs**:\n",
			recordID: "REQ-MCP-006",
			category: DiagnosticMissingRequiredMetadata,
			field:    "work_items",
		},
		{
			name:       "requirement empty work_items list is valid metadata",
			kind:       RecordKindRequirement,
			path:       "docs/requirements/mcp/REQ-MCP-006-test.md",
			content:    "# REQ-MCP-006: Test\n- **id**: REQ-MCP-006\n- **status**: accepted\n- **date**: 2026-06-01\n- **source_refs**:\n- **work_items**:\n",
			recordID:   "REQ-MCP-006",
			validField: "work_items",
		},
		{
			name:     "work item missing date",
			kind:     RecordKindWorkItem,
			path:     "docs/work-items/mcp/WORK-MCP-006-test.md",
			content:  "# WORK-MCP-006: Test\n- **id**: WORK-MCP-006\n- **status**: not_started\n- **source_requirement**: REQ-MCP-006\n- **impact_refs**:\n- **tasks**:\n",
			recordID: "WORK-MCP-006",
			category: DiagnosticMissingRequiredMetadata,
			field:    "date",
		},
		{
			name:         "work item invalid date",
			kind:         RecordKindWorkItem,
			path:         "docs/work-items/mcp/WORK-MCP-006-test.md",
			content:      "# WORK-MCP-006: Test\n- **id**: WORK-MCP-006\n- **status**: not_started\n- **date**: 2026-6-1\n- **source_requirement**: REQ-MCP-006\n- **impact_refs**:\n- **tasks**:\n",
			recordID:     "WORK-MCP-006",
			category:     DiagnosticInvalidMetadataValue,
			field:        "date",
			value:        "2026-6-1",
			valuePresent: true,
		},
		{
			name:     "work item missing source_requirement",
			kind:     RecordKindWorkItem,
			path:     "docs/work-items/mcp/WORK-MCP-006-test.md",
			content:  "# WORK-MCP-006: Test\n- **id**: WORK-MCP-006\n- **status**: not_started\n- **date**: 2026-06-01\n- **impact_refs**:\n- **tasks**:\n",
			recordID: "WORK-MCP-006",
			category: DiagnosticMissingRequiredMetadata,
			field:    "source_requirement",
		},
		{
			name:         "work item empty source_requirement",
			kind:         RecordKindWorkItem,
			path:         "docs/work-items/mcp/WORK-MCP-006-test.md",
			content:      "# WORK-MCP-006: Test\n- **id**: WORK-MCP-006\n- **status**: not_started\n- **date**: 2026-06-01\n- **source_requirement**:\n- **impact_refs**:\n- **tasks**:\n",
			recordID:     "WORK-MCP-006",
			category:     DiagnosticEmptyRequiredMetadata,
			field:        "source_requirement",
			valuePresent: true,
		},
		{
			name:     "work item missing impact_refs",
			kind:     RecordKindWorkItem,
			path:     "docs/work-items/mcp/WORK-MCP-006-test.md",
			content:  "# WORK-MCP-006: Test\n- **id**: WORK-MCP-006\n- **status**: not_started\n- **date**: 2026-06-01\n- **source_requirement**: REQ-MCP-006\n- **tasks**:\n",
			recordID: "WORK-MCP-006",
			category: DiagnosticMissingRequiredMetadata,
			field:    "impact_refs",
		},
		{
			name:       "work item empty impact_refs list is valid metadata",
			kind:       RecordKindWorkItem,
			path:       "docs/work-items/mcp/WORK-MCP-006-test.md",
			content:    "# WORK-MCP-006: Test\n- **id**: WORK-MCP-006\n- **status**: not_started\n- **date**: 2026-06-01\n- **source_requirement**: REQ-MCP-006\n- **impact_refs**:\n- **tasks**:\n",
			recordID:   "WORK-MCP-006",
			validField: "impact_refs",
		},
		{
			name:         "work item impact_refs empty item",
			kind:         RecordKindWorkItem,
			path:         "docs/work-items/mcp/WORK-MCP-006-test.md",
			content:      "# WORK-MCP-006: Test\n- **id**: WORK-MCP-006\n- **status**: not_started\n- **date**: 2026-06-01\n- **source_requirement**: REQ-MCP-006\n- **impact_refs**:\n  -\n- **tasks**:\n",
			recordID:     "WORK-MCP-006",
			category:     DiagnosticEmptyRequiredMetadata,
			field:        "impact_refs",
			valuePresent: true,
		},
		{
			name:     "work item missing tasks",
			kind:     RecordKindWorkItem,
			path:     "docs/work-items/mcp/WORK-MCP-006-test.md",
			content:  "# WORK-MCP-006: Test\n- **id**: WORK-MCP-006\n- **status**: not_started\n- **date**: 2026-06-01\n- **source_requirement**: REQ-MCP-006\n- **impact_refs**:\n",
			recordID: "WORK-MCP-006",
			category: DiagnosticMissingRequiredMetadata,
			field:    "tasks",
		},
		{
			name:       "work item empty tasks list is valid metadata",
			kind:       RecordKindWorkItem,
			path:       "docs/work-items/mcp/WORK-MCP-006-test.md",
			content:    "# WORK-MCP-006: Test\n- **id**: WORK-MCP-006\n- **status**: not_started\n- **date**: 2026-06-01\n- **source_requirement**: REQ-MCP-006\n- **impact_refs**:\n- **tasks**:\n",
			recordID:   "WORK-MCP-006",
			validField: "tasks",
		},
		{
			name:       "in_progress is valid work item status",
			kind:       RecordKindWorkItem,
			path:       "docs/work-items/mcp/WORK-MCP-006-test.md",
			content:    "# WORK-MCP-006: Test\n- **id**: WORK-MCP-006\n- **status**: in_progress\n- **date**: 2026-06-01\n- **source_requirement**: REQ-MCP-006\n- **impact_refs**:\n- **tasks**:\n",
			recordID:   "WORK-MCP-006",
			validField: "status",
		},
		{
			name:     "task missing date",
			kind:     RecordKindTask,
			path:     "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:  "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **work_item**: WORK-MCP-006\n- **source_requirement**: REQ-MCP-006\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n",
			recordID: "TASK-MCP-006-04",
			category: DiagnosticMissingRequiredMetadata,
			field:    "date",
		},
		{
			name:         "task invalid date",
			kind:         RecordKindTask,
			path:         "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:      "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **date**: abc\n- **work_item**: WORK-MCP-006\n- **source_requirement**: REQ-MCP-006\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n",
			recordID:     "TASK-MCP-006-04",
			category:     DiagnosticInvalidMetadataValue,
			field:        "date",
			value:        "abc",
			valuePresent: true,
		},
		{
			name:     "task missing work_item",
			kind:     RecordKindTask,
			path:     "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:  "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **date**: 2026-06-01\n- **source_requirement**: REQ-MCP-006\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n",
			recordID: "TASK-MCP-006-04",
			category: DiagnosticMissingRequiredMetadata,
			field:    "work_item",
		},
		{
			name:         "task empty work_item",
			kind:         RecordKindTask,
			path:         "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:      "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **date**: 2026-06-01\n- **work_item**:\n- **source_requirement**: REQ-MCP-006\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n",
			recordID:     "TASK-MCP-006-04",
			category:     DiagnosticEmptyRequiredMetadata,
			field:        "work_item",
			valuePresent: true,
		},
		{
			name:     "task missing source_requirement",
			kind:     RecordKindTask,
			path:     "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:  "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **date**: 2026-06-01\n- **work_item**: WORK-MCP-006\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n",
			recordID: "TASK-MCP-006-04",
			category: DiagnosticMissingRequiredMetadata,
			field:    "source_requirement",
		},
		{
			name:         "task empty source_requirement",
			kind:         RecordKindTask,
			path:         "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:      "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **date**: 2026-06-01\n- **work_item**: WORK-MCP-006\n- **source_requirement**:\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n",
			recordID:     "TASK-MCP-006-04",
			category:     DiagnosticEmptyRequiredMetadata,
			field:        "source_requirement",
			valuePresent: true,
		},
		{
			name:     "task missing estimate",
			kind:     RecordKindTask,
			path:     "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:  "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **date**: 2026-06-01\n- **work_item**: WORK-MCP-006\n- **source_requirement**: REQ-MCP-006\n- **depends_on**:\n- **outputs**:\n",
			recordID: "TASK-MCP-006-04",
			category: DiagnosticMissingRequiredMetadata,
			field:    "estimate",
		},
		{
			name:         "task empty estimate",
			kind:         RecordKindTask,
			path:         "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:      "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **date**: 2026-06-01\n- **work_item**: WORK-MCP-006\n- **source_requirement**: REQ-MCP-006\n- **estimate**:\n- **depends_on**:\n- **outputs**:\n",
			recordID:     "TASK-MCP-006-04",
			category:     DiagnosticEmptyRequiredMetadata,
			field:        "estimate",
			valuePresent: true,
		},
		{
			name:     "task missing depends_on",
			kind:     RecordKindTask,
			path:     "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:  "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **date**: 2026-06-01\n- **work_item**: WORK-MCP-006\n- **source_requirement**: REQ-MCP-006\n- **estimate**: 0.5d\n- **outputs**:\n",
			recordID: "TASK-MCP-006-04",
			category: DiagnosticMissingRequiredMetadata,
			field:    "depends_on",
		},
		{
			name:       "task empty depends_on list is valid metadata",
			kind:       RecordKindTask,
			path:       "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:    "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **date**: 2026-06-01\n- **work_item**: WORK-MCP-006\n- **source_requirement**: REQ-MCP-006\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n",
			recordID:   "TASK-MCP-006-04",
			validField: "depends_on",
		},
		{
			name:         "task depends_on empty item",
			kind:         RecordKindTask,
			path:         "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:      "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **date**: 2026-06-01\n- **work_item**: WORK-MCP-006\n- **source_requirement**: REQ-MCP-006\n- **estimate**: 0.5d\n- **depends_on**:\n  -\n- **outputs**:\n",
			recordID:     "TASK-MCP-006-04",
			category:     DiagnosticEmptyRequiredMetadata,
			field:        "depends_on",
			valuePresent: true,
		},
		{
			name:     "task missing outputs",
			kind:     RecordKindTask,
			path:     "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:  "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **date**: 2026-06-01\n- **work_item**: WORK-MCP-006\n- **source_requirement**: REQ-MCP-006\n- **estimate**: 0.5d\n- **depends_on**:\n",
			recordID: "TASK-MCP-006-04",
			category: DiagnosticMissingRequiredMetadata,
			field:    "outputs",
		},
		{
			name:       "task empty outputs list is valid metadata",
			kind:       RecordKindTask,
			path:       "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:    "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **date**: 2026-06-01\n- **work_item**: WORK-MCP-006\n- **source_requirement**: REQ-MCP-006\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n",
			recordID:   "TASK-MCP-006-04",
			validField: "outputs",
		},
		{
			name:         "task outputs empty item",
			kind:         RecordKindTask,
			path:         "docs/tasks/mcp/TASK-MCP-006-04-test.md",
			content:      "# TASK-MCP-006-04: Test\n- **id**: TASK-MCP-006-04\n- **status**: not_started\n- **date**: 2026-06-01\n- **work_item**: WORK-MCP-006\n- **source_requirement**: REQ-MCP-006\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n  -\n",
			recordID:     "TASK-MCP-006-04",
			category:     DiagnosticEmptyRequiredMetadata,
			field:        "outputs",
			valuePresent: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := t.TempDir()
			writeTestFile(t, root, tt.path, tt.content)
			idx := buildTestIndex(t, root)
			resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{Kind: tt.kind})
			if err != nil {
				t.Fatalf("ValidateRecords: %v", err)
			}
			if tt.validField != "" {
				assertNoWorkflowMetadataDiagnostic(t, resp.Diagnostics, tt.recordID, tt.validField)
				if tt.validField == "status" && hasDiagnosticForRecord(resp.Diagnostics, DiagnosticInvalidStatusForKind, tt.recordID) {
					t.Fatalf("fixture_pending produced invalid status diagnostic: %#v", resp.Diagnostics)
				}
				return
			}
			assertWorkflowMetadataDiagnostic(t, resp.Diagnostics, tt.category, tt.recordID, tt.field, tt.value, tt.valuePresent)
		})
	}
}

func TestValidateRecordsWorkflowRelationHappyPath(t *testing.T) {
	root := t.TempDir()
	writeWorkflowHappyPathFixture(t, root)
	idx := buildTestIndex(t, root)

	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if !resp.OK {
		t.Fatalf("OK = false, diagnostics = %#v", resp.Diagnostics)
	}
	assertNoWorkflowRelationDiagnostics(t, resp.Diagnostics)
}

func TestValidateRecordsWorkflowRelationUnresolvedTargets(t *testing.T) {
	tests := []struct {
		name     string
		records  []Record
		recordID string
		field    string
		value    string
	}{
		{
			name:     "requirement work_items",
			records:  []Record{requirementTestRecord("REQ-MCP-003", []string{"WORK-MCP-999"})},
			recordID: "REQ-MCP-003",
			field:    "work_items",
			value:    "WORK-MCP-999",
		},
		{
			name:     "work item source_requirement",
			records:  []Record{workItemTestRecord("WORK-MCP-003", "REQ-MCP-999", nil)},
			recordID: "WORK-MCP-003",
			field:    "source_requirement",
			value:    "REQ-MCP-999",
		},
		{
			name:     "work item tasks",
			records:  []Record{workItemTestRecord("WORK-MCP-003", "", []string{"TASK-MCP-003-99"})},
			recordID: "WORK-MCP-003",
			field:    "tasks",
			value:    "TASK-MCP-003-99",
		},
		{
			name:     "task work_item",
			records:  []Record{taskTestRecord("TASK-MCP-003-01", "WORK-MCP-999", "", nil)},
			recordID: "TASK-MCP-003-01",
			field:    "work_item",
			value:    "WORK-MCP-999",
		},
		{
			name:     "task source_requirement",
			records:  []Record{taskTestRecord("TASK-MCP-003-01", "", "REQ-MCP-999", nil)},
			recordID: "TASK-MCP-003-01",
			field:    "source_requirement",
			value:    "REQ-MCP-999",
		},
		{
			name:     "task depends_on",
			records:  []Record{taskTestRecord("TASK-MCP-003-01", "", "", []string{"TASK-MCP-003-99"})},
			recordID: "TASK-MCP-003-01",
			field:    "depends_on",
			value:    "TASK-MCP-003-99",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ValidateRecords(context.Background(), workflowTestIndex(tt.records...), ValidateRecordsRequest{})
			if err != nil {
				t.Fatalf("ValidateRecords: %v", err)
			}
			assertWorkflowDiagnostic(t, resp.Diagnostics, DiagnosticUnresolvedWorkflowRelation, tt.recordID, tt.field, tt.value, "unresolved", tt.value)
		})
	}
}

func TestValidateRecordsWorkflowRelationInvalidTargets(t *testing.T) {
	tests := []struct {
		name     string
		records  []Record
		recordID string
		field    string
		value    string
	}{
		{
			name: "requirement work_items wrong ID form",
			records: []Record{
				requirementTestRecord("REQ-MCP-003", []string{"TASK-MCP-003-01"}),
				taskTestRecord("TASK-MCP-003-01", "", "", nil),
			},
			recordID: "REQ-MCP-003",
			field:    "work_items",
			value:    "TASK-MCP-003-01",
		},
		{
			name: "work item source_requirement wrong ID form",
			records: []Record{
				workItemTestRecord("WORK-MCP-003", "WORK-MCP-004", nil),
				workItemTestRecord("WORK-MCP-004", "", nil),
			},
			recordID: "WORK-MCP-003",
			field:    "source_requirement",
			value:    "WORK-MCP-004",
		},
		{
			name: "work item tasks wrong ID form",
			records: []Record{
				workItemTestRecord("WORK-MCP-003", "", []string{"REQ-MCP-003"}),
				requirementTestRecord("REQ-MCP-003", nil),
			},
			recordID: "WORK-MCP-003",
			field:    "tasks",
			value:    "REQ-MCP-003",
		},
		{
			name: "task work_item wrong ID form",
			records: []Record{
				taskTestRecord("TASK-MCP-003-01", "TASK-MCP-003-02", "", nil),
				taskTestRecord("TASK-MCP-003-02", "", "", nil),
			},
			recordID: "TASK-MCP-003-01",
			field:    "work_item",
			value:    "TASK-MCP-003-02",
		},
		{
			name: "task source_requirement wrong ID form",
			records: []Record{
				taskTestRecord("TASK-MCP-003-01", "", "WORK-MCP-003", nil),
				workItemTestRecord("WORK-MCP-003", "", nil),
			},
			recordID: "TASK-MCP-003-01",
			field:    "source_requirement",
			value:    "WORK-MCP-003",
		},
		{
			name: "task depends_on wrong ID form",
			records: []Record{
				taskTestRecord("TASK-MCP-003-01", "", "", []string{"WORK-MCP-003"}),
				workItemTestRecord("WORK-MCP-003", "", nil),
			},
			recordID: "TASK-MCP-003-01",
			field:    "depends_on",
			value:    "WORK-MCP-003",
		},
		{
			name: "actual indexed target kind must match ID form",
			records: []Record{
				requirementTestRecord("REQ-MCP-003", []string{"WORK-MCP-003"}),
				{ID: "WORK-MCP-003", NormalizedID: "WORK-MCP-003", Kind: RecordKindRequirement, Title: "Wrong actual kind", Status: RecordStatusAccepted, Path: "docs/requirements/mcp/WORK-MCP-003-wrong.md", Requirement: &RequirementDetail{WorkItems: []string{}}},
			},
			recordID: "REQ-MCP-003",
			field:    "work_items",
			value:    "WORK-MCP-003",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ValidateRecords(context.Background(), workflowTestIndex(tt.records...), ValidateRecordsRequest{})
			if err != nil {
				t.Fatalf("ValidateRecords: %v", err)
			}
			assertWorkflowDiagnostic(t, resp.Diagnostics, DiagnosticInvalidWorkflowRelationTarget, tt.recordID, tt.field, tt.value, "invalid_target", tt.value)
		})
	}
}

func TestValidateRecordsWorkflowRelationBidirectionalMismatch(t *testing.T) {
	tests := []struct {
		name     string
		records  []Record
		recordID string
		field    string
		value    string
		targetID string
	}{
		{
			name: "requirement lists work item pointing to another requirement",
			records: []Record{
				requirementTestRecord("REQ-MCP-003", []string{"WORK-MCP-003"}),
				requirementTestRecord("REQ-MCP-004", []string{"WORK-MCP-003"}),
				workItemTestRecord("WORK-MCP-003", "REQ-MCP-004", nil),
			},
			recordID: "REQ-MCP-003",
			field:    "work_items",
			value:    "WORK-MCP-003",
			targetID: "WORK-MCP-003",
		},
		{
			name: "work item points to requirement that does not list it",
			records: []Record{
				requirementTestRecord("REQ-MCP-003", nil),
				workItemTestRecord("WORK-MCP-003", "REQ-MCP-003", nil),
			},
			recordID: "WORK-MCP-003",
			field:    "source_requirement",
			value:    "REQ-MCP-003",
			targetID: "REQ-MCP-003",
		},
		{
			name: "work item lists task pointing to another work item",
			records: []Record{
				workItemTestRecord("WORK-MCP-003", "", []string{"TASK-MCP-003-01"}),
				workItemTestRecord("WORK-MCP-004", "", []string{"TASK-MCP-003-01"}),
				taskTestRecord("TASK-MCP-003-01", "WORK-MCP-004", "", nil),
			},
			recordID: "WORK-MCP-003",
			field:    "tasks",
			value:    "TASK-MCP-003-01",
			targetID: "TASK-MCP-003-01",
		},
		{
			name: "task points to work item that does not list it",
			records: []Record{
				workItemTestRecord("WORK-MCP-003", "", nil),
				taskTestRecord("TASK-MCP-003-01", "WORK-MCP-003", "", nil),
			},
			recordID: "TASK-MCP-003-01",
			field:    "work_item",
			value:    "WORK-MCP-003",
			targetID: "WORK-MCP-003",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ValidateRecords(context.Background(), workflowTestIndex(tt.records...), ValidateRecordsRequest{})
			if err != nil {
				t.Fatalf("ValidateRecords: %v", err)
			}
			assertWorkflowDiagnostic(t, resp.Diagnostics, DiagnosticWorkflowRelationMismatch, tt.recordID, tt.field, tt.value, "mismatch", tt.targetID)
		})
	}
}

func TestValidateRecordsTaskSourceRequirementMismatch(t *testing.T) {
	idx := workflowTestIndex(
		requirementTestRecord("REQ-MCP-003", []string{"WORK-MCP-003"}),
		requirementTestRecord("REQ-MCP-004", nil),
		workItemTestRecord("WORK-MCP-003", "REQ-MCP-003", []string{"TASK-MCP-003-01"}),
		taskTestRecord("TASK-MCP-003-01", "WORK-MCP-003", "REQ-MCP-004", nil),
	)

	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	assertWorkflowDiagnostic(t, resp.Diagnostics, DiagnosticWorkflowSourceReqMismatch, "TASK-MCP-003-01", "source_requirement", "REQ-MCP-004", "mismatch", "REQ-MCP-003")
}

func TestValidateRecordsTaskDependsOnBoundary(t *testing.T) {
	validBase := []Record{
		requirementTestRecord("REQ-MCP-003", []string{"WORK-MCP-003"}),
		workItemTestRecord("WORK-MCP-003", "REQ-MCP-003", []string{"TASK-MCP-003-01", "TASK-MCP-003-02"}),
		taskTestRecord("TASK-MCP-003-01", "WORK-MCP-003", "REQ-MCP-003", nil),
		taskTestRecord("TASK-MCP-003-02", "WORK-MCP-003", "REQ-MCP-003", []string{"TASK-MCP-003-01"}),
	}
	resp, err := ValidateRecords(context.Background(), workflowTestIndex(validBase...), ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords valid dependencies: %v", err)
	}
	assertNoWorkflowRelationDiagnostics(t, resp.Diagnostics)

	selfDependency := []Record{
		requirementTestRecord("REQ-MCP-003", []string{"WORK-MCP-003"}),
		workItemTestRecord("WORK-MCP-003", "REQ-MCP-003", []string{"TASK-MCP-003-01"}),
		taskTestRecord("TASK-MCP-003-01", "WORK-MCP-003", "REQ-MCP-003", []string{"TASK-MCP-003-01"}),
	}
	resp, err = ValidateRecords(context.Background(), workflowTestIndex(selfDependency...), ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords self dependency: %v", err)
	}
	assertNoWorkflowRelationDiagnostics(t, resp.Diagnostics)

	missingDependency := []Record{
		requirementTestRecord("REQ-MCP-003", []string{"WORK-MCP-003"}),
		workItemTestRecord("WORK-MCP-003", "REQ-MCP-003", []string{"TASK-MCP-003-01"}),
		taskTestRecord("TASK-MCP-003-01", "WORK-MCP-003", "REQ-MCP-003", []string{"TASK-MCP-003-99"}),
	}
	resp, err = ValidateRecords(context.Background(), workflowTestIndex(missingDependency...), ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords missing dependency: %v", err)
	}
	assertWorkflowDiagnostic(t, resp.Diagnostics, DiagnosticUnresolvedWorkflowRelation, "TASK-MCP-003-01", "depends_on", "TASK-MCP-003-99", "unresolved", "TASK-MCP-003-99")

	wrongKindDependency := []Record{
		taskTestRecord("TASK-MCP-003-01", "", "", []string{"WORK-MCP-003"}),
		workItemTestRecord("WORK-MCP-003", "", nil),
	}
	resp, err = ValidateRecords(context.Background(), workflowTestIndex(wrongKindDependency...), ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords wrong-kind dependency: %v", err)
	}
	assertWorkflowDiagnostic(t, resp.Diagnostics, DiagnosticInvalidWorkflowRelationTarget, "TASK-MCP-003-01", "depends_on", "WORK-MCP-003", "invalid_target", "WORK-MCP-003")
}

func TestValidateRecordsWorkflowRelationKindFilter(t *testing.T) {
	idx := workflowTestIndex(
		requirementTestRecord("REQ-MCP-003", []string{"WORK-MCP-003"}),
		requirementTestRecord("REQ-MCP-004", nil),
		workItemTestRecord("WORK-MCP-003", "REQ-MCP-004", []string{"TASK-MCP-003-01"}),
		workItemTestRecord("WORK-MCP-004", "REQ-MCP-004", nil),
		taskTestRecord("TASK-MCP-003-01", "WORK-MCP-004", "REQ-MCP-004", nil),
	)

	for _, tt := range []struct {
		name string
		kind RecordKind
		want RecordKind
	}{
		{name: "requirement sources only", kind: RecordKindRequirement, want: RecordKindRequirement},
		{name: "work item sources only", kind: RecordKindWorkItem, want: RecordKindWorkItem},
		{name: "task sources only", kind: RecordKindTask, want: RecordKindTask},
	} {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{Kind: tt.kind})
			if err != nil {
				t.Fatalf("ValidateRecords: %v", err)
			}
			if len(resp.Diagnostics) == 0 {
				t.Fatalf("expected diagnostics for kind %s", tt.kind)
			}
			for _, diagnostic := range resp.Diagnostics {
				if !isWorkflowRelationDiagnostic(diagnostic.Category) {
					continue
				}
				record := findRecord(idx.Records, diagnostic.RecordID)
				if record == nil {
					t.Fatalf("diagnostic record not found: %#v", diagnostic)
				}
				if record.Kind != tt.want {
					t.Fatalf("diagnostic leaked from kind %s into %s filter: %#v", record.Kind, tt.kind, diagnostic)
				}
			}
		})
	}
}

func TestValidateRecordsRepositoryWorkflowRelationBootstrap(t *testing.T) {
	root := findRepoRoot(t)
	idx := buildTestIndex(t, root)
	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	for _, diagnostic := range resp.Diagnostics {
		if !isWorkflowRelationDiagnostic(diagnostic.Category) {
			continue
		}
		if diagnostic.RecordID == "REQ-MCP-003" || diagnostic.RecordID == "WORK-MCP-003" || strings.HasPrefix(diagnostic.RecordID, "TASK-MCP-003-") {
			t.Fatalf("unexpected repository workflow relation diagnostic: %#v", diagnostic)
		}
	}
}

func TestValidateRecordsInvestigationKnownDependencyTargetExists(t *testing.T) {
	root := findRepoRoot(t)
	idx := buildTestIndex(t, root)
	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	for _, diagnostic := range resp.Diagnostics {
		if diagnostic.Category == DiagnosticMissingDependsOnTarget && diagnostic.RecordID == "ADR-086" && diagnostic.TargetID == "INV-DOCS-001" {
			t.Fatalf("ADR-086 -> INV-DOCS-001 still unresolved: %#v", diagnostic)
		}
	}
}

func TestValidateRecordsSemanticRefDeclarationDiagnostics(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/spec/a.md", "---\nstatus: draft\nsemantic_refs:\n  - spec:trace.duplicate\n  - Spec:Bad\nsections:\n  spec:trace.missing: Missing heading\n  spec:trace.ambiguous: Duplicate\n\ndesign_record:\n  id: SPEC-a\n  kind: spec\n  status: draft\n---\n# A\n## Duplicate\n## Duplicate\n")
	writeTestFile(t, root, "docs/spec/b.md", "---\nstatus: draft\nsemantic_refs:\n  - spec:trace.duplicate\nsections:\n  bad/ref: Existing\n\ndesign_record:\n  id: SPEC-b\n  kind: spec\n  status: draft\n---\n# B\n## Existing\n")

	idx := buildTestIndex(t, root)
	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	for _, category := range []DiagnosticCategory{
		DiagnosticInvalidSemanticRefDeclaration,
		DiagnosticMissingSectionTarget,
		DiagnosticAmbiguousSectionTarget,
		DiagnosticDuplicateSemanticRef,
	} {
		if !hasDiagnostic(resp.Diagnostics, category) {
			t.Fatalf("missing %s in %#v", category, resp.Diagnostics)
		}
	}
}

func TestValidateRecordsAllowsRootSemanticRefDeclarations(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/spec/trace.md", "---\nstatus: draft\nsemantic_refs:\n  - spec:trace\n  - spec:trace.semantic-ref\nsections:\n  spec:trace.resolve-and-validation: Resolve and validation\n\ndesign_record:\n  id: SPEC-trace\n  kind: spec\n  status: draft\n---\n# Trace\n## Resolve and validation\n")
	writeTestFile(t, root, "docs/spec/project-artifact-model/index.md", "---\nstatus: draft\nsemantic_refs:\n  - spec:project-artifact-model\nsections:\n  spec:project-artifact-model.responsibilities: Artifact responsibility matrix\n\ndesign_record:\n  id: SPEC-project-artifact-model\n  kind: spec\n  status: draft\n---\n# Project artifact model\n## Artifact responsibility matrix\n")

	idx := buildTestIndex(t, root)
	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if !resp.OK || len(resp.Diagnostics) != 0 {
		t.Fatalf("root semantic refs should be valid: %#v", resp)
	}
}

func TestValidateRecordsInvestigationReferenceDiagnostics(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/001-valid.md", "# 001: Valid\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/spec/trace.md", "---\nstatus: draft\nsemantic_refs:\n  - spec:trace.valid\ndesign_record:\n  id: SPEC-trace\n  kind: spec\n  status: draft\n---\n# Trace\n")
	writeTestFile(t, root, "docs/investigations/docs/INV-DOCS-001-test.md", "# INV-DOCS-001: Test investigation\n- **status**: concluded\n- **date**: 2026-05-19\n- **trigger**: ADR-001\n- **scope**: test\n- **non_scope**: none\n- **source_refs**:\n  - ADR-999\n  - docs/adr/001-valid.md\n  - internal-design:resolver.semantic-ref-index\n- **follow_up_candidates**:\n  - SPEC-missing\n  - docs/spec/future.md\n  - coverage:trace\n- **follow_up_results**:\n  - spec:trace.missing\n  - docs/spec/trace.md\n  - COV-TRACE-001\n")
	idx := buildTestIndex(t, root)

	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if resp.OK {
		t.Fatalf("OK = true, want false: %#v", resp.Diagnostics)
	}
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnresolvedSourceRef, DiagnosticSeverityError, "source_refs", "ADR-999", "unresolved")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticNoncanonicalSourceRef, DiagnosticSeverityError, "source_refs", "docs/adr/001-valid.md", "noncanonical")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnsupportedReference, DiagnosticSeverityError, "source_refs", "internal-design:resolver.semantic-ref-index", "unsupported")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnresolvedFollowUpResult, DiagnosticSeverityError, "follow_up_results", "spec:trace.missing", "unresolved")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticNoncanonicalFollowUpResult, DiagnosticSeverityError, "follow_up_results", "docs/spec/trace.md", "noncanonical")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnsupportedReference, DiagnosticSeverityError, "follow_up_results", "COV-TRACE-001", "unsupported")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnresolvedFollowUpCandidate, DiagnosticSeverityInfo, "follow_up_candidates", "SPEC-missing", "unresolved")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticNoncanonicalFollowUpCandidate, DiagnosticSeverityInfo, "follow_up_candidates", "docs/spec/future.md", "noncanonical")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnsupportedReference, DiagnosticSeverityInfo, "follow_up_candidates", "coverage:trace", "unsupported")
}

func TestValidateRecordsInvestigationWorkflowReferenceBoundary(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/001-valid.md", "# 001: Valid\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-003-valid.md", "# REQ-MCP-003: Valid requirement\n- **id**: REQ-MCP-003\n- **status**: accepted\n- **date**: 2026-05-25\n- **source_refs**:\n  - ADR-001\n- **work_items**:\n  - WORK-MCP-003\n")
	writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-003-valid.md", "# WORK-MCP-003: Valid work item\n- **id**: WORK-MCP-003\n- **status**: not_started\n- **date**: 2026-05-26\n- **source_requirement**: REQ-MCP-003\n- **impact_refs**:\n  - ADR-001\n- **tasks**:\n  - TASK-MCP-003-01\n")
	writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-003-01-valid.md", "# TASK-MCP-003-01: Valid task\n- **id**: TASK-MCP-003-01\n- **status**: not_started\n- **date**: 2026-05-26\n- **work_item**: WORK-MCP-003\n- **source_requirement**: REQ-MCP-003\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n  - test\n")
	writeTestFile(t, root, "docs/investigations/docs/INV-DOCS-001-test.md", "# INV-DOCS-001: Test investigation\n- **status**: concluded\n- **date**: 2026-05-19\n- **trigger**: TASK-MCP-003-01\n- **scope**: test\n- **non_scope**: none\n- **source_refs**:\n  - REQ-MCP-003\n  - WORK-MCP-003\n  - REQ-MCP-999\n  - WORK-MCP-999\n  - TASK-MCP-003-01\n  - TASK-MCP-999-99\n- **follow_up_candidates**:\n  - REQ-MCP-003\n  - WORK-MCP-003\n  - REQ-MCP-998\n  - WORK-MCP-998\n  - TASK-MCP-003-01\n  - TASK-MCP-998-99\n- **follow_up_results**:\n  - REQ-MCP-003\n  - WORK-MCP-003\n  - REQ-MCP-997\n  - WORK-MCP-997\n  - TASK-MCP-003-01\n  - TASK-MCP-997-99\n- **related_requirements**:\n  - REQ-MCP-999\n  - TASK-MCP-003-01\n- **related_work_items**:\n  - WORK-MCP-999\n")
	idx := buildTestIndex(t, root)

	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if resp.OK {
		t.Fatalf("OK = true, want false: %#v", resp.Diagnostics)
	}

	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnresolvedSourceRef, DiagnosticSeverityError, "source_refs", "REQ-MCP-999", "unresolved")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnresolvedSourceRef, DiagnosticSeverityError, "source_refs", "WORK-MCP-999", "unresolved")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnsupportedReference, DiagnosticSeverityError, "source_refs", "TASK-MCP-003-01", "unsupported")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnsupportedReference, DiagnosticSeverityError, "source_refs", "TASK-MCP-999-99", "unsupported")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnresolvedFollowUpResult, DiagnosticSeverityError, "follow_up_results", "REQ-MCP-997", "unresolved")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnresolvedFollowUpResult, DiagnosticSeverityError, "follow_up_results", "WORK-MCP-997", "unresolved")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnsupportedReference, DiagnosticSeverityError, "follow_up_results", "TASK-MCP-003-01", "unsupported")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnsupportedReference, DiagnosticSeverityError, "follow_up_results", "TASK-MCP-997-99", "unsupported")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnresolvedFollowUpCandidate, DiagnosticSeverityInfo, "follow_up_candidates", "REQ-MCP-998", "unresolved")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnresolvedFollowUpCandidate, DiagnosticSeverityInfo, "follow_up_candidates", "WORK-MCP-998", "unresolved")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnsupportedReference, DiagnosticSeverityInfo, "follow_up_candidates", "TASK-MCP-003-01", "unsupported")
	assertInvestigationDiagnostic(t, resp.Diagnostics, DiagnosticUnsupportedReference, DiagnosticSeverityInfo, "follow_up_candidates", "TASK-MCP-998-99", "unsupported")

	for _, value := range []string{"REQ-MCP-003", "WORK-MCP-003"} {
		assertNoInvestigationDiagnosticValue(t, resp.Diagnostics, value)
	}
	for _, diagnostic := range resp.Diagnostics {
		if diagnostic.Field == "trigger" || diagnostic.Field == "related_requirements" || diagnostic.Field == "related_work_items" {
			t.Fatalf("unexpected validation for out-of-scope investigation field: %#v", diagnostic)
		}
	}
}

func TestValidateRecordsInvestigationSourceRefsResolveRootSemanticRefs(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/001-valid.md", "# 001: Valid\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/spec/project-artifact-model/index.md", "---\nstatus: draft\nsemantic_refs:\n  - spec:project-artifact-model\ndesign_record:\n  id: SPEC-project-artifact-model\n  kind: spec\n  status: draft\n---\n# Project artifact model\n")
	writeTestFile(t, root, "docs/investigations/docs/INV-DOCS-001-test.md", "# INV-DOCS-001: Test investigation\n- **status**: concluded\n- **date**: 2026-05-19\n- **trigger**: ADR-001\n- **scope**: test\n- **non_scope**: none\n- **source_refs**:\n  - spec:project-artifact-model\n- **follow_up_candidates**:\n  - SPEC-project-artifact-model\n")

	idx := buildTestIndex(t, root)
	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	for _, diagnostic := range resp.Diagnostics {
		if diagnostic.Category == DiagnosticUnsupportedReference || diagnostic.Category == DiagnosticUnresolvedSourceRef {
			t.Fatalf("root source_ref should resolve without unsupported/unresolved diagnostics: %#v", resp.Diagnostics)
		}
	}
	if !resp.OK {
		t.Fatalf("root source_ref should not produce errors: %#v", resp.Diagnostics)
	}
}

func TestValidateRecordsRejectsInvalidSemanticRefForms(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/spec/invalid.md", "---\nstatus: draft\nsemantic_refs:\n  - Spec:trace\n  - 'spec:'\n  - spec:trace/\n  - spec:trace..broken\n\ndesign_record:\n  id: SPEC-invalid\n  kind: spec\n  status: draft\n---\n# Invalid refs\n")

	idx := buildTestIndex(t, root)
	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if got := countDiagnostics(resp.Diagnostics, DiagnosticInvalidSemanticRefDeclaration); got != 4 {
		t.Fatalf("invalid semantic ref diagnostics = %d, want 4: %#v", got, resp.Diagnostics)
	}
}

func TestValidateRecordsInfoOnlyDiagnosticsKeepOKTrue(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/001-valid.md", "# 001: Valid\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/investigations/docs/INV-DOCS-001-test.md", "# INV-DOCS-001: Test investigation\n- **status**: concluded\n- **date**: 2026-05-19\n- **trigger**: ADR-001\n- **scope**: test\n- **non_scope**: none\n- **source_refs**:\n  - ADR-001\n- **follow_up_candidates**:\n  - SPEC-future\n  - docs/spec/future.md\n  - coverage:trace\n")
	idx := buildTestIndex(t, root)

	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if !resp.OK {
		t.Fatalf("OK = false for info-only diagnostics: %#v", resp.Diagnostics)
	}
	for _, diagnostic := range resp.Diagnostics {
		if diagnostic.Severity != DiagnosticSeverityInfo {
			t.Fatalf("non-info diagnostic in info-only test: %#v", diagnostic)
		}
	}
}

func TestValidateRecordsIgnoresYAMLRefsInInvestigationMetadata(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/001-valid.md", "# 001: Valid\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/investigations/docs/INV-DOCS-001-test.md", "# INV-DOCS-001: Test investigation\n- **status**: concluded\n- **date**: 2026-05-19\n- **trigger**: ADR-001\n- **scope**: test\n- **non_scope**: none\n- **source_refs**:\n  - ADR-001\n  - yaml:uc-001\n- **follow_up_candidates**:\n  - yaml:future\n- **follow_up_results**:\n  - yaml:result\n")
	idx := buildTestIndex(t, root)

	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if !resp.OK || len(resp.Diagnostics) != 0 {
		t.Fatalf("yaml refs should be ignored by M19 validation: %#v", resp)
	}
}

func TestValidateRecordsDuplicateTargetsDoNotAddFieldSpecificAmbiguousDiagnostics(t *testing.T) {
	idx := &Index{
		Records: []Record{
			{ID: "ADR-001", NormalizedID: "ADR-001", Kind: RecordKindDecision, Title: "One", Status: RecordStatusAccepted, Path: "docs/adr/001-one.md"},
			{ID: "ADR-001", NormalizedID: "ADR-001", Kind: RecordKindDecision, Title: "Duplicate", Status: RecordStatusAccepted, Path: "docs/adr/001-duplicate.md"},
			{ID: "SPEC-a", NormalizedID: "SPEC-A", Kind: RecordKindSpec, Title: "A", Status: RecordStatusDraft, Path: "docs/spec/a.md", Spec: &SpecDetail{}, SemanticRefs: []SemanticRefDecl{{Ref: "spec:trace.duplicate", Path: "docs/spec/a.md", TargetType: SemanticTargetDocument}}},
			{ID: "SPEC-b", NormalizedID: "SPEC-B", Kind: RecordKindSpec, Title: "B", Status: RecordStatusDraft, Path: "docs/spec/b.md", Spec: &SpecDetail{}, SemanticRefs: []SemanticRefDecl{{Ref: "spec:trace.duplicate", Path: "docs/spec/b.md", TargetType: SemanticTargetDocument}}},
			{ID: "INV-DOCS-001", NormalizedID: "INV-DOCS-001", Kind: RecordKindInvestigation, Title: "Investigation", Status: RecordStatusConcluded, Path: "docs/investigations/docs/INV-DOCS-001-test.md", Investigation: &InvestigationDetail{SourceRefs: []string{"ADR-001", "spec:trace.duplicate"}, FollowUpCandidates: []string{}}},
		},
	}
	idx.SemanticRefs = collectSemanticRefs(idx.Records)
	idx.SemanticRefSources = []SemanticRefSource{
		{Path: "docs/spec/a.md", RecordID: "SPEC-a", Decls: idx.Records[2].SemanticRefs},
		{Path: "docs/spec/b.md", RecordID: "SPEC-b", Decls: idx.Records[3].SemanticRefs},
	}

	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if !hasDiagnostic(resp.Diagnostics, DiagnosticDuplicateID) || !hasDiagnostic(resp.Diagnostics, DiagnosticDuplicateSemanticRef) {
		t.Fatalf("missing duplicate diagnostics: %#v", resp.Diagnostics)
	}
	for _, diagnostic := range resp.Diagnostics {
		if diagnostic.RecordID == "INV-DOCS-001" && diagnostic.Field != "" {
			t.Fatalf("field-specific duplicate target diagnostic was added: %#v", diagnostic)
		}
	}
}

func TestValidateRecordsSpecStatusMismatchRequiresBothStatuses(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/spec/missing-top-level-status.md", "---\ndesign_record:\n  id: SPEC-missing-top-level-status\n  kind: spec\n  status: draft\n---\n# Missing top-level status\n")

	idx := buildTestIndex(t, root)
	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if hasDiagnosticForRecord(resp.Diagnostics, DiagnosticSpecStatusMismatch, "SPEC-missing-top-level-status") {
		t.Fatalf("unexpected spec_status_mismatch: %#v", resp.Diagnostics)
	}
	if !hasDiagnosticForRecord(resp.Diagnostics, DiagnosticInvalidStatusForKind, "SPEC-missing-top-level-status") {
		t.Fatalf("missing invalid_status_for_kind for empty canonical status: %#v", resp.Diagnostics)
	}
}

func TestValidateRecordsADRFilenameNumberMissingIsMismatch(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/foo.md", "# 001: Example\n- **status**: accepted\n")

	idx := buildTestIndex(t, root)
	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if !hasDiagnosticForRecord(resp.Diagnostics, DiagnosticFilenameIDMismatch, "ADR-001") {
		t.Fatalf("missing filename_id_mismatch for ADR filename without number: %#v", resp.Diagnostics)
	}
}

func TestValidateRecordsKindFilter(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/001-decision.md", "# 001: Decision\n- **status**: draft\n")
	writeTestFile(t, root, "docs/spec/spec.md", "---\nstatus: accepted\ndesign_record:\n  id: SPEC-bad\n  kind: spec\n---\n# Spec\n")
	writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-003-01-task.md", "# TASK-MCP-003-01: Task\n- **id**: TASK-MCP-003-01\n- **status**: accepted\n- **date**: 2026-05-26\n- **work_item**: WORK-MCP-003\n- **source_requirement**: REQ-MCP-003\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n")

	idx := buildTestIndex(t, root)
	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{Kind: RecordKindDecision})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if !hasDiagnosticForRecord(resp.Diagnostics, DiagnosticInvalidStatusForKind, "ADR-001") {
		t.Fatalf("decision diagnostic missing: %#v", resp.Diagnostics)
	}
	if hasDiagnosticForRecord(resp.Diagnostics, DiagnosticInvalidStatusForKind, "SPEC-bad") {
		t.Fatalf("spec diagnostic leaked into decision filter: %#v", resp.Diagnostics)
	}

	resp, err = ValidateRecords(context.Background(), idx, ValidateRecordsRequest{Kind: RecordKindTask})
	if err != nil {
		t.Fatalf("ValidateRecords task: %v", err)
	}
	if !hasDiagnosticForRecord(resp.Diagnostics, DiagnosticInvalidStatusForKind, "TASK-MCP-003-01") {
		t.Fatalf("task diagnostic missing: %#v", resp.Diagnostics)
	}
	if hasDiagnosticForRecord(resp.Diagnostics, DiagnosticInvalidStatusForKind, "ADR-001") || hasDiagnosticForRecord(resp.Diagnostics, DiagnosticInvalidStatusForKind, "SPEC-bad") {
		t.Fatalf("non-task diagnostic leaked into task filter: %#v", resp.Diagnostics)
	}
}

func TestValidateRecordsIDRangeFilter(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/001-one.md", "# 001: One\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/adr/002-two.md", "# 002: Two\n- **status**: draft\n")
	writeTestFile(t, root, "docs/adr/003-three.md", "# 003: Three\n- **status**: draft\n")
	writeTestFile(t, root, "docs/spec/bad.md", "---\nstatus: accepted\ndesign_record:\n  id: SPEC-bad\n  kind: spec\n---\n# Bad\n")

	idx := buildTestIndex(t, root)

	tests := []struct {
		name      string
		req       ValidateRecordsRequest
		wantIDs   []string
		rejectIDs []string
	}{
		{
			name:      "inclusive endpoints",
			req:       ValidateRecordsRequest{Kind: RecordKindDecision, IDRange: &IDRange{From: "ADR-002", To: "ADR-002"}},
			wantIDs:   []string{"ADR-002"},
			rejectIDs: []string{"ADR-003", "SPEC-bad"},
		},
		{
			name:      "one sided from",
			req:       ValidateRecordsRequest{Kind: RecordKindDecision, IDRange: &IDRange{From: "ADR-003"}},
			wantIDs:   []string{"ADR-003"},
			rejectIDs: []string{"ADR-002", "SPEC-bad"},
		},
		{
			name:      "one sided to",
			req:       ValidateRecordsRequest{Kind: RecordKindDecision, IDRange: &IDRange{To: "ADR-002"}},
			wantIDs:   []string{"ADR-002"},
			rejectIDs: []string{"ADR-003", "SPEC-bad"},
		},
		{
			name:      "omitted kind behaves as decision",
			req:       ValidateRecordsRequest{IDRange: &IDRange{From: "ADR-002", To: "ADR-003"}},
			wantIDs:   []string{"ADR-002", "ADR-003"},
			rejectIDs: []string{"SPEC-bad"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ValidateRecords(context.Background(), idx, tt.req)
			if err != nil {
				t.Fatalf("ValidateRecords: %v", err)
			}
			for _, id := range tt.wantIDs {
				if !hasDiagnosticForRecord(resp.Diagnostics, DiagnosticInvalidStatusForKind, id) {
					t.Fatalf("missing diagnostic for %s in %#v", id, resp.Diagnostics)
				}
			}
			for _, id := range tt.rejectIDs {
				if hasDiagnosticForRecord(resp.Diagnostics, DiagnosticInvalidStatusForKind, id) {
					t.Fatalf("unexpected diagnostic for %s in %#v", id, resp.Diagnostics)
				}
			}
		})
	}
}

func TestValidateRecordsWorkflowIDRangeFilter(t *testing.T) {
	records := []Record{
		{ID: "REQ-DATA-001", NormalizedID: "REQ-DATA-001", Kind: RecordKindRequirement, Title: "Req 1", Status: RecordStatusCaptured, Path: "docs/requirements/data/REQ-DATA-001-test.md", Requirement: &RequirementDetail{}},
		{ID: "REQ-DATA-002", NormalizedID: "REQ-DATA-002", Kind: RecordKindRequirement, Title: "Req 2", Status: RecordStatusNotStarted, Path: "docs/requirements/data/REQ-DATA-002-test.md", Requirement: &RequirementDetail{}},
		{ID: "REQ-MCP-001", NormalizedID: "REQ-MCP-001", Kind: RecordKindRequirement, Title: "Req other domain", Status: RecordStatusNotStarted, Path: "docs/requirements/mcp/REQ-MCP-001-test.md", Requirement: &RequirementDetail{}},
		{ID: "WORK-DATA-001", NormalizedID: "WORK-DATA-001", Kind: RecordKindWorkItem, Title: "Work 1", Status: RecordStatusNotStarted, Path: "docs/work-items/data/WORK-DATA-001-test.md", WorkItem: &WorkItemDetail{}},
		{ID: "WORK-DATA-002", NormalizedID: "WORK-DATA-002", Kind: RecordKindWorkItem, Title: "Work 2", Status: RecordStatusAccepted, Path: "docs/work-items/data/WORK-DATA-002-test.md", WorkItem: &WorkItemDetail{}},
		{ID: "WORK-MCP-002", NormalizedID: "WORK-MCP-002", Kind: RecordKindWorkItem, Title: "Work other domain", Status: RecordStatusAccepted, Path: "docs/work-items/mcp/WORK-MCP-002-test.md", WorkItem: &WorkItemDetail{}},
		{ID: "TASK-MCP-007-01", NormalizedID: "TASK-MCP-007-01", Kind: RecordKindTask, Title: "Task 1", Status: RecordStatusNotStarted, Path: "docs/tasks/mcp/TASK-MCP-007-01-test.md", Task: &TaskDetail{}},
		{ID: "TASK-MCP-007-02", NormalizedID: "TASK-MCP-007-02", Kind: RecordKindTask, Title: "Task 2", Status: RecordStatusAccepted, Path: "docs/tasks/mcp/TASK-MCP-007-02-test.md", Task: &TaskDetail{}},
		{ID: "TASK-MCP-008-01", NormalizedID: "TASK-MCP-008-01", Kind: RecordKindTask, Title: "Task other work", Status: RecordStatusAccepted, Path: "docs/tasks/mcp/TASK-MCP-008-01-test.md", Task: &TaskDetail{}},
	}
	idx := workflowTestIndex(records...)

	tests := []struct {
		name      string
		req       ValidateRecordsRequest
		wantIDs   []string
		rejectIDs []string
	}{
		{
			name:      "requirement same domain range",
			req:       ValidateRecordsRequest{Kind: RecordKindRequirement, IDRange: &IDRange{From: "REQ-DATA-001", To: "REQ-DATA-002"}},
			wantIDs:   []string{"REQ-DATA-002"},
			rejectIDs: []string{"REQ-MCP-001", "WORK-DATA-002", "TASK-MCP-007-02"},
		},
		{
			name:      "work item same domain range",
			req:       ValidateRecordsRequest{Kind: RecordKindWorkItem, IDRange: &IDRange{From: "WORK-DATA-001", To: "WORK-DATA-002"}},
			wantIDs:   []string{"WORK-DATA-002"},
			rejectIDs: []string{"REQ-DATA-002", "WORK-MCP-002", "TASK-MCP-007-02"},
		},
		{
			name:      "task same work sequence range",
			req:       ValidateRecordsRequest{Kind: RecordKindTask, IDRange: &IDRange{From: "TASK-MCP-007-01", To: "TASK-MCP-007-02"}},
			wantIDs:   []string{"TASK-MCP-007-02"},
			rejectIDs: []string{"REQ-DATA-002", "WORK-DATA-002", "TASK-MCP-008-01"},
		},
		{
			name:      "empty range with explicit workflow kind behaves as kind filter",
			req:       ValidateRecordsRequest{Kind: RecordKindWorkItem, IDRange: &IDRange{}},
			wantIDs:   []string{"WORK-DATA-002", "WORK-MCP-002"},
			rejectIDs: []string{"REQ-DATA-002", "TASK-MCP-007-02"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ValidateRecords(context.Background(), idx, tt.req)
			if err != nil {
				t.Fatalf("ValidateRecords: %v", err)
			}
			for _, id := range tt.wantIDs {
				if !hasDiagnosticForRecord(resp.Diagnostics, DiagnosticInvalidStatusForKind, id) {
					t.Fatalf("missing diagnostic for %s in %#v", id, resp.Diagnostics)
				}
			}
			for _, id := range tt.rejectIDs {
				if hasDiagnosticForRecord(resp.Diagnostics, DiagnosticInvalidStatusForKind, id) {
					t.Fatalf("unexpected diagnostic for %s in %#v", id, resp.Diagnostics)
				}
			}
		})
	}
}

func TestValidateRecordsRequestErrors(t *testing.T) {
	idx := &Index{}
	tests := []struct {
		name string
		req  ValidateRecordsRequest
		code ErrorCode
	}{
		{
			name: "invalid kind",
			req:  ValidateRecordsRequest{Kind: RecordKind("milestone")},
			code: ErrorCodeInvalidRequest,
		},
		{
			name: "kind spec with id range",
			req:  ValidateRecordsRequest{Kind: RecordKindSpec, IDRange: &IDRange{From: "ADR-001"}},
			code: ErrorCodeInvalidIDRange,
		},
		{
			name: "kind task with ADR id range",
			req:  ValidateRecordsRequest{Kind: RecordKindTask, IDRange: &IDRange{From: "ADR-001"}},
			code: ErrorCodeInvalidIDRange,
		},
		{
			name: "SPEC range endpoint",
			req:  ValidateRecordsRequest{IDRange: &IDRange{From: "SPEC-design-records-mcp-schema"}},
			code: ErrorCodeInvalidIDRange,
		},
		{
			name: "malformed range endpoint",
			req:  ValidateRecordsRequest{IDRange: &IDRange{From: "ADR-x"}},
			code: ErrorCodeInvalidIDRange,
		},
		{
			name: "mixed workflow domains",
			req:  ValidateRecordsRequest{IDRange: &IDRange{From: "WORK-DATA-001", To: "WORK-MCP-010"}},
			code: ErrorCodeInvalidIDRange,
		},
		{
			name: "mixed task work sequences",
			req:  ValidateRecordsRequest{IDRange: &IDRange{From: "TASK-MCP-006-01", To: "TASK-MCP-007-05"}},
			code: ErrorCodeInvalidIDRange,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ValidateRecords(context.Background(), idx, tt.req)
			if err == nil {
				t.Fatal("ValidateRecords error = nil, want ToolError")
			}
			toolErr, ok := err.(*ToolError)
			if !ok {
				t.Fatalf("error = %T %v, want *ToolError", err, err)
			}
			if toolErr.Code != tt.code {
				t.Fatalf("error code = %q, want %q", toolErr.Code, tt.code)
			}
			if resp.OK || len(resp.Diagnostics) != 0 {
				t.Fatalf("request error response = %#v, want zero validation response", resp)
			}
		})
	}
}

func buildTestIndex(t *testing.T, root string) *Index {
	t.Helper()
	cfg, err := NewConfig(root)
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	return idx
}

func hasDiagnostic(diagnostics []Diagnostic, category DiagnosticCategory) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic.Category == category {
			return true
		}
	}
	return false
}

func hasDiagnosticForRecord(diagnostics []Diagnostic, category DiagnosticCategory, recordID string) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic.Category == category && diagnostic.RecordID == recordID {
			return true
		}
	}
	return false
}

func assertDiagnostic(t *testing.T, diagnostics []Diagnostic, category DiagnosticCategory, recordID string, targetID string) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Category == category && diagnostic.RecordID == recordID && diagnostic.TargetID == targetID {
			if diagnostic.Severity != DiagnosticSeverityError {
				t.Fatalf("severity = %q, want error", diagnostic.Severity)
			}
			if diagnostic.Path == "" {
				t.Fatalf("path missing for %#v", diagnostic)
			}
			return
		}
	}
	t.Fatalf("missing diagnostic category=%s record=%s target=%s in %#v", category, recordID, targetID, diagnostics)
}

func assertDiagnosticPath(t *testing.T, diagnostics []Diagnostic, category DiagnosticCategory, path string) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Category == category && diagnostic.Path == path {
			if diagnostic.Severity != DiagnosticSeverityError {
				t.Fatalf("severity = %q, want error", diagnostic.Severity)
			}
			return
		}
	}
	t.Fatalf("missing diagnostic category=%s path=%s in %#v", category, path, diagnostics)
}

func hasTarget(diagnostics []Diagnostic, targetID string) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic.TargetID == targetID {
			return true
		}
	}
	return false
}

func countDiagnostics(diagnostics []Diagnostic, category DiagnosticCategory) int {
	count := 0
	for _, diagnostic := range diagnostics {
		if diagnostic.Category == category {
			count++
		}
	}
	return count
}

func diagnosticsForRecord(diagnostics []Diagnostic, recordID string) []Diagnostic {
	var out []Diagnostic
	for _, diagnostic := range diagnostics {
		if diagnostic.RecordID == recordID {
			out = append(out, diagnostic)
		}
	}
	return out
}

func assertInvestigationDiagnostic(t *testing.T, diagnostics []Diagnostic, category DiagnosticCategory, severity DiagnosticSeverity, field, value, refStatus string) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Category == category && diagnostic.Severity == severity && diagnostic.Field == field && diagnostic.Value == value && diagnostic.RefStatus == refStatus {
			if diagnostic.RecordID == "" || diagnostic.Path == "" {
				t.Fatalf("record/path missing for %#v", diagnostic)
			}
			return
		}
	}
	t.Fatalf("missing investigation diagnostic category=%s severity=%s field=%s value=%s ref_status=%s in %#v", category, severity, field, value, refStatus, diagnostics)
}

func assertNoInvestigationDiagnosticValue(t *testing.T, diagnostics []Diagnostic, value string) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Value == value {
			t.Fatalf("unexpected investigation diagnostic for %s: %#v", value, diagnostic)
		}
	}
}

func assertWorkflowMetadataDiagnostic(t *testing.T, diagnostics []Diagnostic, category DiagnosticCategory, recordID, field, value string, valuePresent bool) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Category == category && diagnostic.RecordID == recordID && diagnostic.Field == field {
			if diagnostic.Severity != DiagnosticSeverityError {
				t.Fatalf("severity = %q, want error for %#v", diagnostic.Severity, diagnostic)
			}
			if diagnostic.Path == "" || diagnostic.Message == "" {
				t.Fatalf("path/message missing for %#v", diagnostic)
			}
			if diagnostic.Value != value || diagnostic.ValuePresent != valuePresent {
				t.Fatalf("value/valuePresent = %q/%v, want %q/%v for %#v", diagnostic.Value, diagnostic.ValuePresent, value, valuePresent, diagnostic)
			}
			return
		}
	}
	t.Fatalf("missing workflow metadata diagnostic category=%s record=%s field=%s in %#v", category, recordID, field, diagnostics)
}

func assertNoWorkflowMetadataDiagnostic(t *testing.T, diagnostics []Diagnostic, recordID, field string) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.RecordID == recordID && diagnostic.Field == field && isWorkflowMetadataDiagnostic(diagnostic.Category) {
			t.Fatalf("unexpected workflow metadata diagnostic: %#v", diagnostic)
		}
	}
}

func isWorkflowMetadataDiagnostic(category DiagnosticCategory) bool {
	switch category {
	case DiagnosticMissingRequiredMetadata, DiagnosticEmptyRequiredMetadata, DiagnosticInvalidMetadataValue:
		return true
	default:
		return false
	}
}

func workflowTestIndex(records ...Record) *Index {
	out := &Index{Records: make([]Record, len(records))}
	copy(out.Records, records)
	return out
}

func requirementTestRecord(id string, workItems []string) Record {
	return Record{
		ID:           id,
		NormalizedID: normalizeRecordID(id),
		Kind:         RecordKindRequirement,
		Title:        id,
		Status:       RecordStatusAccepted,
		Path:         "docs/requirements/mcp/" + id + "-test.md",
		Requirement:  &RequirementDetail{SourceRefs: []string{}, WorkItems: append([]string{}, workItems...)},
	}
}

func workItemTestRecord(id, sourceRequirement string, tasks []string) Record {
	return Record{
		ID:           id,
		NormalizedID: normalizeRecordID(id),
		Kind:         RecordKindWorkItem,
		Title:        id,
		Status:       RecordStatusInProgress,
		Path:         "docs/work-items/mcp/" + id + "-test.md",
		WorkItem: &WorkItemDetail{
			SourceRequirement: sourceRequirement,
			ImpactRefs:        []string{},
			Tasks:             append([]string{}, tasks...),
		},
	}
}

func taskTestRecord(id, workItem, sourceRequirement string, dependsOn []string) Record {
	return Record{
		ID:           id,
		NormalizedID: normalizeRecordID(id),
		Kind:         RecordKindTask,
		Title:        id,
		Status:       RecordStatusNotStarted,
		Path:         "docs/tasks/mcp/" + id + "-test.md",
		Task: &TaskDetail{
			WorkItem:          workItem,
			SourceRequirement: sourceRequirement,
			Estimate:          "0.5d",
			DependsOn:         append([]string{}, dependsOn...),
			Outputs:           []string{},
		},
	}
}

func writeWorkflowHappyPathFixture(t *testing.T, root string) {
	t.Helper()
	writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-003-valid.md", "# REQ-MCP-003: Valid requirement\n- **id**: REQ-MCP-003\n- **status**: accepted\n- **date**: 2026-05-25\n- **source_refs**:\n- **work_items**:\n  - WORK-MCP-003\n\n## Requirement\n\nRequirement content.\n\n## Required Outcome\n\nOutcome content.\n")
	writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-003-valid.md", "# WORK-MCP-003: Valid work item\n- **id**: WORK-MCP-003\n- **status**: in_progress\n- **date**: 2026-05-26\n- **source_requirement**: REQ-MCP-003\n- **impact_refs**:\n- **tasks**:\n  - TASK-MCP-003-01\n  - TASK-MCP-003-02\n")
	writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-003-01-valid.md", "# TASK-MCP-003-01: First task\n- **id**: TASK-MCP-003-01\n- **status**: not_started\n- **date**: 2026-05-26\n- **work_item**: WORK-MCP-003\n- **source_requirement**: REQ-MCP-003\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n  - first\n")
	writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-003-02-valid.md", "# TASK-MCP-003-02: Second task\n- **id**: TASK-MCP-003-02\n- **status**: not_started\n- **date**: 2026-05-26\n- **work_item**: WORK-MCP-003\n- **source_requirement**: REQ-MCP-003\n- **estimate**: 0.5d\n- **depends_on**:\n  - TASK-MCP-003-01\n- **outputs**:\n  - second\n")
}

func isWorkflowRelationDiagnostic(category DiagnosticCategory) bool {
	switch category {
	case DiagnosticUnresolvedWorkflowRelation,
		DiagnosticInvalidWorkflowRelationTarget,
		DiagnosticWorkflowRelationMismatch,
		DiagnosticWorkflowSourceReqMismatch:
		return true
	default:
		return false
	}
}

func assertNoWorkflowRelationDiagnostics(t *testing.T, diagnostics []Diagnostic) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if isWorkflowRelationDiagnostic(diagnostic.Category) {
			t.Fatalf("unexpected workflow relation diagnostic: %#v", diagnostic)
		}
	}
}

func assertWorkflowDiagnostic(t *testing.T, diagnostics []Diagnostic, category DiagnosticCategory, recordID, field, value, refStatus, targetID string) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Category == category && diagnostic.RecordID == recordID && diagnostic.Field == field && diagnostic.Value == value && diagnostic.RefStatus == refStatus && diagnostic.TargetID == targetID {
			if diagnostic.Severity != DiagnosticSeverityError {
				t.Fatalf("severity = %q, want error for %#v", diagnostic.Severity, diagnostic)
			}
			if diagnostic.Path == "" {
				t.Fatalf("path missing for %#v", diagnostic)
			}
			return
		}
	}
	t.Fatalf("missing workflow diagnostic category=%s record=%s field=%s value=%s ref_status=%s target=%s in %#v", category, recordID, field, value, refStatus, targetID, diagnostics)
}

func assertRequiredSectionDiagnostic(t *testing.T, diagnostics []Diagnostic, category DiagnosticCategory, recordID, section, status string) {
	t.Helper()
	for _, d := range diagnostics {
		if d.Category == category && d.RecordID == recordID && d.Section == section && d.Status == status {
			if d.Severity != DiagnosticSeverityError {
				t.Fatalf("severity = %q, want error for %#v", d.Severity, d)
			}
			if d.Path == "" {
				t.Fatalf("path missing for %#v", d)
			}
			if d.Message == "" {
				t.Fatalf("message missing for %#v", d)
			}
			return
		}
	}
	t.Fatalf("missing required section diagnostic category=%s record=%s section=%s status=%s in %#v", category, recordID, section, status, diagnostics)
}

func assertSectionHeadingCaseMismatchDiagnostic(t *testing.T, diagnostics []Diagnostic, recordID, section, actualHeading, status string) {
	t.Helper()
	for _, d := range diagnostics {
		if d.Category != diagnosticSectionHeadingCaseMismatch || d.RecordID != recordID || d.Section != section || d.Status != status {
			continue
		}
		if d.Severity != DiagnosticSeverityInfo {
			t.Fatalf("severity = %q, want info for %#v", d.Severity, d)
		}
		if d.Path == "" {
			t.Fatalf("path missing for %#v", d)
		}
		if d.Message == "" {
			t.Fatalf("message missing for %#v", d)
		}
		var raw map[string]any
		data, err := json.Marshal(d)
		if err != nil {
			t.Fatalf("marshal diagnostic: %v", err)
		}
		if err := json.Unmarshal(data, &raw); err != nil {
			t.Fatalf("unmarshal diagnostic: %v\n%s", err, data)
		}
		if raw["actual_heading"] != actualHeading {
			t.Fatalf("actual_heading = %#v, want %q in %#v", raw["actual_heading"], actualHeading, raw)
		}
		if len(d.CandidateHeadings) == 0 {
			t.Fatalf("candidate_headings missing for %#v", d)
		}
		return
	}
	t.Fatalf("missing section_heading_case_mismatch record=%s section=%s actual=%s status=%s in %#v", recordID, section, actualHeading, status, diagnostics)
}

func assertNoRequiredSectionDiagnosticForRecord(t *testing.T, diagnostics []Diagnostic, recordID string) {
	t.Helper()
	for _, d := range diagnostics {
		if (d.Category == DiagnosticMissingRequiredSection || d.Category == DiagnosticEmptyRequiredSection) && d.RecordID == recordID {
			t.Fatalf("unexpected required section diagnostic for %s: %#v", recordID, d)
		}
	}
}

func TestValidateRecordsRequiredNarrativeSections(t *testing.T) {
	const (
		workItemMetaHeader = "- **id**: WORK-MCP-007\n- **status**: done\n- **date**: 2026-06-03\n- **source_requirement**: REQ-MCP-007\n- **impact_refs**:\n- **tasks**:\n"
		taskMetaHeader     = "- **id**: TASK-MCP-007-01\n- **status**: done\n- **date**: 2026-06-03\n- **work_item**: WORK-MCP-007\n- **source_requirement**: REQ-MCP-007\n- **estimate**: 1d\n- **depends_on**:\n- **outputs**:\n"
		reqMetaHeader      = "- **id**: REQ-MCP-007\n- **status**: accepted\n- **date**: 2026-06-03\n- **source_refs**:\n- **work_items**:\n"
	)

	t.Run("work_item done missing Goal", func(t *testing.T) {
		root := t.TempDir()
		content := "# WORK-MCP-007: Test\n" + workItemMetaHeader + "\n## Boundary\n\nBoundary text.\n\n## Evidence\n\nEvidence text.\n"
		writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindWorkItem})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		assertRequiredSectionDiagnostic(t, resp.Diagnostics, DiagnosticMissingRequiredSection, "WORK-MCP-007", "Goal", "done")
	})

	t.Run("work_item done empty Goal", func(t *testing.T) {
		root := t.TempDir()
		content := "# WORK-MCP-007: Test\n" + workItemMetaHeader + "\n## Goal\n\n## Boundary\n\nBoundary text.\n\n## Evidence\n\nEvidence text.\n"
		writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindWorkItem})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		assertRequiredSectionDiagnostic(t, resp.Diagnostics, DiagnosticEmptyRequiredSection, "WORK-MCP-007", "Goal", "done")
	})

	t.Run("work_item done empty Boundary", func(t *testing.T) {
		root := t.TempDir()
		content := "# WORK-MCP-007: Test\n" + workItemMetaHeader + "\n## Goal\n\nGoal text.\n\n## Boundary\n\n## Evidence\n\nEvidence text.\n"
		writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindWorkItem})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		assertRequiredSectionDiagnostic(t, resp.Diagnostics, DiagnosticEmptyRequiredSection, "WORK-MCP-007", "Boundary", "done")
	})

	t.Run("work_item done empty Evidence", func(t *testing.T) {
		root := t.TempDir()
		content := "# WORK-MCP-007: Test\n" + workItemMetaHeader + "\n## Goal\n\nGoal text.\n\n## Boundary\n\nBoundary text.\n\n## Evidence\n"
		writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindWorkItem})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		assertRequiredSectionDiagnostic(t, resp.Diagnostics, DiagnosticEmptyRequiredSection, "WORK-MCP-007", "Evidence", "done")
	})

	t.Run("work_item done all sections non-empty", func(t *testing.T) {
		root := t.TempDir()
		content := "# WORK-MCP-007: Test\n" + workItemMetaHeader + "\n## Goal\n\nGoal text.\n\n## Boundary\n\nBoundary text.\n\n## Evidence\n\nEvidence text.\n"
		writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindWorkItem})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		assertNoRequiredSectionDiagnosticForRecord(t, resp.Diagnostics, "WORK-MCP-007")
	})

	t.Run("task done required sections enforced", func(t *testing.T) {
		for _, sectionName := range []string{"Goal", "Work", "Done condition", "Verification", "Evidence"} {
			sectionName := sectionName
			t.Run("missing "+sectionName, func(t *testing.T) {
				root := t.TempDir()
				allSections := map[string]string{
					"Goal":           "Goal text.",
					"Work":           "Work text.",
					"Done condition": "Done condition text.",
					"Verification":   "Verification text.",
					"Evidence":       "Evidence text.",
				}
				var body strings.Builder
				body.WriteString("# TASK-MCP-007-01: Test\n")
				body.WriteString(taskMetaHeader)
				for _, s := range []string{"Goal", "Work", "Done condition", "Verification", "Evidence"} {
					if s == sectionName {
						continue
					}
					body.WriteString("\n## " + s + "\n\n" + allSections[s] + "\n")
				}
				writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-007-01-test.md", body.String())
				resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindTask})
				if err != nil {
					t.Fatalf("ValidateRecords: %v", err)
				}
				assertRequiredSectionDiagnostic(t, resp.Diagnostics, DiagnosticMissingRequiredSection, "TASK-MCP-007-01", sectionName, "done")
			})
		}
	})

	t.Run("requirement accepted Requirement required", func(t *testing.T) {
		root := t.TempDir()
		content := "# REQ-MCP-007: Test\n" + reqMetaHeader + "\n## Required Outcome\n\nOutcome text.\n"
		writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindRequirement})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		assertRequiredSectionDiagnostic(t, resp.Diagnostics, DiagnosticMissingRequiredSection, "REQ-MCP-007", "Requirement", "accepted")
	})

	t.Run("requirement accepted Required Outcome required", func(t *testing.T) {
		root := t.TempDir()
		content := "# REQ-MCP-007: Test\n" + reqMetaHeader + "\n## Requirement\n\nRequirement text.\n"
		writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindRequirement})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		assertRequiredSectionDiagnostic(t, resp.Diagnostics, DiagnosticMissingRequiredSection, "REQ-MCP-007", "Required Outcome", "accepted")
	})

	t.Run("requirement accepted Evidence not required", func(t *testing.T) {
		root := t.TempDir()
		content := "# REQ-MCP-007: Test\n" + reqMetaHeader + "\n## Requirement\n\nRequirement text.\n\n## Required Outcome\n\nOutcome text.\n"
		writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindRequirement})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		assertNoRequiredSectionDiagnosticForRecord(t, resp.Diagnostics, "REQ-MCP-007")
		for _, d := range resp.Diagnostics {
			if d.RecordID == "REQ-MCP-007" && d.Section == "Evidence" {
				t.Fatalf("Evidence incorrectly required for REQ accepted: %#v", d)
			}
		}
	})

	t.Run("requirement accepted Boundary not required", func(t *testing.T) {
		root := t.TempDir()
		content := "# REQ-MCP-007: Test\n" + reqMetaHeader + "\n## Requirement\n\nRequirement text.\n\n## Required Outcome\n\nOutcome text.\n"
		writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindRequirement})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		for _, d := range resp.Diagnostics {
			if d.RecordID == "REQ-MCP-007" && d.Section == "Boundary" {
				t.Fatalf("Boundary incorrectly required for REQ accepted: %#v", d)
			}
		}
	})

	t.Run("requirement accepted Explicitly Excluded Scope not required", func(t *testing.T) {
		root := t.TempDir()
		content := "# REQ-MCP-007: Test\n" + reqMetaHeader + "\n## Requirement\n\nRequirement text.\n\n## Required Outcome\n\nOutcome text.\n"
		writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindRequirement})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		for _, d := range resp.Diagnostics {
			if d.RecordID == "REQ-MCP-007" && d.Section == "Explicitly Excluded Scope" {
				t.Fatalf("Explicitly Excluded Scope incorrectly required for REQ accepted: %#v", d)
			}
		}
	})

	t.Run("work_item not_started empty Evidence not checked", func(t *testing.T) {
		root := t.TempDir()
		content := "# WORK-MCP-007: Test\n- **id**: WORK-MCP-007\n- **status**: not_started\n- **date**: 2026-06-03\n- **source_requirement**: REQ-MCP-007\n- **impact_refs**:\n- **tasks**:\n\n## Goal\n\n## Evidence\n"
		writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindWorkItem})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		assertNoRequiredSectionDiagnosticForRecord(t, resp.Diagnostics, "WORK-MCP-007")
	})

	t.Run("task todo empty Evidence not checked", func(t *testing.T) {
		root := t.TempDir()
		content := "# TASK-MCP-007-01: Test\n- **id**: TASK-MCP-007-01\n- **status**: not_started\n- **date**: 2026-06-03\n- **work_item**: WORK-MCP-007\n- **source_requirement**: REQ-MCP-007\n- **estimate**: 1d\n- **depends_on**:\n- **outputs**:\n\n## Goal\n\n## Evidence\n"
		writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-007-01-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindTask})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		assertNoRequiredSectionDiagnosticForRecord(t, resp.Diagnostics, "TASK-MCP-007-01")
	})

	t.Run("requirement captured empty Required Outcome not checked", func(t *testing.T) {
		root := t.TempDir()
		content := "# REQ-MCP-007: Test\n- **id**: REQ-MCP-007\n- **status**: captured\n- **date**: 2026-06-03\n- **source_refs**:\n- **work_items**:\n\n## Requirement\n\n## Required Outcome\n"
		writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindRequirement})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		assertNoRequiredSectionDiagnosticForRecord(t, resp.Diagnostics, "REQ-MCP-007")
	})

	t.Run("missing_required_section diagnostic fields", func(t *testing.T) {
		root := t.TempDir()
		content := "# WORK-MCP-007: Test\n" + workItemMetaHeader + "\n## Boundary\n\nBoundary text.\n\n## Evidence\n\nEvidence text.\n"
		writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindWorkItem})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		var found *Diagnostic
		for i := range resp.Diagnostics {
			d := &resp.Diagnostics[i]
			if d.Category == DiagnosticMissingRequiredSection && d.RecordID == "WORK-MCP-007" && d.Section == "Goal" {
				found = d
				break
			}
		}
		if found == nil {
			t.Fatalf("missing_required_section diagnostic not found: %#v", resp.Diagnostics)
		}
		if found.Severity != DiagnosticSeverityError {
			t.Fatalf("severity = %q, want error", found.Severity)
		}
		if found.RecordID != "WORK-MCP-007" {
			t.Fatalf("record_id = %q, want WORK-MCP-007", found.RecordID)
		}
		if found.Path == "" {
			t.Fatal("path is empty")
		}
		if found.Section != "Goal" {
			t.Fatalf("section = %q, want Goal", found.Section)
		}
		if found.Status != "done" {
			t.Fatalf("status = %q, want done", found.Status)
		}
		if found.Message == "" {
			t.Fatal("message is empty")
		}
	})

	t.Run("empty_required_section diagnostic fields", func(t *testing.T) {
		root := t.TempDir()
		content := "# WORK-MCP-007: Test\n" + workItemMetaHeader + "\n## Goal\n\n## Boundary\n\nBoundary text.\n\n## Evidence\n\nEvidence text.\n"
		writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-007-test.md", content)
		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindWorkItem})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		var found *Diagnostic
		for i := range resp.Diagnostics {
			d := &resp.Diagnostics[i]
			if d.Category == DiagnosticEmptyRequiredSection && d.RecordID == "WORK-MCP-007" && d.Section == "Goal" {
				found = d
				break
			}
		}
		if found == nil {
			t.Fatalf("empty_required_section diagnostic not found: %#v", resp.Diagnostics)
		}
		if found.Severity != DiagnosticSeverityError {
			t.Fatalf("severity = %q, want error", found.Severity)
		}
		if found.RecordID != "WORK-MCP-007" {
			t.Fatalf("record_id = %q, want WORK-MCP-007", found.RecordID)
		}
		if found.Path == "" {
			t.Fatal("path is empty")
		}
		if found.Section != "Goal" {
			t.Fatalf("section = %q, want Goal", found.Section)
		}
		if found.Status != "done" {
			t.Fatalf("status = %q, want done", found.Status)
		}
		if found.Message == "" {
			t.Fatal("message is empty")
		}
	})
}

func TestRequiredSectionHeadingCaseMismatchDiagnostics(t *testing.T) {
	const doneTaskMetadata = "- **id**: TASK-MCP-021-01\n- **status**: done\n- **date**: 2026-06-05\n- **work_item**: WORK-MCP-021\n- **source_requirement**: REQ-MCP-021\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n"

	t.Run("gated task emits strict missing error plus repair info", func(t *testing.T) {
		root := t.TempDir()
		writeHeadingCaseMismatchWorkflowParents(t, root)
		content := "# TASK-MCP-021-01: Heading case mismatch\n" + doneTaskMetadata + "\n## Goal\n\nGoal text.\n\n## Work\n\nWork text.\n\n## Done Condition\n\nDone text.\n\n## Verification\n\nVerification text.\n\n## Evidence\n\nEvidence text.\n"
		writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-021-01-heading-case-mismatch.md", content)

		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindTask})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		if resp.OK {
			t.Fatal("OK = true, want false because canonical required section is still missing")
		}
		assertRequiredSectionDiagnostic(t, resp.Diagnostics, DiagnosticMissingRequiredSection, "TASK-MCP-021-01", "Done condition", "done")
		assertSectionHeadingCaseMismatchDiagnostic(t, resp.Diagnostics, "TASK-MCP-021-01", "Done condition", "Done Condition", "done")
	})

	t.Run("non-gated task does not emit repair info", func(t *testing.T) {
		root := t.TempDir()
		writeHeadingCaseMismatchWorkflowParents(t, root)
		content := "# TASK-MCP-021-01: Heading case mismatch\n" +
			strings.Replace(doneTaskMetadata, "- **status**: done", "- **status**: todo", 1) +
			"\n## Goal\n\nGoal text.\n\n## Work\n\nWork text.\n\n## Done Condition\n\nDone text.\n\n## Verification\n\nVerification text.\n\n## Evidence\n\nEvidence text.\n"
		writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-021-01-heading-case-mismatch.md", content)

		resp, err := ValidateRecords(context.Background(), buildTestIndex(t, root), ValidateRecordsRequest{Kind: RecordKindTask})
		if err != nil {
			t.Fatalf("ValidateRecords: %v", err)
		}
		if hasDiagnosticForRecord(resp.Diagnostics, diagnosticSectionHeadingCaseMismatch, "TASK-MCP-021-01") {
			t.Fatalf("unexpected section_heading_case_mismatch for non-gated task: %#v", resp.Diagnostics)
		}
	})
}

func writeHeadingCaseMismatchWorkflowParents(t *testing.T, root string) {
	t.Helper()
	writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-021-heading-case-mismatch.md", "# REQ-MCP-021: Heading case mismatch\n\n- **id**: REQ-MCP-021\n- **status**: captured\n- **date**: 2026-06-05\n- **source_refs**:\n- **work_items**:\n  - WORK-MCP-021\n")
	writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-021-heading-case-mismatch.md", "# WORK-MCP-021: Heading case mismatch\n\n- **id**: WORK-MCP-021\n- **status**: not_started\n- **date**: 2026-06-05\n- **source_requirement**: REQ-MCP-021\n- **impact_refs**:\n- **tasks**:\n  - TASK-MCP-021-01\n")
}
