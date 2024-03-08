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
	stdinReader := bufio.NewScanner(os.Stdin)
	stdinReader.Split(rpc.Scan)

	logger := getLogger()
	logger.Println("starting LSP server: " + strconv.Itoa(int(time.Now().Unix())))

	serverState := lsp.NewServerState()
	for stdinReader.Scan() {
		message := stdinReader.Bytes()
		method, contents, err := rpc.DecodeMessageMethod(message)
		if err != nil {
			logger.Println("Unable to decode message method: " + fmt.Sprintf("%+v", err))
			continue
		}

		msg, err := lsp.ToMessage(method, contents)
		if err != nil {
			logger.Println("Unable to decode message: " + fmt.Sprintf("%+v", err))
			continue
		}

		switch msg := msg.(type) {
		case lsp.InitializeMessage:
			logger.Printf("got initialize request: %d\n", msg.ID)
			serverState.Initialize(&msg)

		case lsp.TextDocumentDidOpen:
			logger.Printf("textDocument/didOpen: %s\n", msg.Params.TextDocument.URI)
			serverState.TextDocumentDidOpen(&msg)

		case lsp.TextDocumentDidChange:
			uri := msg.Params.TextDocument.URI
			textLen := len(msg.Params.ContentChanges)
			logger.Printf("textDocument/didChange: %s -> size=%d\n", uri, textLen)

			serverState.TextDocumentDidChange(&msg)

		case lsp.TextDocumentHover:
			logger.Printf("textDocument/hover: %d\n", msg.ID)
			serverState.TextDocumentHover(&msg)

		case lsp.TextDocumentCodeAction:
			logger.Printf("textDocument/codeAction: %d\n", msg.ID)
			serverState.TextDocumentCodeAction(&msg)

		case lsp.TextDocumentCompletion:
			logger.Printf("textDocument/completion: %d\n", msg.ID)
			serverState.TextDocumentCompletion(&msg)

		default:
			logger.Printf("This shouldn't happen: %s\n", message)
		}
	}
}

func getLogger() *log.Logger {
	// Open or create the log file
	logFile, err := os.OpenFile("/home/tjdevries/git/go-lsp/log.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}

	// Create a logger that writes to the file, with a time-stamp prefix and log.LstdFlags for standard logging flags
	return log.New(logFile, "[educationlsp] ", log.Ldate|log.Ltime|log.Lshortfile)
}
