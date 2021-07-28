package tcp

import (
	"fmt"
	"log"
	"net"
)

// Server TCP
type Server struct {
	Addr string
}

// Run TCP
func (server *Server) Run() {
	listener, err := net.Listen("tcp", server.Addr)
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	fmt.Printf("Running on: %s\n", server.Addr)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Connection error %s\n", err.Error())
			continue
		}

		go echo(conn)
	}
}
