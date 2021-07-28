package tcp

import (
	"bufio"
	"net"
	"testing"
)

const addr = ":8090"

func init() {
	server := &Server{Addr: addr}
	go server.Run()
}

func TestTCPEchoServer(t *testing.T) {
	payload := []byte("hello world\n")

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Errorf("Connection error: %s", err.Error())
	}
	defer conn.Close()

	if _, err := conn.Write(payload); err != nil {
		t.Errorf("Error writing data: %s", err.Error())
	}

	res := bufio.NewReader(conn)
	got, err := res.ReadString('\n')
	if err != nil {
		t.Errorf("Error reading response: %s", err.Error())
	}

	want := string(payload)
	if got != want {
		t.Errorf("Error: got %s but we want %s", got, want)
	}
}
