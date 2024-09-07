package rpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

func EncodeMessage(msg any) string {
	content, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("Content-Length: %d\r\n\r\n%s", len(content), content)
}

type BaseMessage struct {
	Method string `json:"method"`
}

func DecodeMessage(msg []byte) (string, []byte, error) {
	header, content, found := bytes.Cut(msg, []byte("\r\n\r\n"))
	if !found {
		return "", nil, errors.New("No separator found")
	}

	// Content-Length: <number>
	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		return "", nil, err
	}

	// TODO: We'll get to this eventually
	_ = content

	var baseMessage BaseMessage
	if err := json.Unmarshal(content[:contentLength], &baseMessage); err != nil {
		return "", nil, err
	}

	return baseMessage.Method, content[:contentLength], nil
}

func Split(data []byte, _ bool) (advance int, token []byte, err error) {
	header, content, found := bytes.Cut(data, []byte("\r\n\r\n"))
	if !found {
		// \r\n\r\n was not found.
		// Full header hasn't been recieved.
		// Wait for more data.
		return 0, nil, nil
	}

	contentLengthBytes := header[len("Content-Length: "):]
	contentLength, err := strconv.Atoi(string(contentLengthBytes))
	if err != nil {
		// There was an error converting value of Content-Length to int
		return 0, nil, err
	}

	if len(content) < contentLength {
		// Length of given content is less than length of expected content.
		// The full message hasn't been recieved.
		// Wait for more data.
		return 0, nil, nil
	}

	totalLength := len(header) + 4 + contentLength
	return totalLength, data[:totalLength], nil
}
