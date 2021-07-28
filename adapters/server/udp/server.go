package udp

import (
	"errors"
	"fmt"
	"net"

	"github.com/andrewesteves/echo/adapters"
)

// Server UDP
type Server struct {
	Addr string
}

// NewServer with settings
func NewServer(addr string) adapters.App {
	return &Server{Addr: addr}
}

// Run UDP
func (server *Server) Run() error {
	if server.Addr == "" {
		return errors.New("Address is required")
	}

	listener, err := net.ListenPacket("udp", ":"+server.Addr)
	if err != nil {
		return err
	}

	defer listener.Close()

	fmt.Printf("Running on %s\n", server.Addr)

	echo(listener)

	return nil
}
