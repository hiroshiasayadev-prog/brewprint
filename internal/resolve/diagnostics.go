package resolve

import (
	"sort"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

func sortedDiagnostics(diagnostics []semantic.Diagnostic) []semantic.Diagnostic {
	out := append([]semantic.Diagnostic(nil), diagnostics...)
	sort.SliceStable(out, func(i, j int) bool {
		return diagnosticLess(out[i], out[j])
	})
	return out
}

func diagnosticLess(left semantic.Diagnostic, right semantic.Diagnostic) bool {
	if severityRank(left.Severity) != severityRank(right.Severity) {
		return severityRank(left.Severity) < severityRank(right.Severity)
	}
	if left.FileID != right.FileID {
		return left.FileID < right.FileID
	}
	if left.Code != right.Code {
		return left.Code < right.Code
	}
	return left.Message < right.Message
}

func severityRank(severity semantic.Severity) int {
	switch severity {
	case semantic.SeverityError:
		return 0
	case semantic.SeverityWarning:
		return 1
	default:
		return 2
	}
}
