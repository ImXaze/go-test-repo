```go
package main

import (
	"encoding/json"
	"os"
	"./network"
	"./protection"
	"./tunnel"
)

type Config struct {
	BindAddress     string `json:"bind_address"`
	InternalAddress string `json:"internal_address"`
	Ports           []int  `json:"ports"`
}

func LoadConfig() Config {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		panic(err)
	}
	return config
}

func main() {
	config := LoadConfig()

	bindAddress := network.Bind(config.BindAddress)
	internalAddress := network.InternalAddress(config.InternalAddress)
	ports := network.Ports(config.Ports)

	greTunnel := tunnel.GRE(config.BindAddress, config.InternalAddress, config.Ports)
	greTunnel.StartGRE()

	protection.Layer7Protection()

	for _, port := range ports {
		greTunnel.MoveTraffic(bindAddress, internalAddress, port)
	}
}
```