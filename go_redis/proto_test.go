package main

import (
	"fmt"
	"testing"
)

func TestProtocol(t *testing.T) {
	// raw := "*3\r\n$3\r\nSET\r\n$6\r\nleader\r\n$6\r\nksysio\r\n"
	raw := "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$3\r\nbar\r\n"
	cmd, err := parseCommand(raw)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(cmd)
}
