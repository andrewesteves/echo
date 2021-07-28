package udp

import (
	"bufio"
	"net"
	"testing"
)

const addr = "9090"

func init() {
	server := &Server{Addr: addr}
	go server.Run()
}

func TestUDPEchoServer(t *testing.T) {
	payload := []byte("hello world\n")

	conn, err := net.Dial("udp", ":"+addr)
	if err != nil {
		t.Errorf("Connection error: %s", err.Error())
	}
	defer conn.Close()

	if _, err := conn.Write(payload); err != nil {
		t.Errorf("Error writing data: %s", err.Error())
	}

	got, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Errorf("Error reading response: %s", err.Error())
	}

	want := string(payload)
	if got != want {
		t.Errorf("Error: got %s but we want %s", got, want)
	}
}

func TestUDPEchoServerInvalidSettings(t *testing.T) {
	if NewServer("").Run() == nil {
		t.Error("it should be an invalid server")
	}
}
