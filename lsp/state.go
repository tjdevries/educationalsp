package lsp

import (
	"educationalsp/rpc"
	"fmt"
	"io"
	"os"
	"strings"
)

type ServerState struct {
	Writer    io.Writer
	Documents map[string]string
}

func NewServerState() ServerState {
	return ServerState{
		Writer:    os.Stdout,
		Documents: map[string]string{},
	}
}

func (s *ServerState) Initialize(msg *InitializeMessage) {
	// Reply to the initialize request
	response := rpc.EncodeMessage(NewInitializeResponse(msg.ID))
	s.Writer.Write([]byte(response))
}

// emitDiagnosticsForFile sends the diagnostics for the given file to the client,
// if any are found.
func (s *ServerState) emitDiagnosticsForFile(uri string) {
	text := s.Documents[uri]
	diagnostics := []Diagnostic{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			diagnostics = append(diagnostics, Diagnostic{
				Severity: DiagnosticSeverityError,
				Range:    LineRange(row, idx, idx+len("VS Code")),
				Message:  "NEVER, EVER WRITE THAT (uncensored) IN MY EDITOR AGAIN",
			})
		}

		idx = strings.Index(line, "VS C*de")
		if idx >= 0 {
			diagnostics = append(diagnostics, Diagnostic{
				Severity: DiagnosticSeverityWarning,
				Range:    LineRange(row, idx, idx+len("VS C*de")),
				Message:  "This is ok... but be careful",
			})
		}

		idx = strings.Index(line, "What do I do now?")
		if idx >= 0 {
			diagnostics = append(diagnostics, Diagnostic{
				Severity: DiagnosticSeverityInformation,
				Range:    LineRange(row, idx, idx+len("What do I do now?")),
				Message:  "SMASH THAT LIKE BUTTON, CLICK SUBSCRIBE AND RING THAT BELL",
			})
		}
	}

	publishDiags := PublishDiagnostics{
		Notification: Notification{
			RPC:    "2.0",
			Method: "textDocument/publishDiagnostics",
		},
		Params: PublishDiagnosticsParams{
			URI:         uri,
			Diagnostics: diagnostics,
		},
	}

	s.Writer.Write([]byte(rpc.EncodeMessage(publishDiags)))
}

func (s *ServerState) TextDocumentDidOpen(msg *TextDocumentDidOpen) {
	s.Documents[msg.Params.TextDocument.URI] = msg.Params.TextDocument.Text

	s.emitDiagnosticsForFile(msg.Params.TextDocument.URI)
}

func (s *ServerState) TextDocumentDidChange(msg *TextDocumentDidChange) {
	// We only do full text sync, so this is easy
	text := msg.Params.ContentChanges[0].Text
	s.Documents[msg.Params.TextDocument.URI] = text

	s.emitDiagnosticsForFile(msg.Params.TextDocument.URI)
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
	s.Writer.Write([]byte(rpc.EncodeMessage(response)))
}

func (s *ServerState) TextDocumentCodeAction(msg *TextDocumentCodeAction) {
	text := s.Documents[msg.Params.TextDocument.URI]

	actions := []CodeAction{}
	for row, line := range strings.Split(text, "\n") {
		idx := strings.Index(line, "VS Code")
		if idx >= 0 {
			replaceChange := map[string][]TextEdit{}
			replaceChange[msg.Params.TextDocument.URI] = []TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "Neovim",
				},
			}

			actions = append(actions, CodeAction{
				Title: "Replace VS C*de with a superior editor",
				Edit:  &WorkspaceEdit{Changes: replaceChange},
			})

			censorChange := map[string][]TextEdit{}
			censorChange[msg.Params.TextDocument.URI] = []TextEdit{
				{
					Range:   LineRange(row, idx, idx+len("VS Code")),
					NewText: "VS C*de",
				},
			}

			actions = append(actions, CodeAction{
				Title: "Censor to VS C*de",
				Edit:  &WorkspaceEdit{Changes: censorChange},
			})
		}
	}

	response := TextDocumentCodeActionResponse{
		Response: Response{
			RPC: "2.0",
			ID:  msg.ID,
		},
		Result: actions,
	}

	s.Writer.Write([]byte(rpc.EncodeMessage(response)))
}

func (s *ServerState) TextDocumentCompletion(msg *TextDocumentCompletion) {
	// In real life, you'd ask the compiler for a bunch of things that work!
	// or make some guesses about what might work via context of your file/imports/etc.
	//
	// Instead, we'll just return hello world, because we're cool like that

	response := TextDocumentCompletionResponse{
		Response: Response{
			RPC: "2.0",
			ID:  msg.ID,
		},
		Result: []CompletionItem{
			{
				Label:         "Neovim (BTW)",
				Detail:        "Details?",
				Documentation: "This truly do be documentation",
			},
		},
	}
	s.Writer.Write([]byte(rpc.EncodeMessage(response)))
}
