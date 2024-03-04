package educationlsp

type PublishDiagnostics struct {
	Notification
	Params PublishDiagnosticsParams `json:"params"`
}

type PublishDiagnosticsParams struct {
	URI         string       `json:"uri"`
	Diagnostics []Diagnostic `json:"diagnostics"`
}

type Diagnostic struct {
	Range           Range              `json:"range"`
	Severity        DiagnosticSeverity `json:"severity,omitempty"`
	Code            interface{}        `json:"code,omitempty"`
	CodeDescription *CodeDescription   `json:"codeDescription,omitempty"`
	Source          *string            `json:"source,omitempty"`
	Message         string             `json:"message"`
	Tags            []DiagnosticTag    `json:"tags,omitempty"`
}

type DiagnosticSeverity int

const (
	DiagnosticSeverityError       DiagnosticSeverity = 1
	DiagnosticSeverityWarning     DiagnosticSeverity = 2
	DiagnosticSeverityInformation DiagnosticSeverity = 3
	DiagnosticSeverityHint        DiagnosticSeverity = 4
)

type CodeDescription struct {
	Href string `json:"href"`
}

type DiagnosticTag int
