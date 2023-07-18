1. "Config": This is the configuration data that will be shared across all files. It will be loaded from the "config.json" file and will contain the settings for the bind address, internal address, and ports. 

2. "GRE Tunnel": This is the main object that will be used to forward traffic. It will be defined in "tunnel/gre.go" and used in "main.go".

3. "Layer7Protection": This is the function that will provide Layer 7 protection. It will be defined in "protection/layer7.go" and used in "main.go".

4. "BindAddress": This is the function that will bind the address. It will be defined in "network/bind_address.go" and used in "main.go".

5. "InternalAddress": This is the function that will shift the packets to the internal address. It will be defined in "network/internal_address.go" and used in "main.go".

6. "Ports": This is the function that will handle the specific ports for moving traffic. It will be defined in "network/ports.go" and used in "main.go".

7. "LoadConfig": This is the function that will load the configuration from the JSON file. It will be defined and used in "main.go".

8. "StartGRE": This is the function that will start the GRE tunnel. It will be defined in "tunnel/gre.go" and used in "main.go".

9. "ProtectLayer7": This is the function that will apply the Layer 7 protection. It will be defined in "protection/layer7.go" and used in "main.go".

10. "Bind": This is the function that will bind the address. It will be defined in "network/bind_address.go" and used in "main.go".

11. "ShiftPackets": This is the function that will shift the packets to the internal address. It will be defined in "network/internal_address.go" and used in "main.go".

12. "MoveTraffic": This is the function that will move the traffic from specific ports. It will be defined in "network/ports.go" and used in "main.go".