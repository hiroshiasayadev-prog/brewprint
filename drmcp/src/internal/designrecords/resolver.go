package designrecords

import (
	"context"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	activeSpecRefPattern      = regexp.MustCompile(`^spec:[a-z0-9-]+(?:\.[a-z0-9-]+)*$`)
	currentSpecRefPattern     = regexp.MustCompile(`^spec:[a-z0-9-]+(?:\.[a-z0-9_]+)*$`)
	currentNamespaceIDPattern = regexp.MustCompile(`^[A-Z][A-Z0-9]*(?:-[A-Z0-9]+)*-(ADR-[A-Z][A-Z0-9]*-\d{3}|INV-[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?-\d{3}|REQ-[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?-\d{3}|WORK-[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?-\d{3}|TASK-[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?-\d{3}-\d{2})$`)
	unsupportedIDPattern      = regexp.MustCompile(`^COV-`)
)

const (
	refKindCurrentSpecRef  = "current_spec_ref"
	refKindCurrentRecordID = "current_record_id"
	refKindUnsupported     = "unsupported"

	resolveStatusResolved    = "resolved"
	resolveStatusUnresolved  = "unresolved"
	resolveStatusUnsupported = "unsupported"
)

func ResolveReference(ctx context.Context, idx *Index, req ResolveReferenceRequest) (ResolveReferenceResponse, error) {
	if err := ctx.Err(); err != nil {
		return ResolveReferenceResponse{}, err
	}
	if idx == nil {
		return ResolveReferenceResponse{}, newToolError(ErrorCodeInvalidRequest, "index is nil")
	}
	if req.Ref == "" {
		return ResolveReferenceResponse{}, newToolError(ErrorCodeInvalidRequest, "ref is required")
	}
	return resolveReference(idx, req.Ref), nil
}

func resolveReference(idx *Index, ref string) ResolveReferenceResponse {
	kind := classifyReference(ref)
	switch kind {
	case refKindCurrentSpecRef:
		if hasCurrentConflict(idx, ref) {
			return ambiguousResponse(ref, kind)
		}
		if response, ok := resolveExactCurrentReference(idx, ref, kind); ok {
			return response
		}
		if isCurrentSpecAliasOrPartial(idx, ref) {
			return unsupportedResponse(ref)
		}
		return unresolvedResponse(ref, kind)
	case refKindCurrentRecordID:
		return resolveCurrentReference(idx, ref, kind)
	default:
		return unsupportedResponse(ref)
	}
}

func classifyReference(ref string) string {
	if currentSpecRefPattern.MatchString(ref) {
		return refKindCurrentSpecRef
	}
	if currentNamespaceIDPattern.MatchString(ref) {
		return refKindCurrentRecordID
	}
	return refKindUnsupported
}

func isSupportedInvestigationRecordIDReference(ref string) bool {
	return currentNamespaceIDPattern.MatchString(ref) && strings.Contains(ref, "-INV-")
}

func resolveCurrentReference(idx *Index, ref, refKind string) ResolveReferenceResponse {
	if hasCurrentConflict(idx, ref) {
		return ambiguousResponse(ref, refKind)
	}
	if response, ok := resolveExactCurrentReference(idx, ref, refKind); ok {
		return response
	}
	return unresolvedResponse(ref, refKind)
}

func resolveExactCurrentReference(idx *Index, ref, refKind string) (ResolveReferenceResponse, bool) {
	for _, record := range idx.Records {
		if record.ID != ref {
			continue
		}
		return ResolveReferenceResponse{
			Ref:     ref,
			RefKind: refKind,
			Status:  resolveStatusResolved,
			Target: &ResolvedTarget{
				TargetType: targetTypeForRecord(record),
				RecordID:   record.ID,
				RecordKind: record.Kind,
				Title:      record.Title,
				Status:     record.Status,
			},
			Diagnostics: []Diagnostic{},
		}, true
	}
	return ResolveReferenceResponse{}, false
}

func hasCurrentConflict(idx *Index, ref string) bool {
	for _, conflict := range idx.ConflictGroups {
		if conflict.Ref == ref {
			return true
		}
	}
	return false
}

func targetTypeForRecord(record Record) string {
	if record.Kind == RecordKindSpec {
		return "spec"
	}
	return "record"
}

func isCurrentSpecAliasOrPartial(idx *Index, ref string) bool {
	for _, record := range idx.Records {
		if record.Kind == RecordKindSpec && strings.HasPrefix(ref, record.ID+".") {
			return true
		}
	}
	for _, conflict := range idx.ConflictGroups {
		if strings.HasPrefix(ref, conflict.Ref+".") {
			return true
		}
	}
	return false
}

func unsupportedResponse(ref string) ResolveReferenceResponse {
	return ResolveReferenceResponse{
		Ref:     ref,
		RefKind: refKindUnsupported,
		Status:  resolveStatusUnsupported,
		Target:  nil,
		Diagnostics: []Diagnostic{{
			Category: DiagnosticUnsupportedReference,
			Severity: DiagnosticSeverityInfo,
			Message:  "reference form is outside the MVP resolver contract",
		}},
	}
}

func unresolvedResponse(ref, refKind string) ResolveReferenceResponse {
	return ResolveReferenceResponse{
		Ref:     ref,
		RefKind: refKind,
		Status:  resolveStatusUnresolved,
		Target:  nil,
		Diagnostics: []Diagnostic{{
			Category: DiagnosticUnresolvedReference,
			Severity: DiagnosticSeverityError,
			Message:  fmt.Sprintf("reference %s did not resolve to a target", ref),
		}},
	}
}

func ambiguousResponse(ref, refKind string) ResolveReferenceResponse {
	return ResolveReferenceResponse{
		Ref:     ref,
		RefKind: refKind,
		Status:  resolveStatusUnresolved,
		Target:  nil,
		Diagnostics: []Diagnostic{{
			Category: DiagnosticAmbiguousReference,
			Severity: DiagnosticSeverityError,
			Message:  fmt.Sprintf("reference %s resolves to multiple targets", ref),
		}},
	}
}

func semanticTargetsByRef(idx *Index) map[string][]SemanticRefDecl {
	out := map[string][]SemanticRefDecl{}
	for _, decl := range idx.SemanticRefs {
		if decl.Ref == "" || !activeSpecRefPattern.MatchString(decl.Ref) {
			continue
		}
		out[decl.Ref] = append(out[decl.Ref], decl)
	}
	for ref := range out {
		sort.Slice(out[ref], func(i, j int) bool {
			if out[ref][i].Path == out[ref][j].Path {
				return out[ref][i].Section < out[ref][j].Section
			}
			return out[ref][i].Path < out[ref][j].Path
		})
	}
	return out
}

func isPhysicalPathReference(value string) bool {
	return strings.Contains(value, "/") || strings.Contains(value, `\`) || strings.HasSuffix(strings.ToLower(value), ".md")
}

func isExplicitUnsupportedReference(value string) bool {
	return strings.HasPrefix(value, "internal-design:") ||
		strings.HasPrefix(value, "coverage:") ||
		unsupportedIDPattern.MatchString(value)
}
