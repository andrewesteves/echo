package cmd

import (
	"flag"
	"log"
	"os"

	"github.com/andrewesteves/echo/adapters"
	"github.com/andrewesteves/echo/adapters/client"
	"github.com/andrewesteves/echo/adapters/server/tcp"
	"github.com/andrewesteves/echo/adapters/server/udp"
)

var opt string

func init() {
	flag.StringVar(&opt, "opt", "", "Which application to run?")
}

// Run application based on option
func Run() {
	flag.Parse()

	var addr = os.Getenv("ADDR")
	var protocol = os.Getenv("PROTOCOL")
	var app adapters.App

	switch opt {
	case "tcp":
		app = &tcp.Server{Addr: addr}
	case "udp":
		app = &udp.Server{Addr: addr}
	case "client":
		app = &client.Client{Protocol: protocol, Addr: addr}
	default:
		log.Fatal("Option not available")
	}

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
