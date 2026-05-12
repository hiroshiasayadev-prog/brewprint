package designrecords

import (
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

var (
	adrH1Pattern       = regexp.MustCompile(`^#\s+(\d{3}):\s+(.+?)\s*$`)
	specH1Pattern      = regexp.MustCompile(`^#\s+(.+?)\s*$`)
	atxHeadingPattern  = regexp.MustCompile(`^(#{1,6})\s+(.+?)\s*$`)
	metadataPattern    = regexp.MustCompile(`^-\s+\*\*([^*]+)\*\*:\s*(.*)$`)
	filenameNumPattern = regexp.MustCompile(`^(\d{3})(?:-|\.md$)`)
	datePattern        = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
)

type adrMetadata struct {
	Status                  string
	DependsOn               []string
	Supersedes              []string
	MigratedToSpec          *string
	MigratedToSpecRaw       string
	MigratedToSpecSpecified bool
}

type specFrontMatter struct {
	Status       string `yaml:"status"`
	DesignRecord *struct {
		ID             string   `yaml:"id"`
		Kind           string   `yaml:"kind"`
		Status         string   `yaml:"status"`
		DependsOn      []string `yaml:"depends_on"`
		Supersedes     []string `yaml:"supersedes"`
		MigratedToSpec string   `yaml:"migrated_to_spec"`
	} `yaml:"design_record"`
}

func parseADRRecord(path, raw string) (*Record, RecordCandidate, []ParseIssue) {
	lines := splitMarkdownLines(raw)
	h1Line := firstLine(lines)
	filenameNumber := filenameNumber(path)
	candidate := RecordCandidate{
		Path:           path,
		Kind:           RecordKindDecision,
		H1Line:         h1Line,
		FilenameNumber: filenameNumber,
		Included:       true,
	}
	var issues []ParseIssue

	num, title, h1OK := parseADRH1(h1Line)
	candidate.H1Valid = h1OK
	candidate.H1Number = num
	if !h1OK {
		candidate.Included = false
		candidate.SkipReason = "invalid_adr_h1"
		issues = append(issues, ParseIssue{
			Category: DiagnosticInvalidH1Title,
			Path:     path,
			Message:  "ADR H1 is missing or invalid",
			Details:  map[string]string{"h1": h1Line},
		})
		return nil, candidate, issues
	}

	id := "ADR-" + num
	candidate.ID = id
	candidate.NormalizedID = normalizeRecordID(id)
	if filenameNumber != num {
		candidate.FilenameIDMismatch = true
		issues = append(issues, ParseIssue{
			Category: DiagnosticFilenameIDMismatch,
			Path:     path,
			RecordID: id,
			Message:  "ADR filename number is missing or does not match H1 number",
			Details: map[string]string{
				"h1_number":       num,
				"filename_number": filenameNumber,
			},
		})
	}

	metadata, metadataIssues := parseADRMetadata(lines, path, id)
	issues = append(issues, metadataIssues...)
	record := &Record{
		ID:             id,
		Kind:           RecordKindDecision,
		Title:          title,
		Status:         RecordStatus(metadata.Status),
		Path:           path,
		DependsOn:      metadata.DependsOn,
		Supersedes:     metadata.Supersedes,
		MigratedToSpec: metadata.MigratedToSpec,
		Headings:       extractHeadings(raw),
		RawBody:        raw,
		NormalizedID:   normalizeRecordID(id),
	}
	return record, candidate, issues
}

func parseADRH1(line string) (string, string, bool) {
	match := adrH1Pattern.FindStringSubmatch(trimLineEnd(line))
	if match == nil {
		return "", "", false
	}
	title := strings.TrimSpace(match[2])
	if title == "" {
		return "", "", false
	}
	return match[1], title, true
}

func parseADRMetadata(lines []string, path, recordID string) (adrMetadata, []ParseIssue) {
	metadata := adrMetadata{
		DependsOn:  []string{},
		Supersedes: []string{},
	}
	var issues []ParseIssue
	for _, line := range metadataBlock(lines) {
		match := metadataPattern.FindStringSubmatch(trimLineEnd(line))
		if match == nil {
			continue
		}
		key := match[1]
		value := strings.TrimSpace(match[2])
		switch key {
		case "status":
			if value != "" {
				metadata.Status = value
			}
		case "date":
			continue
		case "depends_on":
			metadata.DependsOn = splitCommaList(value)
		case "supersedes":
			metadata.Supersedes = splitCommaList(value)
		case "migrated_to_spec":
			metadata.MigratedToSpecRaw = value
			if value == "" {
				metadata.MigratedToSpec = nil
				continue
			}
			metadata.MigratedToSpecSpecified = true
			metadata.MigratedToSpec = stringPtr(value)
			if !validDateOnly(value) {
				issues = append(issues, ParseIssue{
					Category: DiagnosticInvalidMigratedToSpec,
					Path:     path,
					RecordID: recordID,
					Message:  "ADR migrated_to_spec is not YYYY-MM-DD",
					Details:  map[string]string{"value": value},
				})
			}
		}
	}
	return metadata, issues
}

func metadataBlock(lines []string) []string {
	if len(lines) <= 1 {
		return nil
	}
	var block []string
	for _, line := range lines[1:] {
		line = trimLineEnd(line)
		if strings.HasPrefix(line, "##") || strings.HasPrefix(line, ">") {
			break
		}
		block = append(block, line)
	}
	return block
}

func parseSpecRecord(path, raw string) (*Record, RecordCandidate, []ParseIssue) {
	fmBytes, _, ok := extractFrontMatter(raw)
	if !ok {
		return nil, RecordCandidate{}, nil
	}
	var fm specFrontMatter
	if err := yaml.Unmarshal([]byte(fmBytes), &fm); err != nil {
		return nil, RecordCandidate{}, nil
	}
	if fm.DesignRecord == nil || fm.DesignRecord.ID == "" || fm.DesignRecord.Kind == "" {
		return nil, RecordCandidate{}, nil
	}
	kind := RecordKind(fm.DesignRecord.Kind)
	candidate := RecordCandidate{
		Path:         path,
		Kind:         kind,
		ID:           fm.DesignRecord.ID,
		NormalizedID: normalizeRecordID(fm.DesignRecord.ID),
		Included:     true,
	}
	if kind != RecordKindDecision && kind != RecordKindSpec {
		candidate.Included = false
		candidate.SkipReason = "unsupported_design_record_kind"
		return nil, candidate, nil
	}

	title, h1Line, h1OK := parseSpecH1(raw)
	candidate.H1Line = h1Line
	candidate.H1Valid = h1OK
	var issues []ParseIssue
	if !h1OK {
		issues = append(issues, ParseIssue{
			Category: DiagnosticInvalidH1Title,
			Path:     path,
			RecordID: fm.DesignRecord.ID,
			Message:  "spec H1 is missing or invalid",
			Details:  map[string]string{"h1": h1Line},
		})
	}
	if fm.Status != "" && fm.DesignRecord.Status != "" && fm.DesignRecord.Status != fm.Status {
		issues = append(issues, ParseIssue{
			Category: DiagnosticSpecStatusMismatch,
			Path:     path,
			RecordID: fm.DesignRecord.ID,
			Message:  "spec top-level status does not match design_record.status",
			Details: map[string]string{
				"status":               fm.Status,
				"design_record.status": fm.DesignRecord.Status,
			},
		})
	}

	record := &Record{
		ID:             fm.DesignRecord.ID,
		Kind:           kind,
		Title:          title,
		Status:         RecordStatus(fm.Status),
		Path:           path,
		DependsOn:      append([]string(nil), fm.DesignRecord.DependsOn...),
		Supersedes:     []string{},
		MigratedToSpec: nil,
		Headings:       extractHeadings(raw),
		RawBody:        raw,
		NormalizedID:   normalizeRecordID(fm.DesignRecord.ID),
	}
	return record, candidate, issues
}

func parseSpecH1(raw string) (string, string, bool) {
	for _, line := range contentLinesOutsideFrontMatterAndFences(raw) {
		line = trimLineEnd(line)
		if !strings.HasPrefix(line, "#") {
			continue
		}
		match := specH1Pattern.FindStringSubmatch(line)
		if match == nil {
			return "", line, false
		}
		title := strings.TrimSpace(match[1])
		if title == "" {
			return "", line, false
		}
		return title, line, true
	}
	return "", "", false
}

func extractHeadings(raw string) []Heading {
	var headings []Heading
	for _, line := range contentLinesOutsideFrontMatterAndFences(raw) {
		match := atxHeadingPattern.FindStringSubmatch(trimLineEnd(line))
		if match == nil {
			continue
		}
		text := strings.TrimSpace(match[2])
		if text == "" {
			continue
		}
		headings = append(headings, Heading{Level: len(match[1]), Text: text})
	}
	return headings
}

func contentLinesOutsideFrontMatterAndFences(raw string) []string {
	lines := splitMarkdownLines(raw)
	start := 0
	if hasOpeningFrontMatter(lines) {
		for i := 1; i < len(lines); i++ {
			if strings.TrimSpace(trimLineEnd(lines[i])) == "---" {
				start = i + 1
				break
			}
		}
	}
	var out []string
	inFence := false
	fenceMarker := ""
	for _, line := range lines[start:] {
		trimmed := strings.TrimSpace(trimLineEnd(line))
		if isFenceLine(trimmed) {
			marker := fencePrefix(trimmed)
			if !inFence {
				inFence = true
				fenceMarker = marker
			} else if marker == fenceMarker {
				inFence = false
				fenceMarker = ""
			}
			continue
		}
		if inFence {
			continue
		}
		out = append(out, line)
	}
	return out
}

func extractFrontMatter(raw string) (string, string, bool) {
	lines := splitMarkdownLines(raw)
	if !hasOpeningFrontMatter(lines) {
		return "", raw, false
	}
	var fm []string
	for i := 1; i < len(lines); i++ {
		line := trimLineEnd(lines[i])
		if strings.TrimSpace(line) == "---" {
			return strings.Join(fm, "\n"), strings.Join(lines[i+1:], "\n"), true
		}
		fm = append(fm, line)
	}
	return "", raw, false
}

func hasOpeningFrontMatter(lines []string) bool {
	return len(lines) > 0 && strings.TrimSpace(trimLineEnd(lines[0])) == "---"
}

func splitMarkdownLines(raw string) []string {
	return strings.Split(raw, "\n")
}

func firstLine(lines []string) string {
	if len(lines) == 0 {
		return ""
	}
	return trimLineEnd(lines[0])
}

func trimLineEnd(line string) string {
	return strings.TrimSuffix(line, "\r")
}

func filenameNumber(path string) string {
	base := filepath.Base(path)
	match := filenameNumPattern.FindStringSubmatch(base)
	if match == nil {
		return ""
	}
	return match[1]
}

func splitCommaList(value string) []string {
	value = strings.TrimSpace(value)
	if value == "" {
		return []string{}
	}
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		item := strings.TrimSpace(part)
		if item != "" {
			out = append(out, item)
		}
	}
	return out
}

func validDateOnly(value string) bool {
	if !datePattern.MatchString(value) {
		return false
	}
	parsed, err := time.Parse("2006-01-02", value)
	return err == nil && parsed.Format("2006-01-02") == value
}

func isFenceLine(trimmed string) bool {
	return strings.HasPrefix(trimmed, "```") || strings.HasPrefix(trimmed, "~~~")
}

func fencePrefix(trimmed string) string {
	if strings.HasPrefix(trimmed, "```") {
		return "```"
	}
	return "~~~"
}

func normalizeRecordID(id string) string {
	return strings.ToUpper(strings.TrimSpace(id))
}

func stringPtr(value string) *string {
	v := value
	return &v
}
