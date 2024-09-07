package rpc_test

import (
	"testing"

	"github.com/agent-e11/educationalsp/rpc"
)

type EncodingExample struct {
	Method string
}

func TestEncode(t *testing.T) {
	expected := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	actual := rpc.EncodeMessage(EncodingExample{Method:"hi"})

	if expected != actual {
		t.Fatalf("Expected: %s, Actual: %s", expected, actual)
	}
}

func TestDecode(t *testing.T) {
	incomingMessage := "Content-Length: 15\r\n\r\n{\"Method\":\"hi\"}"
	method, content, err := rpc.DecodeMessage([]byte(incomingMessage))
	contentLength := len(content)
	if err != nil {
		t.Fatal(err)
	}

	if contentLength != 15 {
		t.Fatalf("Expected: 15, Actual: %d", contentLength)
	}

	if method != "hi" {
		t.Fatalf("Expected: 'hi', Actual: %s", method)
	}
}
