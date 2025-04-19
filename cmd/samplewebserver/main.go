package main

import (
	"fmt"
	"net"
	"strings"
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
		go handleConnection(conn)
	}
}
func handleConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 1024)

	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}

	bufStr := strings.Fields(string(buf))
	println(bufStr[0], bufStr[1], bufStr[2])
	responseStr := bufStr[2] + " 200 OK\r\n\r\n" + "Requested path: " + bufStr[1] + "\r\n"

	conn.Write([]byte(responseStr))
}
