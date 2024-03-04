package lsp

type TextDocumentHover struct {
	Request
	Params TextDocumentHoverParams `json:"params"`
}

type TextDocumentHoverParams struct {
	TextDocument TextDocumentIdentifier `json:"textDocument"`
	Position     Position               `json:"position"`
}

type TextDocumentHoverResponse struct {
	Response
	Result TextDocumentHoverResponseResult `json:"result"`
}

type TextDocumentHoverResponseResult struct {
	Contents string `json:"contents"`
}
