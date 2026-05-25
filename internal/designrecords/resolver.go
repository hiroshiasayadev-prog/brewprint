package designrecords

import (
	"fmt"
	"regexp"
	"sort"
	"strings"
)

var (
	activeSpecRefPattern = regexp.MustCompile(`^spec:[a-z0-9-]+(?:\.[a-z0-9-]+)+$`)
	recordIDRefPattern   = regexp.MustCompile(`^(ADR-\d{3}|SPEC-[A-Za-z0-9][A-Za-z0-9-]*|INV-[A-Z0-9-]+-\d{3})$`)
	unsupportedIDPattern = regexp.MustCompile(`^(COV|REQ|WORK)-`)
)

const (
	refKindSemanticRef = "semantic_ref"
	refKindRecordID    = "record_id"
	refKindUnsupported = "unsupported"

	resolveStatusResolved    = "resolved"
	resolveStatusUnresolved  = "unresolved"
	resolveStatusUnsupported = "unsupported"
)

func resolveReference(idx *Index, ref string) ResolveReferenceResponse {
	switch classifyReference(ref) {
	case refKindSemanticRef:
		return resolveSemanticReference(idx, ref)
	case refKindRecordID:
		return resolveRecordReference(idx, ref)
	default:
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
}

func classifyReference(ref string) string {
	if activeSpecRefPattern.MatchString(ref) {
		return refKindSemanticRef
	}
	if recordIDRefPattern.MatchString(ref) {
		return refKindRecordID
	}
	return refKindUnsupported
}

func resolveSemanticReference(idx *Index, ref string) ResolveReferenceResponse {
	targets := semanticTargetsByRef(idx)[ref]
	switch len(targets) {
	case 0:
		return unresolvedResponse(ref, refKindSemanticRef)
	case 1:
		target := targets[0]
		out := &ResolvedTarget{
			TargetType: string(target.TargetType),
			Path:       target.Path,
		}
		if target.TargetType == SemanticTargetSection {
			out.Section = target.Section
		}
		return ResolveReferenceResponse{Ref: ref, RefKind: refKindSemanticRef, Status: resolveStatusResolved, Target: out, Diagnostics: []Diagnostic{}}
	default:
		return ambiguousResponse(ref, refKindSemanticRef)
	}
}

func resolveRecordReference(idx *Index, ref string) ResolveReferenceResponse {
	records := recordsByNormalizedID(idx)[normalizeRecordID(ref)]
	switch len(records) {
	case 0:
		return unresolvedResponse(ref, refKindRecordID)
	case 1:
		record := records[0]
		return ResolveReferenceResponse{
			Ref:     ref,
			RefKind: refKindRecordID,
			Status:  resolveStatusResolved,
			Target: &ResolvedTarget{
				TargetType: "record",
				Path:       record.Path,
				RecordID:   record.ID,
				RecordKind: record.Kind,
				Title:      record.Title,
				Status:     record.Status,
			},
			Diagnostics: []Diagnostic{},
		}
	default:
		return ambiguousResponse(ref, refKindRecordID)
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

func recordsByNormalizedID(idx *Index) map[string][]Record {
	out := make(map[string][]Record, len(idx.Records))
	for _, record := range idx.Records {
		if record.NormalizedID != "" {
			out[record.NormalizedID] = append(out[record.NormalizedID], record)
		}
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
