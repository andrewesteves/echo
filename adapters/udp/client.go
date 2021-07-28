package udp

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

// Client TCP
type Client struct {
	Addr string
}

// Run client application
func (client *Client) Run() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("Type the data to be sent: ")
		data, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Input error: %s\n", err.Error())
			continue
		}

		conn, err := net.Dial("udp", client.Addr)
		if err != nil {
			log.Fatalf("Connection error: %s\n", err.Error())
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
