package query

import (
	"strings"

	"github.com/hiroshiasayadev-prog/brewprint/internal/semantic"
)

type mechanicalJudgementInput struct {
	From                    string
	To                      string
	Source                  *SourceLocation
	YAMLStructurePreserving bool
	ReferenceStable         bool
}

type mechanicalJudgementResult struct {
	Mechanical bool
	Reason     string
}

func (s *Service) mechanicalJudgementGate(input mechanicalJudgementInput) mechanicalJudgementResult {
	if input.Source == nil || input.Source.File == "" || input.Source.Line <= 0 || input.Source.Column <= 0 {
		return mechanicalJudgementResult{Reason: "source location is not precise"}
	}
	if input.From == "" || input.To == "" {
		return mechanicalJudgementResult{Reason: "replacement token is not fully specified"}
	}
	if !input.YAMLStructurePreserving {
		return mechanicalJudgementResult{Reason: "replacement may change YAML structure"}
	}
	if !input.ReferenceStable {
		return mechanicalJudgementResult{Reason: "replacement reference stability is not proven"}
	}
	content, ok := s.fileContent(semantic.FileID(input.Source.File))
	if !ok {
		return mechanicalJudgementResult{Reason: "source file is unavailable"}
	}
	if strings.Count(content, input.From) != 1 {
		return mechanicalJudgementResult{Reason: "replacement token is not unique in source file"}
	}
	lines := splitSourceLines(content)
	lineIndex := input.Source.Line - 1
	if lineIndex < 0 || lineIndex >= len(lines) {
		return mechanicalJudgementResult{Reason: "source line is out of range"}
	}
	line := lines[lineIndex]
	if strings.Count(line, input.From) != 1 {
		return mechanicalJudgementResult{Reason: "replacement token is not unique on source line"}
	}
	return mechanicalJudgementResult{Mechanical: true}
}

func sourceLocationPrecise(source *SourceLocation) bool {
	return source != nil && source.File != "" && source.Line > 0 && source.Column > 0
}
