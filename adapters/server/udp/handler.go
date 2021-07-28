package udp

import (
	"fmt"
	"net"
)

func echo(conn net.PacketConn) {
	buffer := make([]byte, 2048)

	for {
		_, addr, err := conn.ReadFrom(buffer)
		if err != nil {
			fmt.Println("Error reading packet")
			continue
		}

		if _, err := conn.WriteTo(buffer, addr); err != nil {
			fmt.Println("Error echoing data back")
		}
	}
}
