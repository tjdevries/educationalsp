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
		method, contents, err := rpc.DecodeMessageMethod(message)
		if err != nil {
			logger.Println("got err: " + fmt.Sprintf("%+v", err))
			continue
		}

		msg, err := lsp.ToMessage(method, contents)

		switch msg := msg.(type) {
		case lsp.InitializeMessage:
			logger.Printf("Original Message: %s\n", string(message))
			logger.Println(fmt.Sprintf("got initialize request: %d\n", msg.ID))
			serverState.Initialize(&msg)

		case lsp.TextDocumentDidOpen:
			logger.Println(fmt.Sprintf("textDocument/didOpen: %s\n", msg.Params.TextDocument.URI))
			serverState.TextDocumentDidOpen(&msg)

		case lsp.TextDocumentDidChange:
			logger.Println(fmt.Sprintf(
				"textDocument/didChange: %s-> size=%d\n",
				msg.Params.TextDocument.URI,
				len(msg.Params.ContentChanges)))

			serverState.TextDocumentDidChange(&msg)

		case lsp.TextDocumentHover:
			logger.Println(fmt.Sprintf("textDocument/hover: %d\n", msg.ID))
			serverState.TextDocumentHover(&msg)

		case lsp.TextDocumentCodeAction:
			logger.Println(fmt.Sprintf("textDocument/codeAction: %d\n", msg.ID))
			serverState.TextDocumentCodeAction(&msg)

		case lsp.TextDocumentCompletion:
			logger.Println(fmt.Sprintf("textDocument/completion: %d\n", msg.ID))
			serverState.TextDocumentCompletion(&msg)

		default:
			logger.Println("This shouldn't happen:" + string(message))
		}
	}
}
