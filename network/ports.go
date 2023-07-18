```go
package network

import (
	"net"
	"strconv"
	"github.com/pkg/errors"
	"../config"
)

func MoveTraffic(fromPort int, toPort int) error {
	ln, err := net.Listen("tcp", ":"+strconv.Itoa(fromPort))
	if err != nil {
		return errors.Wrap(err, "Failed to listen on port")
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			return errors.Wrap(err, "Failed to accept connection")
		}

		go handleConnection(conn, toPort)
	}
}

func handleConnection(conn net.Conn, toPort int) {
	defer conn.Close()

	remote, err := net.Dial("tcp", config.Config.InternalAddress+":"+strconv.Itoa(toPort))
	if err != nil {
		return
	}
	defer remote.Close()

	go copyIO(conn, remote)
	go copyIO(remote, conn)
}

func copyIO(src net.Conn, dst net.Conn) {
	defer src.Close()
	defer dst.Close()

	buf := make([]byte, 1024)
	for {
		n, err := src.Read(buf)
		if err != nil {
			return
		}

		_, err = dst.Write(buf[:n])
		if err != nil {
			return
		}
	}
}
```