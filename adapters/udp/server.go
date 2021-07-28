package udp

import (
	"fmt"
	"log"
	"net"
)

// Server UDP
type Server struct {
	Addr string
}

// Run UDP
func (server *Server) Run() {
	listener, err := net.ListenPacket("udp", server.Addr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Printf("Running on %s\n", server.Addr)

	echo(listener)
}
