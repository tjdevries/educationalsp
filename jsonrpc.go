package educationlsp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

// Takes a message and returns a string that can be sent over the wire
func EncodeMessage(msg string) string {
	msgLength := len(msg)
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", msgLength, msg)
}

func EncodeMessageStruct(message interface{}) string {
	msg, err := json.Marshal(message)
	if err != nil {
		panic("hahaha lul this won't happen")
	}

	return EncodeMessage(string(msg))
}

type BaseMessage struct {
	ID     int    `json:"id"`
	Method string `json:"method"`
}

const CONTENT_LENGTH = len("Content-Length: ")

func parseContentLength(msg []byte) ([]byte, int, error) {
	before, after, found := bytes.Cut(msg, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return nil, 0, fmt.Errorf("no content length found")
	}

	contentLengthStr := string(before[CONTENT_LENGTH:])
	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		return nil, 0, err
	}

	return after, contentLength, nil
}
func Scan(data []byte, atEOF bool) (advance int, token []byte, err error) {
	before, after, found := bytes.Cut(data, []byte{'\r', '\n', '\r', '\n'})
	if !found {
		return 0, nil, nil
	}

	contentLengthStr := string(before[CONTENT_LENGTH:])
	contentLength, err := strconv.Atoi(contentLengthStr)
	if err != nil {
		return 0, nil, err
	}

	// We don't have enough data, got wait til later
	if len(after) < contentLength {
		return 0, nil, nil
	}

	totalLength := len(before) + 4 + contentLength
	return totalLength, data[:totalLength], nil
}

func DecodeMessage(contents []byte) (any, error) {
	remaining, contentLength, err := parseContentLength(contents)
	if err != nil {
		return nil, err
	}

	// TODO: ... kind of annoying, what if two messages? why not just read til end? no one knows
	byteContents := remaining[:contentLength]

	var message BaseMessage
	if err := json.Unmarshal(byteContents, &message); err != nil {
		return nil, err
	}

	switch message.Method {
	case "initialize":
		var initializeMessage InitializeMessage
		err := json.Unmarshal(byteContents, &initializeMessage)
		return initializeMessage, err
	case "textDocument/didOpen":
		var parsed TextDocumentDidOpen
		err := json.Unmarshal(byteContents, &parsed)
		return parsed, err
	case "textDocument/didChange":
		var parsed TextDocumentDidChange
		err := json.Unmarshal(byteContents, &parsed)
		return parsed, err
	case "textDocument/hover":
		var parsed TextDocumentHover
		err := json.Unmarshal(byteContents, &parsed)
		return parsed, err
	case "textDocument/codeAction":
		var parsed TextDocumentCodeAction
		err := json.Unmarshal(byteContents, &parsed)
		return parsed, err
	case "textDocument/completion":
		var parsed TextDocumentCompletion
		err := json.Unmarshal(byteContents, &parsed)
		return parsed, err
	default:
		return message, nil
	}

}
