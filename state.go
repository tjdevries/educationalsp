package educationlsp

import (
	"fmt"
	"io"
	"strings"
)

type ServerState struct {
	Writer    io.Writer
	Documents map[string]string
}

func (s *ServerState) Initialize(msg *InitializeMessage) {
	// Reply to the initialize request
	response := EncodeMessageStruct(NewInitializeResponse(msg.ID))
	s.Writer.Write([]byte(response))
}

func (s *ServerState) TextDocumentDidOpen(msg *TextDocumentDidOpen) {
	s.Documents[msg.Params.TextDocument.URI] = msg.Params.TextDocument.Text
}

func (s *ServerState) TextDocumentDidChange(msg *TextDocumentDidChange) {
	// We only do full text sync, so this is easy
	text := msg.Params.ContentChanges[0].Text
	s.Documents[msg.Params.TextDocument.URI] = text

	diagnostics := []Diagnostic{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			diagnostics = append(diagnostics, Diagnostic{
				Severity: DiagnosticSeverityError,
				Range: Range{
					Start: Position{
						Line:      row,
						Character: idx,
					},
					End: Position{
						Line:      row,
						Character: idx + len("VS Code"),
					},
				},
				Message: "NEVER, EVER WRITE THAT (uncensored) IN MY EDITOR AGAIN",
			})
		}

		idx = strings.Index(line, "VS C*de")
		if idx >= 0 {
			diagnostics = append(diagnostics, Diagnostic{
				Severity: DiagnosticSeverityWarning,
				Range: Range{
					Start: Position{
						Line:      row,
						Character: idx,
					},
					End: Position{
						Line:      row,
						Character: idx + len("VS C*de"),
					},
				},
				Message: "This is ok... but be careful",
			})
		}
	}

	if len(diagnostics) > 0 {
		publishDiags := PublishDiagnostics{
			Notification: Notification{
				RPC:    "2.0",
				Method: "textDocument/publishDiagnostics",
			},
			Params: PublishDiagnosticsParams{
				URI:         msg.Params.TextDocument.URI,
				Diagnostics: diagnostics,
			},
		}

		s.Writer.Write([]byte(EncodeMessageStruct(publishDiags)))
	}
	// s.Writer.Write
	// s.Writer.Write()
}

func (s *ServerState) TextDocumentHover(msg *TextDocumentHover) {
	documentURI := msg.Params.TextDocument.URI
	contents := s.Documents[documentURI]

	response := TextDocumentHoverResponse{
		Response: Response{
			RPC: "2.0",
			ID:  msg.ID,
		},
		Result: TextDocumentHoverResponseResult{
			Contents: fmt.Sprintf("This is from the LSP: Document has %d characters", len(contents)),
		},
	}
	s.Writer.Write([]byte(EncodeMessageStruct(response)))
}
