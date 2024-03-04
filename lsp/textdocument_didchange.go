package lsp

type TextDocumentDidChange struct {
	Notification
	Params TextDocumentDidChangeParams `json:"params"`
}

type TextDocumentDidChangeParams struct {
	TextDocument   VersionedTextDocumentIdentifier `json:"textDocument"`
	ContentChanges []TextDocumentDidChangeEvent    `json:"contentChanges"`
}

type TextDocumentDidChangeEvent struct {
	Text string `json:"text"`
}
