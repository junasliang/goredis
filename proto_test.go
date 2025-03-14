package main

import (
	"testing"
)

func TestProtocol(t *testing.T) {
	msg := "*3\r\n$3\r\nSET\r\n$3\r\nkey\r\n$5\r\nvalue\r\n"
	parseCommand(msg)
}
