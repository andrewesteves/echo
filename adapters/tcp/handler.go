package tcp

import (
	"fmt"
	"io"
	"net"
)

func echo(conn net.Conn) {
	defer conn.Close()

	_, err := io.Copy(conn, conn)
	if err != nil {
		fmt.Printf("Error copying data in the echo handler: %s\n", err.Error())
		return
	}
}
