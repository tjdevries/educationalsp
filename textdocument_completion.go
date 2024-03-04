package educationlsp

type TextDocumentCompletion struct {
	Request
	Params TextDocumentCompletionParams `json:"params"`
}

type TextDocumentCompletionParams struct {
	TextDocumentPositionParams
}

type TextDocumentCompletionResponse struct {
	Response
	Result []CompletionItem `json:"result"`
}

type CompletionItem struct {
	Label         string      `json:"label"`
	Detail        string      `json:"detail,omitempty"`
	Documentation interface{} `json:"documentation,omitempty"`
}
