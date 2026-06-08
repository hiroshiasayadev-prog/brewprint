package designrecords

import (
	"context"
	"testing"
)

func TestResolveReferenceSemanticAndRecordTargets(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/adr/088-test.md", "# 088: Test ADR\n- **status**: accepted\n")
	writeTestFile(t, root, "docs/spec/trace.md", "---\nstatus: draft\nsemantic_refs:\n  - spec:trace\n  - spec:trace.resolve-and-validation\nsections:\n  spec:trace.resolve: Resolve\n  spec:trace.validation: Validation\n\ndesign_record:\n  id: SPEC-trace\n  kind: spec\n  status: draft\n---\n# Trace spec\n## Resolve\n## Validation\n")
	writeTestFile(t, root, "docs/spec/project-artifact-model/index.md", "---\nstatus: draft\nsemantic_refs:\n  - spec:project-artifact-model\nsections:\n  spec:project-artifact-model.responsibilities: Artifact responsibility matrix\n\ndesign_record:\n  id: SPEC-project-artifact-model\n  kind: spec\n  status: draft\n---\n# Project artifact model\n## Artifact responsibility matrix\n")
	writeTestFile(t, root, "docs/investigations/docs/INV-DOCS-001-test.md", "# INV-DOCS-001: Test investigation\n- **status**: concluded\n- **date**: 2026-05-19\n- **trigger**: ADR-088\n- **scope**: test\n- **non_scope**: none\n- **source_refs**:\n  - ADR-088\n- **follow_up_candidates**:\n  - SPEC-trace\n")
	idx := buildTestIndex(t, root)

	for _, tt := range []struct {
		ref  string
		path string
	}{
		{ref: "spec:trace", path: "docs/spec/trace.md"},
		{ref: "spec:project-artifact-model", path: "docs/spec/project-artifact-model/index.md"},
		{ref: "spec:trace.resolve-and-validation", path: "docs/spec/trace.md"},
	} {
		document, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: tt.ref})
		if err != nil {
			t.Fatalf("ResolveReference document %s: %v", tt.ref, err)
		}
		if document.Status != resolveStatusResolved || document.RefKind != refKindSemanticRef || document.Target == nil || document.Target.TargetType != "document" || document.Target.Path != tt.path {
			t.Fatalf("%s document resolve = %#v", tt.ref, document)
		}
	}

	section, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: "spec:trace.resolve"})
	if err != nil {
		t.Fatalf("ResolveReference section: %v", err)
	}
	if section.Status != resolveStatusResolved || section.Target == nil || section.Target.TargetType != "section" || section.Target.Section != "Resolve" {
		t.Fatalf("section resolve = %#v", section)
	}

	for _, id := range []string{"ADR-088", "SPEC-trace", "INV-DOCS-001"} {
		resp, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: id})
		if err != nil {
			t.Fatalf("ResolveReference %s: %v", id, err)
		}
		if resp.Status != resolveStatusResolved || resp.RefKind != refKindRecordID || resp.Target == nil || resp.Target.TargetType != "record" || resp.Target.RecordID != id {
			t.Fatalf("%s resolve = %#v", id, resp)
		}
	}
}

func TestResolveReferenceWorkflowRecordTargets(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/requirements/mcp/REQ-MCP-003-test.md", "# REQ-MCP-003: Test requirement\n- **id**: REQ-MCP-003\n- **status**: accepted\n- **date**: 2026-05-25\n- **source_refs**:\n- **work_items**:\n  - WORK-MCP-003\n")
	writeTestFile(t, root, "docs/work-items/mcp/WORK-MCP-003-test.md", "# WORK-MCP-003: Test work item\n- **id**: WORK-MCP-003\n- **status**: implementation_pending\n- **date**: 2026-05-26\n- **source_requirement**: REQ-MCP-003\n- **impact_refs**:\n- **tasks**:\n  - TASK-MCP-003-01\n")
	writeTestFile(t, root, "docs/tasks/mcp/TASK-MCP-003-01-test.md", "# TASK-MCP-003-01: Test task\n- **id**: TASK-MCP-003-01\n- **status**: todo\n- **date**: 2026-05-26\n- **work_item**: WORK-MCP-003\n- **source_requirement**: REQ-MCP-003\n- **estimate**: 0.5d\n- **depends_on**:\n- **outputs**:\n  - test\n")
	idx := buildTestIndex(t, root)

	for _, tt := range []struct {
		ref  string
		kind RecordKind
	}{
		{ref: "REQ-MCP-003", kind: RecordKindRequirement},
		{ref: "WORK-MCP-003", kind: RecordKindWorkItem},
		{ref: "TASK-MCP-003-01", kind: RecordKindTask},
	} {
		resp, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: tt.ref})
		if err != nil {
			t.Fatalf("ResolveReference %s: %v", tt.ref, err)
		}
		if resp.Status != resolveStatusResolved || resp.RefKind != refKindRecordID || resp.Target == nil || resp.Target.TargetType != "record" || resp.Target.RecordID != tt.ref || resp.Target.RecordKind != tt.kind {
			t.Fatalf("%s resolve = %#v", tt.ref, resp)
		}
	}
}

func TestResolveReferenceUsesSemanticRefsFromNonRecordSpec(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/spec/non-record.md", "---\nstatus: draft\nsemantic_refs:\n  - spec:non-record.doc\nsections:\n  spec:non-record.section: Target Section\n---\n# Non-record spec\n## Target Section\n")
	idx := buildTestIndex(t, root)

	listResp, err := ListRecords(context.Background(), idx, ListRecordsRequest{})
	if err != nil {
		t.Fatalf("ListRecords: %v", err)
	}
	if len(listResp.Records) != 0 {
		t.Fatalf("non-record semantic source leaked into list_records: %#v", listResp.Records)
	}
	if _, err := GetRecord(context.Background(), idx, GetRecordRequest{ID: "SPEC-non-record"}); err == nil {
		t.Fatalf("GetRecord succeeded for non-record semantic source")
	}

	document, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: "spec:non-record.doc"})
	if err != nil {
		t.Fatalf("ResolveReference document: %v", err)
	}
	if document.Status != resolveStatusResolved || document.Target == nil || document.Target.TargetType != "document" || document.Target.Path != "docs/spec/non-record.md" {
		t.Fatalf("document resolve = %#v", document)
	}

	section, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: "spec:non-record.section"})
	if err != nil {
		t.Fatalf("ResolveReference section: %v", err)
	}
	if section.Status != resolveStatusResolved || section.Target == nil || section.Target.TargetType != "section" || section.Target.Section != "Target Section" {
		t.Fatalf("section resolve = %#v", section)
	}
}

func TestValidateRecordsSemanticRefDiagnosticsFromNonRecordSpec(t *testing.T) {
	root := t.TempDir()
	writeTestFile(t, root, "docs/spec/non-record.md", "---\nstatus: draft\nsemantic_refs:\n  - Bad:Ref\nsections:\n  spec:non-record.missing: Missing\n  spec:non-record.ambiguous: Duplicate\n---\n# Non-record spec\n## Duplicate\n## Duplicate\n")
	idx := buildTestIndex(t, root)

	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	for _, category := range []DiagnosticCategory{
		DiagnosticInvalidSemanticRefDeclaration,
		DiagnosticMissingSectionTarget,
		DiagnosticAmbiguousSectionTarget,
	} {
		if !hasDiagnostic(resp.Diagnostics, category) {
			t.Fatalf("missing %s in %#v", category, resp.Diagnostics)
		}
	}
	for _, diagnostic := range resp.Diagnostics {
		if diagnostic.Path != "docs/spec/non-record.md" {
			t.Fatalf("diagnostic path = %q, want non-record spec path: %#v", diagnostic.Path, diagnostic)
		}
		if diagnostic.RecordID != "" {
			t.Fatalf("non-record spec diagnostic should not have record_id: %#v", diagnostic)
		}
	}
}

func TestResolveReferenceRepositoryBootstrapIDs(t *testing.T) {
	root := findRepoRoot(t)
	idx := buildTestIndex(t, root)
	for _, id := range []string{"ADR-088", "INV-DOCS-001", "INV-DOCS-002", "INV-DOCS-003", "REQ-MCP-003", "WORK-MCP-003", "TASK-MCP-003-01"} {
		resp, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: id})
		if err != nil {
			t.Fatalf("ResolveReference %s: %v", id, err)
		}
		if resp.Status != resolveStatusResolved || resp.Target == nil || resp.Target.RecordID != id {
			t.Fatalf("%s resolve = %#v", id, resp)
		}
	}
}

func TestResolveReferenceUnresolvedAmbiguousAndUnsupported(t *testing.T) {
	idx := &Index{
		Records: []Record{
			{ID: "ADR-001", NormalizedID: "ADR-001", Kind: RecordKindDecision, Title: "One", Status: RecordStatusAccepted, Path: "docs/adr/001-one.md"},
			{ID: "ADR-001", NormalizedID: "ADR-001", Kind: RecordKindDecision, Title: "Duplicate", Status: RecordStatusAccepted, Path: "docs/adr/001-duplicate.md"},
		},
		SemanticRefs: []SemanticRefDecl{
			{Ref: "spec:trace.duplicate", Path: "docs/spec/a.md", TargetType: SemanticTargetDocument},
			{Ref: "spec:trace.duplicate", Path: "docs/spec/b.md", TargetType: SemanticTargetDocument},
		},
	}

	unresolved, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: "ADR-999"})
	if err != nil {
		t.Fatalf("ResolveReference unresolved: %v", err)
	}
	if unresolved.Status != resolveStatusUnresolved || unresolved.Target != nil || !hasDiagnostic(unresolved.Diagnostics, DiagnosticUnresolvedReference) {
		t.Fatalf("unresolved = %#v", unresolved)
	}

	for _, ref := range []string{"REQ-MCP-999", "WORK-MCP-999", "TASK-MCP-999-99"} {
		resp, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: ref})
		if err != nil {
			t.Fatalf("ResolveReference unresolved %s: %v", ref, err)
		}
		if resp.Status != resolveStatusUnresolved || resp.RefKind != refKindRecordID || resp.Target != nil || !hasDiagnostic(resp.Diagnostics, DiagnosticUnresolvedReference) {
			t.Fatalf("%s unresolved response = %#v", ref, resp)
		}
	}

	ambiguousRecord, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: "ADR-001"})
	if err != nil {
		t.Fatalf("ResolveReference ambiguous record: %v", err)
	}
	if ambiguousRecord.Status != resolveStatusUnresolved || ambiguousRecord.Target != nil || !hasDiagnostic(ambiguousRecord.Diagnostics, DiagnosticAmbiguousReference) {
		t.Fatalf("ambiguous record = %#v", ambiguousRecord)
	}

	ambiguousSemantic, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: "spec:trace.duplicate"})
	if err != nil {
		t.Fatalf("ResolveReference ambiguous semantic: %v", err)
	}
	if ambiguousSemantic.Status != resolveStatusUnresolved || ambiguousSemantic.Target != nil || !hasDiagnostic(ambiguousSemantic.Diagnostics, DiagnosticAmbiguousReference) {
		t.Fatalf("ambiguous semantic = %#v", ambiguousSemantic)
	}

	for _, ref := range []string{"internal-design:resolver.semantic-ref-index", "coverage:trace", "COV-TRACE-001", "docs/spec/trace.md", "ADR-abc", "TASK-MCP-003-1", "REQ-mcp-003", "WORK-MCP-003-extra-01"} {
		resp, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: ref})
		if err != nil {
			t.Fatalf("ResolveReference unsupported %s: %v", ref, err)
		}
		if resp.Status != resolveStatusUnsupported || resp.Target != nil || !hasDiagnostic(resp.Diagnostics, DiagnosticUnsupportedReference) {
			t.Fatalf("%s unsupported response = %#v", ref, resp)
		}
	}
}
