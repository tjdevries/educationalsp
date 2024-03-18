package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
)

type baseMessage struct {
	ID     int    `json:"id,omitempty"`
	Method string `json:"method"`
}

const CONTENT_LENGTH = len("Content-Length: ")

// Takes a message and returns a string that can be sent over the wire
func EncodeMessage(message interface{}) string {
	msg, err := json.Marshal(message)
	if err != nil {
		panic("Unable to encode message: " + fmt.Sprintf("%+v", err))
	}

	msgLength := len(msg)
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", msgLength, msg)
}

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

// DecodeMessageMethod takes a message and returns the method and the contents of the message
func DecodeMessageMethod(message []byte) (string, []byte, error) {
	remaining, contentLength, err := parseContentLength(message)
	if err != nil {
		return "", nil, err
	}

	// Contents of the actual message
	content := remaining[:contentLength]

	// Decode the message
	var base baseMessage
	if err := json.Unmarshal(content, &base); err != nil {
		return "", nil, err
	}

	return base.Method, content, err
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
