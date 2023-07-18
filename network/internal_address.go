package network

import (
	"net"
	"log"
)

type InternalAddress struct {
	Address string
}

func (ia *InternalAddress) ShiftPackets(conn net.Conn) {
	_, err := conn.Write([]byte(ia.Address))
	if err != nil {
		log.Fatal(err)
	}
}