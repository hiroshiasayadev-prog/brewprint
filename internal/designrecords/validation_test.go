package designrecords

import (
	"context"
	"errors"
	"testing"
)

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

	idx := buildTestIndex(t, root)
	resp, err := ValidateRecords(context.Background(), idx, ValidateRecordsRequest{})
	if err != nil {
		t.Fatalf("ValidateRecords: %v", err)
	}
	if !resp.OK || len(resp.Diagnostics) != 0 {
		t.Fatalf("response = %#v, want ok with no diagnostics", resp)
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

func TestValidateRecordsRequestErrors(t *testing.T) {
	idx := &Index{}
	tests := []struct {
		name string
		req  ValidateRecordsRequest
		code ErrorCode
	}{
		{
			name: "invalid kind",
			req:  ValidateRecordsRequest{Kind: RecordKind("task")},
			code: ErrorCodeInvalidRequest,
		},
		{
			name: "kind spec with id range",
			req:  ValidateRecordsRequest{Kind: RecordKindSpec, IDRange: &IDRange{From: "ADR-001"}},
			code: ErrorCodeIDRangeRequiresDecisionKind,
		},
		{
			name: "SPEC range endpoint",
			req:  ValidateRecordsRequest{IDRange: &IDRange{From: "SPEC-design-records-mcp-schema"}},
			code: ErrorCodeIDRangeRequiresDecisionKind,
		},
		{
			name: "malformed range endpoint",
			req:  ValidateRecordsRequest{IDRange: &IDRange{From: "ADR-x"}},
			code: ErrorCodeIDRangeRequiresDecisionKind,
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
