package lsp

type TextDocumentDidOpen struct {
	Notification
	Params TextDocumentDidOpenParams `json:"params"`
}

type TextDocumentDidOpenParams struct {
	TextDocument TextDocumentItem `json:"textDocument"`
}
