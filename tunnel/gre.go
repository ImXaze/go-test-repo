```go
package tunnel

import (
	"net"
	"golang.org/x/net/ipv4"
)

type GRETunnel struct {
	Conn *ipv4.RawConn
}

func NewGRETunnel() *GRETunnel {
	return &GRETunnel{}
}

func (t *GRETunnel) StartGRE(bindAddress string, internalAddress string) error {
	conn, err := net.ListenPacket("ip4:gre", bindAddress)
	if err != nil {
		return err
	}

	rawConn, err := ipv4.NewRawConn(conn)
	if err != nil {
		return err
	}

	t.Conn = rawConn
	return nil
}

func (t *GRETunnel) ForwardTraffic(srcPort, destPort int, internalAddress string) error {
	packet := make([]byte, 4096)
	for {
		_, _, err := t.Conn.ReadFrom(packet)
		if err != nil {
			return err
		}

		// Forward the packet to the internal address
		_, err = t.Conn.WriteTo(packet, &net.IPAddr{IP: net.ParseIP(internalAddress)})
		if err != nil {
			return err
		}
	}
}
```