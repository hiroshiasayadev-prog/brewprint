package designrecords

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestBuildIndexEmptyState(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "test", "records"), 0o755); err != nil {
		t.Fatalf("MkdirAll: %v", err)
	}
	cfg, err := NewConfig(root, "test/records")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	if idx.Root != cfg.Root {
		t.Fatalf("Root = %q, want %q", idx.Root, cfg.Root)
	}
	if len(idx.Records) != 0 || len(idx.Candidates) != 0 || len(idx.ConflictGroups) != 0 {
		t.Fatalf("empty index = records:%d candidates:%d conflicts:%d", len(idx.Records), len(idx.Candidates), len(idx.ConflictGroups))
	}
}

func TestBuildIndexRejectsInvalidMandatoryRoot(t *testing.T) {
	root := t.TempDir()
	cfg := Config{
		Root: root,
		RecordsRoots: []RecordsEntry{{
			AppNamespace:    "product",
			RecordsRoot:     "product/records",
			NamespacePrefix: "PRODUCT-",
		}},
	}
	if _, err := BuildIndex(context.Background(), cfg); err == nil {
		t.Fatal("BuildIndex missing mandatory current root: want error, got nil")
	}
}

func TestCurrentActiveIndexMultiRootOrderIndependent(t *testing.T) {
	root := t.TempDir()

	writeTestFile(t, root, "product/records/adr/spec/PRODUCT-ADR-SPEC-901-first.md", currentADRSource("PRODUCT-ADR-SPEC-901", "First"))
	writeTestFile(t, root, "product/records/adr/PRODUCT-ADR-OPS-902-flat.md", currentADRSource("PRODUCT-ADR-OPS-902", "Flat compatibility"))
	writeTestFile(t, root, "product/records/spec/alpha/index.md", currentSpecSource("spec:product.alpha", "Alpha index", "spec:product"))
	writeTestFile(t, root, "product/records/spec/alpha.md", currentSpecSource("spec:product.alpha", "Alpha leaf", "spec:product"))
	writeTestFile(t, root, "product/records/spec/beta/index.md", currentSpecSource("spec:product.beta", "Beta", "spec:product"))

	writeTestFile(t, root, "drmcp/records/investigations/mcp/DRMCP-INV-MCP-001-first.md", currentInvestigationSource("DRMCP-INV-MCP-001", "First"))
	writeTestFile(t, root, "drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-first.md", currentRequirementSource("DRMCP-REQ-MCP-001", "First"))
	writeTestFile(t, root, "drmcp/records/requirements/mcp/DRMCP-REQ-MCP-003-filename.md", currentRequirementSource("DRMCP-REQ-MCP-004", "Filename mismatch"))
	writeTestFile(t, root, "drmcp/records/work-items/mcp/DRMCP-WORK-MCP-001-first.md", currentWorkItemSource("DRMCP-WORK-MCP-001", "First"))
	writeTestFile(t, root, "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-01-first.md", currentTaskSource("DRMCP-TASK-MCP-001-01", "First"))

	writeTestFile(t, root, "drmcp/records/requirements/mcp/DRMCP-REQ-MCP-980-duplicate-b.md", currentRequirementSource("DRMCP-REQ-MCP-980", "Duplicate B"))
	writeTestFile(t, root, "drmcp/records/requirements/mcp/DRMCP-REQ-MCP-980-duplicate-a.md", currentRequirementSource("DRMCP-REQ-MCP-980", "Duplicate A"))
	writeTestFile(t, root, "drmcp/records/requirements/mcp/DRMCP-REQ-MCP-990-duplicate-b.md", currentRequirementSource("DRMCP-REQ-MCP-990", "Duplicate B"))
	writeTestFile(t, root, "drmcp/records/requirements/mcp/DRMCP-REQ-MCP-990-duplicate-a.md", currentRequirementSource("DRMCP-REQ-MCP-990", "Duplicate A"))

	rootsAB := []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: "product/records"},
		{AppNamespace: "drmcp", RecordsRoot: "drmcp/records"},
	}
	rootsBA := []CurrentRoot{
		{AppNamespace: "drmcp", RecordsRoot: "drmcp/records"},
		{AppNamespace: "product", RecordsRoot: "product/records"},
	}
	cfgAB, err := NormalizeConfig(root, rootsAB)
	if err != nil {
		t.Fatalf("NormalizeConfig AB: %v", err)
	}
	cfgBA, err := NormalizeConfig(root, rootsBA)
	if err != nil {
		t.Fatalf("NormalizeConfig BA: %v", err)
	}
	idxAB, err := BuildIndex(context.Background(), cfgAB)
	if err != nil {
		t.Fatalf("BuildIndex AB: %v", err)
	}
	idxBA, err := BuildIndex(context.Background(), cfgBA)
	if err != nil {
		t.Fatalf("BuildIndex BA: %v", err)
	}

	if !reflect.DeepEqual(idxAB.Records, idxBA.Records) {
		t.Fatalf("records depend on root order:\nAB=%#v\nBA=%#v", idxAB.Records, idxBA.Records)
	}
	if !reflect.DeepEqual(idxAB.Candidates, idxBA.Candidates) {
		t.Fatalf("candidates depend on root order:\nAB=%#v\nBA=%#v", idxAB.Candidates, idxBA.Candidates)
	}
	if !reflect.DeepEqual(idxAB.ConflictGroups, idxBA.ConflictGroups) {
		t.Fatalf("conflicts depend on root order:\nAB=%#v\nBA=%#v", idxAB.ConflictGroups, idxBA.ConflictGroups)
	}
	if !reflect.DeepEqual(idxAB.ParseIssues, idxBA.ParseIssues) {
		t.Fatalf("parse issues depend on root order:\nAB=%#v\nBA=%#v", idxAB.ParseIssues, idxBA.ParseIssues)
	}
	if !reflect.DeepEqual(idxAB.PathIssues, idxBA.PathIssues) {
		t.Fatalf("path issues depend on root order:\nAB=%#v\nBA=%#v", idxAB.PathIssues, idxBA.PathIssues)
	}

	assertIndexStrings(t, recordIndexKeys(idxAB.Records), []string{
		"DRMCP-INV-MCP-001|drmcp/records/investigations/mcp/DRMCP-INV-MCP-001-first.md",
		"DRMCP-REQ-MCP-001|drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-first.md",
		"DRMCP-REQ-MCP-004|drmcp/records/requirements/mcp/DRMCP-REQ-MCP-003-filename.md",
		"DRMCP-TASK-MCP-001-01|drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-01-first.md",
		"DRMCP-WORK-MCP-001|drmcp/records/work-items/mcp/DRMCP-WORK-MCP-001-first.md",
		"PRODUCT-ADR-OPS-902|product/records/adr/PRODUCT-ADR-OPS-902-flat.md",
		"PRODUCT-ADR-SPEC-901|product/records/adr/spec/PRODUCT-ADR-SPEC-901-first.md",
		"spec:product.beta|product/records/spec/beta/index.md",
	})
	assertIndexStrings(t, candidateIndexKeys(idxAB.Candidates), []string{
		"DRMCP-INV-MCP-001|drmcp/records/investigations/mcp/DRMCP-INV-MCP-001-first.md",
		"DRMCP-REQ-MCP-001|drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-first.md",
		"DRMCP-REQ-MCP-004|drmcp/records/requirements/mcp/DRMCP-REQ-MCP-003-filename.md",
		"DRMCP-REQ-MCP-980|drmcp/records/requirements/mcp/DRMCP-REQ-MCP-980-duplicate-a.md",
		"DRMCP-REQ-MCP-980|drmcp/records/requirements/mcp/DRMCP-REQ-MCP-980-duplicate-b.md",
		"DRMCP-REQ-MCP-990|drmcp/records/requirements/mcp/DRMCP-REQ-MCP-990-duplicate-a.md",
		"DRMCP-REQ-MCP-990|drmcp/records/requirements/mcp/DRMCP-REQ-MCP-990-duplicate-b.md",
		"DRMCP-TASK-MCP-001-01|drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-01-first.md",
		"DRMCP-WORK-MCP-001|drmcp/records/work-items/mcp/DRMCP-WORK-MCP-001-first.md",
		"PRODUCT-ADR-OPS-902|product/records/adr/PRODUCT-ADR-OPS-902-flat.md",
		"PRODUCT-ADR-SPEC-901|product/records/adr/spec/PRODUCT-ADR-SPEC-901-first.md",
		"spec:product.alpha|product/records/spec/alpha.md",
		"spec:product.alpha|product/records/spec/alpha/index.md",
		"spec:product.beta|product/records/spec/beta/index.md",
	})
	if len(idxAB.ConflictGroups) != 3 {
		t.Fatalf("ConflictGroups len = %d, want 3", len(idxAB.ConflictGroups))
	}
	if idxAB.ConflictGroups[0].Ref != "DRMCP-REQ-MCP-980" || idxAB.ConflictGroups[1].Ref != "DRMCP-REQ-MCP-990" || idxAB.ConflictGroups[2].Ref != "spec:product.alpha" {
		t.Fatalf("ConflictGroups order = %#v", idxAB.ConflictGroups)
	}
	assertIndexStrings(t, idxAB.ConflictGroups[0].Sources, []string{
		"drmcp/records/requirements/mcp/DRMCP-REQ-MCP-980-duplicate-a.md",
		"drmcp/records/requirements/mcp/DRMCP-REQ-MCP-980-duplicate-b.md",
	})
	assertIndexStrings(t, idxAB.ConflictGroups[1].Sources, []string{
		"drmcp/records/requirements/mcp/DRMCP-REQ-MCP-990-duplicate-a.md",
		"drmcp/records/requirements/mcp/DRMCP-REQ-MCP-990-duplicate-b.md",
	})
	assertIndexStrings(t, idxAB.ConflictGroups[2].Sources, []string{
		"product/records/spec/alpha.md",
		"product/records/spec/alpha/index.md",
	})
	assertIndexStrings(t, parseIssueIndexKeys(idxAB.ParseIssues), []string{
		"drmcp/records/requirements/mcp/DRMCP-REQ-MCP-003-filename.md|DRMCP-REQ-MCP-004|filename_id_mismatch",
	})
	if len(idxAB.PathIssues) != 0 {
		t.Fatalf("PathIssues = %#v, want empty", idxAB.PathIssues)
	}
}

func TestDuplicateCurrentCanonicalIdentityHasNoWinner(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-duplicate-b.md", currentADRSource("DRMCP-ADR-MCP-001", "Duplicate B"))
	writeTestFile(t, root, "drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-duplicate-a.md", currentADRSource("DRMCP-ADR-MCP-001", "Duplicate A"))

	cfg, err := NewConfig(root, "drmcp/records")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	if len(idx.Records) != 0 {
		t.Fatalf("Records = %#v, want no duplicate winner", idx.Records)
	}
	assertIndexStrings(t, candidateIndexKeys(idx.Candidates), []string{
		"DRMCP-ADR-MCP-001|drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-duplicate-a.md",
		"DRMCP-ADR-MCP-001|drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-duplicate-b.md",
	})
	if len(idx.ConflictGroups) != 1 || idx.ConflictGroups[0].Ref != "DRMCP-ADR-MCP-001" {
		t.Fatalf("ConflictGroups = %#v", idx.ConflictGroups)
	}
	assertIndexStrings(t, idx.ConflictGroups[0].Sources, []string{
		"drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-duplicate-a.md",
		"drmcp/records/adr/mcp/DRMCP-ADR-MCP-001-duplicate-b.md",
	})
}

func TestIndexDeterministicIssueOrdering(t *testing.T) {
	idx := &Index{
		ParseIssues: []ParseIssue{
			{Path: "b.md", RecordID: "B", Category: DiagnosticCategory("c"), Message: "m"},
			{Path: "a.md", RecordID: "A", Category: DiagnosticCategory("d"), Message: "m"},
			{Path: "a.md", RecordID: "A", Category: DiagnosticCategory("c"), Message: "m", Details: map[string]string{"detail": "b"}},
			{Path: "a.md", RecordID: "A", Category: DiagnosticCategory("c"), Message: "m", Details: map[string]string{"detail": "a"}},
		},
		PathIssues: []PathIssue{
			{Path: "b.md", Operation: "read", Err: errors.New("m")},
			{Path: "a.md", Operation: "stat", Err: errors.New("m")},
			{Path: "a.md", Operation: "read", Err: errors.New("z")},
			{Path: "a.md", Operation: "read", Err: errors.New("a")},
		},
	}

	sortIndexDeterministically(idx)

	assertIndexStrings(t, parseIssueDetailIndexKeys(idx.ParseIssues), []string{
		"a.md|A|c|m|a",
		"a.md|A|c|m|b",
		"a.md|A|d|m|",
		"b.md|B|c|m|",
	})
	assertIndexStrings(t, pathIssueIndexKeys(idx.PathIssues), []string{
		"a.md|read|a",
		"a.md|read|z",
		"a.md|stat|m",
		"b.md|read|m",
	})
}

func TestBuildIndexDoesNotIndexCurrentSemanticAliases(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "product/records/spec/legacy-alias.md", "---\nsemantic_refs:\n  - spec:legacy.alias\nsections:\n  spec:legacy.section: Section\ndesign_record:\n  id: PRODUCT-ADR-SPEC-999\n  kind: decision\n  status: accepted\n---\n# Legacy alias\n")
	cfg, err := NewConfig(root, "product/records")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	if len(idx.SemanticRefs) != 0 || len(idx.SemanticRefSources) != 0 {
		t.Fatalf("semantic aliases leaked into current index: refs=%#v sources=%#v", idx.SemanticRefs, idx.SemanticRefSources)
	}
	if len(idx.Records) != 0 {
		t.Fatalf("YAML-front-matter current spec entered active records: %#v", idx.Records)
	}
}

func TestBuildIndexSkipsFileSymlinksForEveryCurrentKind(t *testing.T) {
	tests := []struct {
		name    string
		linkRel string
		content string
	}{
		{"adr", "test/records/adr/mcp/TEST-ADR-MCP-001-linked.md", currentADRSource("TEST-ADR-MCP-001", "Linked")},
		{"investigation", "test/records/investigations/mcp/TEST-INV-MCP-001-linked.md", currentInvestigationSource("TEST-INV-MCP-001", "Linked")},
		{"requirement", "test/records/requirements/mcp/TEST-REQ-MCP-001-linked.md", currentRequirementSource("TEST-REQ-MCP-001", "Linked")},
		{"work item", "test/records/work-items/mcp/TEST-WORK-MCP-001-linked.md", currentWorkItemSource("TEST-WORK-MCP-001", "Linked")},
		{"task", "test/records/tasks/mcp/TEST-TASK-MCP-001-01-linked.md", currentTaskSource("TEST-TASK-MCP-001-01", "Linked")},
		{"spec", "test/records/spec/linked.md", currentSpecSource("spec:test.linked", "Linked", "spec:test")},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			if err := os.MkdirAll(filepath.Join(root, "test", "records"), 0o755); err != nil {
				t.Fatalf("MkdirAll records root: %v", err)
			}
			targetRel := "targets/" + filepath.Base(tc.linkRel)
			writeTestFile(t, root, targetRel, tc.content)
			link := filepath.Join(root, filepath.FromSlash(tc.linkRel))
			if err := os.MkdirAll(filepath.Dir(link), 0o755); err != nil {
				t.Fatalf("MkdirAll link parent: %v", err)
			}
			if err := os.Symlink(filepath.Join(root, filepath.FromSlash(targetRel)), link); err != nil {
				t.Skipf("file symlink unavailable: %v", err)
			}
			assertSymlinkOnlyIndexEmpty(t, root)
		})
	}
}

func TestBuildIndexSkipsDirectorySymlinksForEveryCurrentKind(t *testing.T) {
	tests := []struct {
		name          string
		linkRel       string
		targetFileRel string
		content       string
	}{
		{"adr", "test/records/adr/mcp", "TEST-ADR-MCP-001-linked.md", currentADRSource("TEST-ADR-MCP-001", "Linked")},
		{"investigation", "test/records/investigations/mcp", "TEST-INV-MCP-001-linked.md", currentInvestigationSource("TEST-INV-MCP-001", "Linked")},
		{"requirement", "test/records/requirements/mcp", "TEST-REQ-MCP-001-linked.md", currentRequirementSource("TEST-REQ-MCP-001", "Linked")},
		{"work item", "test/records/work-items/mcp", "TEST-WORK-MCP-001-linked.md", currentWorkItemSource("TEST-WORK-MCP-001", "Linked")},
		{"task", "test/records/tasks/mcp", "TEST-TASK-MCP-001-01-linked.md", currentTaskSource("TEST-TASK-MCP-001-01", "Linked")},
		{"spec", "test/records/spec/linked", "index.md", currentSpecSource("spec:test.linked", "Linked", "spec:test")},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := t.TempDir()
			if err := os.MkdirAll(filepath.Join(root, "test", "records"), 0o755); err != nil {
				t.Fatalf("MkdirAll records root: %v", err)
			}
			targetDir := filepath.Join(root, "targets", filepath.FromSlash(tc.name))
			writeTestFile(t, targetDir, tc.targetFileRel, tc.content)
			link := filepath.Join(root, filepath.FromSlash(tc.linkRel))
			if err := os.MkdirAll(filepath.Dir(link), 0o755); err != nil {
				t.Fatalf("MkdirAll link parent: %v", err)
			}
			if err := os.Symlink(targetDir, link); err != nil {
				t.Skipf("directory symlink unavailable: %v", err)
			}
			assertSymlinkOnlyIndexEmpty(t, root)
		})
	}
}

func TestBuildIndexRejectsSymlinkedMandatoryRoot(t *testing.T) {
	t.Run("records root", func(t *testing.T) {
		root := t.TempDir()
		target := filepath.Join(root, "actual-records")
		if err := os.MkdirAll(target, 0o755); err != nil {
			t.Fatalf("MkdirAll target: %v", err)
		}
		if err := os.MkdirAll(filepath.Join(root, "test"), 0o755); err != nil {
			t.Fatalf("MkdirAll namespace: %v", err)
		}
		link := filepath.Join(root, "test", "records")
		if err := os.Symlink(target, link); err != nil {
			t.Skipf("directory symlink unavailable: %v", err)
		}
		assertMandatoryRootRejected(t, root)
	})

	t.Run("namespace ancestor", func(t *testing.T) {
		root := t.TempDir()
		target := filepath.Join(root, "actual-test")
		if err := os.MkdirAll(filepath.Join(target, "records"), 0o755); err != nil {
			t.Fatalf("MkdirAll target: %v", err)
		}
		if err := os.Symlink(target, filepath.Join(root, "test")); err != nil {
			t.Skipf("directory symlink unavailable: %v", err)
		}
		assertMandatoryRootRejected(t, root)
	})
}

func assertMandatoryRootRejected(t *testing.T, root string) {
	t.Helper()
	cfg, err := NewConfig(root, "test/records")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	if _, err := BuildIndex(context.Background(), cfg); err == nil {
		t.Fatal("BuildIndex symlinked mandatory root: want error, got nil")
	}
}

func assertSymlinkOnlyIndexEmpty(t *testing.T, root string) {
	t.Helper()
	cfg, err := NewConfig(root, "test/records")
	if err != nil {
		t.Fatalf("NewConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	if len(idx.Records) != 0 || len(idx.Candidates) != 0 || len(idx.ParseIssues) != 0 || len(idx.PathIssues) != 0 {
		t.Fatalf("symlink contributed to index: records=%#v candidates=%#v parse=%#v path=%#v", idx.Records, idx.Candidates, idx.ParseIssues, idx.PathIssues)
	}
}

func currentADRSource(id, title string) string {
	return "# " + id + ": " + title + "\n- **status**: accepted\n- **date**: 2026-06-28\n- **depends_on**: []\n- **supersedes**: []\n- **migrated_to_spec**: null\n"
}

func currentInvestigationSource(id, title string) string {
	return "# " + id + ": " + title + "\n- **status**: concluded\n- **date**: 2026-06-28\n- **trigger**: test\n- **scope**: test\n- **non_scope**: none\n- **source_refs**: []\n- **follow_up_candidates**: []\n"
}

func currentRequirementSource(id, title string) string {
	return "# " + id + ": " + title + "\n- **id**: " + id + "\n- **status**: captured\n- **date**: 2026-06-28\n- **source_refs**: []\n- **work_items**: []\n"
}

func currentWorkItemSource(id, title string) string {
	return "# " + id + ": " + title + "\n- **id**: " + id + "\n- **status**: in_progress\n- **date**: 2026-06-28\n- **source_requirement**: DRMCP-REQ-MCP-001\n- **impact_refs**: []\n- **tasks**: []\n"
}

func currentTaskSource(id, title string) string {
	return "# " + id + ": " + title + "\n- **id**: " + id + "\n- **status**: in_progress\n- **date**: 2026-06-28\n- **work_item**: DRMCP-WORK-MCP-001\n- **source_requirement**: DRMCP-REQ-MCP-001\n- **estimate**: 0.5d\n- **depends_on**: []\n- **outputs**: []\n"
}

func currentSpecSource(id, title, parent string) string {
	return "# " + title + "\n- **id**: `" + id + "`\n- **status**: accepted\n- **date**: 2026-06-28\n- **parent**: `" + parent + "`\n"
}

func recordIndexKeys(records []Record) []string {
	keys := make([]string, 0, len(records))
	for _, record := range records {
		keys = append(keys, record.ID+"|"+record.Path)
	}
	return keys
}

func candidateIndexKeys(candidates []RecordCandidate) []string {
	keys := make([]string, 0, len(candidates))
	for _, candidate := range candidates {
		keys = append(keys, candidate.ID+"|"+candidate.Path)
	}
	return keys
}

func parseIssueIndexKeys(issues []ParseIssue) []string {
	keys := make([]string, 0, len(issues))
	for _, issue := range issues {
		keys = append(keys, issue.Path+"|"+issue.RecordID+"|"+string(issue.Category))
	}
	return keys
}

func parseIssueDetailIndexKeys(issues []ParseIssue) []string {
	keys := make([]string, 0, len(issues))
	for _, issue := range issues {
		keys = append(keys, issue.Path+"|"+issue.RecordID+"|"+string(issue.Category)+"|"+issue.Message+"|"+issue.Details["detail"])
	}
	return keys
}

func pathIssueIndexKeys(issues []PathIssue) []string {
	keys := make([]string, 0, len(issues))
	for _, issue := range issues {
		keys = append(keys, issue.Path+"|"+issue.Operation+"|"+errorText(issue.Err))
	}
	return keys
}

func assertIndexStrings(t *testing.T, got, want []string) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}
