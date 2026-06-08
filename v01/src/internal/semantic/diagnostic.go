package semantic

type Severity string

const (
	SeverityError   Severity = "error"
	SeverityWarning Severity = "warning"
)

type Diagnostic struct {
	Severity Severity `json:"severity"`
	Code     string   `json:"code,omitempty"`
	FileID   FileID   `json:"file,omitempty"`
	Message  string   `json:"message"`
}
