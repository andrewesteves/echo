package cmd

import (
	"flag"
	"log"
	"os"

	"github.com/andrewesteves/echo/adapters/tcp"
)

var opt string

func init() {
	flag.StringVar(&opt, "opt", "", "Which application to run?")
}

// Run application based on option
func Run() {
	flag.Parse()

	switch opt {
	case "tcp_server":
		server := &tcp.Server{Addr: os.Getenv("ADDR")}
		server.Run()
	case "tcp_client":
		client := &tcp.Client{Addr: os.Getenv("ADDR")}
		client.Run()
	default:
		log.Fatal("Option not available")
	}
}
