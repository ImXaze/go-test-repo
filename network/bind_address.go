```go
package network

import (
	"net"
	"log"
)

type BindAddress struct {
	Address string
}

func (b *BindAddress) Bind() {
	listener, err := net.Listen("tcp", b.Address)
	if err != nil {
		log.Fatal("Failed to bind address: ", err)
	}
	defer listener.Close()

	log.Println("Successfully binded to address: ", b.Address)
}
```