package tcp

import (
	"errors"
	"fmt"
	"net"

	"github.com/andrewesteves/echo/adapters"
)

// Server TCP
type Server struct {
	Addr string
}

// NewServer with settings
func NewServer(addr string) adapters.App {
	return &Server{Addr: addr}
}

// Run TCP
func (server *Server) Run() error {
	if server.Addr == "" {
		return errors.New("Address is required")
	}

	listener, err := net.Listen("tcp", ":"+server.Addr)
	if err != nil {
		return err
	}

	defer listener.Close()

	fmt.Printf("Running on %s\n", server.Addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Connection error %s\n", err.Error())
			continue
		}

		go echo(conn)
	}
}
