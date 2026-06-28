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
	// adrH1Pattern matches the bare ADR H1 form after namespace prefix stripping.
	// Accepts "# ADR-076: title" (ns stripped from "# V01-ADR-076: title").
	adrH1Pattern                   = regexp.MustCompile(`^#\s+ADR-(\d{3}):\s+(.+?)\s*$`)
	adrFilenamePattern             = regexp.MustCompile(`^ADR-(\d{3})(?:-|\.md$)`)
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
	Subdomain  string
	Workflow   WorkflowMetadata
}

type workItemMetadata struct {
	ID                string
	Status            string
	SourceRequirement string
	ImpactRefs        []string
	Tasks             []string
	Subdomain         string
	Workflow          WorkflowMetadata
}

type taskMetadata struct {
	ID                string
	Status            string
	WorkItem          string
	SourceRequirement string
	Estimate          string
	DependsOn         []string
	Outputs           []string
	Subdomain         string
	Workflow          WorkflowMetadata
}

// stripNamespacePrefix removes ns from the start of s. Returns s unchanged when
// ns is empty or s does not start with ns.
func stripNamespacePrefix(s, ns string) string {
	if ns != "" && strings.HasPrefix(s, ns) {
		return s[len(ns):]
	}
	return s
}

func parseADRRecord(path, raw, ns string) (*Record, RecordCandidate, []ParseIssue) {
	lines := splitMarkdownLines(raw)
	h1Line := firstLine(lines)
	filenameNum := filenameNumber(path, ns)
	candidate := RecordCandidate{
		Path:           path,
		Kind:           RecordKindDecision,
		H1Line:         h1Line,
		FilenameNumber: filenameNum,
		Included:       true,
	}
	var issues []ParseIssue

	bareNum, title, h1OK := parseADRH1(h1Line, ns)
	candidate.H1Valid = h1OK
	candidate.H1Number = bareNum
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

	id := ns + "ADR-" + bareNum
	candidate.ID = id
	candidate.NormalizedID = normalizeRecordID(id)
	if filenameNum != bareNum {
		candidate.FilenameIDMismatch = true
		issues = append(issues, ParseIssue{
			Category: DiagnosticFilenameIDMismatch,
			Path:     path,
			RecordID: id,
			Message:  "ADR filename number is missing or does not match H1 number",
			Details: map[string]string{
				"h1_number":       bareNum,
				"filename_number": filenameNum,
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

// parseADRH1 strips the namespace prefix from line and matches the bare ADR H1
// form "# ADR-NNN: title". Returns (bareNum, title, ok).
func parseADRH1(line, ns string) (string, string, bool) {
	line = trimLineEnd(line)
	// Strip "# <ns>" from "# V01-ADR-076: title" → "# ADR-076: title"
	if ns != "" {
		stripped := "# " + stripNamespacePrefix(strings.TrimPrefix(line, "# "), ns)
		line = stripped
	}
	match := adrH1Pattern.FindStringSubmatch(line)
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

func parseInvestigationRecord(path, raw, ns string) (*Record, RecordCandidate, []ParseIssue) {
	lines := splitMarkdownLines(raw)
	h1Line := firstLine(lines)
	filenameID := investigationFilenameID(path, ns)
	candidate := RecordCandidate{
		Path:           path,
		Kind:           RecordKindInvestigation,
		H1Line:         h1Line,
		FilenameNumber: filenameID,
		Included:       true,
	}
	var issues []ParseIssue

	bareID, title, h1OK := parseInvestigationH1(h1Line, ns)
	candidate.H1Valid = h1OK
	candidate.H1Number = bareID
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

	publicID := ns + bareID
	candidate.ID = publicID
	candidate.NormalizedID = normalizeRecordID(publicID)
	if filenameID != bareID {
		candidate.FilenameIDMismatch = true
		issues = append(issues, ParseIssue{
			Category: DiagnosticFilenameIDMismatch,
			Path:     path,
			RecordID: publicID,
			Message:  "investigation filename ID is missing or does not match H1 ID",
			Details: map[string]string{
				"h1_id":       bareID,
				"filename_id": filenameID,
			},
		})
	}

	metadata := parseInvestigationMetadata(lines)
	record := &Record{
		ID:     publicID,
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
		NormalizedID: normalizeRecordID(publicID),
	}
	return record, candidate, issues
}

// parseInvestigationH1 strips ns and matches the bare INV H1 form.
// Returns (bareID, title, ok) where bareID is e.g. "INV-MCP-001".
func parseInvestigationH1(line, ns string) (string, string, bool) {
	line = trimLineEnd(line)
	if ns != "" {
		stripped := "# " + stripNamespacePrefix(strings.TrimPrefix(line, "# "), ns)
		line = stripped
	}
	match := investigationH1Pattern.FindStringSubmatch(line)
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

func parseRequirementRecord(path, raw, ns string) (*Record, RecordCandidate, []ParseIssue) {
	lines := splitMarkdownLines(raw)
	return parseWorkflowRecord(path, raw, lines, RecordKindRequirement, ns)
}

func parseWorkItemRecord(path, raw, ns string) (*Record, RecordCandidate, []ParseIssue) {
	lines := splitMarkdownLines(raw)
	return parseWorkflowRecord(path, raw, lines, RecordKindWorkItem, ns)
}

func parseTaskRecord(path, raw, ns string) (*Record, RecordCandidate, []ParseIssue) {
	lines := splitMarkdownLines(raw)
	return parseWorkflowRecord(path, raw, lines, RecordKindTask, ns)
}

func parseWorkflowRecord(path, raw string, lines []string, kind RecordKind, ns string) (*Record, RecordCandidate, []ParseIssue) {
	h1Line := firstLine(lines)
	filenameBareID := workflowFilenameID(path, kind, ns)
	candidate := RecordCandidate{
		Path:           path,
		Kind:           kind,
		H1Line:         h1Line,
		FilenameNumber: filenameBareID,
		Included:       true,
	}
	var issues []ParseIssue

	h1BareID, title, h1FormOK := parseWorkflowH1(h1Line, ns)
	candidate.H1Valid = h1FormOK
	candidate.H1Number = h1BareID
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

	metadataIDRaw, status := "", ""
	var workflowMeta *WorkflowMetadata
	var requirement *RequirementDetail
	var workItem *WorkItemDetail
	var task *TaskDetail
	switch kind {
	case RecordKindRequirement:
		metadata := parseRequirementMetadata(lines)
		metadataIDRaw = metadata.ID
		status = metadata.Status
		workflowMeta = &metadata.Workflow
		requirement = &RequirementDetail{
			SourceRefs: metadata.SourceRefs,
			WorkItems:  metadata.WorkItems,
			Subdomain:  optionalString(metadata.Subdomain),
		}
	case RecordKindWorkItem:
		metadata := parseWorkItemMetadata(lines)
		metadataIDRaw = metadata.ID
		status = metadata.Status
		workflowMeta = &metadata.Workflow
		workItem = &WorkItemDetail{
			SourceRequirement: metadata.SourceRequirement,
			ImpactRefs:        metadata.ImpactRefs,
			Tasks:             metadata.Tasks,
			Subdomain:         optionalString(metadata.Subdomain),
		}
	case RecordKindTask:
		metadata := parseTaskMetadata(lines)
		metadataIDRaw = metadata.ID
		status = metadata.Status
		workflowMeta = &metadata.Workflow
		task = &TaskDetail{
			WorkItem:          metadata.WorkItem,
			SourceRequirement: metadata.SourceRequirement,
			Estimate:          metadata.Estimate,
			DependsOn:         metadata.DependsOn,
			Outputs:           metadata.Outputs,
			Subdomain:         optionalString(metadata.Subdomain),
		}
	}

	// Strip ns from metadata id for bare-form validation and comparison.
	metadataBareID := stripNamespacePrefix(metadataIDRaw, ns)

	if !validWorkflowIDForKind(h1BareID, kind) || (metadataBareID != "" && !validWorkflowIDForKind(metadataBareID, kind)) || (filenameBareID != "" && !validWorkflowIDForKind(filenameBareID, kind)) {
		candidate.Included = false
		candidate.SkipReason = "invalid_workflow_id"
		issues = append(issues, ParseIssue{
			Category: DiagnosticInvalidWorkflowID,
			Path:     path,
			RecordID: ns + h1BareID,
			Message:  "workflow ID does not match the required grammar",
			Details: map[string]string{
				"h1_id":       h1BareID,
				"metadata_id": metadataBareID,
				"filename_id": filenameBareID,
			},
		})
		return nil, candidate, issues
	}

	publicID := ns + h1BareID
	candidate.ID = publicID
	candidate.NormalizedID = normalizeRecordID(publicID)
	if metadataBareID == "" || metadataBareID != h1BareID || filenameBareID == "" || filenameBareID != h1BareID {
		candidate.FilenameIDMismatch = true
		issues = append(issues, ParseIssue{
			Category: DiagnosticFilenameIDMismatch,
			Path:     path,
			RecordID: publicID,
			Message:  "workflow metadata ID, H1 ID, and filename ID must match",
			Details: map[string]string{
				"h1_id":       h1BareID,
				"metadata_id": metadataBareID,
				"filename_id": filenameBareID,
			},
		})
	}

	record := &Record{
		ID:           publicID,
		Kind:         kind,
		Title:        title,
		Status:       RecordStatus(status),
		Path:         path,
		Requirement:  requirement,
		WorkItem:     workItem,
		Task:         task,
		Headings:     extractHeadings(raw),
		RawBody:      raw,
		NormalizedID: normalizeRecordID(publicID),
		WorkflowMeta: workflowMeta,
	}
	return record, candidate, issues
}

// parseWorkflowH1 strips ns from the H1 ID and validates the bare prefix.
// Returns (bareID, title, ok) where bareID is e.g. "WORK-DRMCP-001".
func parseWorkflowH1(line, ns string) (string, string, bool) {
	line = trimLineEnd(line)
	if !strings.HasPrefix(line, "# ") {
		return "", "", false
	}
	rest := strings.TrimSpace(strings.TrimPrefix(line, "# "))
	rawID, title, ok := strings.Cut(rest, ": ")
	if !ok {
		return "", "", false
	}
	title = strings.TrimSpace(title)
	bareID := stripNamespacePrefix(rawID, ns)
	if strings.TrimSpace(bareID) != bareID || bareID == "" || title == "" {
		return "", "", false
	}
	if !strings.HasPrefix(bareID, "REQ-") && !strings.HasPrefix(bareID, "WORK-") && !strings.HasPrefix(bareID, "TASK-") {
		return "", "", false
	}
	return bareID, title, true
}

func parseRequirementMetadata(lines []string) requirementMetadata {
	metadata := requirementMetadata{SourceRefs: []string{}, WorkItems: []string{}, Workflow: newWorkflowMetadata()}
	block := metadataBlock(lines)
	for i := 0; i < len(block); i++ {
		key, value, ok := parseMetadataLine(block[i])
		if !ok {
			continue
		}
		switch key {
		case "id":
			metadata.Workflow.setScalar(key, value)
			metadata.ID = value
		case "status":
			metadata.Workflow.setScalar(key, value)
			metadata.Status = value
		case "date":
			metadata.Workflow.setScalar(key, value)
		case "source_refs":
			values, emptyItems := metadataListValue(block, i, value)
			metadata.Workflow.setList(key, value, emptyItems)
			metadata.SourceRefs = values
		case "work_items":
			values, emptyItems := metadataListValue(block, i, value)
			metadata.Workflow.setList(key, value, emptyItems)
			metadata.WorkItems = values
		case "subdomain":
			metadata.Subdomain = value
		}
	}
	return metadata
}

func parseWorkItemMetadata(lines []string) workItemMetadata {
	metadata := workItemMetadata{ImpactRefs: []string{}, Tasks: []string{}, Workflow: newWorkflowMetadata()}
	block := metadataBlock(lines)
	for i := 0; i < len(block); i++ {
		key, value, ok := parseMetadataLine(block[i])
		if !ok {
			continue
		}
		switch key {
		case "id":
			metadata.Workflow.setScalar(key, value)
			metadata.ID = value
		case "status":
			metadata.Workflow.setScalar(key, value)
			metadata.Status = value
		case "date":
			metadata.Workflow.setScalar(key, value)
		case "source_requirement":
			metadata.Workflow.setScalar(key, value)
			metadata.SourceRequirement = value
		case "impact_refs":
			values, emptyItems := metadataListValue(block, i, value)
			metadata.Workflow.setList(key, value, emptyItems)
			metadata.ImpactRefs = values
		case "tasks":
			values, emptyItems := metadataListValue(block, i, value)
			metadata.Workflow.setList(key, value, emptyItems)
			metadata.Tasks = values
		case "subdomain":
			metadata.Subdomain = value
		}
	}
	return metadata
}

func parseTaskMetadata(lines []string) taskMetadata {
	metadata := taskMetadata{DependsOn: []string{}, Outputs: []string{}, Workflow: newWorkflowMetadata()}
	block := metadataBlock(lines)
	for i := 0; i < len(block); i++ {
		key, value, ok := parseMetadataLine(block[i])
		if !ok {
			continue
		}
		switch key {
		case "id":
			metadata.Workflow.setScalar(key, value)
			metadata.ID = value
		case "status":
			metadata.Workflow.setScalar(key, value)
			metadata.Status = value
		case "date":
			metadata.Workflow.setScalar(key, value)
		case "work_item":
			metadata.Workflow.setScalar(key, value)
			metadata.WorkItem = value
		case "source_requirement":
			metadata.Workflow.setScalar(key, value)
			metadata.SourceRequirement = value
		case "estimate":
			metadata.Workflow.setScalar(key, value)
			metadata.Estimate = value
		case "depends_on":
			values, emptyItems := metadataListValue(block, i, value)
			metadata.Workflow.setList(key, value, emptyItems)
			metadata.DependsOn = values
		case "outputs":
			values, emptyItems := metadataListValue(block, i, value)
			metadata.Workflow.setList(key, value, emptyItems)
			metadata.Outputs = values
		case "subdomain":
			metadata.Subdomain = value
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

func optionalString(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func newWorkflowMetadata() WorkflowMetadata {
	return WorkflowMetadata{Fields: map[string]WorkflowMetadataField{}}
}

func (m WorkflowMetadata) setScalar(key, value string) {
	m.setField(key, WorkflowMetadataField{Present: true, Value: value})
}

func (m WorkflowMetadata) setList(key, value string, emptyItems []string) {
	m.setField(key, WorkflowMetadataField{Present: true, Value: value, EmptyItems: emptyItems})
}

func (m WorkflowMetadata) setField(key string, field WorkflowMetadataField) {
	if m.Fields == nil {
		return
	}
	m.Fields[key] = field
}

func metadataListValue(block []string, index int, value string) ([]string, []string) {
	if strings.TrimSpace(value) != "" {
		return splitCommaListWithEmptyItems(value)
	}
	return collectIndentedListWithEmptyItems(block, index)
}

func collectIndentedList(block []string, index int) []string {
	values, _ := collectIndentedListWithEmptyItems(block, index)
	return values
}

func collectIndentedListWithEmptyItems(block []string, index int) ([]string, []string) {
	out := []string{}
	emptyItems := []string{}
	for _, line := range block[index+1:] {
		raw := trimLineEnd(line)
		if metadataPattern.MatchString(raw) {
			break
		}
		trimmedLeft := strings.TrimLeft(raw, " \t")
		if len(raw) == len(trimmedLeft) {
			continue
		}
		if strings.TrimSpace(trimmedLeft) == "-" {
			emptyItems = append(emptyItems, "")
			continue
		}
		if !strings.HasPrefix(trimmedLeft, "- ") {
			continue
		}
		item := strings.TrimSpace(strings.TrimPrefix(trimmedLeft, "- "))
		if item != "" {
			out = append(out, item)
		} else {
			emptyItems = append(emptyItems, "")
		}
	}
	return out, emptyItems
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

// filenameNumber extracts the 3-digit ADR number from the filename,
// stripping ns first. "V01-ADR-076-slug.md" with ns="V01-" → "076".
func filenameNumber(path, ns string) string {
	base := filepath.Base(path)
	base = stripNamespacePrefix(base, ns)
	// Try "ADR-NNN" form (post-ns-strip)
	if m := adrFilenamePattern.FindStringSubmatch(base); m != nil {
		return m[1]
	}
	// Fall back to legacy "NNN-slug.md" form
	match := filenameNumPattern.FindStringSubmatch(base)
	if match == nil {
		return ""
	}
	return match[1]
}

// investigationFilenameID extracts the bare INV ID from the filename,
// stripping ns first. "V01-INV-MCP-001-slug.md" with ns="V01-" → "INV-MCP-001".
func investigationFilenameID(path, ns string) string {
	base := filepath.Base(path)
	base = stripNamespacePrefix(base, ns)
	match := investigationFilenameIDPattern.FindStringSubmatch(base)
	if match == nil {
		return ""
	}
	return match[1]
}

// workflowFilenameID extracts the bare workflow ID from the filename,
// stripping ns first. "V01-WORK-DRMCP-001-slug.md" with ns="V01-" → "WORK-DRMCP-001".
func workflowFilenameID(path string, kind RecordKind, ns string) string {
	base := filepath.Base(path)
	base = strings.TrimSuffix(base, filepath.Ext(base))
	base = stripNamespacePrefix(base, ns)
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

func splitCommaListWithEmptyItems(value string) ([]string, []string) {
	value = strings.TrimSpace(value)
	if value == "" || value == "[]" {
		return []string{}, nil
	}
	parts := strings.Split(value, ",")
	out := make([]string, 0, len(parts))
	emptyItems := []string{}
	for _, part := range parts {
		item := strings.TrimSpace(part)
		if item != "" {
			out = append(out, item)
		} else {
			emptyItems = append(emptyItems, "")
		}
	}
	return out, emptyItems
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

// ── Current-format parsers ────────────────────────────────────────────────────
//
// These functions implement current sequential and spec record parsers.
// All existing parser functions are preserved unchanged for authoring compatibility.

var (
	// currentADRH1Pattern matches the post-ns-strip ADR ID form in H1.
	// Accepts "ADR-SPEC-901" (from "# PRODUCT-ADR-SPEC-901: title" after stripping "PRODUCT-").
	currentADRH1Pattern       = regexp.MustCompile(`^#\s+(ADR-[A-Z][A-Z0-9]*-\d{3}):\s+(.+?)\s*$`)
	currentADRFilenamePattern = regexp.MustCompile(`^(ADR-[A-Z][A-Z0-9]*-\d{3})(?:-|\.md$)`)
)

type currentSpecMetadata struct {
	ID     string
	Status string
	Date   string
	Parent string
}

type currentADRMetadata struct {
	Status         string
	DependsOn      []string
	Supersedes     []string
	MigratedToSpec *string
}

// parseCurrentADRRecord parses a current-format ADR.
// The H1 must include the full canonical ID with the app namespace prefix.
// No identity repair: an H1 without the ns prefix is rejected.
func parseCurrentADRRecord(path, raw, ns string) (*Record, RecordCandidate, []ParseIssue) {
	lines := splitMarkdownLines(raw)
	h1Line := firstLine(lines)
	candidate := RecordCandidate{
		Path:     path,
		Kind:     RecordKindDecision,
		H1Line:   h1Line,
		Included: true,
	}
	var issues []ParseIssue

	bareID, title, h1OK := parseCurrentADRH1(h1Line, ns)
	candidate.H1Valid = h1OK
	candidate.H1Number = bareID
	if !h1OK {
		candidate.Included = false
		candidate.SkipReason = "invalid_adr_h1"
		issues = append(issues, ParseIssue{
			Category: DiagnosticInvalidH1Title,
			Path:     path,
			Message:  "ADR H1 is missing or does not contain a valid current canonical ID",
			Details:  map[string]string{"h1": h1Line},
		})
		return nil, candidate, issues
	}

	publicID := ns + bareID
	candidate.ID = publicID
	candidate.NormalizedID = normalizeRecordID(publicID)

	filenameID := currentADRFilenameID(path, ns)
	candidate.FilenameNumber = filenameID
	if filenameID != bareID {
		candidate.FilenameIDMismatch = true
		issues = append(issues, ParseIssue{
			Category: DiagnosticFilenameIDMismatch,
			Path:     path,
			RecordID: publicID,
			Message:  "ADR filename ID does not match H1 ID",
			Details: map[string]string{
				"h1_id":       bareID,
				"filename_id": filenameID,
			},
		})
	}

	metadata, metadataIssues := parseCurrentADRMetadata(lines, path, publicID)
	issues = append(issues, metadataIssues...)
	record := &Record{
		ID:     publicID,
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
		NormalizedID: normalizeRecordID(publicID),
	}
	return record, candidate, issues
}

// parseCurrentADRH1 extracts the bare ADR ID (without ns) and title from an H1 line.
// The ns prefix must be explicitly present; no repair is performed.
func parseCurrentADRH1(line, ns string) (string, string, bool) {
	line = trimLineEnd(line)
	if !strings.HasPrefix(line, "# ") {
		return "", "", false
	}
	content := line[2:]
	if ns != "" && !strings.HasPrefix(content, ns) {
		return "", "", false
	}
	bareContent := content[len(ns):]
	match := currentADRH1Pattern.FindStringSubmatch("# " + bareContent)
	if match == nil {
		return "", "", false
	}
	title := strings.TrimSpace(match[2])
	if title == "" {
		return "", "", false
	}
	return match[1], title, true
}

// currentADRFilenameID extracts the bare current ADR ID from a filename after stripping ns.
func currentADRFilenameID(path, ns string) string {
	base := filepath.Base(path)
	base = stripNamespacePrefix(base, ns)
	m := currentADRFilenamePattern.FindStringSubmatch(base)
	if m != nil {
		return m[1]
	}
	return ""
}

// parseCurrentADRMetadata reads H1-adjacent ADR metadata.
// Handles current conventions: "null" → nil for migrated_to_spec, "[]" → empty for list fields.
func parseCurrentADRMetadata(lines []string, path, recordID string) (currentADRMetadata, []ParseIssue) {
	metadata := currentADRMetadata{
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
			metadata.DependsOn = parseCurrentListField(value)
		case "supersedes":
			metadata.Supersedes = parseCurrentListField(value)
		case "migrated_to_spec":
			if value == "" || value == "null" {
				metadata.MigratedToSpec = nil
				continue
			}
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

// parseCurrentListField parses a current-format list metadata value.
// "[]" and empty string both yield an empty slice; otherwise comma-split.
func parseCurrentListField(value string) []string {
	value = strings.TrimSpace(value)
	if value == "" || value == "[]" {
		return []string{}
	}
	return splitCommaList(value)
}

// parseCurrentInvestigationRecord parses a current-format investigation record.
// The H1 must include the full canonical ID with ns prefix; no repair is performed.
func parseCurrentInvestigationRecord(path, raw, ns string) (*Record, RecordCandidate, []ParseIssue) {
	lines := splitMarkdownLines(raw)
	h1Line := firstLine(lines)
	candidate := RecordCandidate{
		Path:     path,
		Kind:     RecordKindInvestigation,
		H1Line:   h1Line,
		Included: true,
	}
	var issues []ParseIssue

	bareID, title, h1OK := parseCurrentInvestigationH1(h1Line, ns)
	candidate.H1Valid = h1OK
	candidate.H1Number = bareID
	if !h1OK {
		candidate.Included = false
		candidate.SkipReason = "invalid_investigation_h1"
		issues = append(issues, ParseIssue{
			Category: DiagnosticInvalidH1Title,
			Path:     path,
			Message:  "investigation H1 is missing or does not contain a valid current canonical ID",
			Details:  map[string]string{"h1": h1Line},
		})
		return nil, candidate, issues
	}

	publicID := ns + bareID
	candidate.ID = publicID
	candidate.NormalizedID = normalizeRecordID(publicID)

	filenameID := investigationFilenameID(path, ns)
	candidate.FilenameNumber = filenameID
	if filenameID != bareID {
		candidate.FilenameIDMismatch = true
		issues = append(issues, ParseIssue{
			Category: DiagnosticFilenameIDMismatch,
			Path:     path,
			RecordID: publicID,
			Message:  "investigation filename ID does not match H1 ID",
			Details: map[string]string{
				"h1_id":       bareID,
				"filename_id": filenameID,
			},
		})
	}

	metadata := parseInvestigationMetadata(lines)
	record := &Record{
		ID:     publicID,
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
		NormalizedID: normalizeRecordID(publicID),
	}
	return record, candidate, issues
}

// parseCurrentInvestigationH1 extracts the bare INV ID and title from an H1 line.
// The ns prefix must be explicitly present; no repair is performed.
func parseCurrentInvestigationH1(line, ns string) (string, string, bool) {
	line = trimLineEnd(line)
	if !strings.HasPrefix(line, "# ") {
		return "", "", false
	}
	content := line[2:]
	if ns != "" && !strings.HasPrefix(content, ns) {
		return "", "", false
	}
	bareContent := content[len(ns):]
	match := investigationH1Pattern.FindStringSubmatch("# " + bareContent)
	if match == nil {
		return "", "", false
	}
	title := strings.TrimSpace(match[2])
	if title == "" {
		return "", "", false
	}
	return match[1], title, true
}

// parseCurrentRequirementRecord parses a current-format requirement record.
func parseCurrentRequirementRecord(path, raw, ns string) (*Record, RecordCandidate, []ParseIssue) {
	return parseCurrentWorkflowRecord(path, raw, splitMarkdownLines(raw), RecordKindRequirement, ns)
}

// parseCurrentWorkItemRecord parses a current-format work-item record.
func parseCurrentWorkItemRecord(path, raw, ns string) (*Record, RecordCandidate, []ParseIssue) {
	return parseCurrentWorkflowRecord(path, raw, splitMarkdownLines(raw), RecordKindWorkItem, ns)
}

// parseCurrentTaskRecord parses a current-format task record.
func parseCurrentTaskRecord(path, raw, ns string) (*Record, RecordCandidate, []ParseIssue) {
	return parseCurrentWorkflowRecord(path, raw, splitMarkdownLines(raw), RecordKindTask, ns)
}

// parseCurrentWorkflowRecord parses a current-format workflow record (REQ/WORK/TASK).
// The H1 must include the full canonical ID with ns prefix; no repair is performed.
func parseCurrentWorkflowRecord(path, raw string, lines []string, kind RecordKind, ns string) (*Record, RecordCandidate, []ParseIssue) {
	h1Line := firstLine(lines)
	filenameBareID := workflowFilenameID(path, kind, ns)
	candidate := RecordCandidate{
		Path:           path,
		Kind:           kind,
		H1Line:         h1Line,
		FilenameNumber: filenameBareID,
		Included:       true,
	}
	var issues []ParseIssue

	h1BareID, title, h1FormOK := parseCurrentWorkflowH1(h1Line, ns)
	candidate.H1Valid = h1FormOK
	candidate.H1Number = h1BareID
	if !h1FormOK {
		candidate.Included = false
		candidate.SkipReason = "invalid_workflow_h1"
		issues = append(issues, ParseIssue{
			Category: DiagnosticInvalidH1Title,
			Path:     path,
			Message:  "workflow H1 is missing or does not contain a valid current canonical ID",
			Details:  map[string]string{"h1": h1Line},
		})
		return nil, candidate, issues
	}

	metadataIDRaw, status := "", ""
	var workflowMeta *WorkflowMetadata
	var requirement *RequirementDetail
	var workItem *WorkItemDetail
	var task *TaskDetail
	switch kind {
	case RecordKindRequirement:
		metadata := parseRequirementMetadata(lines)
		metadataIDRaw = metadata.ID
		status = metadata.Status
		workflowMeta = &metadata.Workflow
		requirement = &RequirementDetail{
			SourceRefs: metadata.SourceRefs,
			WorkItems:  metadata.WorkItems,
			Subdomain:  optionalString(metadata.Subdomain),
		}
	case RecordKindWorkItem:
		metadata := parseWorkItemMetadata(lines)
		metadataIDRaw = metadata.ID
		status = metadata.Status
		workflowMeta = &metadata.Workflow
		workItem = &WorkItemDetail{
			SourceRequirement: metadata.SourceRequirement,
			ImpactRefs:        metadata.ImpactRefs,
			Tasks:             metadata.Tasks,
			Subdomain:         optionalString(metadata.Subdomain),
		}
	case RecordKindTask:
		metadata := parseTaskMetadata(lines)
		metadataIDRaw = metadata.ID
		status = metadata.Status
		workflowMeta = &metadata.Workflow
		task = &TaskDetail{
			WorkItem:          metadata.WorkItem,
			SourceRequirement: metadata.SourceRequirement,
			Estimate:          metadata.Estimate,
			DependsOn:         metadata.DependsOn,
			Outputs:           metadata.Outputs,
			Subdomain:         optionalString(metadata.Subdomain),
		}
	}

	metadataBareID := stripNamespacePrefix(metadataIDRaw, ns)

	if !validWorkflowIDForKind(h1BareID, kind) || (metadataBareID != "" && !validWorkflowIDForKind(metadataBareID, kind)) || (filenameBareID != "" && !validWorkflowIDForKind(filenameBareID, kind)) {
		candidate.Included = false
		candidate.SkipReason = "invalid_workflow_id"
		issues = append(issues, ParseIssue{
			Category: DiagnosticInvalidWorkflowID,
			Path:     path,
			RecordID: ns + h1BareID,
			Message:  "workflow ID does not match the required grammar",
			Details: map[string]string{
				"h1_id":       h1BareID,
				"metadata_id": metadataBareID,
				"filename_id": filenameBareID,
			},
		})
		return nil, candidate, issues
	}

	publicID := ns + h1BareID
	candidate.ID = publicID
	candidate.NormalizedID = normalizeRecordID(publicID)
	if metadataBareID == "" || metadataBareID != h1BareID || filenameBareID == "" || filenameBareID != h1BareID {
		candidate.FilenameIDMismatch = true
		issues = append(issues, ParseIssue{
			Category: DiagnosticFilenameIDMismatch,
			Path:     path,
			RecordID: publicID,
			Message:  "workflow metadata ID, H1 ID, and filename ID must match",
			Details: map[string]string{
				"h1_id":       h1BareID,
				"metadata_id": metadataBareID,
				"filename_id": filenameBareID,
			},
		})
	}

	record := &Record{
		ID:           publicID,
		Kind:         kind,
		Title:        title,
		Status:       RecordStatus(status),
		Path:         path,
		Requirement:  requirement,
		WorkItem:     workItem,
		Task:         task,
		Headings:     extractHeadings(raw),
		RawBody:      raw,
		NormalizedID: normalizeRecordID(publicID),
		WorkflowMeta: workflowMeta,
	}
	return record, candidate, issues
}

// parseCurrentWorkflowH1 extracts the bare workflow ID and title from an H1 line.
// The ns prefix must be explicitly present; no repair is performed.
func parseCurrentWorkflowH1(line, ns string) (string, string, bool) {
	line = trimLineEnd(line)
	if !strings.HasPrefix(line, "# ") {
		return "", "", false
	}
	content := line[2:]
	if ns != "" && !strings.HasPrefix(content, ns) {
		return "", "", false
	}
	bareContent := content[len(ns):]
	rawID, title, ok := strings.Cut(bareContent, ": ")
	if !ok {
		return "", "", false
	}
	rawID = strings.TrimSpace(rawID)
	title = strings.TrimSpace(title)
	if rawID == "" || title == "" {
		return "", "", false
	}
	if !strings.HasPrefix(rawID, "REQ-") && !strings.HasPrefix(rawID, "WORK-") && !strings.HasPrefix(rawID, "TASK-") {
		return "", "", false
	}
	return rawID, title, true
}

// parseCurrentSpecRecord parses a current-format spec record using H1-adjacent metadata.
// YAML front matter files are rejected (case R06). The canonical ref is path-derived and
// cross-checked against the metadata id field. Invalid sources with missing required
// metadata are retained as path-addressable (case C15).
// recordsRoot is the ancestor records directory path (same base as path).
// appNamespace is the lowercase app namespace identifier (e.g., "product").
func parseCurrentSpecRecord(path, raw, recordsRoot, appNamespace string) (*Record, RecordCandidate, []ParseIssue) {
	lines := splitMarkdownLines(raw)
	derivedRef := deriveSpecRef(path, recordsRoot, appNamespace)

	candidate := RecordCandidate{
		Path:     path,
		Kind:     RecordKindSpec,
		Included: true,
	}
	if derivedRef != "" {
		candidate.ID = derivedRef
		candidate.NormalizedID = normalizeRecordID(derivedRef)
	}

	var issues []ParseIssue

	if hasOpeningFrontMatter(lines) {
		candidate.H1Valid = false
		candidate.Included = false
		candidate.SkipReason = "yaml_front_matter_current_spec"
		issues = append(issues, ParseIssue{
			Category: DiagnosticInvalidH1Title,
			Path:     path,
			RecordID: derivedRef,
			Message:  "current spec must not use YAML front matter",
		})
		return nil, candidate, issues
	}

	h1Line := firstLine(lines)
	candidate.H1Line = h1Line
	title, _, h1OK := parseSpecH1(raw)
	candidate.H1Valid = h1OK
	if !h1OK {
		issues = append(issues, ParseIssue{
			Category: DiagnosticInvalidH1Title,
			Path:     path,
			RecordID: derivedRef,
			Message:  "current spec H1 is missing or invalid",
			Details:  map[string]string{"h1": h1Line},
		})
	}

	specMeta := parseCurrentSpecMetadata(lines)

	if specMeta.ID != "" && derivedRef != "" && specMeta.ID != derivedRef {
		issues = append(issues, ParseIssue{
			Category: DiagnosticFilenameIDMismatch,
			Path:     path,
			RecordID: derivedRef,
			Message:  "current spec metadata id does not match path-derived ref",
			Details: map[string]string{
				"path_ref":    derivedRef,
				"metadata_id": specMeta.ID,
			},
		})
	}

	if specMeta.Parent == "" {
		issues = append(issues, ParseIssue{
			Category: DiagnosticMissingRequiredMetadata,
			Path:     path,
			RecordID: derivedRef,
			Message:  "current spec is missing required 'parent' metadata",
		})
	}

	recordID := derivedRef
	if recordID == "" {
		recordID = specMeta.ID
	}

	record := &Record{
		ID:           recordID,
		Kind:         RecordKindSpec,
		Title:        title,
		Status:       RecordStatus(specMeta.Status),
		Date:         specMeta.Date,
		Path:         path,
		Spec:         &SpecDetail{DependsOn: []string{}},
		Headings:     extractHeadings(raw),
		RawBody:      raw,
		NormalizedID: normalizeRecordID(recordID),
	}
	return record, candidate, issues
}

// parseCurrentSpecMetadata reads H1-adjacent spec metadata fields.
// Values for id and parent may use backtick quoting.
func parseCurrentSpecMetadata(lines []string) currentSpecMetadata {
	var meta currentSpecMetadata
	for _, line := range metadataBlock(lines) {
		key, value, ok := parseMetadataLine(line)
		if !ok {
			continue
		}
		switch key {
		case "id":
			meta.ID = stripBackticks(value)
		case "status":
			meta.Status = value
		case "date":
			meta.Date = value
		case "parent":
			meta.Parent = stripBackticks(value)
		}
	}
	return meta
}

// stripBackticks removes surrounding backtick quoting from a metadata value.
func stripBackticks(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 2 && s[0] == '`' && s[len(s)-1] == '`' {
		return s[1 : len(s)-1]
	}
	return s
}

// deriveSpecRef derives the canonical spec ref from a current spec source path.
// path and recordsRoot must share the same base (both relative or both absolute).
// Hyphens in path segments are converted to underscores; "index" as the final
// segment collapses to the parent level.
func deriveSpecRef(path, recordsRoot, appNamespace string) string {
	path = filepath.ToSlash(path)
	recordsRoot = filepath.ToSlash(recordsRoot)
	if !strings.HasSuffix(recordsRoot, "/") {
		recordsRoot += "/"
	}
	if !strings.HasPrefix(path, recordsRoot) {
		return ""
	}
	relPath := path[len(recordsRoot):]
	if !strings.HasPrefix(relPath, "spec/") {
		return ""
	}
	specRelPath := relPath[len("spec/"):]
	if strings.HasSuffix(specRelPath, ".md") {
		specRelPath = specRelPath[:len(specRelPath)-3]
	}
	segments := strings.Split(specRelPath, "/")
	if len(segments) > 0 && segments[len(segments)-1] == "index" {
		segments = segments[:len(segments)-1]
	}
	if len(segments) == 0 {
		return ""
	}
	for i, seg := range segments {
		segments[i] = strings.ReplaceAll(seg, "-", "_")
	}
	return "spec:" + appNamespace + "." + strings.Join(segments, ".")
}
