package main

import (
	"bufio"
	"educationalsp/lsp"
	"educationalsp/rpc"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	// Open or create the log file
	logFile, err := os.OpenFile("/home/tjdevries/git/go-lsp/log.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer logFile.Close()

	// Create a logger that writes to the file, with a time-stamp prefix and log.LstdFlags for standard logging flags
	logger := log.New(logFile, "[educationlsp] ", log.Ldate|log.Ltime|log.Lshortfile)

	stdinReader := bufio.NewScanner(os.Stdin)
	stdinReader.Split(rpc.Scan)

	logger.Println("starting LSP server: " + strconv.Itoa(int(time.Now().Unix())))

	serverState := lsp.ServerState{
		Writer:    os.Stdout,
		Documents: make(map[string]string),
	}

	for stdinReader.Scan() {
		message := stdinReader.Bytes()
		base, err := rpc.DecodeMessage(message)
		if err != nil {
			logger.Println("got err: " + fmt.Sprintf("%+v", err))
			continue
		}

		msg, err := lsp.ToMessage(message, base)

		switch v := msg.(type) {
		case lsp.InitializeMessage:
			logger.Println(fmt.Sprintf("got initialize request: %d\n", v.ID))
			serverState.Initialize(&v)

		case lsp.TextDocumentDidOpen:
			logger.Println(fmt.Sprintf("textDocument/didOpen: %s\n", v.Params.TextDocument.URI))
			serverState.TextDocumentDidOpen(&v)

		case lsp.TextDocumentDidChange:
			logger.Println(fmt.Sprintf("textDocument/didChange: %s\n", v.Params.ContentChanges))
			serverState.TextDocumentDidChange(&v)

		case lsp.TextDocumentHover:
			logger.Println(fmt.Sprintf("textDocument/hover: %d\n", v.ID))
			serverState.TextDocumentHover(&v)

		case lsp.TextDocumentCodeAction:
			logger.Println(fmt.Sprintf("textDocument/codeAction: %d\n", v.ID))
			serverState.TextDocumentCodeAction(&v)

		case lsp.TextDocumentCompletion:
			logger.Println(fmt.Sprintf("textDocument/completion: %d\n", v.ID))
			serverState.TextDocumentCompletion(&v)

		case rpc.BaseMessage:
			logger.Println(fmt.Sprintf("Not properly decoded: %d %s\n", v.ID, v.Method))

		default:
			logger.Println("This shouldn't happen:" + string(message))
		}
	}
}
