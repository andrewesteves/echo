package cmd

import (
	"flag"
	"log"
	"os"

	"github.com/andrewesteves/echo/adapters/tcp"
	"github.com/andrewesteves/echo/adapters/udp"
)

var opt string

func init() {
	flag.StringVar(&opt, "opt", "", "Which application to run?")
}

// Run application based on option
func Run() {
	flag.Parse()

	var addr = ":" + os.Getenv("ADDR")

	switch opt {
	case "tcp_server":
		server := &tcp.Server{Addr: addr}
		server.Run()
	case "tcp_client":
		client := &tcp.Client{Addr: addr}
		client.Run()
	case "udp_server":
		server := &udp.Server{Addr: addr}
		server.Run()
	case "udp_client":
		server := &udp.Client{Addr: addr}
		server.Run()
	default:
		log.Fatal("Option not available")
	}
}
