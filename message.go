package educationlsp

type Notification struct {
	RPC    string `json:"jsonrpc"`
	Method string `json:"method"`
}

type Request struct {
	RPC    string `json:"jsonrpc"`
	ID     int    `json:"id"`
	Method string `json:"method"`
}

type Response struct {
	RPC string `json:"jsonrpc"`
	ID  int    `json:"id"`

	// TODO: Handle errors maybe?
	// Error any    `json:"error"`
}

type Position struct {
	Line      int `json:"line"`
	Character int `json:"character"`
}

type Range struct {
	Start Position `json:"start"`
	End   Position `json:"end"`
}

type Location struct {
	URI   string `json:"uri"`
	Range Range  `json:"range"`
}

type WorkspaceEdit struct {
	Changes map[string][]TextEdit `json:"changes"`
}

type TextEdit struct {
	Range   Range  `json:"range"`
	NewText string `json:"newText"`
}

func LineRange(line, start, end int) Range {
	return Range{
		Start: Position{
			Line:      line,
			Character: start,
		},
		End: Position{
			Line:      line,
			Character: end,
		},
	}
}
