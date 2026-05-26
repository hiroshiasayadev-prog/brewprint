package designrecords

import (
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"time"

	"gopkg.in/yaml.v3"
)

var (
	adrH1Pattern                   = regexp.MustCompile(`^#\s+(\d{3}):\s+(.+?)\s*$`)
	investigationH1Pattern         = regexp.MustCompile(`^#\s+(INV-[A-Z0-9-]+-\d{3}):\s+(.+?)\s*$`)
	specH1Pattern                  = regexp.MustCompile(`^#\s+(.+?)\s*$`)
	atxHeadingPattern              = regexp.MustCompile(`^(#{1,6})\s+(.+?)\s*$`)
	metadataPattern                = regexp.MustCompile(`^-\s+\*\*([^*]+)\*\*:\s*(.*)$`)
	filenameNumPattern             = regexp.MustCompile(`^(\d{3})(?:-|\.md$)`)
	investigationFilenameIDPattern = regexp.MustCompile(`^(INV-[A-Z0-9-]+-\d{3})(?:-|\.md$)`)
	requirementIDPattern           = regexp.MustCompile(`^REQ-[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?-\d{3}$`)
	workItemIDPattern              = regexp.MustCompile(`^WORK-[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?-\d{3}$`)
	taskIDPattern                  = regexp.MustCompile(`^TASK-[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?-\d{3}-\d{2}$`)
	requirementFilenameIDPattern   = regexp.MustCompile(`^(REQ-[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?-\d{3})(?:-|$)`)
	workItemFilenameIDPattern      = regexp.MustCompile(`^(WORK-[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?-\d{3})(?:-|$)`)
	taskFilenameIDPattern          = regexp.MustCompile(`^(TASK-[A-Z0-9](?:[A-Z0-9-]*[A-Z0-9])?-\d{3}-\d{2})(?:-|$)`)
	datePattern                    = regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
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
	Status       string            `yaml:"status"`
	SemanticRefs []string          `yaml:"semantic_refs"`
	Sections     map[string]string `yaml:"sections"`
	DesignRecord *struct {
		ID             string   `yaml:"id"`
		Kind           string   `yaml:"kind"`
		Status         string   `yaml:"status"`
		DependsOn      []string `yaml:"depends_on"`
		Supersedes     []string `yaml:"supersedes"`
		MigratedToSpec string   `yaml:"migrated_to_spec"`
	} `yaml:"design_record"`
}

type investigationMetadata struct {
	Status                string
	Trigger               string
	Scope                 string
	NonScope              string
	SourceRefs            []string
	FollowUpCandidates    []string
	Supersedes            []string
	RelatedRequirements   []string
	RelatedWorkItems      []string
	RelatedADRs           []string
	RelatedSpecs          []string
	RelatedInternalDesign []string
	RelatedCoverage       []string
	FollowUpResults       []string
}

type requirementMetadata struct {
	ID         string
	Status     string
	SourceRefs []string
	WorkItems  []string
}

type workItemMetadata struct {
	ID                string
	Status            string
	SourceRequirement string
	ImpactRefs        []string
	Tasks             []string
}

type taskMetadata struct {
	ID                string
	Status            string
	WorkItem          string
	SourceRequirement string
	Estimate          string
	DependsOn         []string
	Outputs           []string
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
		ID:     id,
		Kind:   RecordKindDecision,
		Title:  title,
		Status: RecordStatus(metadata.Status),
		Path:   path,
		Decision: &DecisionDetail{
			DependsOn:      metadata.DependsOn,
			Supersedes:     metadata.Supersedes,
			MigratedToSpec: metadata.MigratedToSpec,
		},
		Headings:     extractHeadings(raw),
		RawBody:      raw,
		NormalizedID: normalizeRecordID(id),
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
		ID:           fm.DesignRecord.ID,
		Kind:         kind,
		Title:        title,
		Status:       RecordStatus(fm.Status),
		Path:         path,
		Spec:         &SpecDetail{DependsOn: append([]string(nil), fm.DesignRecord.DependsOn...)},
		SemanticRefs: semanticRefDecls(path, fm.SemanticRefs, fm.Sections),
		Headings:     extractHeadings(raw),
		RawBody:      raw,
		NormalizedID: normalizeRecordID(fm.DesignRecord.ID),
	}
	return record, candidate, issues
}

func parseSpecSemanticRefSource(path, raw string) (SemanticRefSource, bool) {
	fmBytes, _, ok := extractFrontMatter(raw)
	if !ok {
		return SemanticRefSource{}, false
	}
	var fm specFrontMatter
	if err := yaml.Unmarshal([]byte(fmBytes), &fm); err != nil {
		return SemanticRefSource{}, false
	}
	decls := semanticRefDecls(path, fm.SemanticRefs, fm.Sections)
	if len(decls) == 0 {
		return SemanticRefSource{}, false
	}
	recordID := ""
	if fm.DesignRecord != nil {
		recordID = fm.DesignRecord.ID
	}
	return SemanticRefSource{
		Path:     path,
		RecordID: recordID,
		Decls:    decls,
		Headings: extractHeadings(raw),
	}, true
}

func parseInvestigationRecord(path, raw string) (*Record, RecordCandidate, []ParseIssue) {
	lines := splitMarkdownLines(raw)
	h1Line := firstLine(lines)
	filenameID := investigationFilenameID(path)
	candidate := RecordCandidate{
		Path:           path,
		Kind:           RecordKindInvestigation,
		H1Line:         h1Line,
		FilenameNumber: filenameID,
		Included:       true,
	}
	var issues []ParseIssue

	id, title, h1OK := parseInvestigationH1(h1Line)
	candidate.H1Valid = h1OK
	candidate.H1Number = id
	if !h1OK {
		candidate.Included = false
		candidate.SkipReason = "invalid_investigation_h1"
		issues = append(issues, ParseIssue{
			Category: DiagnosticInvalidH1Title,
			Path:     path,
			Message:  "investigation H1 is missing or invalid",
			Details:  map[string]string{"h1": h1Line},
		})
		return nil, candidate, issues
	}
	candidate.ID = id
	candidate.NormalizedID = normalizeRecordID(id)
	if filenameID != id {
		candidate.FilenameIDMismatch = true
		issues = append(issues, ParseIssue{
			Category: DiagnosticFilenameIDMismatch,
			Path:     path,
			RecordID: id,
			Message:  "investigation filename ID is missing or does not match H1 ID",
			Details: map[string]string{
				"h1_id":       id,
				"filename_id": filenameID,
			},
		})
	}

	metadata := parseInvestigationMetadata(lines)
	record := &Record{
		ID:     id,
		Kind:   RecordKindInvestigation,
		Title:  title,
		Status: RecordStatus(metadata.Status),
		Path:   path,
		Investigation: &InvestigationDetail{
			Trigger:               metadata.Trigger,
			Scope:                 metadata.Scope,
			NonScope:              metadata.NonScope,
			SourceRefs:            metadata.SourceRefs,
			FollowUpCandidates:    metadata.FollowUpCandidates,
			Supersedes:            metadata.Supersedes,
			RelatedRequirements:   metadata.RelatedRequirements,
			RelatedWorkItems:      metadata.RelatedWorkItems,
			RelatedADRs:           metadata.RelatedADRs,
			RelatedSpecs:          metadata.RelatedSpecs,
			RelatedInternalDesign: metadata.RelatedInternalDesign,
			RelatedCoverage:       metadata.RelatedCoverage,
			FollowUpResults:       metadata.FollowUpResults,
		},
		Headings:     extractHeadings(raw),
		RawBody:      raw,
		NormalizedID: normalizeRecordID(id),
	}
	return record, candidate, issues
}

func parseInvestigationH1(line string) (string, string, bool) {
	match := investigationH1Pattern.FindStringSubmatch(trimLineEnd(line))
	if match == nil {
		return "", "", false
	}
	title := strings.TrimSpace(match[2])
	if title == "" {
		return "", "", false
	}
	return match[1], title, true
}

func parseInvestigationMetadata(lines []string) investigationMetadata {
	metadata := investigationMetadata{
		SourceRefs:         []string{},
		FollowUpCandidates: []string{},
	}
	block := metadataBlock(lines)
	for i := 0; i < len(block); i++ {
		line := trimLineEnd(block[i])
		match := metadataPattern.FindStringSubmatch(line)
		if match == nil {
			continue
		}
		key := match[1]
		value := strings.TrimSpace(match[2])
		switch key {
		case "status":
			metadata.Status = value
		case "date":
			continue
		case "trigger":
			metadata.Trigger = value
		case "scope":
			metadata.Scope = value
		case "non_scope":
			metadata.NonScope = value
		case "source_refs":
			metadata.SourceRefs = collectIndentedList(block, i)
		case "follow_up_candidates":
			metadata.FollowUpCandidates = collectIndentedList(block, i)
		case "supersedes":
			metadata.Supersedes = collectIndentedList(block, i)
		case "related_requirements":
			metadata.RelatedRequirements = collectIndentedList(block, i)
		case "related_work_items":
			metadata.RelatedWorkItems = collectIndentedList(block, i)
		case "related_adrs":
			metadata.RelatedADRs = collectIndentedList(block, i)
		case "related_specs":
			metadata.RelatedSpecs = collectIndentedList(block, i)
		case "related_internal_design":
			metadata.RelatedInternalDesign = collectIndentedList(block, i)
		case "related_coverage":
			metadata.RelatedCoverage = collectIndentedList(block, i)
		case "follow_up_results":
			metadata.FollowUpResults = collectIndentedList(block, i)
		}
	}
	return metadata
}

func parseRequirementRecord(path, raw string) (*Record, RecordCandidate, []ParseIssue) {
	lines := splitMarkdownLines(raw)
	return parseWorkflowRecord(path, raw, lines, RecordKindRequirement)
}

func parseWorkItemRecord(path, raw string) (*Record, RecordCandidate, []ParseIssue) {
	lines := splitMarkdownLines(raw)
	return parseWorkflowRecord(path, raw, lines, RecordKindWorkItem)
}

func parseTaskRecord(path, raw string) (*Record, RecordCandidate, []ParseIssue) {
	lines := splitMarkdownLines(raw)
	return parseWorkflowRecord(path, raw, lines, RecordKindTask)
}

func parseWorkflowRecord(path, raw string, lines []string, kind RecordKind) (*Record, RecordCandidate, []ParseIssue) {
	h1Line := firstLine(lines)
	filenameID := workflowFilenameID(path, kind)
	candidate := RecordCandidate{
		Path:           path,
		Kind:           kind,
		H1Line:         h1Line,
		FilenameNumber: filenameID,
		Included:       true,
	}
	var issues []ParseIssue

	h1ID, title, h1FormOK := parseWorkflowH1(h1Line)
	candidate.H1Valid = h1FormOK
	candidate.H1Number = h1ID
	if !h1FormOK {
		candidate.Included = false
		candidate.SkipReason = "invalid_workflow_h1"
		issues = append(issues, ParseIssue{
			Category: DiagnosticInvalidH1Title,
			Path:     path,
			Message:  "workflow H1 is missing or invalid",
			Details:  map[string]string{"h1": h1Line},
		})
		return nil, candidate, issues
	}

	metadataID, status := "", ""
	var requirement *RequirementDetail
	var workItem *WorkItemDetail
	var task *TaskDetail
	switch kind {
	case RecordKindRequirement:
		metadata := parseRequirementMetadata(lines)
		metadataID = metadata.ID
		status = metadata.Status
		requirement = &RequirementDetail{
			SourceRefs: metadata.SourceRefs,
			WorkItems:  metadata.WorkItems,
		}
	case RecordKindWorkItem:
		metadata := parseWorkItemMetadata(lines)
		metadataID = metadata.ID
		status = metadata.Status
		workItem = &WorkItemDetail{
			SourceRequirement: metadata.SourceRequirement,
			ImpactRefs:        metadata.ImpactRefs,
			Tasks:             metadata.Tasks,
		}
	case RecordKindTask:
		metadata := parseTaskMetadata(lines)
		metadataID = metadata.ID
		status = metadata.Status
		task = &TaskDetail{
			WorkItem:          metadata.WorkItem,
			SourceRequirement: metadata.SourceRequirement,
			Estimate:          metadata.Estimate,
			DependsOn:         metadata.DependsOn,
			Outputs:           metadata.Outputs,
		}
	}

	if !validWorkflowIDForKind(h1ID, kind) || (metadataID != "" && !validWorkflowIDForKind(metadataID, kind)) || (filenameID != "" && !validWorkflowIDForKind(filenameID, kind)) {
		candidate.Included = false
		candidate.SkipReason = "invalid_workflow_id"
		issues = append(issues, ParseIssue{
			Category: DiagnosticInvalidWorkflowID,
			Path:     path,
			RecordID: h1ID,
			Message:  "workflow ID does not match the required grammar",
			Details: map[string]string{
				"h1_id":       h1ID,
				"metadata_id": metadataID,
				"filename_id": filenameID,
			},
		})
		return nil, candidate, issues
	}

	id := h1ID
	candidate.ID = id
	candidate.NormalizedID = normalizeRecordID(id)
	if metadataID == "" || metadataID != h1ID || filenameID == "" || filenameID != h1ID {
		candidate.FilenameIDMismatch = true
		issues = append(issues, ParseIssue{
			Category: DiagnosticFilenameIDMismatch,
			Path:     path,
			RecordID: id,
			Message:  "workflow metadata ID, H1 ID, and filename ID must match",
			Details: map[string]string{
				"h1_id":       h1ID,
				"metadata_id": metadataID,
				"filename_id": filenameID,
			},
		})
	}

	record := &Record{
		ID:           id,
		Kind:         kind,
		Title:        title,
		Status:       RecordStatus(status),
		Path:         path,
		Requirement:  requirement,
		WorkItem:     workItem,
		Task:         task,
		Headings:     extractHeadings(raw),
		RawBody:      raw,
		NormalizedID: normalizeRecordID(id),
	}
	return record, candidate, issues
}

func parseWorkflowH1(line string) (string, string, bool) {
	line = trimLineEnd(line)
	if !strings.HasPrefix(line, "# ") {
		return "", "", false
	}
	rest := strings.TrimSpace(strings.TrimPrefix(line, "# "))
	id, title, ok := strings.Cut(rest, ": ")
	if !ok {
		return "", "", false
	}
	title = strings.TrimSpace(title)
	if strings.TrimSpace(id) != id || id == "" || title == "" {
		return "", "", false
	}
	if !strings.HasPrefix(id, "REQ-") && !strings.HasPrefix(id, "WORK-") && !strings.HasPrefix(id, "TASK-") {
		return "", "", false
	}
	return id, title, true
}

func parseRequirementMetadata(lines []string) requirementMetadata {
	metadata := requirementMetadata{SourceRefs: []string{}, WorkItems: []string{}}
	block := metadataBlock(lines)
	for i := 0; i < len(block); i++ {
		key, value, ok := parseMetadataLine(block[i])
		if !ok {
			continue
		}
		switch key {
		case "id":
			metadata.ID = value
		case "status":
			metadata.Status = value
		case "date":
			continue
		case "source_refs":
			metadata.SourceRefs = metadataListValue(block, i, value)
		case "work_items":
			metadata.WorkItems = metadataListValue(block, i, value)
		}
	}
	return metadata
}

func parseWorkItemMetadata(lines []string) workItemMetadata {
	metadata := workItemMetadata{ImpactRefs: []string{}, Tasks: []string{}}
	block := metadataBlock(lines)
	for i := 0; i < len(block); i++ {
		key, value, ok := parseMetadataLine(block[i])
		if !ok {
			continue
		}
		switch key {
		case "id":
			metadata.ID = value
		case "status":
			metadata.Status = value
		case "date":
			continue
		case "source_requirement":
			metadata.SourceRequirement = value
		case "impact_refs":
			metadata.ImpactRefs = metadataListValue(block, i, value)
		case "tasks":
			metadata.Tasks = metadataListValue(block, i, value)
		}
	}
	return metadata
}

func parseTaskMetadata(lines []string) taskMetadata {
	metadata := taskMetadata{DependsOn: []string{}, Outputs: []string{}}
	block := metadataBlock(lines)
	for i := 0; i < len(block); i++ {
		key, value, ok := parseMetadataLine(block[i])
		if !ok {
			continue
		}
		switch key {
		case "id":
			metadata.ID = value
		case "status":
			metadata.Status = value
		case "date":
			continue
		case "work_item":
			metadata.WorkItem = value
		case "source_requirement":
			metadata.SourceRequirement = value
		case "estimate":
			metadata.Estimate = value
		case "depends_on":
			metadata.DependsOn = metadataListValue(block, i, value)
		case "outputs":
			metadata.Outputs = metadataListValue(block, i, value)
		}
	}
	return metadata
}

func parseMetadataLine(line string) (string, string, bool) {
	match := metadataPattern.FindStringSubmatch(trimLineEnd(line))
	if match == nil {
		return "", "", false
	}
	return match[1], strings.TrimSpace(match[2]), true
}

func metadataListValue(block []string, index int, value string) []string {
	if strings.TrimSpace(value) != "" {
		return splitCommaList(value)
	}
	return collectIndentedList(block, index)
}

func collectIndentedList(block []string, index int) []string {
	out := []string{}
	for _, line := range block[index+1:] {
		raw := trimLineEnd(line)
		if metadataPattern.MatchString(raw) {
			break
		}
		trimmedLeft := strings.TrimLeft(raw, " \t")
		if len(raw) == len(trimmedLeft) {
			continue
		}
		if !strings.HasPrefix(trimmedLeft, "- ") {
			continue
		}
		item := strings.TrimSpace(strings.TrimPrefix(trimmedLeft, "- "))
		if item != "" {
			out = append(out, item)
		}
	}
	return out
}

func semanticRefDecls(path string, semanticRefs []string, sections map[string]string) []SemanticRefDecl {
	out := make([]SemanticRefDecl, 0, len(semanticRefs)+len(sections))
	for _, ref := range semanticRefs {
		out = append(out, SemanticRefDecl{
			Ref:        strings.TrimSpace(ref),
			Path:       path,
			TargetType: SemanticTargetDocument,
		})
	}
	keys := make([]string, 0, len(sections))
	for ref := range sections {
		keys = append(keys, ref)
	}
	sort.Strings(keys)
	for _, ref := range keys {
		out = append(out, SemanticRefDecl{
			Ref:        strings.TrimSpace(ref),
			Path:       path,
			TargetType: SemanticTargetSection,
			Section:    strings.TrimSpace(sections[ref]),
		})
	}
	return out
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

func investigationFilenameID(path string) string {
	base := filepath.Base(path)
	match := investigationFilenameIDPattern.FindStringSubmatch(base)
	if match == nil {
		return ""
	}
	return match[1]
}

func workflowFilenameID(path string, kind RecordKind) string {
	base := filepath.Base(path)
	base = strings.TrimSuffix(base, filepath.Ext(base))
	switch kind {
	case RecordKindRequirement:
		return filenameIDMatch(requirementFilenameIDPattern, base)
	case RecordKindWorkItem:
		return filenameIDMatch(workItemFilenameIDPattern, base)
	case RecordKindTask:
		return filenameIDMatch(taskFilenameIDPattern, base)
	default:
		return ""
	}
}

func filenameIDMatch(pattern *regexp.Regexp, base string) string {
	match := pattern.FindStringSubmatch(base)
	if match != nil {
		return match[1]
	}
	if strings.HasPrefix(base, "REQ-") || strings.HasPrefix(base, "WORK-") || strings.HasPrefix(base, "TASK-") {
		return base
	}
	return ""
}

func validWorkflowIDForKind(id string, kind RecordKind) bool {
	switch kind {
	case RecordKindRequirement:
		return requirementIDPattern.MatchString(id)
	case RecordKindWorkItem:
		return workItemIDPattern.MatchString(id)
	case RecordKindTask:
		return taskIDPattern.MatchString(id)
	default:
		return false
	}
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
