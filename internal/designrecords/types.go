package designrecords

type RecordKind string

const (
	RecordKindDecision RecordKind = "decision"
	RecordKindSpec     RecordKind = "spec"
)

type RecordStatus string

const (
	RecordStatusProposed   RecordStatus = "proposed"
	RecordStatusAccepted   RecordStatus = "accepted"
	RecordStatusSuperseded RecordStatus = "superseded"
	RecordStatusConfirmed  RecordStatus = "confirmed"
	RecordStatusDraft      RecordStatus = "draft"
	RecordStatusWIP        RecordStatus = "wip"
)

type Heading struct {
	Level int    `json:"level"`
	Text  string `json:"text"`
}

type Record struct {
	ID             string       `json:"id"`
	Kind           RecordKind   `json:"kind"`
	Title          string       `json:"title"`
	Status         RecordStatus `json:"status"`
	Path           string       `json:"path"`
	DependsOn      []string     `json:"depends_on"`
	Supersedes     []string     `json:"supersedes"`
	MigratedToSpec *string      `json:"migrated_to_spec"`
	Headings       []Heading    `json:"headings,omitempty"`
	Body           *string      `json:"body,omitempty"`
	RawBody        string       `json:"-"`
	NormalizedID   string       `json:"-"`
}

type Index struct {
	Root        string            `json:"root"`
	Records     []Record          `json:"records"`
	Diagnostics []Diagnostic      `json:"diagnostics,omitempty"`
	Candidates  []RecordCandidate `json:"-"`
	ParseIssues []ParseIssue      `json:"-"`
	PathIssues  []PathIssue       `json:"-"`
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
	DiagnosticDuplicateID             DiagnosticCategory = "duplicate_id"
	DiagnosticFilenameIDMismatch      DiagnosticCategory = "filename_id_mismatch"
	DiagnosticInvalidH1Title          DiagnosticCategory = "invalid_h1_title"
	DiagnosticInvalidStatusForKind    DiagnosticCategory = "invalid_status_for_kind"
	DiagnosticSpecStatusMismatch      DiagnosticCategory = "spec_status_mismatch"
	DiagnosticMissingDependsOnTarget  DiagnosticCategory = "missing_depends_on_target"
	DiagnosticMissingSupersedesTarget DiagnosticCategory = "missing_supersedes_target"
	DiagnosticInvalidMigratedToSpec   DiagnosticCategory = "invalid_migrated_to_spec"
	DiagnosticMissingRecordPath       DiagnosticCategory = "missing_record_path"
)

type DiagnosticSeverity string

const (
	DiagnosticSeverityError DiagnosticSeverity = "error"
)

type Diagnostic struct {
	Category DiagnosticCategory `json:"category"`
	Severity DiagnosticSeverity `json:"severity"`
	RecordID string             `json:"record_id,omitempty"`
	Path     string             `json:"path,omitempty"`
	Message  string             `json:"message"`
	TargetID string             `json:"target_id,omitempty"`
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
	ID             string       `json:"id"`
	Kind           RecordKind   `json:"kind"`
	Title          string       `json:"title"`
	Status         RecordStatus `json:"status"`
	Path           string       `json:"path"`
	DependsOn      []string     `json:"depends_on"`
	Supersedes     []string     `json:"supersedes"`
	MigratedToSpec *string      `json:"migrated_to_spec"`
}

type ListRecordsResponse struct {
	Records []ListedRecord `json:"records"`
}

type GetRecordRequest struct {
	ID          string `json:"id"`
	IncludeBody bool   `json:"include_body,omitempty"`
}

type GetRecordResponse struct {
	Record Record `json:"record"`
}

type ValidateRecordsRequest struct {
	Kind    RecordKind `json:"kind,omitempty"`
	IDRange *IDRange   `json:"id_range,omitempty"`
}

type ValidateRecordsResponse struct {
	OK          bool         `json:"ok"`
	Diagnostics []Diagnostic `json:"diagnostics"`
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
