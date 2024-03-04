package main

import (
	"bufio"
	"educationlsp"
	"fmt"
	"os"
	"strconv"
	"time"
)

func log(s string) {}

func main() {
	stdinReader := bufio.NewScanner(os.Stdin)
	stdinReader.Split(educationlsp.Scan)

	log("starting LSP server: " + strconv.Itoa(int(time.Now().Unix())) + "\n")

	serverState := educationlsp.ServerState{
		Writer:    os.Stdout,
		Documents: make(map[string]string),
	}

	for stdinReader.Scan() {
		message := stdinReader.Bytes()
		msg, err := educationlsp.DecodeMessage(message)
		if err != nil {
			log("got err: " + fmt.Sprintf("%+v", err) + "\n")
			continue
		}

		switch v := msg.(type) {
		case educationlsp.InitializeMessage:
			log(fmt.Sprintf("got initialize request: %d\n", v.ID))
			serverState.Initialize(&v)

		case educationlsp.TextDocumentDidOpen:
			log(fmt.Sprintf("textDocument/didOpen: %s\n", v.Params.TextDocument.URI))
			serverState.TextDocumentDidOpen(&v)

		case educationlsp.TextDocumentDidChange:
			log(fmt.Sprintf("textDocument/didChange: %s\n", v.Params.ContentChanges))
			serverState.TextDocumentDidChange(&v)

		case educationlsp.TextDocumentHover:
			log(fmt.Sprintf("textDocument/hover: %d\n", v.ID))
			serverState.TextDocumentHover(&v)

		case educationlsp.BaseMessage:
			log(fmt.Sprintf("Not properly decoded: %d %s\n", v.ID, v.Method))

		default:
			log("This shouldn't happen:" + string(message) + "\n")
		}
	}
}
