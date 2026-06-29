package designrecords

import (
	"context"
	"encoding/json"
	"strings"
	"testing"
)

func TestResolveReferenceCurrentUniqueTargets(t *testing.T) {
	idx := buildCurrentResolverIndex(t)

	tests := []struct {
		ref        string
		refKind    string
		targetType string
		kind       RecordKind
		title      string
		status     RecordStatus
	}{
		{ref: "PRODUCT-ADR-SPEC-901", refKind: refKindCurrentRecordID, targetType: "record", kind: RecordKindDecision, title: "Product ADR", status: RecordStatusAccepted},
		{ref: "DRMCP-INV-MCP-001", refKind: refKindCurrentRecordID, targetType: "record", kind: RecordKindInvestigation, title: "DRMCP investigation", status: RecordStatusConcluded},
		{ref: "DRMCP-REQ-MCP-001", refKind: refKindCurrentRecordID, targetType: "record", kind: RecordKindRequirement, title: "DRMCP requirement", status: RecordStatusCaptured},
		{ref: "DRMCP-WORK-MCP-001", refKind: refKindCurrentRecordID, targetType: "record", kind: RecordKindWorkItem, title: "DRMCP work item", status: RecordStatusInProgress},
		{ref: "DRMCP-TASK-MCP-001-01", refKind: refKindCurrentRecordID, targetType: "record", kind: RecordKindTask, title: "DRMCP task", status: RecordStatusInProgress},
		{ref: "spec:product.beta", refKind: refKindCurrentSpecRef, targetType: "spec", kind: RecordKindSpec, title: "Product beta", status: RecordStatusAccepted},
	}

	for _, tt := range tests {
		t.Run(tt.ref, func(t *testing.T) {
			resp, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: tt.ref})
			if err != nil {
				t.Fatalf("ResolveReference: %v", err)
			}
			if resp.Ref != tt.ref || resp.RefKind != tt.refKind || resp.Status != resolveStatusResolved {
				t.Fatalf("classification = %#v", resp)
			}
			if resp.Target == nil {
				t.Fatalf("target is nil")
			}
			if resp.Target.TargetType != tt.targetType || resp.Target.RecordID != tt.ref || resp.Target.RecordKind != tt.kind || resp.Target.Title != tt.title || resp.Target.Status != tt.status {
				t.Fatalf("target = %#v", resp.Target)
			}
			if resp.Target.Path != "" || resp.Target.Section != "" {
				t.Fatalf("resolved target leaked path or section: %#v", resp.Target)
			}
			raw, err := json.Marshal(resp)
			if err != nil {
				t.Fatalf("Marshal: %v", err)
			}
			if strings.Contains(string(raw), ".md") || strings.Contains(string(raw), "/records/") || strings.Contains(string(raw), `\`) {
				t.Fatalf("resolved JSON leaked physical path: %s", raw)
			}
		})
	}
}

func TestResolveReferenceCurrentMissingAndConflict(t *testing.T) {
	idx := buildCurrentResolverIndex(t)

	tests := []struct {
		name       string
		ref        string
		refKind    string
		category   DiagnosticCategory
		status     string
		targetMust bool
	}{
		{name: "missing current record", ref: "DRMCP-REQ-MCP-999", refKind: refKindCurrentRecordID, category: DiagnosticUnresolvedReference, status: resolveStatusUnresolved},
		{name: "missing current spec", ref: "spec:product.missing", refKind: refKindCurrentSpecRef, category: DiagnosticUnresolvedReference, status: resolveStatusUnresolved},
		{name: "conflicting current record", ref: "DRMCP-REQ-MCP-980", refKind: refKindCurrentRecordID, category: DiagnosticAmbiguousReference, status: resolveStatusUnresolved},
		{name: "conflicting current spec", ref: "spec:product.alpha", refKind: refKindCurrentSpecRef, category: DiagnosticAmbiguousReference, status: resolveStatusUnresolved},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: tt.ref})
			if err != nil {
				t.Fatalf("ResolveReference: %v", err)
			}
			if resp.Ref != tt.ref || resp.RefKind != tt.refKind || resp.Status != tt.status || resp.Target != nil || !hasDiagnostic(resp.Diagnostics, tt.category) {
				t.Fatalf("response = %#v", resp)
			}
		})
	}
}

func TestResolveReferenceNestedCurrentSpecs(t *testing.T) {
	const childRef = "spec:product.beta.resolve"

	t.Run("exact child resolves when parent exists", func(t *testing.T) {
		idx := buildNestedCurrentSpecResolverIndex(t, 1)

		resp, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: childRef})
		if err != nil {
			t.Fatalf("ResolveReference: %v", err)
		}
		if resp.Ref != childRef || resp.RefKind != refKindCurrentSpecRef || resp.Status != resolveStatusResolved {
			t.Fatalf("response = %#v", resp)
		}
		if resp.Target == nil || resp.Target.RecordID != childRef || resp.Target.RecordKind != RecordKindSpec {
			t.Fatalf("target = %#v", resp.Target)
		}
		if resp.Target.Path != "" || resp.Target.Section != "" {
			t.Fatalf("resolved target leaked path or section: %#v", resp.Target)
		}
	})

	t.Run("exact child conflict does not select winner", func(t *testing.T) {
		idx := buildNestedCurrentSpecResolverIndex(t, 2)

		resp, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: childRef})
		if err != nil {
			t.Fatalf("ResolveReference: %v", err)
		}
		if resp.Ref != childRef || resp.RefKind != refKindCurrentSpecRef || resp.Status != resolveStatusUnresolved || resp.Target != nil || !hasDiagnostic(resp.Diagnostics, DiagnosticAmbiguousReference) {
			t.Fatalf("response = %#v", resp)
		}
	})

	t.Run("missing child below existing parent is unsupported", func(t *testing.T) {
		idx := buildNestedCurrentSpecResolverIndex(t, 0)

		resp, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: childRef})
		if err != nil {
			t.Fatalf("ResolveReference: %v", err)
		}
		if resp.Ref != childRef || resp.RefKind != refKindUnsupported || resp.Status != resolveStatusUnsupported || resp.Target != nil || !hasDiagnostic(resp.Diagnostics, DiagnosticUnsupportedReference) {
			t.Fatalf("response = %#v", resp)
		}
	})
}

func TestResolveReferenceRejectsUnsupportedForms(t *testing.T) {
	idx := buildCurrentResolverIndex(t)

	for _, ref := range []string{
		"ADR-SPEC-901",
		"REQ-MCP-001",
		"WORK-MCP-001",
		"TASK-MCP-001-01",
		"PRODUCT-REQ-MCP",
		"PRODUCT-REQ-MCP-00",
		"product-req-mcp-001",
		"PRODUCT-REQ-mcp-001",
		"PRODUCT-SPEC-beta",
		"V01-ADR-088",
		"product/records/spec/beta/index.md",
		"records/spec/beta.md",
		"spec:product..beta",
		"spec:product.beta/",
		"spec:Product.beta",
		"internal-design:resolver.semantic-ref-index",
		"coverage:trace",
		"COV-TRACE-001",
	} {
		t.Run(ref, func(t *testing.T) {
			resp, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: ref})
			if err != nil {
				t.Fatalf("ResolveReference: %v", err)
			}
			if resp.Ref != ref || resp.RefKind != refKindUnsupported || resp.Status != resolveStatusUnsupported || resp.Target != nil || !hasDiagnostic(resp.Diagnostics, DiagnosticUnsupportedReference) {
				t.Fatalf("unsupported response = %#v", resp)
			}
		})
	}
}

func TestResolveReferenceIgnoresSemanticAliasesAndLegacyFallback(t *testing.T) {
	idx := &Index{
		NamespacePrefix: "V01-",
		Records: []Record{
			{ID: "V01-ADR-088", Kind: RecordKindDecision, Title: "Legacy ADR", Status: RecordStatusAccepted, Path: "v01/records/ADR/ADR-088.md"},
			{ID: "spec:product.beta", Kind: RecordKindSpec, Title: "Product beta", Status: RecordStatusAccepted, Path: "product/records/spec/beta/index.md"},
		},
		SemanticRefs: []SemanticRefDecl{
			{Ref: "spec:legacy.alias", Path: "product/records/spec/beta/index.md", TargetType: SemanticTargetDocument},
			{Ref: "spec:legacy.alias.section", Path: "product/records/spec/beta/index.md", TargetType: SemanticTargetSection, Section: "Section"},
		},
	}

	legacy, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: "V01-ADR-088"})
	if err != nil {
		t.Fatalf("ResolveReference legacy: %v", err)
	}
	if legacy.Status != resolveStatusUnsupported || legacy.RefKind != refKindUnsupported || legacy.Target != nil {
		t.Fatalf("legacy fallback response = %#v", legacy)
	}

	for _, ref := range []string{"spec:legacy.alias", "spec:legacy.alias.section"} {
		resp, err := ResolveReference(context.Background(), idx, ResolveReferenceRequest{Ref: ref})
		if err != nil {
			t.Fatalf("ResolveReference %s: %v", ref, err)
		}
		if resp.Status != resolveStatusUnresolved || resp.RefKind != refKindCurrentSpecRef || resp.Target != nil || !hasDiagnostic(resp.Diagnostics, DiagnosticUnresolvedReference) {
			t.Fatalf("semantic alias response = %#v", resp)
		}
	}
}

func buildNestedCurrentSpecResolverIndex(t *testing.T, childCopies int) *Index {
	t.Helper()

	root := t.TempDir()
	writeTestFile(t, root, "product/records/spec/beta/index.md", currentSpecSource("spec:product.beta", "Product beta", "spec:product"))
	if childCopies >= 1 {
		writeTestFile(t, root, "product/records/spec/beta/resolve/index.md", currentSpecSource("spec:product.beta.resolve", "Product beta resolve index", "spec:product.beta"))
	}
	if childCopies >= 2 {
		writeTestFile(t, root, "product/records/spec/beta/resolve.md", currentSpecSource("spec:product.beta.resolve", "Product beta resolve leaf", "spec:product.beta"))
	}

	cfg, err := NormalizeConfig(root, []CurrentRoot{
		{AppNamespace: "product", RecordsRoot: "product/records"},
	})
	if err != nil {
		t.Fatalf("NormalizeConfig: %v", err)
	}
	idx, err := BuildIndex(context.Background(), cfg)
	if err != nil {
		t.Fatalf("BuildIndex: %v", err)
	}
	return idx
}

func buildCurrentResolverIndex(t *testing.T) *Index {
	t.Helper()

	root := t.TempDir()
	writeTestFile(t, root, "product/records/adr/spec/PRODUCT-ADR-SPEC-901-product.md", currentADRSource("PRODUCT-ADR-SPEC-901", "Product ADR"))
	writeTestFile(t, root, "product/records/spec/alpha.md", currentSpecSource("spec:product.alpha", "Product alpha leaf", "spec:product"))
	writeTestFile(t, root, "product/records/spec/alpha/index.md", currentSpecSource("spec:product.alpha", "Product alpha index", "spec:product"))
	writeTestFile(t, root, "product/records/spec/beta/index.md", currentSpecSource("spec:product.beta", "Product beta", "spec:product"))

	writeTestFile(t, root, "drmcp/records/investigations/mcp/DRMCP-INV-MCP-001-current.md", currentInvestigationSource("DRMCP-INV-MCP-001", "DRMCP investigation"))
	writeTestFile(t, root, "drmcp/records/requirements/mcp/DRMCP-REQ-MCP-001-current.md", currentRequirementSource("DRMCP-REQ-MCP-001", "DRMCP requirement"))
	writeTestFile(t, root, "drmcp/records/requirements/mcp/DRMCP-REQ-MCP-980-duplicate-a.md", currentRequirementSource("DRMCP-REQ-MCP-980", "Duplicate A"))
	writeTestFile(t, root, "drmcp/records/requirements/mcp/DRMCP-REQ-MCP-980-duplicate-b.md", currentRequirementSource("DRMCP-REQ-MCP-980", "Duplicate B"))
	writeTestFile(t, root, "drmcp/records/work-items/mcp/DRMCP-WORK-MCP-001-current.md", currentWorkItemSource("DRMCP-WORK-MCP-001", "DRMCP work item"))
	writeTestFile(t, root, "drmcp/records/tasks/mcp/DRMCP-TASK-MCP-001-01-current.md", currentTaskSource("DRMCP-TASK-MCP-001-01", "DRMCP task"))

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
	return idx
}
