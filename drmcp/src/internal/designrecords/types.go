package designrecords

import "encoding/json"

type RecordKind string

const (
	RecordKindDecision      RecordKind = "decision"
	RecordKindSpec          RecordKind = "spec"
	RecordKindInvestigation RecordKind = "investigation"
	RecordKindRequirement   RecordKind = "requirement"
	RecordKindWorkItem      RecordKind = "work_item"
	RecordKindTask          RecordKind = "task"
)

type RecordStatus string

const (
	RecordStatusProposed       RecordStatus = "proposed"
	RecordStatusAccepted       RecordStatus = "accepted"
	RecordStatusSuperseded     RecordStatus = "superseded"
	RecordStatusConfirmed      RecordStatus = "confirmed"
	RecordStatusDraft          RecordStatus = "draft"
	RecordStatusWIP            RecordStatus = "wip"
	RecordStatusInvestigating  RecordStatus = "investigating"
	RecordStatusConcluded      RecordStatus = "concluded"
	RecordStatusCaptured       RecordStatus = "captured"
	RecordStatusDecisionNeeded RecordStatus = "decision_needed"
	RecordStatusDeferred       RecordStatus = "deferred"
	RecordStatusRejected       RecordStatus = "rejected"
	RecordStatusNotStarted     RecordStatus = "not_started"
	RecordStatusInProgress     RecordStatus = "in_progress"
	RecordStatusBlocked        RecordStatus = "blocked"
	RecordStatusDone           RecordStatus = "done"
)

type Heading struct {
	Level int    `json:"level"`
	Text  string `json:"text"`
}

type CandidateHeading struct {
	Heading string `json:"heading"`
	Level   int    `json:"level"`
	Ordinal int    `json:"ordinal"`
}

type Record struct {
	ID            string               `json:"id"`
	Kind          RecordKind           `json:"kind"`
	Title         string               `json:"title"`
	Status        RecordStatus         `json:"status"`
	Path          string               `json:"path"`
	Decision      *DecisionDetail      `json:"decision,omitempty"`
	Spec          *SpecDetail          `json:"spec,omitempty"`
	Investigation *InvestigationDetail `json:"investigation,omitempty"`
	Requirement   *RequirementDetail   `json:"requirement,omitempty"`
	WorkItem      *WorkItemDetail      `json:"work_item,omitempty"`
	Task          *TaskDetail          `json:"task,omitempty"`
	SemanticRefs  []SemanticRefDecl    `json:"-"`
	Headings      []Heading            `json:"headings"`
	Body          *string              `json:"body,omitempty"`
	RawBody       string               `json:"-"`
	NormalizedID  string               `json:"-"`
	WorkflowMeta  *WorkflowMetadata    `json:"-"`
}

type DecisionDetail struct {
	DependsOn      []string `json:"depends_on"`
	Supersedes     []string `json:"supersedes"`
	MigratedToSpec *string  `json:"migrated_to_spec"`
}

type SpecDetail struct {
	DependsOn []string `json:"depends_on"`
}

type InvestigationDetail struct {
	Trigger               string   `json:"trigger"`
	Scope                 string   `json:"scope"`
	NonScope              string   `json:"non_scope"`
	SourceRefs            []string `json:"source_refs"`
	FollowUpCandidates    []string `json:"follow_up_candidates"`
	Supersedes            []string `json:"supersedes,omitempty"`
	RelatedRequirements   []string `json:"related_requirements,omitempty"`
	RelatedWorkItems      []string `json:"related_work_items,omitempty"`
	RelatedADRs           []string `json:"related_adrs,omitempty"`
	RelatedSpecs          []string `json:"related_specs,omitempty"`
	RelatedInternalDesign []string `json:"related_internal_design,omitempty"`
	RelatedCoverage       []string `json:"related_coverage,omitempty"`
	FollowUpResults       []string `json:"follow_up_results,omitempty"`
}

type RequirementDetail struct {
	SourceRefs []string `json:"source_refs"`
	WorkItems  []string `json:"work_items"`
	Subdomain  *string  `json:"subdomain,omitempty"`
}

type WorkItemDetail struct {
	SourceRequirement string   `json:"source_requirement"`
	ImpactRefs        []string `json:"impact_refs"`
	Tasks             []string `json:"tasks"`
	Subdomain         *string  `json:"subdomain,omitempty"`
}

type TaskDetail struct {
	WorkItem          string   `json:"work_item"`
	SourceRequirement string   `json:"source_requirement"`
	Estimate          string   `json:"estimate"`
	DependsOn         []string `json:"depends_on"`
	Outputs           []string `json:"outputs"`
	Subdomain         *string  `json:"subdomain,omitempty"`
}

type WorkflowMetadata struct {
	Fields map[string]WorkflowMetadataField
}

type WorkflowMetadataField struct {
	Present    bool
	Value      string
	EmptyItems []string
}

type SemanticTargetType string

const (
	SemanticTargetDocument SemanticTargetType = "document"
	SemanticTargetSection  SemanticTargetType = "section"
)

type SemanticRefDecl struct {
	Ref        string
	Path       string
	TargetType SemanticTargetType
	Section    string
}

type SemanticRefSource struct {
	Path     string
	RecordID string
	Decls    []SemanticRefDecl
	Headings []Heading
}

type Index struct {
	Root               string              `json:"root"`
	NamespacePrefix    string              `json:"-"`
	RecordsRoot        string              `json:"-"`
	Records            []Record            `json:"records"`
	Diagnostics        []Diagnostic        `json:"diagnostics,omitempty"`
	Candidates         []RecordCandidate   `json:"-"`
	ParseIssues        []ParseIssue        `json:"-"`
	PathIssues         []PathIssue         `json:"-"`
	SemanticRefs       []SemanticRefDecl   `json:"-"`
	SemanticRefSources []SemanticRefSource `json:"-"`
}

type RecordCandidate struct {
	Path               string
	Kind               RecordKind
	ID                 string
	NormalizedID       string
	H1Line             string
	H1Valid            bool
	H1Number           string
	FilenameNumber     string
	FilenameIDMismatch bool
	Included           bool
	SkipReason         string
}

type ParseIssue struct {
	Category DiagnosticCategory
	Path     string
	RecordID string
	Message  string
	Details  map[string]string
}

type PathIssue struct {
	Path      string
	Operation string
	Err       error
}

type DiagnosticCategory string

const (
	DiagnosticDuplicateID                           DiagnosticCategory = "duplicate_id"
	DiagnosticFilenameIDMismatch                    DiagnosticCategory = "filename_id_mismatch"
	DiagnosticInvalidH1Title                        DiagnosticCategory = "invalid_h1_title"
	DiagnosticInvalidWorkflowID                     DiagnosticCategory = "invalid_workflow_id"
	DiagnosticInvalidStatusForKind                  DiagnosticCategory = "invalid_status_for_kind"
	DiagnosticSpecStatusMismatch                    DiagnosticCategory = "spec_status_mismatch"
	DiagnosticMissingDependsOnTarget                DiagnosticCategory = "missing_depends_on_target"
	DiagnosticMissingSupersedesTarget               DiagnosticCategory = "missing_supersedes_target"
	DiagnosticInvalidMigratedToSpec                 DiagnosticCategory = "invalid_migrated_to_spec"
	DiagnosticMissingRecordPath                     DiagnosticCategory = "missing_record_path"
	DiagnosticInvalidSemanticRefDeclaration         DiagnosticCategory = "invalid_semantic_ref_declaration"
	DiagnosticMissingSectionTarget                  DiagnosticCategory = "missing_section_target"
	DiagnosticAmbiguousSectionTarget                DiagnosticCategory = "ambiguous_section_target"
	DiagnosticDuplicateSemanticRef                  DiagnosticCategory = "duplicate_semantic_ref"
	DiagnosticUnresolvedSourceRef                   DiagnosticCategory = "unresolved_source_ref"
	DiagnosticUnresolvedFollowUpResult              DiagnosticCategory = "unresolved_follow_up_result"
	DiagnosticUnresolvedFollowUpCandidate           DiagnosticCategory = "unresolved_follow_up_candidate"
	DiagnosticNoncanonicalSourceRef                 DiagnosticCategory = "noncanonical_source_ref"
	DiagnosticNoncanonicalFollowUpResult            DiagnosticCategory = "noncanonical_follow_up_result"
	DiagnosticNoncanonicalFollowUpCandidate         DiagnosticCategory = "noncanonical_follow_up_candidate"
	DiagnosticUnsupportedReference                  DiagnosticCategory = "unsupported_reference"
	DiagnosticUnresolvedReference                   DiagnosticCategory = "unresolved_reference"
	DiagnosticAmbiguousReference                    DiagnosticCategory = "ambiguous_reference"
	DiagnosticUnresolvedWorkflowRelation            DiagnosticCategory = "unresolved_workflow_relation"
	DiagnosticInvalidWorkflowRelationTarget         DiagnosticCategory = "invalid_workflow_relation_target"
	DiagnosticWorkflowRelationMismatch              DiagnosticCategory = "workflow_relation_mismatch"
	DiagnosticWorkflowSourceReqMismatch             DiagnosticCategory = "workflow_source_requirement_mismatch"
	DiagnosticMissingRequiredMetadata               DiagnosticCategory = "missing_required_metadata"
	DiagnosticEmptyRequiredMetadata                 DiagnosticCategory = "empty_required_metadata"
	DiagnosticMissingRequiredSection                DiagnosticCategory = "missing_required_section"
	DiagnosticEmptyRequiredSection                  DiagnosticCategory = "empty_required_section"
	DiagnosticSectionHeadingCaseMismatch            DiagnosticCategory = "section_heading_case_mismatch"
	DiagnosticSectionReplacementBodyHeadingStripped DiagnosticCategory = "section_replacement_body_heading_stripped"
	DiagnosticInvalidMetadataValue                  DiagnosticCategory = "invalid_metadata_value"
	DiagnosticNoOpUpdate                            DiagnosticCategory = "no_op_update"
	DiagnosticRecordNotFound                        DiagnosticCategory = "record_not_found"
	DiagnosticDuplicateRequestedIDIgnored           DiagnosticCategory = "duplicate_requested_id_ignored"
	DiagnosticExactIDSequenceGap                    DiagnosticCategory = "exact_id_sequence_gap"
	DiagnosticMissingRequiredMetadataBatch          DiagnosticCategory = "missing_required_metadata_batch"
	DiagnosticReciprocalFollowUpModeRequired        DiagnosticCategory = "reciprocal_follow_up_mode_required"
	DiagnosticReciprocalUpdateIncluded              DiagnosticCategory = "reciprocal_update_included"
	DiagnosticNewSubdomainValue                     DiagnosticCategory = "new_subdomain_value"
)

type DiagnosticSeverity string

const (
	DiagnosticSeverityError   DiagnosticSeverity = "error"
	DiagnosticSeverityWarning DiagnosticSeverity = "warning"
	DiagnosticSeverityInfo    DiagnosticSeverity = "info"
)

type Diagnostic struct {
	Category          DiagnosticCategory `json:"category"`
	Severity          DiagnosticSeverity `json:"severity"`
	RecordID          string             `json:"record_id,omitempty"`
	Path              string             `json:"path,omitempty"`
	Message           string             `json:"message"`
	TargetID          string             `json:"target_id,omitempty"`
	Field             string             `json:"field,omitempty"`
	Value             string             `json:"value,omitempty"`
	RefStatus         string             `json:"ref_status,omitempty"`
	Section           string             `json:"section,omitempty"`
	Status            string             `json:"status,omitempty"`
	RequestedID       string             `json:"requested_id,omitempty"`
	FirstIndex        *int               `json:"first_index,omitempty"`
	DuplicateIndexes  []int              `json:"duplicate_indexes,omitempty"`
	CandidateHeadings []CandidateHeading `json:"candidate_headings,omitempty"`
	// Authoring repair guidance fields (REQ-MCP-024 / REQ-MCP-028)
	AllowedValues    []string       `json:"allowed_values,omitempty"`
	RequiredFields   []string       `json:"required_fields,omitempty"`
	TargetKind       string         `json:"target_kind,omitempty"`
	RepairSuggestion map[string]any `json:"repair_suggestion,omitempty"`
	// Internal-only fields (not serialized via json tags; handled by MarshalJSON)
	ActualHeading string `json:"-"`
	StrippedHeading string `json:"-"`
	StrippedLevel   int    `json:"-"`
	ValuePresent    bool   `json:"-"`
}

func (d Diagnostic) MarshalJSON() ([]byte, error) {
	type diagnosticJSON struct {
		Category          DiagnosticCategory `json:"category"`
		Severity          DiagnosticSeverity `json:"severity"`
		RecordID          string             `json:"record_id,omitempty"`
		Path              string             `json:"path,omitempty"`
		Message           string             `json:"message"`
		TargetID          string             `json:"target_id,omitempty"`
		Field             string             `json:"field,omitempty"`
		Value             *string            `json:"value,omitempty"`
		RefStatus         string             `json:"ref_status,omitempty"`
		Section           string             `json:"section,omitempty"`
		Status            string             `json:"status,omitempty"`
		ActualHeading     string             `json:"actual_heading,omitempty"`
		StrippedHeading   string             `json:"stripped_heading,omitempty"`
		StrippedLevel     *int               `json:"stripped_level,omitempty"`
		RequestedID       string             `json:"requested_id,omitempty"`
		FirstIndex        *int               `json:"first_index,omitempty"`
		DuplicateIndexes  []int              `json:"duplicate_indexes,omitempty"`
		CandidateHeadings []CandidateHeading `json:"candidate_headings,omitempty"`
		AllowedValues     []string           `json:"allowed_values,omitempty"`
		RequiredFields    []string           `json:"required_fields,omitempty"`
		TargetKind        string             `json:"target_kind,omitempty"`
		RepairSuggestion  map[string]any     `json:"repair_suggestion,omitempty"`
	}
	var value *string
	if d.Value != "" || d.ValuePresent {
		value = &d.Value
	}
	var strippedLevel *int
	if d.StrippedLevel != 0 {
		sl := d.StrippedLevel
		strippedLevel = &sl
	}
	return json.Marshal(diagnosticJSON{
		Category:          d.Category,
		Severity:          d.Severity,
		RecordID:          d.RecordID,
		Path:              d.Path,
		Message:           d.Message,
		TargetID:          d.TargetID,
		Field:             d.Field,
		Value:             value,
		RefStatus:         d.RefStatus,
		Section:           d.Section,
		Status:            d.Status,
		ActualHeading:     d.ActualHeading,
		StrippedHeading:   d.StrippedHeading,
		StrippedLevel:     strippedLevel,
		RequestedID:       d.RequestedID,
		FirstIndex:        d.FirstIndex,
		DuplicateIndexes:  d.DuplicateIndexes,
		CandidateHeadings: d.CandidateHeadings,
		AllowedValues:     d.AllowedValues,
		RequiredFields:    d.RequiredFields,
		TargetKind:        d.TargetKind,
		RepairSuggestion:  d.RepairSuggestion,
	})
}

type ErrorCode string

const (
	ErrorCodeRecordNotFound               ErrorCode = "record_not_found"
	ErrorCodeGuideNotFound                ErrorCode = "guide_not_found"
	ErrorCodeInvalidRequest               ErrorCode = "invalid_request"
	ErrorCodeUnsupportedKind              ErrorCode = "unsupported_kind"
	ErrorCodeInvalidIDRange               ErrorCode = "invalid_id_range"
	ErrorCodeIDRangeRequiresDecisionKind  ErrorCode = "id_range_requires_decision_kind"
	ErrorCodeProposalNotFound             ErrorCode = "proposal_not_found"
	ErrorCodeProposalExpired              ErrorCode = "proposal_expired"
	ErrorCodeProposalDiscarded            ErrorCode = "proposal_discarded"
	ErrorCodeProposalAlreadyAccepted      ErrorCode = "proposal_already_accepted"
	ErrorCodeProposalStale                ErrorCode = "proposal_stale"
	ErrorCodeTargetChanged                ErrorCode = "target_changed"
	ErrorCodeIDCollision                  ErrorCode = "id_collision"
	ErrorCodeRequiredFollowUpNotSatisfied ErrorCode = "required_follow_up_not_satisfied"
	ErrorCodeInvalidBodySource            ErrorCode = "invalid_body_source"
	ErrorCodeBodyCacheNotFound            ErrorCode = "body_cache_not_found"
	ErrorCodeBodyCacheExpired             ErrorCode = "body_cache_expired"
	ErrorCodeProposalPreparationFailed    ErrorCode = "proposal_preparation_failed"
	ErrorCodeSectionSelectorNoMatch              ErrorCode = "section_selector_no_match"
	ErrorCodeSectionSelectorAmbiguous            ErrorCode = "section_selector_ambiguous"
	ErrorCodeConflictingOperations               ErrorCode = "conflicting_operations"
	ErrorCodeMultipleSectionReplaceNotSupported  ErrorCode = "multiple_section_replace_not_supported"
)

type ToolError struct {
	Code    ErrorCode `json:"code"`
	Message string    `json:"message"`
}

func (e *ToolError) Error() string {
	if e == nil {
		return ""
	}
	return e.Message
}

func newToolError(code ErrorCode, message string) *ToolError {
	return &ToolError{Code: code, Message: message}
}

type IDRange struct {
	From string `json:"from,omitempty"`
	To   string `json:"to,omitempty"`
}

type ListRecordsRequest struct {
	Kind    RecordKind   `json:"kind,omitempty"`
	Status  RecordStatus `json:"status,omitempty"`
	ID      string       `json:"id,omitempty"`
	IDRange *IDRange     `json:"id_range,omitempty"`
	OrderBy string       `json:"order_by,omitempty"`
	Order   string       `json:"order,omitempty"`
	Limit   *int         `json:"limit,omitempty"`
}

type ListedRecord struct {
	ID            string               `json:"id"`
	Kind          RecordKind           `json:"kind"`
	Title         string               `json:"title"`
	Status        RecordStatus         `json:"status"`
	Path          string               `json:"path"`
	Decision      *DecisionDetail      `json:"decision,omitempty"`
	Spec          *SpecDetail          `json:"spec,omitempty"`
	Investigation *InvestigationDetail `json:"investigation,omitempty"`
	Requirement   *RequirementDetail   `json:"requirement,omitempty"`
	WorkItem      *WorkItemDetail      `json:"work_item,omitempty"`
	Task          *TaskDetail          `json:"task,omitempty"`
}

type ListRecordsResponse struct {
	Records []ListedRecord `json:"records"`
}

type GetRecordRequest struct {
	ID          string `json:"id"`
	IncludeBody bool   `json:"include_body,omitempty"`
}

type GetRecordRecord struct {
	ID            string               `json:"id"`
	Kind          RecordKind           `json:"kind"`
	Title         string               `json:"title"`
	Status        RecordStatus         `json:"status"`
	Path          string               `json:"path"`
	Decision      *DecisionDetail      `json:"decision,omitempty"`
	Spec          *SpecDetail          `json:"spec,omitempty"`
	Investigation *InvestigationDetail `json:"investigation,omitempty"`
	Requirement   *RequirementDetail   `json:"requirement,omitempty"`
	WorkItem      *WorkItemDetail      `json:"work_item,omitempty"`
	Task          *TaskDetail          `json:"task,omitempty"`
	Headings      []Heading            `json:"headings"`
	Body          *string              `json:"body,omitempty"`
}

type GetRecordResponse struct {
	Record GetRecordRecord `json:"record"`
}

type GetRecordsRequest struct {
	IDs         []string `json:"ids"`
	IncludeBody bool     `json:"include_body,omitempty"`
}

type RetrievalStatus string

const (
	RetrievalStatusFound    RetrievalStatus = "found"
	RetrievalStatusNotFound RetrievalStatus = "not_found"
)

type GetRecordsItem struct {
	ID              string           `json:"id"`
	RetrievalStatus RetrievalStatus  `json:"retrieval_status"`
	Record          *GetRecordRecord `json:"record"`
	Diagnostics     []Diagnostic     `json:"diagnostics"`
}

type GetRecordsResponse struct {
	Items       []GetRecordsItem `json:"items"`
	Diagnostics []Diagnostic     `json:"diagnostics"`
}

type ValidateRecordsRequest struct {
	Kind    RecordKind `json:"kind,omitempty"`
	IDRange *IDRange   `json:"id_range,omitempty"`
}

type ValidateRecordsResponse struct {
	OK          bool         `json:"ok"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type ResolveReferenceRequest struct {
	Ref string `json:"ref"`
}

type ResolveReferenceResponse struct {
	Ref         string          `json:"ref"`
	RefKind     string          `json:"ref_kind"`
	Status      string          `json:"status"`
	Target      *ResolvedTarget `json:"target"`
	Diagnostics []Diagnostic    `json:"diagnostics"`
}

type ResolvedTarget struct {
	TargetType string       `json:"target_type"`
	Path       string       `json:"path"`
	Section    string       `json:"section,omitempty"`
	RecordID   string       `json:"record_id,omitempty"`
	RecordKind RecordKind   `json:"record_kind,omitempty"`
	Title      string       `json:"title,omitempty"`
	Status     RecordStatus `json:"status,omitempty"`
}

type SuggestNextRecordRequest struct {
	Kind  RecordKind `json:"kind"`
	Title string     `json:"title"`
}

type SuggestNextRecordResponse struct {
	Kind          RecordKind `json:"kind"`
	Title         string     `json:"title"`
	NextID        string     `json:"next_id"`
	NextNumber    int        `json:"next_number"`
	SuggestedPath string     `json:"suggested_path"`
	ExistingMaxID string     `json:"existing_max_id,omitempty"`
}

type ListAuthoringGuidesRequest struct{}

type AuthoringGuideSummary struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Abstract string `json:"abstract"`
}

type ListAuthoringGuidesResponse struct {
	Guides []AuthoringGuideSummary `json:"guides"`
}

type GetAuthoringGuidanceRequest struct {
	ID string `json:"id"`
}

type GetAuthoringGuidanceResponse struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}
