package main

import (
	"fmt"
	"net"
	"www.github.com/noahra/samplewebserver/internal/server_utils"
)

func main() {
	ln, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		server_utils.HandleConnection(conn)
	}
}
