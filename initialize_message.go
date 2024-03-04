package educationlsp

type InitializeMessage struct {
	RPC    string           `json:"jsonrpc"`
	ID     int              `json:"id"`
	Method string           `json:"method"`
	Params InitializeParams `json:"params"`
}

// InitializeParams defines the parameters for the initialize request.
type InitializeParams struct {
	ProcessID             int                    `json:"processId"`
	RootPath              string                 `json:"rootPath,omitempty"`
	RootURI               string                 `json:"rootUri"`
	Capabilities          ClientCapabilities     `json:"capabilities"`
	InitializationOptions map[string]interface{} `json:"initializationOptions,omitempty"`
}

// ClientCapabilities defines the capabilities provided by the client.
type ClientCapabilities struct {
	Workspace    WorkspaceCapabilities    `json:"workspace,omitempty"`
	TextDocument TextDocumentCapabilities `json:"textDocument,omitempty"`
	// Other capabilities can be added here
}

// WorkspaceCapabilities defines workspace-specific capabilities.
type WorkspaceCapabilities struct {
	ApplyEdit              bool                               `json:"applyEdit"`
	DidChangeConfiguration DidChangeConfigurationCapabilities `json:"didChangeConfiguration,omitempty"`
	// Add more workspace capabilities as needed
}

// DidChangeConfigurationCapabilities defines capabilities for configuration change.
type DidChangeConfigurationCapabilities struct {
	DynamicRegistration bool `json:"dynamicRegistration"`
}

// TextDocumentCapabilities defines text document-specific capabilities.
type TextDocumentCapabilities struct {
	Synchronization SynchronizationCapabilities `json:"synchronization,omitempty"`
	// Add more text document capabilities as needed
}

// SynchronizationCapabilities defines capabilities for document synchronization.
type SynchronizationCapabilities struct {
	WillSave          bool `json:"willSave"`
	WillSaveWaitUntil bool `json:"willSaveWaitUntil"`
	DidSave           bool `json:"didSave"`
}

type InitializeResponse struct {
	RPC    string           `json:"jsonrpc"`
	ID     int              `json:"id"`
	Result InitializeResult `json:"result"`
}

type InitializeResult struct {
	Capabilities ServerCapabilities `json:"capabilities"`
	ServerInfo   ServerInfo         `json:"serverInfo"`
}

type ServerCapabilities struct {
	TextDocumentSync   int            `json:"textDocumentSync"`
	HoverProvider      bool           `json:"hoverProvider"`
	CodeActionProvider bool           `json:"codeActionProvider"`
	CompletionProvider map[string]any `json:"completionProvider"`
}

type ServerInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewInitializeResponse(id int) InitializeResponse {
	return InitializeResponse{
		RPC: "2.0",
		ID:  id,
		Result: InitializeResult{
			Capabilities: ServerCapabilities{
				// You can only send me "Full document updates"
				//   If we were a smarter LSP, we would ask for incremental updates
				//   But this is just for education
				TextDocumentSync: 1,

				// We can provide the following features
				HoverProvider:      true,
				CodeActionProvider: true,
				CompletionProvider: map[string]any{},
			},
			ServerInfo: ServerInfo{
				Name:    "educationalsp",
				Version: "0.0.0.0.0-beta",
			},
		},
	}
}
