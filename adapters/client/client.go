package client

import (
	"bufio"
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/andrewesteves/echo/adapters"
)

// Client TCP
type Client struct {
	Protocol string
	Addr     string
}

// NewClient with settings
func NewClient(protocol string, addr string) adapters.App {
	return &Client{
		Protocol: protocol,
		Addr:     addr,
	}
}

// Run client application
func (client *Client) Run() error {
	if client.Protocol == "" {
		return errors.New("Protocol is required")
	}

	if client.Addr == "" {
		return errors.New("Address is required")
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Type the data to be sent: ")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Input error: %s\n", err.Error())
			continue
		}

		conn, err := net.Dial(client.Protocol, ":"+client.Addr)
		if err != nil {
			return fmt.Errorf("Connection error: %s", err.Error())
		}

		defer conn.Close()

		fmt.Fprintf(conn, data)

		response, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Printf("Response error: %s\n", err.Error())
		}

		fmt.Printf("Received: %s", response)
	}
}
