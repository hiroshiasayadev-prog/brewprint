package designrecords

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestValidateRecordsSelectorMatrix(t *testing.T) {
	idx := buildCurrentValidationIndex(t)

	all, err := ValidateCurrentRecords(context.Background(), idx, CurrentValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateCurrentRecords all: %v", err)
	}
	if all.Scope != "all" || all.OK {
		t.Fatalf("all response scope/ok = %q/%v, want all/false", all.Scope, all.OK)
	}
	assertHasDiagnostic(t, all.Diagnostics, DiagnosticMissingRequiredMetadata, "spec:product.invalid.missing_parent")
	assertHasDiagnostic(t, all.Diagnostics, DiagnosticCategory("current_conflict"), "PRODUCT-REQ-MCP-777")
	assertHasDiagnostic(t, all.Diagnostics, DiagnosticUnresolvedWorkflowRelation, "DRMCP-TASK-MCP-001-01")

	product, err := ValidateCurrentRecords(context.Background(), idx, CurrentValidateRecordsRequest{AppNamespace: "product"})
	if err != nil {
		t.Fatalf("ValidateCurrentRecords app_namespace: %v", err)
	}
	if product.Scope != "app_namespace" {
		t.Fatalf("scope = %q, want app_namespace", product.Scope)
	}
	assertHasDiagnostic(t, product.Diagnostics, DiagnosticMissingRequiredMetadata, "spec:product.invalid.missing_parent")
	assertHasDiagnostic(t, product.Diagnostics, DiagnosticCategory("current_conflict"), "PRODUCT-REQ-MCP-777")
	assertNoDiagnosticForRecord(t, product.Diagnostics, "DRMCP-TASK-MCP-001-01")

	byRef, err := ValidateCurrentRecords(context.Background(), idx, CurrentValidateRecordsRequest{Ref: "spec:product.invalid.missing_parent"})
	if err != nil {
		t.Fatalf("ValidateCurrentRecords ref: %v", err)
	}
	if byRef.Scope != "ref" || byRef.Summary.Total != 1 || byRef.Summary.Invalid != 1 {
		t.Fatalf("ref response = %#v", byRef)
	}
	assertHasOnlyRecords(t, byRef.Diagnostics, "spec:product.invalid.missing_parent")

	var mcpReq ValidateRecordsRequest
	if err := json.Unmarshal([]byte(`{"app_namespace":"product"}`), &mcpReq); err != nil {
		t.Fatalf("unmarshal app_namespace selector: %v", err)
	}
	compat, err := ValidateRecords(context.Background(), idx, mcpReq)
	if err != nil {
		t.Fatalf("ValidateRecords compat app_namespace: %v", err)
	}
	assertHasDiagnostic(t, compat.Diagnostics, DiagnosticMissingRequiredMetadata, "spec:product.invalid.missing_parent")
	assertNoDiagnosticForRecord(t, compat.Diagnostics, "DRMCP-TASK-MCP-001-01")
}

func TestValidateRecordsRejectsUnsupportedSelectors(t *testing.T) {
	idx := buildCurrentValidationIndex(t)
	currentRejects := []CurrentValidateRecordsRequest{
		{AppNamespace: "product", Ref: "spec:product.valid"},
		{AppNamespace: "missing"},
		{Ref: "V01-ADR-001"},
		{Ref: "product/records/spec/valid/index.md"},
		{Ref: "REQ-MCP-001"},
	}
	for _, req := range currentRejects {
		if _, err := ValidateCurrentRecords(context.Background(), idx, req); err == nil {
			t.Fatalf("ValidateCurrentRecords(%#v) error = nil, want rejection", req)
		}
	}

	legacyRejects := []ValidateRecordsRequest{
		{Kind: RecordKindRequirement},
		{IDRange: &IDRange{From: "PRODUCT-REQ-MCP-001"}},
	}
	for _, req := range legacyRejects {
		if _, err := ValidateRecords(context.Background(), idx, req); err == nil {
			t.Fatalf("ValidateRecords(%#v) error = nil, want rejection", req)
		}
	}

	for _, raw := range []string{
		`null`,
		`{"kind":"requirement"}`,
		`{"domain":"mcp"}`,
		`{"id_range":{"from":"PRODUCT-REQ-MCP-001"}}`,
		`{"app_namespace":""}`,
		`{"app_namespace":null}`,
		`{"ref":""}`,
		`{"ref":null}`,
		`{"app_namespace":"product","ref":"spec:product.valid"}`,
		`{"app_namespace":"product","ref":""}`,
		`{"app_namespace":"product","ref":null}`,
		`{"app_namespace":"","ref":"spec:product.valid"}`,
		`{"app_namespace":null,"ref":"spec:product.valid"}`,
	} {
		var req ValidateRecordsRequest
		if err := json.Unmarshal([]byte(raw), &req); err == nil {
			t.Fatalf("Unmarshal(%s) error = nil, want rejection", raw)
		}
	}

	for _, raw := range []string{
		`{}`,
		`{"app_namespace":"product"}`,
		`{"ref":"spec:product.valid"}`,
	} {
		var req ValidateRecordsRequest
		if err := json.Unmarshal([]byte(raw), &req); err != nil {
			t.Fatalf("Unmarshal(%s): %v", raw, err)
		}
	}
}

func TestValidateRecordsDiagnosticCategorySeverityLocationMatrix(t *testing.T) {
	idx := buildCurrentValidationIndex(t)
	resp, err := ValidateCurrentRecords(context.Background(), idx, CurrentValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateCurrentRecords: %v", err)
	}
	assertDiagnosticShape(t, resp.Diagnostics, DiagnosticFilenameIDMismatch, DiagnosticSeverityError, "PRODUCT-ADR-MCP-010", "product", "product/records", "adr/mcp/PRODUCT-ADR-MCP-011-mismatch.md")
	assertDiagnosticShape(t, resp.Diagnostics, DiagnosticMissingRequiredMetadata, DiagnosticSeverityError, "spec:product.invalid.missing_parent", "product", "product/records", "spec/invalid/missing-parent.md")
	assertDiagnosticShape(t, resp.Diagnostics, DiagnosticMissingRequiredSection, DiagnosticSeverityError, "DRMCP-TASK-MCP-002-01", "drmcp", "drmcp/records", "tasks/mcp/DRMCP-TASK-MCP-002-01-done-missing-sections.md")
	assertDiagnosticShape(t, resp.Diagnostics, DiagnosticUnresolvedWorkflowRelation, DiagnosticSeverityError, "DRMCP-TASK-MCP-001-01", "drmcp", "drmcp/records", "tasks/mcp/DRMCP-TASK-MCP-001-01-cross-namespace.md")
	assertDiagnosticShape(t, resp.Diagnostics, DiagnosticCategory("current_conflict"), DiagnosticSeverityError, "PRODUCT-REQ-MCP-777", "product", "product/records", "requirements/mcp/PRODUCT-REQ-MCP-777-duplicate-a.md")
}

func TestValidateRecordsCurrentRelationExactLookup(t *testing.T) {
	idx := buildCurrentValidationIndex(t)
	resp, err := ValidateCurrentRecords(context.Background(), idx, CurrentValidateRecordsRequest{Ref: "DRMCP-WORK-MCP-001"})
	if err != nil {
		t.Fatalf("ValidateCurrentRecords: %v", err)
	}
	if hasDiagnostic(resp.Diagnostics, DiagnosticUnresolvedWorkflowRelation) || hasDiagnostic(resp.Diagnostics, DiagnosticWorkflowRelationMismatch) {
		t.Fatalf("cross-namespace exact relation should resolve cleanly: %#v", resp.Diagnostics)
	}

	resp, err = ValidateCurrentRecords(context.Background(), idx, CurrentValidateRecordsRequest{Ref: "DRMCP-TASK-MCP-001-01"})
	if err != nil {
		t.Fatalf("ValidateCurrentRecords task: %v", err)
	}
	assertWorkflowDiagnostic(t, resp.Diagnostics, DiagnosticUnresolvedWorkflowRelation, "DRMCP-TASK-MCP-001-01", "depends_on", "DRMCP-TASK-MCP-999-01", "unresolved", "DRMCP-TASK-MCP-999-01")
}

func TestValidateRecordsOrderDeterministicAndDedup(t *testing.T) {
	idx := buildCurrentValidationIndex(t)
	idx.ParseIssues = append(idx.ParseIssues, idx.ParseIssues...)
	idx.SemanticRefs = []SemanticRefDecl{
		{Ref: "spec:product.semantic_duplicate", Path: "product/records/spec/a.md", TargetType: SemanticTargetDocument},
		{Ref: "spec:product.semantic_duplicate", Path: "product/records/spec/b.md", TargetType: SemanticTargetDocument},
	}

	var first []string
	for i := 0; i < 5; i++ {
		resp, err := ValidateCurrentRecords(context.Background(), idx, CurrentValidateRecordsRequest{})
		if err != nil {
			t.Fatalf("ValidateCurrentRecords run %d: %v", i, err)
		}
		if hasDiagnostic(resp.Diagnostics, DiagnosticDuplicateSemanticRef) {
			t.Fatalf("semantic duplicate validator must be suppressed in T07: %#v", resp.Diagnostics)
		}
		keys := diagnosticKeys(resp.Diagnostics)
		if i == 0 {
			first = keys
			continue
		}
		if !reflect.DeepEqual(keys, first) {
			t.Fatalf("diagnostic order changed:\nfirst=%#v\nrun%d=%#v", first, i, keys)
		}
		if len(keys) != len(uniqueStrings(keys)) {
			t.Fatalf("diagnostics were not deduplicated: %#v", keys)
		}
	}
}

func TestValidateRecordsSemanticPerFileDiagnosticsExcluded(t *testing.T) {
	idx := buildCurrentValidationIndex(t)
	idx.SemanticRefSources = append(idx.SemanticRefSources, SemanticRefSource{
		RecordID: "spec:product.valid",
		Path:     "product/records/spec/valid/index.md",
		Headings: []Heading{
			{Text: "Duplicate", Level: 2},
			{Text: "Duplicate", Level: 3},
		},
		Decls: []SemanticRefDecl{
			{Ref: "not-a-semantic-ref", Path: "product/records/spec/valid/index.md", TargetType: SemanticTargetDocument},
			{Ref: "spec:product.semantic.missing", Path: "product/records/spec/valid/index.md", TargetType: SemanticTargetSection, Section: "Missing"},
			{Ref: "spec:product.semantic.ambiguous", Path: "product/records/spec/valid/index.md", TargetType: SemanticTargetSection, Section: "Duplicate"},
		},
	})

	all, err := ValidateCurrentRecords(context.Background(), idx, CurrentValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateCurrentRecords all: %v", err)
	}
	for _, category := range []DiagnosticCategory{
		DiagnosticInvalidSemanticRefDeclaration,
		DiagnosticMissingSectionTarget,
		DiagnosticAmbiguousSectionTarget,
	} {
		if hasDiagnostic(all.Diagnostics, category) {
			t.Fatalf("excluded semantic diagnostic %s was emitted: %#v", category, all.Diagnostics)
		}
	}

	bySemanticRef, err := ValidateCurrentRecords(context.Background(), idx, CurrentValidateRecordsRequest{Ref: "spec:product.semantic.missing"})
	if err != nil {
		t.Fatalf("ValidateCurrentRecords semantic ref selector: %v", err)
	}
	if bySemanticRef.Summary.Total != 0 || len(bySemanticRef.Diagnostics) != 0 {
		t.Fatalf("semantic declaration ref selected its source: %#v", bySemanticRef)
	}
}

func TestValidateRecordsDiagnosticStableKeyDedupIgnoresMessage(t *testing.T) {
	diagnostics := []Diagnostic{
		{
			Category: DiagnosticMissingRequiredMetadata,
			Severity: DiagnosticSeverityError,
			RecordID: "PRODUCT-REQ-MCP-001",
			Path:     "product/records/requirements/mcp/PRODUCT-REQ-MCP-001.md",
			Field:    "date",
			Message:  "first wording",
		},
		{
			Category: DiagnosticMissingRequiredMetadata,
			Severity: DiagnosticSeverityError,
			RecordID: "PRODUCT-REQ-MCP-001",
			Path:     "product/records/requirements/mcp/PRODUCT-REQ-MCP-001.md",
			Field:    "date",
			Message:  "second wording",
		},
	}

	got := sortAndDedupeDiagnostics(&Index{}, diagnostics)
	if len(got) != 1 {
		t.Fatalf("same structured diagnostic with different messages produced %d entries: %#v", len(got), got)
	}
}

func TestValidateRecordsDiagnosticOrderIgnoresMessage(t *testing.T) {
	first := []Diagnostic{
		{Category: DiagnosticMissingRequiredMetadata, Severity: DiagnosticSeverityError, RecordID: "PRODUCT-REQ-MCP-002", Field: "date", Message: "alpha"},
		{Category: DiagnosticMissingRequiredMetadata, Severity: DiagnosticSeverityError, RecordID: "PRODUCT-REQ-MCP-001", Field: "status", Message: "zulu"},
	}
	second := []Diagnostic{
		{Category: DiagnosticMissingRequiredMetadata, Severity: DiagnosticSeverityError, RecordID: "PRODUCT-REQ-MCP-002", Field: "date", Message: "zulu"},
		{Category: DiagnosticMissingRequiredMetadata, Severity: DiagnosticSeverityError, RecordID: "PRODUCT-REQ-MCP-001", Field: "status", Message: "alpha"},
	}

	first = sortAndDedupeDiagnostics(nil, first)
	second = sortAndDedupeDiagnostics(nil, second)
	if !reflect.DeepEqual(diagnosticKeys(first), diagnosticKeys(second)) {
		t.Fatalf("message-only changes altered stable ordering:\nfirst=%#v\nsecond=%#v", diagnosticKeys(first), diagnosticKeys(second))
	}
}

func TestValidateRecordsPortableLocationFailsClosed(t *testing.T) {
	cases := []struct {
		name        string
		recordsRoot string
		path        string
	}{
		{name: "records root", recordsRoot: "product/records", path: "product/records"},
		{name: "parent segment", recordsRoot: "product/records", path: "product/records/../outside.md"},
		{name: "dot segment", recordsRoot: "product/records", path: "product/records/./spec.md"},
		{name: "drive slash", recordsRoot: "C:/repo/product/records", path: "C:/repo/product/records/spec.md"},
		{name: "drive backslash", recordsRoot: `C:\repo\product\records`, path: `C:\repo\product\records`},
		{name: "UNC slash", recordsRoot: "//server/share/records", path: "//server/share/records/spec.md"},
		{name: "UNC backslash", recordsRoot: `\\server\share\records`, path: `\\server\share\records`},
		{name: "URI", recordsRoot: "file://repo/records", path: "file://repo/records/spec.md"},
		{name: "opaque file URI-like", recordsRoot: "file:repo/records", path: "file:repo/records/spec.md"},
		{name: "URN", recordsRoot: "urn:brewprint:records", path: "urn:brewprint:records/spec.md"},
		{name: "colon path", recordsRoot: "product:records", path: "product:records/spec.md"},
		{name: "absolute", recordsRoot: "/repo/records", path: "/repo/records/spec.md"},
		{name: "backslash", recordsRoot: "product/records", path: `product/records/spec\bad.md`},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			idx := &Index{RecordsEntries: []RecordsEntry{{
				AppNamespace: "product",
				RecordsRoot:  tc.recordsRoot,
			}}}
			if location := currentDiagnosticLocation(idx, tc.path); location != nil {
				t.Fatalf("currentDiagnosticLocation(%q, %q) = %#v, want nil", tc.recordsRoot, tc.path, location)
			}
		})
	}

	idx := &Index{RecordsEntries: []RecordsEntry{{
		AppNamespace: "product",
		RecordsRoot:  "product/records",
	}}}
	location := currentDiagnosticLocation(idx, "product/records/spec/valid/index.md")
	if location == nil || location.Path != "spec/valid/index.md" {
		t.Fatalf("valid portable location = %#v", location)
	}
}

func buildCurrentValidationIndex(t *testing.T) *Index {
	t.Helper()
	root := t.TempDir()
	writeValidationFile(t, root, "product/records/requirements/mcp/PRODUCT-REQ-MCP-001-product.md", "# PRODUCT-REQ-MCP-001: Product requirement\n\n- **id**: PRODUCT-REQ-MCP-001\n- **status**: accepted\n- **date**: 2026-06-29\n- **source_refs**: []\n- **work_items**: DRMCP-WORK-MCP-001\n\n## Requirement\n\nNeed product requirement.\n\n## Required Outcome\n\nOutcome.\n")
	writeValidationFile(t, root, "drmcp/records/work-items/mcp/DRMCP-WORK-MCP-001-cross-namespace.md", "# DRMCP-WORK-MCP-001: Cross namespace work\n\n- **id**: DRMCP-WORK-MCP-001\n- **status**: in_progress\n- **date**: 2026-06-29\n- **source_requirement**: PRODUCT-REQ-MCP-001\n- **impact_refs**: []\n- **tasks**: DRMCP-TASK-MCP-001-01, DRMCP-TASK-MCP-002-01\n")
	writeValidationFile(t, root, "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-01-cross-namespace.md", "# DRMCP-TASK-MCP-001-01: Cross namespace task\n\n- **id**: DRMCP-TASK-MCP-001-01\n- **status**: not_started\n- **date**: 2026-06-29\n- **work_item**: DRMCP-WORK-MCP-001\n- **source_requirement**: PRODUCT-REQ-MCP-001\n- **estimate**: 0.5d\n- **depends_on**: DRMCP-TASK-MCP-999-01\n- **outputs**: []\n")
	writeValidationFile(t, root, "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-002-01-done-missing-sections.md", "# DRMCP-TASK-MCP-002-01: Done missing sections\n\n- **id**: DRMCP-TASK-MCP-002-01\n- **status**: done\n- **date**: 2026-06-29\n- **work_item**: DRMCP-WORK-MCP-001\n- **source_requirement**: PRODUCT-REQ-MCP-001\n- **estimate**: 0.5d\n- **depends_on**: []\n- **outputs**: []\n")
	writeValidationFile(t, root, "product/records/spec/valid/index.md", "# Valid spec\n\n- **id**: `spec:product.valid`\n- **status**: accepted\n- **date**: 2026-06-29\n- **parent**: `spec:product`\n")
	writeValidationFile(t, root, "product/records/spec/invalid/missing-parent.md", "# Missing parent\n\n- **id**: `spec:product.invalid.missing_parent`\n- **status**: accepted\n- **date**: 2026-06-29\n")
	writeValidationFile(t, root, "product/records/adr/mcp/PRODUCT-ADR-MCP-011-mismatch.md", "# PRODUCT-ADR-MCP-010: Filename mismatch\n\n- **status**: accepted\n- **depends_on**: []\n- **supersedes**: []\n- **migrated_to_spec**: null\n")

	cfg, err := NormalizeConfig(root, []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: "product/records"},
		{AppNamespace: "drmcp", RecordsRoot: "drmcp/records"},
	})
	if err != nil {
		t.Fatalf("NormalizeConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	idx.ConflictGroups = append(idx.ConflictGroups, CurrentConflict{
		Ref: "PRODUCT-REQ-MCP-777",
		Sources: []string{
			"product/records/requirements/mcp/PRODUCT-REQ-MCP-777-duplicate-b.md",
			"product/records/requirements/mcp/PRODUCT-REQ-MCP-777-duplicate-a.md",
		},
	})
	return idx
}

func buildTestIndex(t *testing.T, root string) *Index {
	t.Helper()
	cfg := Config{
		Root: root,
		RecordsRoots: []RecordsEntry{{
			AppNamespace:    "",
			NamespacePrefix: "",
			RecordsRoot:     "records",
		}},
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	semanticSourcePaths := map[string]bool{}
	specRoot := filepath.Join(root, "records", "spec")
	_ = filepath.WalkDir(specRoot, func(path string, d os.DirEntry, err error) error {
		if err != nil || d == nil || d.IsDir() || filepath.Ext(path) != ".md" {
			return nil
		}
		data, readErr := os.ReadFile(path)
		if readErr != nil {
			return nil
		}
		rel, relErr := filepath.Rel(root, path)
		if relErr != nil {
			return nil
		}
		source, ok := parseSpecSemanticRefSource(filepath.ToSlash(rel), string(data))
		if !ok {
			return nil
		}
		idx.SemanticRefSources = append(idx.SemanticRefSources, source)
		idx.SemanticRefs = append(idx.SemanticRefs, source.Decls...)
		semanticSourcePaths[source.Path] = true
		return nil
	})
	if len(semanticSourcePaths) > 0 {
		filtered := idx.ParseIssues[:0]
		for _, issue := range idx.ParseIssues {
			if semanticSourcePaths[issue.Path] {
				continue
			}
			filtered = append(filtered, issue)
		}
		idx.ParseIssues = filtered
	}
	return idx
}

func writeValidationFile(t *testing.T, root, rel, content string) {
	t.Helper()
	writeTestFile(t, root, rel, content)
}

func assertHasDiagnostic(t *testing.T, diagnostics []Diagnostic, category DiagnosticCategory, recordID string) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Category == category && diagnostic.RecordID == recordID {
			return
		}
	}
	t.Fatalf("missing diagnostic category=%s record=%s in %#v", category, recordID, diagnostics)
}

func assertNoDiagnosticForRecord(t *testing.T, diagnostics []Diagnostic, recordID string) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.RecordID == recordID {
			t.Fatalf("unexpected diagnostic for %s: %#v", recordID, diagnostic)
		}
	}
}

func assertHasOnlyRecords(t *testing.T, diagnostics []Diagnostic, recordID string) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.RecordID != recordID {
			t.Fatalf("unexpected diagnostic record %q, want only %q: %#v", diagnostic.RecordID, recordID, diagnostics)
		}
	}
}

func assertDiagnosticShape(t *testing.T, diagnostics []Diagnostic, category DiagnosticCategory, severity DiagnosticSeverity, recordID, appNamespace, recordsRoot, relPath string) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Category != category || diagnostic.Severity != severity || diagnostic.RecordID != recordID {
			continue
		}
		if diagnostic.Location == nil {
			t.Fatalf("location missing for %#v", diagnostic)
		}
		if diagnostic.Location.SourceScope != "current" || diagnostic.Location.AppNamespace != appNamespace || diagnostic.Location.RecordsRoot != recordsRoot || diagnostic.Location.Path != relPath {
			t.Fatalf("location = %#v, want current/%s/%s/%s", diagnostic.Location, appNamespace, recordsRoot, relPath)
		}
		if strings.Contains(diagnostic.Location.Path, `\`) || strings.HasPrefix(diagnostic.Location.Path, "/") {
			t.Fatalf("location path is not portable repository-relative: %#v", diagnostic.Location)
		}
		return
	}
	t.Fatalf("missing diagnostic shape category=%s record=%s in %#v", category, recordID, diagnostics)
}

func hasDiagnostic(diagnostics []Diagnostic, category DiagnosticCategory) bool {
	for _, diagnostic := range diagnostics {
		if diagnostic.Category == category {
			return true
		}
	}
	return false
}

func assertWorkflowDiagnostic(t *testing.T, diagnostics []Diagnostic, category DiagnosticCategory, recordID, field, value, refStatus, targetID string) {
	t.Helper()
	for _, diagnostic := range diagnostics {
		if diagnostic.Category == category && diagnostic.RecordID == recordID && diagnostic.Field == field && diagnostic.Value == value && diagnostic.RefStatus == refStatus && diagnostic.TargetID == targetID {
			if diagnostic.Severity != DiagnosticSeverityError {
				t.Fatalf("severity = %q, want error for %#v", diagnostic.Severity, diagnostic)
			}
			return
		}
	}
	t.Fatalf("missing workflow diagnostic category=%s record=%s field=%s value=%s ref_status=%s target=%s in %#v", category, recordID, field, value, refStatus, targetID, diagnostics)
}

func diagnosticKeys(diagnostics []Diagnostic) []string {
	keys := make([]string, 0, len(diagnostics))
	for _, diagnostic := range diagnostics {
		keys = append(keys, diagnosticStableKey(diagnostic))
	}
	return keys
}

func uniqueStrings(values []string) map[string]bool {
	out := map[string]bool{}
	for _, value := range values {
		out[value] = true
	}
	return out
}
