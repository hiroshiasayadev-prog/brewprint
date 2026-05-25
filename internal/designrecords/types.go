package designrecords

type RecordKind string

const (
	RecordKindDecision      RecordKind = "decision"
	RecordKindSpec          RecordKind = "spec"
	RecordKindInvestigation RecordKind = "investigation"
)

type RecordStatus string

const (
	RecordStatusProposed      RecordStatus = "proposed"
	RecordStatusAccepted      RecordStatus = "accepted"
	RecordStatusSuperseded    RecordStatus = "superseded"
	RecordStatusConfirmed     RecordStatus = "confirmed"
	RecordStatusDraft         RecordStatus = "draft"
	RecordStatusWIP           RecordStatus = "wip"
	RecordStatusInvestigating RecordStatus = "investigating"
	RecordStatusConcluded     RecordStatus = "concluded"
)

type Heading struct {
	Level int    `json:"level"`
	Text  string `json:"text"`
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
	SemanticRefs  []SemanticRefDecl    `json:"-"`
	Headings      []Heading            `json:"headings"`
	Body          *string              `json:"body,omitempty"`
	RawBody       string               `json:"-"`
	NormalizedID  string               `json:"-"`
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
	DiagnosticDuplicateID                   DiagnosticCategory = "duplicate_id"
	DiagnosticFilenameIDMismatch            DiagnosticCategory = "filename_id_mismatch"
	DiagnosticInvalidH1Title                DiagnosticCategory = "invalid_h1_title"
	DiagnosticInvalidStatusForKind          DiagnosticCategory = "invalid_status_for_kind"
	DiagnosticSpecStatusMismatch            DiagnosticCategory = "spec_status_mismatch"
	DiagnosticMissingDependsOnTarget        DiagnosticCategory = "missing_depends_on_target"
	DiagnosticMissingSupersedesTarget       DiagnosticCategory = "missing_supersedes_target"
	DiagnosticInvalidMigratedToSpec         DiagnosticCategory = "invalid_migrated_to_spec"
	DiagnosticMissingRecordPath             DiagnosticCategory = "missing_record_path"
	DiagnosticInvalidSemanticRefDeclaration DiagnosticCategory = "invalid_semantic_ref_declaration"
	DiagnosticMissingSectionTarget          DiagnosticCategory = "missing_section_target"
	DiagnosticAmbiguousSectionTarget        DiagnosticCategory = "ambiguous_section_target"
	DiagnosticDuplicateSemanticRef          DiagnosticCategory = "duplicate_semantic_ref"
	DiagnosticUnresolvedSourceRef           DiagnosticCategory = "unresolved_source_ref"
	DiagnosticUnresolvedFollowUpResult      DiagnosticCategory = "unresolved_follow_up_result"
	DiagnosticUnresolvedFollowUpCandidate   DiagnosticCategory = "unresolved_follow_up_candidate"
	DiagnosticNoncanonicalSourceRef         DiagnosticCategory = "noncanonical_source_ref"
	DiagnosticNoncanonicalFollowUpResult    DiagnosticCategory = "noncanonical_follow_up_result"
	DiagnosticNoncanonicalFollowUpCandidate DiagnosticCategory = "noncanonical_follow_up_candidate"
	DiagnosticUnsupportedReference          DiagnosticCategory = "unsupported_reference"
	DiagnosticUnresolvedReference           DiagnosticCategory = "unresolved_reference"
	DiagnosticAmbiguousReference            DiagnosticCategory = "ambiguous_reference"
)

type DiagnosticSeverity string

const (
	DiagnosticSeverityError DiagnosticSeverity = "error"
	DiagnosticSeverityInfo  DiagnosticSeverity = "info"
)

type Diagnostic struct {
	Category  DiagnosticCategory `json:"category"`
	Severity  DiagnosticSeverity `json:"severity"`
	RecordID  string             `json:"record_id,omitempty"`
	Path      string             `json:"path,omitempty"`
	Message   string             `json:"message"`
	TargetID  string             `json:"target_id,omitempty"`
	Field     string             `json:"field,omitempty"`
	Value     string             `json:"value,omitempty"`
	RefStatus string             `json:"ref_status,omitempty"`
}

type ErrorCode string

const (
	ErrorCodeRecordNotFound              ErrorCode = "record_not_found"
	ErrorCodeInvalidRequest              ErrorCode = "invalid_request"
	ErrorCodeUnsupportedKind             ErrorCode = "unsupported_kind"
	ErrorCodeIDRangeRequiresDecisionKind ErrorCode = "id_range_requires_decision_kind"
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
	Headings      []Heading            `json:"headings"`
	Body          *string              `json:"body,omitempty"`
}

type GetRecordResponse struct {
	Record GetRecordRecord `json:"record"`
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
